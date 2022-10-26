package agent

import "text/template"

var (
	keepalievd = template.Must(template.ParseFiles("./keepalievd.conf"))
)

type ToConf interface {
	ToConf(interface{}) (string, error)
}

func KeepalivedConf() ToConf {
	return &keepAlivedConf{}
}
