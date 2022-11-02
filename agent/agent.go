package agent

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"my.domain/lb/agent/common"
	lbv1 "my.domain/lb/api/v1"
	"my.domain/lb/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type agentReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	VirtualServers map[string]*common.VirtualServer
	keepalived     *Keepalived
}

func (r *agentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var vm lbv1.VM
	r.Get(ctx, req.NamespacedName, &vm)

	//
	virtualServer := r.getlb(vm)
	//update := func(applyfunc func(any) bool, index types.NamespacedName, some client.Object) {
	//	r.Get(context.Background(), index, some)
	//	ischange := applyfunc(some)
	//	if ischange {
	//		r.Status().Update(context.Background(), some)
	//	}
	//}

	_, err := r.keepalived.UpdateConfig(vm.Status.Hostname, virtualServer)
	if err != nil {
		fmt.Printf("failed to generate config %s, error: %s", req.Name, err.Error())
	}
	return ctrl.Result{}, nil

}

func (r *agentReconciler) getlb(vm lbv1.VM) []*common.VirtualServer {
	vses := []*common.VirtualServer{}
	for _, vs := range vm.Spec.VirtualServers {
		var virSer lbv1.VirtualServer

		r.Get(context.Background(), util.String2NamespacedName(vs.Index), &virSer)
		if virSer.Status.VIP == "" {

		}
		poolName := virSer.Spec.DefaultServerPool

		serverPool := lbv1.ServerPool{}
		r.Get(context.Background(), poolName, &serverPool)
		members := []common.PoolMember{}
		for _, member := range serverPool.Spec.Members {
			members = append(members, common.PoolMember{
				ServerAddr: member.ServerAddr,
				Port:       member.ServerPort,
				Weight:     member.Weight,
			})
		}
		//server := &common.ServerPool{Members: members}
		//r.ServerPools[poolName] = server

		v := &common.VirtualServer{
			VirtualAddr: virSer.Status.VIP,
			Role:        vs.Role,
			//	Server:      server,
		}
		r.VirtualServers[vs.Index] = v
		vses = append(vses, v)
	}
	return vses
}
