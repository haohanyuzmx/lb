package common

type L7LbConfig struct {
}

// for L4 template config convertion
type L4LbConfig struct {
	RouteID        string
	Vrrps          []Vrrp
	VirtualServers []L4VirtualServer
}

type Vrrp struct {
	Name      string
	Interface string
	RouteID   int
	Priority  int
	Advert    int
	VIPs      []string
}

type L4VirtualServer struct {
	Addr               string
	DelayLoop          int
	LBAlgo             string
	LBKind             string
	PersistenceTimeout int
	Protocol           string
	RealServers        []RealServer
}

type RealServer struct {
	Addr       string
	Weight     int
	TCPTimeout int
}
