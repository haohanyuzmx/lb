package agent

import "strings"

type keepAlivedConf struct {
	RouteID        string
	Vrrps          []vrrp
	VirtualServers []virtualServer
}
type vrrp struct {
	Name      string
	Interface string
	RouteID   int
	Priority  int
	Advert    int
	VIPs      []string
}
type virtualServer struct {
	Addr               string
	DelayLoop          int
	LBAlgo             string
	LBKind             string
	PersistenceTimeout int
	Protocol           string
	RealServers        []realServers
}
type realServers struct {
	Addr       string
	Weight     int
	TCPTimeout int
}

func (k *keepAlivedConf) ToConf(interface{}) (string, error) {
	k = &keepAlivedConf{
		RouteID: "VM1",
		Vrrps: []vrrp{
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
		VirtualServers: []virtualServer{
			{
				Addr:               "4 6",
				DelayLoop:          2,
				LBAlgo:             "hh",
				LBKind:             "aa",
				PersistenceTimeout: 0,
				Protocol:           "dd",
				RealServers: []realServers{
					{
						Addr:       "6 7",
						Weight:     2,
						TCPTimeout: 4,
					},
				},
			},
		},
	}

	s := &strings.Builder{}
	err := keepalievd.Execute(s, k)
	return s.String(), err
}
