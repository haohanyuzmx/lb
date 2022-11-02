package template

import (
	"fmt"
	"testing"

	"my.domain/lb/agent/common"
	"my.domain/lb/agent/constants"
)

func TestL4Conf(t *testing.T) {
	var k *common.L4LbConfig
	k = &common.L4LbConfig{
		RouteID: "VM1",
		Vrrps: []common.Vrrp{
			{
				Interface: "eth0",
				RouteID:   12,
				Priority:  100,
				Advert:    3,
				VIPs:      []string{"1", "2", "3"},
			},
			{
				Interface: "eth2",
				RouteID:   13,
				Priority:  100,
				Advert:    3,
				VIPs:      []string{"1", "2", "3"},
			},
		},
		VirtualServers: []common.L4VirtualServer{
			{
				Addr:               "4 6",
				DelayLoop:          2,
				LBAlgo:             "hh",
				LBKind:             "aa",
				PersistenceTimeout: 0,
				Protocol:           "dd",
				RealServers: []common.RealServer{
					{
						Addr:       "6 7",
						Weight:     2,
						TCPTimeout: 4,
					},
				},
			},
		},
	}
	var executor *L4TemplateExecuter
	var tmpl *Template
	var err error
	tmpl, err = NewL4Template(constants.LbL4TemplatePath)
	if err != nil {
		println(err)
	}
	executor = NewL4TemplateExecuter(tmpl, k)
	result, err := ExecuteL4(executor)
	fmt.Printf(result)
}
