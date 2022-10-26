package util

import (
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

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
