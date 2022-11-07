// package constants ...
package constants

const (
	DefaultMaxConcurrentReconciles = 4
)

const (
	LbL4TemplatePath    = "/root/lb/agent/template/keepalived.tmpl"
	LbL4ConfigFilePath  = "/root/lb/agent/configs/L4/keepalived.conf"
	KeepalivedStartCmd  = "systemctl start keepalived"
	KeepalivedStopCmd   = "systemctl stop keepalived"
	KeepalivedReloadCmd = "systemctl reload keepalived"
)

const (
	LbTypeL4 = "L4"
	LbTypeL7 = "L7"
)

type RoleType string

const (
	Master RoleType = "MASTER"
	Backup RoleType = "BACKUP"
)

func (role RoleType) String() string {
	switch role {
	case Master:
		return "MASTER"
	case Backup:
		return "BACKUP"
	}
	return "(unknown)"
}

type PoolAlgorithm string

const (
	RoundRobin       PoolAlgorithm = "rr"
	WeightRoundRobin PoolAlgorithm = "wrr"
)

func (alg PoolAlgorithm) String() string {
	switch alg {
	case RoundRobin:
		return "rr"
	case WeightRoundRobin:
		return "wrr"
	}
	return "(unknown)"
}
