package layer4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"my.domain/lb/agent/common"
)

var (
	nicConfig = common.NIC{
		Index:  "nic1",
		Name:   "nic",
		Number: 1,
		LIP:    []string{"1.1.1.1"},
		VIP:    []string{"2.2.2.2"},
	}

	basicL4FullConfig = &common.VirtualServer{
		Index:       "1",
		Enabled:     true,
		Name:        "L4_test",
		Protocol:    "TCP",
		Port:        80,
		VirtualAddr: "12.12.12.12",
		Role:        "master",
		Nic:         nicConfig,
		AppProfile: common.ApplicationProfile{
			Index:              "app1",
			SessionPersistence: true,
			AccessControl:      "\"black\": \"2.1.2.1\"",
			TrafficControl:     []int{10, 30},
		},
		ServerPools: []common.ServerPool{
			{
				Index:     "pool1",
				Name:      "test_pool",
				Algorithm: "rr",
				Monitor: common.Monitor{
					Index: "monitor1",
					Name:  "monitor",
					Type:  "TCP",
				},
				Members: []common.PoolMember{

					{
						ServerAddr: "3.3.3.1",
						ServerPort: 80,
						Weight:     100,
					},
					{
						ServerAddr: "3.3.3.2",
						ServerPort: 80,
						Weight:     100,
					},
				},
			},
		},
	}
)

func TestL4ConfGenerate(t *testing.T) {
	var virtualServers []*common.VirtualServer

	k := NewKeepalived()

	virtualServer1 := basicL4FullConfig
	virtualServers = append(virtualServers, virtualServer1)

	_, err := k.GenerateL4Config(virtualServers)
	assert.Equal(t, err, nil)
}
