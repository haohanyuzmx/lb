package controller_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "my.domain/lb/api/v1"
	"my.domain/lb/common"
	"time"
)

const (
	timeout  = time.Second * 10
	interval = time.Millisecond * 250
)

var _ = Describe("new virtualServer", func() {
	ctx := context.Background()
	nic1 := newNic("nic1", "nic2")
	nic2 := newNic("nic2", "nic1")
	nic3 := newNic("nic3", "nic4")
	nic4 := newNic("nic4", "nic3")
	vm1 := newVM("vm1", nic1, nic3)
	vm2 := newVM("vm2", nic2, nic4)
	vn1 := newVirNet("vn1", nic1, nic2)
	vn2 := newVirNet("vn2", nic3, nic4)

	BeforeEach(func() {
		k8sClient.Create(ctx, nic1)
		k8sClient.Create(ctx, nic2)
		k8sClient.Create(ctx, nic3)
		k8sClient.Create(ctx, nic4)
		k8sClient.Create(ctx, vm1)
		k8sClient.Create(ctx, vm2)
		k8sClient.Create(ctx, vn1)
		k8sClient.Create(ctx, vn2)
	})
	AfterEach(func() {
		//k8sClient.Delete(ctx, nic1)
		//k8sClient.Delete(ctx, nic2)
		//k8sClient.Delete(ctx, nic3)
		//k8sClient.Delete(ctx, nic4)
		//k8sClient.Delete(ctx, vm1)
		//k8sClient.Delete(ctx, vm2)
		//k8sClient.Delete(ctx, vn1)
		//k8sClient.Delete(ctx, vn2)
	})
	It("add virSer", func() {
		vs := v1.VirtualServer{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "vs1",
			},
			Spec: v1.VirtualServerSpec{
				Enabled:        true,
				VirtualNetwork: getNamespaceName(vn1.ObjectMeta),
				Port:           80,
				DefaultServerPool: common.NamespacedName{
					Namespace: "", Name: "sp1",
				},
				ApplicationProfile: common.NamespacedName{
					Namespace: "", Name: "app1",
				},
			},
		}
		k8sClient.Create(ctx, &vs)

		Eventually(func() bool {
			myvs := v1.VirtualServer{}
			k8sClient.Get(ctx, getNamespaceName(vs.ObjectMeta).Into(), &myvs)
			if myvs.Status.NowVirNet.String() == "/" {
				return false
			}

			master := v1.VM{}
			k8sClient.Get(ctx, myvs.Status.Master.Into(), &master)
			if master.Spec.Master[0] != getNamespaceName(vs.ObjectMeta) {
				fmt.Println("wrong!!!!")
				return false
			}

			backup := v1.VM{}
			k8sClient.Get(ctx, myvs.Status.Backup.Into(), &backup)
			if backup.Spec.Backup[0] != getNamespaceName(vs.ObjectMeta) {
				fmt.Println("wrong!!!!")
				return false
			}

			return true
		}, timeout, interval).Should(BeTrue())
	})
})

func newVM(vmName string, nics ...*v1.NIC) *v1.VM {
	vm := &v1.VM{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      vmName,
		},
		Spec: v1.VMSpec{
			NICs: nil,
		},
	}
	for _, nic := range nics {
		vm.Spec.NICs = append(vm.Spec.NICs, getNamespaceName(nic.ObjectMeta))
		nic.Spec.VM = getNamespaceName(vm.ObjectMeta)
	}
	return vm
}

func newVirNet(netName string, nics ...*v1.NIC) *v1.VirNet {
	vn := &v1.VirNet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      netName,
		},
		Spec: v1.VirNetSpec{
			VIPPool: "0.0.0.0/0",
		},
	}
	for _, nic := range nics {
		vn.Spec.NICs = append(vn.Spec.NICs, getNamespaceName(nic.ObjectMeta))
		nic.Labels = map[string]string{"net": netName}
	}
	return vn
}

func newNic(name, master string) *v1.NIC {
	return &v1.NIC{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      name,
		},
		Spec: v1.NICSpec{
			Master: common.NamespacedName{
				Namespace: "default",
				Name:      master,
			},
		},
	}
}

func getNamespaceName(data metav1.ObjectMeta) common.NamespacedName {
	return common.NamespacedName{
		Namespace: data.Namespace,
		Name:      data.Name,
	}
}
