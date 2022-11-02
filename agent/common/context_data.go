package common

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
	Port       int
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
