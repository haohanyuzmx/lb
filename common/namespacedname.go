package common

import (
	"strings"

	"k8s.io/apimachinery/pkg/types"
)

type NamespacedName struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (n NamespacedName) ToString() string {
	return n.Namespace + `/` + n.Name
}

func (n NamespacedName) IntoTypes() types.NamespacedName {
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
