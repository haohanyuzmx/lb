package agent

import (
	"fmt"

	"my.domain/lb/agent/common"
	"my.domain/lb/agent/template"

	"my.domain/lb/agent/constants"
)

type KeepalivedConfig struct {
	config *common.L4LbConfig
}

var (
	keepalvedTmpl = template.NewL4Template(constants.LbL4TemplatePath)
)

func NewKeepalivedConfig() *KeepalivedConfig {
	k := new(KeepalivedConfig)
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

func (k *KeepalivedConfig) GenerateL4Config(vmHostname string, virtualServers []*common.VirtualServer) (*common.L4LbConfig, error) {

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
	k.config.RouteID = vmHostname
	k.config.Vrrps = vrrps
	k.config.VirtualServers = l4VirtualServers

	executor := template.NewL4TemplateExecuter(keepalvedTmpl, k.config)
	result, err := template.ExecuteL4(executor)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return k.config, nil
}
