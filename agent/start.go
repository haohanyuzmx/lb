package agent

import (
	"k8s.io/apimachinery/pkg/types"
	"my.domain/lb/controllers"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type updatefun func(applyfunc func(any) bool, index types.NamespacedName, some client.Object)

func apply(virSers []controllers.VirtualServer, update updatefun) {

}
