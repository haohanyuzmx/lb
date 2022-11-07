package agent

import (
	"context"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	agentcommon "my.domain/lb/pkg/agent/common"
	"my.domain/lb/pkg/agent/constants"
	lbv1 "my.domain/lb/pkg/apis/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type AgentReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	LbService   *LBService
	Initialized bool
}

var (
	hostname, _ = os.Hostname()
)

func (a *AgentReconciler) getServerPool(NamespacedName types.NamespacedName) *agentcommon.ServerPool {

	var serverPool lbv1.ServerPool
	a.Get(context.Background(), NamespacedName, &serverPool)
	return agentcommon.TransferPoolFromCrd(NamespacedName.String(), serverPool)
}

func (a *AgentReconciler) getVirtualServer(NamespacedName types.NamespacedName, role string) *agentcommon.VirtualServer {

	var virtualServer lbv1.VirtualServer
	a.Get(context.Background(), NamespacedName, &virtualServer)
	if virtualServer.Status.VIP == "" {
		return nil
	}
	poolNamespacedName := virtualServer.Spec.DefaultServerPool
	defaultServerPool := a.getServerPool(poolNamespacedName.IntoTypes())
	var serverPools []agentcommon.ServerPool
	serverPools = append(serverPools, *defaultServerPool)
	//TODO multiple server pools may exist for l7 lb
	return agentcommon.TransferVSFromCrd(NamespacedName.String(), role, virtualServer, serverPools)
}

func (a *AgentReconciler) getLbToUpdate(vm lbv1.VM) ([]*agentcommon.VirtualServer, []string) {
	virtualServers := []*agentcommon.VirtualServer{}

	var existedVirtualServerIds []string
	for _, vsNamespacedName := range vm.Status.MasterLBs {
		if a.LbService.isExisted(vsNamespacedName.ToString()) {
			existedVirtualServerIds = append(existedVirtualServerIds, vsNamespacedName.ToString())
			continue
		}
		vs := a.getVirtualServer(vsNamespacedName.IntoTypes(), constants.Master.String())
		if vs == nil {
			continue
		}
		virtualServers = append(virtualServers, vs)
	}

	for _, vsNamespacedName := range vm.Status.BackupLBs {
		if a.LbService.isExisted(vsNamespacedName.ToString()) {
			existedVirtualServerIds = append(existedVirtualServerIds, vsNamespacedName.ToString())
			continue
		}
		vs := a.getVirtualServer(vsNamespacedName.IntoTypes(), constants.Backup.String())
		if vs == nil {
			continue
		}
		virtualServers = append(virtualServers, vs)
	}
	deleteVSIds := a.LbService.GetVSToDelete(existedVirtualServerIds)

	return virtualServers, deleteVSIds
}

func (a *AgentReconciler) fullUpdate(vm lbv1.VM) (ctrl.Result, error) {
	err := a.LbService.CleanInstances()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to initialize current LBS.")
	}
	lbs := a.GetLbByVm(vm)
	lbIds, err := a.LbService.CreateInstances(lbs)
	if err != nil {
		fmt.Printf("failed to create lbs %s, error: %s", lbIds, err.Error())
		return ctrl.Result{}, err
	}
	fmt.Printf("succeed to create lbs %s", lbIds)

	return ctrl.Result{}, nil
}

func (a *AgentReconciler) GetLbByVm(vm lbv1.VM) []*agentcommon.VirtualServer {
	virtualServers := []*agentcommon.VirtualServer{}

	for _, vsNamespacedName := range vm.Status.MasterLBs {
		vs := a.getVirtualServer(vsNamespacedName.IntoTypes(), constants.Master.String())
		if vs == nil {
			continue
		}
		virtualServers = append(virtualServers, vs)
	}

	for _, vsNamespacedName := range vm.Status.BackupLBs {
		vs := a.getVirtualServer(vsNamespacedName.IntoTypes(), constants.Backup.String())
		if vs == nil {
			continue
		}
		virtualServers = append(virtualServers, vs)
	}
	return virtualServers
}

func (a *AgentReconciler) ReconcileServerPool(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var serverPool lbv1.ServerPool
	a.Get(ctx, req.NamespacedName, &serverPool)
	//TODO
	return ctrl.Result{}, nil
}

func (a *AgentReconciler) ReconcileVirtualServer(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var virtualServer lbv1.VirtualServer
	a.Get(ctx, req.NamespacedName, &virtualServer)

	vsNamespacedName := req.NamespacedName.String()
	if !a.LbService.isExisted(vsNamespacedName) {
		fmt.Printf("skip to update LBS %s.", vsNamespacedName)
		return ctrl.Result{}, nil
	}

	vs := a.getVirtualServer(req.NamespacedName, constants.Master.String())
	if vs == nil {
		return ctrl.Result{}, fmt.Errorf("failed to get crd for virtual server: %s", vsNamespacedName)
	}
	var virtualServers []*agentcommon.VirtualServer
	virtualServers = append(virtualServers, vs)
	err := a.LbService.UpdateInstances(virtualServers)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to get update virtual server: %s", vsNamespacedName)
	}

	return ctrl.Result{}, nil
}

func (a *AgentReconciler) ReconcileVM(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var vm lbv1.VM
	a.Get(ctx, req.NamespacedName, &vm)
	if vm.Spec.Hostname != hostname {
		return ctrl.Result{}, nil
	}
	if a.Initialized == false {
		_, err := a.fullUpdate(vm)
		if err != nil {
			return ctrl.Result{}, err
		}
		a.Initialized = true
		return ctrl.Result{}, nil
	}
	lbsToCreate, lbsToDelete := a.getLbToUpdate(vm) //for L4, create and delete operations can be merged
	if lbsToCreate != nil {
		lbIds, err := a.LbService.CreateInstances(lbsToCreate)
		if err != nil {
			fmt.Printf("failed to create lbs %s, error: %s", lbIds, err.Error())
			return ctrl.Result{}, err
		}
		fmt.Printf("succeed to create lbs %s", lbIds)
	}
	if len(lbsToDelete) != 0 {
		err := a.LbService.DeleteInstances(lbsToDelete)
		if err != nil {
			fmt.Printf("failed to delete lbs %s, error: %s", lbsToDelete, err.Error())
			return ctrl.Result{}, err
		}
		fmt.Printf("succeed to delete lbs %s", lbsToDelete)
	}
	return ctrl.Result{}, nil

}

func (a *AgentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if mgr == nil {
		return fmt.Errorf("can't setup with nil manager")
	}

	c, err := controller.New("VM-controller", mgr, controller.Options{
		MaxConcurrentReconciles: constants.DefaultMaxConcurrentReconciles,
		Reconciler:              reconcile.Func(a.ReconcileVM),
	})
	if err != nil {
		return err
	}

	if err := c.Watch(&source.Kind{Type: &lbv1.VM{}}, &handler.Funcs{}); err != nil {
		return err
	}
	c, err = controller.New("virtualServer-controller", mgr, controller.Options{
		MaxConcurrentReconciles: constants.DefaultMaxConcurrentReconciles,
		Reconciler:              reconcile.Func(a.ReconcileVirtualServer),
	})
	if err != nil {
		return err
	}
	if err := c.Watch(&source.Kind{Type: &lbv1.VirtualServer{}}, &handler.Funcs{}); err != nil {
		return err
	}
	c, err = controller.New("serverPool-controller", mgr, controller.Options{
		MaxConcurrentReconciles: constants.DefaultMaxConcurrentReconciles,
		Reconciler:              reconcile.Func(a.ReconcileServerPool),
	})
	if err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &lbv1.ServerPool{}}, &handler.Funcs{}); err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &lbv1.Monitor{}}, &handler.Funcs{}); err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &lbv1.ApplicationProfile{}}, &handler.Funcs{}); err != nil {
		return err
	}
	return nil
}
