package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	lbv1 "my.domain/lb/api/v1"
	"my.domain/lb/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type agentReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	ServerPools    map[string]*Server
	VirtualServers map[string]*VirtualServer
}

type VirtualServer struct {
	VAddr  string
	Status string
	Server *Server
}
type Server struct {
	Members []Member
}
type Member struct {
	IPAddr      string
	MonitorPort int
	Weight      int
}

func (r *agentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//var vm lbv1.VM
	//r.Get(ctx, req.NamespacedName, &vm)
	//
	//VirSer := r.getlb(vm)
	//update := func(applyfunc func(any) bool, index types.NamespacedName, some client.Object) {
	//	r.Get(context.Background(), index, some)
	//	ischange := applyfunc(some)
	//	if ischange {
	//		r.Status().Update(context.Background(), some)
	//	}
	//}

	//apply(VirSer, update)
	return ctrl.Result{}, nil

}

func (r *agentReconciler) getlb(vm lbv1.VM) []*VirtualServer {
	vses := []*VirtualServer{}
	for _, vs := range vm.Spec.VirServers {
		var virSer lbv1.VirtualServer

		r.Get(context.Background(), util.String2NamespacedName(vs.VS), &virSer)
		if virSer.Status.VIP == "" {

		}
		poolName := virSer.Spec.ServerPools

		serverPool := lbv1.ServerPool{}
		r.Get(context.Background(), util.String2NamespacedName(poolName), &serverPool)
		members := []Member{}
		for _, member := range serverPool.Spec.Members {
			members = append(members, Member{
				IPAddr:      member.IPAddr,
				MonitorPort: member.MonitorPort,
				Weight:      member.Weight,
			})
		}
		server := &Server{Members: members}
		r.ServerPools[poolName] = server

		v := &VirtualServer{
			VAddr:  virSer.Status.VIP,
			Status: vs.Identity,
			Server: server,
		}
		r.VirtualServers[vs.VS] = v
		vses = append(vses, v)
	}
	return vses
}
