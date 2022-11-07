package agent

import (
	"fmt"

	"my.domain/lb/pkg/agent/common"
	"my.domain/lb/pkg/agent/constants"
	"my.domain/lb/pkg/agent/layer4"
	"my.domain/lb/pkg/util"
)

type LBService struct {
	currentL4VirtualServerIds []string
	currentL7VirtualServerIds []string
	currentL4VirtualServers   []*common.VirtualServer
	currentL7VirtualServers   []*common.VirtualServer

	keepalived *layer4.Keepalived
}

func (l *LBService) isExisted(vsIndex string) bool {
	return util.Contains(vsIndex, l.currentL4VirtualServerIds) || util.Contains(vsIndex, l.currentL7VirtualServerIds)
}

func (l *LBService) GetVSToDelete(reservedVSIds []string) []string {
	var deleteVSIds []string
	for _, vsId := range l.currentL4VirtualServerIds {
		if !util.Contains(vsId, reservedVSIds) {
			deleteVSIds = append(deleteVSIds, vsId)
		}
	}
	for _, vsId := range l.currentL7VirtualServerIds {
		if !util.Contains(vsId, reservedVSIds) {
			deleteVSIds = append(deleteVSIds, vsId)
		}
	}
	return deleteVSIds
}

func classifyVirtualServer(virtualServers []*common.VirtualServer) ([]*common.VirtualServer, []*common.VirtualServer, []string, []string) {
	var l4VirtualServerIds []string
	var l7VirtualServerIds []string
	var l4VirtualServers []*common.VirtualServer
	var l7VirtualServers []*common.VirtualServer

	for _, virtualServer := range virtualServers {
		if virtualServer.Protocol == constants.LbTypeL4 {
			l4VirtualServerIds = append(l4VirtualServerIds, virtualServer.Index)
			l4VirtualServers = append(l4VirtualServers, virtualServer)
		}
	}
	return l4VirtualServers, l7VirtualServers, l4VirtualServerIds, l7VirtualServerIds
}

func (l *LBService) classifyVirtualServerById(virtualServerIds []string) ([]string, []string) {
	var l4VirtualServerIds []string
	var l7VirtualServerIds []string

	for _, virtualServerId := range virtualServerIds {
		if util.Contains(virtualServerId, l.currentL4VirtualServerIds) {
			l4VirtualServerIds = append(l4VirtualServerIds, virtualServerId)
		} else if util.Contains(virtualServerId, l.currentL7VirtualServerIds) {
			l7VirtualServerIds = append(l7VirtualServerIds, virtualServerId)
		}
	}
	return l4VirtualServerIds, l7VirtualServerIds
}

func (l *LBService) DeleteInstances(virtualServerIds []string) error {
	l4VirtualServerIds, _ := l.classifyVirtualServerById(virtualServerIds)

	if len(l4VirtualServerIds) != 0 {
		var newL4VirtualServers []*common.VirtualServer
		var newL4VirtualServerIds []string
		for _, vs := range l.currentL4VirtualServers {
			if !util.Contains(vs.Index, l4VirtualServerIds) {
				newL4VirtualServers = append(newL4VirtualServers, vs)
				newL4VirtualServerIds = append(newL4VirtualServerIds, vs.Index)
			}
		}
		_, err := l.keepalived.Update(newL4VirtualServers)
		if err != nil {
			fmt.Printf("failed to delete L4 LB %s, error: %s", newL4VirtualServerIds, err.Error())
			return err
		}
		l.currentL4VirtualServerIds = newL4VirtualServerIds
		l.currentL4VirtualServers = newL4VirtualServers
	}

	return nil
}

func (l *LBService) CreateInstances(virtualServers []*common.VirtualServer) ([]string, error) {
	l4VirtualServers, _, l4VirtualServerIds, l7VirtualServerIds := classifyVirtualServer(virtualServers)

	if len(l4VirtualServerIds) != 0 {
		var newL4VirtualServers []*common.VirtualServer
		for _, vs := range l4VirtualServers {
			newL4VirtualServers = append(l.currentL4VirtualServers, vs)
		}
		newL4VirtualServerIds := append(l.currentL4VirtualServerIds, l4VirtualServerIds...)
		_, err := l.keepalived.Update(newL4VirtualServers)
		if err != nil {
			fmt.Printf("failed to create L4 LB %s, error: %s", newL4VirtualServerIds, err.Error())
			return nil, err
		}
		l.currentL4VirtualServerIds = newL4VirtualServerIds
		l.currentL4VirtualServers = newL4VirtualServers
	}
	newCreatedVSIds := append(l4VirtualServerIds, l7VirtualServerIds...)

	return newCreatedVSIds, nil
}

func (l *LBService) UpdateInstances(virtualServers []*common.VirtualServer) error {
	l4VirtualServers, _, l4VirtualServerIds, _ := classifyVirtualServer(virtualServers)
	if len(l4VirtualServers) != 0 {
		var newL4VirtualServers []*common.VirtualServer
		//delete first
		for _, curVs := range l.currentL4VirtualServers {
			for _, vs := range l4VirtualServers {
				if vs.Index != curVs.Index {
					newL4VirtualServers = append(newL4VirtualServers, curVs)
					break
				}
			}
		}
		// add new vs
		for _, vs := range l4VirtualServers {
			newL4VirtualServers = append(newL4VirtualServers, vs)
		}
		_, err := l.keepalived.Update(newL4VirtualServers)
		if err != nil {
			fmt.Printf("failed to update L4 LB %s, error: %s", l4VirtualServerIds, err.Error())
			return err
		}
		l.currentL4VirtualServers = newL4VirtualServers
	}
	return nil
}

func (l *LBService) CleanInstances() error {

	if len(l.currentL4VirtualServerIds) != 0 {
		_, err := l.keepalived.Remove()
		if err != nil {
			fmt.Printf("failed to clean L4 LB, error: %s", err.Error())
			return err
		}
		l.currentL4VirtualServerIds = nil
		l.currentL4VirtualServers = nil
	}
	return nil
}
