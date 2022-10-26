package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"math"
	lbv1 "my.domain/lb/api/v1"
	"my.domain/lb/util"
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

	var vs lbv1.VirtualServer

	if err := r.Get(ctx, req.NamespacedName, &vs); err != nil {
		return ctrl.Result{}, err
	}

	net := vs.Spec.VIPVirNet
	var vn *lbv1.VirNet

	if err := r.Get(ctx, util.String2NamespacedName(net), vn); err != nil {
		return ctrl.Result{}, err
	}
	vip := vn.GetVIP()
	if err := r.Status().Update(ctx, vn); err != nil {
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
	if err := r.Status().Update(ctx, &vs); err != nil {
		return ctrl.Result{}, err
	}

	masters, backups := r.getvm(vn)
	var master, backup *lbv1.VM

	if err := r.Get(ctx, util.String2NamespacedName(masters), master); err != nil {
		return ctrl.Result{}, err
	}
	//master.AddMaster()
	if err := r.Update(ctx, master); err != nil {
		return ctrl.Result{}, err
	}

	if err := r.Get(ctx, util.String2NamespacedName(backups), backup); err != nil {
		return ctrl.Result{}, err
	}
	//backup.AddBackup()
	if err := r.Update(ctx, backup); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ControllerReconciler) getvm(net *lbv1.VirNet) (string, string) {
	nics := lbv1.NICList{}
	r.List(context.Background(), &nics, client.MatchingLabels{"net": net.Name})
	var min, index = math.MaxInt8, -1
	for i, nic := range nics.Items {
		sum := 0
		for _, num := range nic.Status.VIP {
			sum += num
		}
		if sum < min {
			index = i
		}
	}
	master := nics.Items[index].Labels["vm"]
	return master, ""
}
