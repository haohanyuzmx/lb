package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"my.domain/lb/pkg/agent/common"
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

func TestCreateLbInstance(t *testing.T) {

	var virtualServers []*common.VirtualServer
	virtualServer1 := basicL4FullConfig
	virtualServers = append(virtualServers, virtualServer1)

	service := NewLBService()
	service.currentL4VirtualServerIds = append(service.currentL4VirtualServerIds, "2")
	currentL4Service := basicL4FullConfig
	currentL4Service.Index = "2"
	service.currentL4VirtualServers = append(service.currentL4VirtualServers, currentL4Service)
	_, err := service.CreateInstances(virtualServers)

	assert.Equal(t, err, nil)
}

func TestUpdateLbInstance(t *testing.T) {

}
