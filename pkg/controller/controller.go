package controller

import (
	"context"
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"math"
	lbv1 "my.domain/lb/pkg/apis/v1"
	"my.domain/lb/pkg/common"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ControllerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *ControllerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	var err error

	vs := &lbv1.VirtualServer{}
	if err = r.Get(ctx, req.NamespacedName, vs); err != nil {
		return ctrl.Result{}, err
	}
	if vs.DeletionTimestamp != nil {
		return ctrl.Result{}, r.deleteVS(ctx, vs) //delete出错可以重试
	}

	if vs.Status.NowVirNet != vs.Spec.VirtualNetwork { //更改了vip所在虚拟机网络
		//wait.Until()

		if err = r.deleteVS(ctx, vs); err != nil {
			fmt.Println(err)
		}
	}

	if vs.Status.Backup.IntoString() != `/` && vs.Status.Master.IntoString() != `/` { //master和backup分配成功
		return ctrl.Result{}, nil
	}

	if err = r.addVS(ctx, vs); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ControllerReconciler) deleteVS(ctx context.Context, vs *lbv1.VirtualServer) error { //幂等
	var err error
	vsindex := vs.Namespace + `/` + vs.Name
	if vs.Status.NowVirNet.IntoString() != `/` {
		vip := vs.Status.VIP
		vn := &lbv1.VirNet{}
		if err = r.Get(ctx, vs.Status.NowVirNet.IntoTypes(), vn); err != nil {
			return err
		}
		vn.Free(vip)
		r.Status().Update(ctx, vn)
	}
	if vs.Status.Master.IntoString() != `/` {
		master := &lbv1.VM{}
		if err = r.Get(ctx, vs.Status.Master.IntoTypes(), master); err != nil {
			return err
		}
		index := -1
		for i, namespacedName := range master.Status.MasterLBs {
			if namespacedName.IntoString() == vsindex {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Println("vm master wrong, should have vs")
		} else {
			master.Status.MasterLBs = append(master.Status.MasterLBs[:index], master.Status.MasterLBs[index+1:]...)
		}
		r.Status().Update(ctx, master)
	}
	if vs.Status.Backup.IntoString() != `/` {
		backup := &lbv1.VM{}
		if err = r.Get(ctx, vs.Status.Backup.IntoTypes(), backup); err != nil {
			return err
		}
		index := -1
		for i, namespacedName := range backup.Status.BackupLBs {
			if namespacedName.IntoString() == vsindex {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Println("vm backup wrong, should have vs")
		} else {
			backup.Status.BackupLBs = append(backup.Status.BackupLBs[:index], backup.Status.BackupLBs[index+1:]...)
		}
		r.Status().Update(ctx, backup)
	}
	vs.Status = lbv1.VirtualServerStatus{}
	return nil
}

func (r *ControllerReconciler) addVS(ctx context.Context, vs *lbv1.VirtualServer) error { //
	//todo:事务
	//onFail:= func() {  todo:失败后的恢复函数
	//
	//}
	var err error
	net := vs.Spec.VirtualNetwork
	vn := &lbv1.VirNet{}

	if err = r.Get(ctx, net.IntoTypes(), vn); err != nil {
		return err
	}
	vip := vn.Alloc()
	if vip == "" {
		return errors.New("not enough vip")
	}
	if err = r.Status().Update(ctx, vn); err != nil {
		//todo:后续失败需要考虑free vip
		return err
	}
	defer func() {
		if err != nil {

		}
	}()

	//stop:=make(chan struct{})
	//wait.Until(func() {
	//	if err := r.Get(ctx, util.String2NamespacedName(net), vn);err != nil {
	//		return
	//	}
	//	if err:=r.Status().Update(ctx, vn);err!=nil{
	//		return
	//	}
	//	stop<- struct{}{}
	//},time.Second,stop)

	vs.Status.VIP = vip
	vs.Status.NowVirNet = vs.Spec.VirtualNetwork
	//todo:考虑两段性更新(先更新vip，vn;然后更新master，backup)

	masters, backups, err := r.getvm(vn)
	if err != nil {
		return err
	}
	master, backup := &lbv1.VM{}, &lbv1.VM{}

	if err = r.Get(ctx, masters, master); err != nil {
		return err
	}

	master.AddMaster(types.NamespacedName{Namespace: vs.Namespace, Name: vs.Name})
	if err = r.Status().Update(ctx, master); err != nil {
		//master没成功更新不用考虑backup更新
		return err
	}
	vs.Status.Master.FromTypes(masters)

	if backups.String() != `/` { //todo: 标记重试更新backup
		if err = r.Get(ctx, backups, backup); err != nil {
			return err
		}
		backup.AddBackup(types.NamespacedName{Namespace: vs.Namespace, Name: vs.Name})
		if err = r.Status().Update(ctx, backup); err != nil {
			return err
		}
		vs.Status.Backup.FromTypes(backups)
	}

	if err = r.Status().Update(ctx, vs); err != nil {
		return err
	}
	return nil
}

func (r *ControllerReconciler) getvm(net *lbv1.VirNet) (types.NamespacedName, types.NamespacedName, error) {
	nics := lbv1.NICList{}
	r.List(context.Background(), &nics, client.MatchingLabels{"net": net.Name})
	var min, index = math.MaxInt8, -1
	for i, nic := range nics.Items {
		if nic.Status.Load < min {
			index = i
		}
	}
	if index == -1 {
		return types.NamespacedName{}, types.NamespacedName{}, errors.New("no nic")
	}

	master := nics.Items[index].Spec.VM
	backupnic := nics.Items[index].Spec.Master
	nic := &lbv1.NIC{}
	var backup common.NamespacedName
	if err := r.Get(context.Background(), backupnic.IntoTypes(), nic); err != nil {
		fmt.Println(err)
	} else {
		backup = nic.Spec.VM
	}

	return master.IntoTypes(), backup.IntoTypes(), nil
}

func (r *ControllerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&lbv1.VirtualServer{}).
		Complete(r)
}
