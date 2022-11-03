package layer4

import (
	"fmt"
	"os"

	"my.domain/lb/agent/common"
	"my.domain/lb/agent/template"
	"my.domain/lb/util"

	"my.domain/lb/agent/constants"
)

type Keepalived struct {
	config *common.L4LbConfig
}

var (
	hostname, _ = os.Hostname()
)

var (
	keepalvedTmpl = template.NewL4Template(constants.LbL4TemplatePath)
)

func NewKeepalived() *Keepalived {
	k := new(Keepalived)
	k.config = new(common.L4LbConfig)
	return k
}

func generateVrrpConfig(nic common.NIC, virtualServer *common.VirtualServer) (*common.Vrrp, error) {
	var vrrpConfig common.Vrrp

	vrrpConfig = common.Vrrp{
		Interface: nic.Name,
		RouteID:   nic.Number,
		Priority:  100,
		Advert:    3,
		VIPs:      []string{virtualServer.VirtualAddr},
	}
	return &vrrpConfig, nil
}

func generateVSConfig(virtualServer *common.VirtualServer) (*common.L4VirtualServer, error) {

	serverPools := virtualServer.ServerPools

	vsConfig := common.L4VirtualServer{
		Addr:               virtualServer.VirtualAddr,
		LBAlgo:             serverPools[0].Algorithm,
		PersistenceTimeout: 0,
		Protocol:           virtualServer.Protocol,
	}

	var realServers []common.RealServer
	for _, server := range serverPools[0].Members {
		temp := server
		realServer := common.RealServer{
			Addr:   temp.ServerAddr,
			Weight: temp.Weight,
		}
		realServers = append(realServers, realServer)
	}
	vsConfig.RealServers = realServers

	return &vsConfig, nil
}

func (k *Keepalived) Create(virtualServers []*common.VirtualServer) (bool, error) {
	_, err := k.GenerateL4Config(virtualServers)
	if err != nil {
		return false, fmt.Errorf("Failed to generate keepalived config, for vm %s", hostname)
	}

	err = util.ExecuteCommand(constants.KeepalivedStartCmd, "")
	if err != nil {
		return false, fmt.Errorf("Failed to start keepalived, for vm %s", hostname)
	}
	return true, err
}

func (k *Keepalived) Remove() (bool, error) {

	err := util.ExecuteCommand(constants.KeepalivedStopCmd, "")
	if err != nil {
		return false, fmt.Errorf("Failed to stop keepalived, for vm %s", hostname)
	}
	return true, err
}

func (k *Keepalived) Update(virtualServers []*common.VirtualServer) (bool, error) {
	_, err := k.GenerateL4Config(virtualServers)
	if err != nil {
		return false, fmt.Errorf("Failed to generate keepalived config, for vm %s", hostname)
	}
	err = util.ExecuteCommand(constants.KeepalivedReloadCmd, "")
	if err != nil {
		return false, fmt.Errorf("Failed to reload keepalived, for vm %s", hostname)
	}
	return true, nil
}

func (k *Keepalived) GenerateL4Config(virtualServers []*common.VirtualServer) (string, error) {

	var vrrpsMap map[string]*common.Vrrp
	vrrpsMap = make(map[string]*common.Vrrp)

	var l4VirtualServers []common.L4VirtualServer

	for _, virtualServer := range virtualServers {
		nic := virtualServer.Nic
		exsited := false
		if len(vrrpsMap) != 0 {
			_, ok := vrrpsMap[nic.Index]
			if ok {
				exsited = true
			}
		}
		if exsited == false {
			vrrp, _ := generateVrrpConfig(nic, virtualServer)
			vrrpsMap[nic.Index] = vrrp
		} else {
			vrrp := vrrpsMap[nic.Index]
			vrrp.VIPs = append(vrrp.VIPs, virtualServer.VirtualAddr)
		}

		vs, _ := generateVSConfig(virtualServer)
		l4VirtualServers = append(l4VirtualServers, *vs)
	}
	var vrrps []common.Vrrp
	for _, v := range vrrpsMap {
		vrrps = append(vrrps, *v)
	}
	k.config.RouteID = hostname
	k.config.Vrrps = vrrps
	k.config.VirtualServers = l4VirtualServers

	executor := template.NewL4TemplateExecuter(keepalvedTmpl, k.config)
	result, err := template.ExecuteL4(executor)
	if err != nil {
		return "", fmt.Errorf("Failed to generate keepalived config by template, for vm %s", hostname)
	}

	err = util.WriteToFile(constants.LbL4ConfigFilePath, result)
	if err != nil {
		return "", fmt.Errorf("Failed to write keepalived config, for vm %s", hostname)
	}

	return result, nil
}
