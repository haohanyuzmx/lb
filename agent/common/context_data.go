package common

import (
	lbv1 "my.domain/lb/api/v1"
)

// the full config from crd and runtime data for one virtual server
type VirtualServer struct {
	Index       string
	Enabled     bool
	Name        string
	Protocol    string
	Port        int
	ServerPools []ServerPool
	AppProfile  ApplicationProfile
	VirtualAddr string
	Role        string
	Nic         NIC
}

type ServerPool struct {
	Index     string
	Name      string
	Algorithm string
	Members   []PoolMember
	Monitor   Monitor
}

type PoolMember struct {
	ServerAddr string
	ServerPort int
	Weight     int
}

type Monitor struct {
	Index    string
	Name     string
	Type     string
	Interval int
	Timeout  int
	Method   string
	URL      string
}

type ApplicationProfile struct {
	Index              string
	SessionPersistence bool
	AccessControl      string
	TrafficControl     []int
}

type VM struct {
	Index    string
	Hostname string
	NICs     []NIC
}

type NIC struct {
	Index  string
	Name   string
	Number int
	LIP    []string
	VIP    []string
}

func TransferPoolFromCrd(namespaceName string, src lbv1.ServerPool) *ServerPool {

	members := []PoolMember{}
	for _, member := range src.Spec.Members {
		members = append(members, PoolMember{
			ServerAddr: member.ServerAddr,
			ServerPort: member.ServerPort,
			Weight:     member.Weight,
		})
	}
	pool := &ServerPool{
		Index:     namespaceName,
		Name:      src.Spec.Name,
		Algorithm: src.Spec.Algorithm,
		Members:   members,
	}
	return pool
}

func TransferVSFromCrd(namespaceName string, role string, src lbv1.VirtualServer, serverPools []ServerPool) *VirtualServer {

	return &VirtualServer{
		Index:       namespaceName,
		Enabled:     src.Spec.Enabled,
		Name:        src.Spec.Name,
		Protocol:    src.Spec.Protocol,
		Port:        src.Spec.Port,
		VirtualAddr: src.Status.VIP,
		Role:        role,
		ServerPools: serverPools,
	}
}
