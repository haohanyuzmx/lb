package common

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/types"
	"testing"
)

func TestNamespacedName(t *testing.T) {
	{
		n := NamespacedName{}
		tn := types.NamespacedName{
			Namespace: "default",
			Name:      "vm1",
		}
		n.FromTypes(tn)
		assert.Equal(t, n.String(), tn.String())
	}

	{
		n := NamespacedName{
			Namespace: "default",
			Name:      "vm1",
		}
		tn := types.NamespacedName{
			Namespace: "default",
			Name:      "vm1",
		}
		assert.Equal(t, n.Into(), tn)
	}
	{
		tn := types.NamespacedName{
			Namespace: "default",
			Name:      "vm1",
		}
		n := NamespacedName{}
		n.FromString(tn.String())
		assert.Equal(t, tn.String(), n.String())
	}
}
