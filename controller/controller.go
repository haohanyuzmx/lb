package controller

import (
	"context"
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"math"
	"my.domain/lb/common"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"k8s.io/apimachinery/pkg/runtime"
	lbv1 "my.domain/lb/api/v1"
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

	}

	//check(vs)
	if vs.Status.VIP != "" && vs.Status.Backup.String() != `/` && vs.Status.Master.String() != `/` {
		return ctrl.Result{}, nil
	}

	if err = r.addVS(ctx, vs); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ControllerReconciler) deleteVS(ctx context.Context, vs *lbv1.VirtualServer) error {
	var err error
	vsindex := vs.Namespace + `/` + vs.Name
	if vs.Status.NowVirNet.String() != `/` {
		vip := vs.Status.VIP
		vn := &lbv1.VirNet{}
		if err = r.Get(ctx, vs.Status.NowVirNet.Into(), vn); err != nil {
			return err
		}
		vn.Free(vip)
	}
	if vs.Status.Master.String() != `/` {
		master := &lbv1.VM{}
		if err = r.Get(ctx, vs.Status.Master.Into(), master); err != nil {
			return err
		}
		index := -1
		for i, namespacedName := range master.Spec.Master {
			if namespacedName.String() == vsindex {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Println("vm master wrong, should have vs")
		} else {
			master.Spec.Master = append(master.Spec.Master[:index], master.Spec.Master[index+1:]...)
		}
	}
	if vs.Status.Backup.String() != `/` {
		backup := &lbv1.VM{}
		if err = r.Get(ctx, vs.Status.Backup.Into(), backup); err != nil {
			return err
		}
		index := -1
		for i, namespacedName := range backup.Spec.Master {
			if namespacedName.String() == vsindex {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Println("vm backup wrong, should have vs")
		} else {
			backup.Spec.Master = append(backup.Spec.Master[:index], backup.Spec.Master[index+1:]...)
		}
	}
	return nil
}

func (r *ControllerReconciler) addVS(ctx context.Context, vs *lbv1.VirtualServer) error { //todo:事务
	//onFail:= func() {  todo:失败后的恢复函数
	//
	//}
	var err error
	net := vs.Spec.VirtualNetwork
	vn := &lbv1.VirNet{}

	if err = r.Get(ctx, net.Into(), vn); err != nil {
		return err
	}
	vip := vn.Alloc()
	if vip == "" {
		return errors.New("not enough vip")
	}
	if err = r.Status().Update(ctx, vn); err != nil {
		return err
	}

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

	masters, backups, err := r.getvm(vn)
	if err != nil {
		return err
	}
	master, backup := &lbv1.VM{}, &lbv1.VM{}

	if err = r.Get(ctx, masters, master); err != nil {
		return err
	}

	master.AddMaster(types.NamespacedName{Namespace: vs.Namespace, Name: vs.Name})
	if err = r.Update(ctx, master); err != nil {
		//master没成功更新不用考虑backup更新
		return err
	}
	vs.Status.Master.FromTypes(masters)

	if backups.String() != `/` { //todo: 标记重试更新backup
		if err = r.Get(ctx, backups, backup); err != nil {
			return err
		}
		backup.AddBackup(types.NamespacedName{Namespace: vs.Namespace, Name: vs.Name})
		if err = r.Update(ctx, backup); err != nil {
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
	if err := r.Get(context.Background(), backupnic.Into(), nic); err != nil {
		fmt.Println(err)
	} else {
		backup = nic.Spec.VM
	}

	return master.Into(), backup.Into(), nil
}

func (r *ControllerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if mgr == nil {
		return fmt.Errorf("can't setup with nil manager")
	}

	c, err := controller.New("ctr-controller", mgr, controller.Options{
		Reconciler: r,
	})
	if err != nil {
		return err
	}

	//if err = c.Watch(&source.Kind{Type: &lbv1.VirNet{}}, &handler.Funcs{}); err != nil {
	//	return err
	//}
	//if err = c.Watch(&source.Kind{Type: &lbv1.NIC{}}, &handler.Funcs{}); err != nil {
	//	return err
	//}
	if err = c.Watch(&source.Kind{Type: &lbv1.VirtualServer{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}

	return nil
}
