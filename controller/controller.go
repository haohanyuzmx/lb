package controller

import (
	"context"
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"math"
	"my.domain/lb/util"
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

func (r *ControllerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) { //todo:事务
	_ = log.FromContext(ctx)
	var err error
	fmt.Println("run")
	//onFail:= func() {  todo:失败后的恢复函数
	//
	//}

	vs := &lbv1.VirtualServer{}
	if err = r.Get(ctx, req.NamespacedName, vs); err != nil {
		return ctrl.Result{}, err
	}
	//check(vs)
	if vs.Status.VIP != "" && vs.Status.Backup.String() != `/` && vs.Status.Master.String() != `/` {
		return ctrl.Result{}, nil
	}

	net := vs.Spec.VirtualNetwork
	vn := &lbv1.VirNet{}

	if err = r.Get(ctx, util.NamespacedName2types(net), vn); err != nil {
		return ctrl.Result{}, err
	}
	vip := vn.Alloc()
	if vip == "" {
		return ctrl.Result{}, errors.New("not enough vip")
	}
	if err = r.Status().Update(ctx, vn); err != nil {
		return ctrl.Result{}, err
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
	if err = r.Status().Update(ctx, vs); err != nil {
		return ctrl.Result{}, err
	}

	masters, backups, err := r.getvm(vn)
	if err != nil {
		return ctrl.Result{}, err
	}
	master, backup := &lbv1.VM{}, &lbv1.VM{}

	if err = r.Get(ctx, masters, master); err != nil {
		return ctrl.Result{}, err
	}
	master.AddMaster(req.NamespacedName)
	if err = r.Update(ctx, master); err != nil {
		return ctrl.Result{}, err
	}
	vs.Status.Master = util.Types2NamespacedName(masters)

	if backups.String() != `/` {
		if err = r.Get(ctx, backups, backup); err != nil {
			return ctrl.Result{}, err
		}
		backup.AddBackup(req.NamespacedName)
		if err = r.Update(ctx, backup); err != nil {
			return ctrl.Result{}, err
		}
		vs.Status.Backup = util.Types2NamespacedName(backups)
	}

	if err = r.Status().Update(ctx, vs); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
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
	var backup util.NamespacedName
	if err := r.Get(context.Background(), util.NamespacedName2types(backupnic), nic); err != nil {
		fmt.Println(err)
	} else {
		backup = nic.Spec.VM
	}

	return util.NamespacedName2types(master), util.NamespacedName2types(backup), nil
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

	if err = c.Watch(&source.Kind{Type: &lbv1.VirNet{}}, &handler.Funcs{}); err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &lbv1.NIC{}}, &handler.Funcs{}); err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &lbv1.VirtualServer{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}

	return nil
}
