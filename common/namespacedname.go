package common

import (
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

type NamespacedName struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (n NamespacedName) String() string {
	return n.Namespace + `/` + n.Name
}

func (n NamespacedName) Into() types.NamespacedName {
	return types.NamespacedName{
		Namespace: n.Namespace,
		Name:      n.Name,
	}
}

func (n *NamespacedName) FromTypes(tn types.NamespacedName) {
	n.Namespace = tn.Namespace
	n.Name = tn.Name
}

func (n *NamespacedName) FromString(s string) {
	split := strings.Split(s, "/")
	if len(split) != 2 {
		n.Namespace = ""
		n.Name = s
	} else {
		n.Namespace = split[0]
		n.Name = split[1]
	}
}
