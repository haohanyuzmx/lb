//package constants ...
package constants

const (
	LbL4TemplatePath = "/root/lb/agent/template/keepalived.tmpl"
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

type PoolAlgorithm string

const (
	RoundRobin       PoolAlgorithm = "rr"
	WeightRoundRobin PoolAlgorithm = "wrr"
)
