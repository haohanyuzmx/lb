package util

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

func NamespacedName2types(n NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Namespace: n.Namespace,
		Name:      n.Name,
	}
}

func Types2NamespacedName(n types.NamespacedName) NamespacedName {
	return NamespacedName{
		Namespace: n.Namespace,
		Name:      n.Name,
	}
}

func String2NamespacedName(s string) types.NamespacedName {

	split := strings.Split(s, "/")
	if len(split) != 2 {
		return types.NamespacedName{
			Namespace: "",
			Name:      s,
		}
	}
	return types.NamespacedName{
		Namespace: split[0],
		Name:      split[1],
	}
}
