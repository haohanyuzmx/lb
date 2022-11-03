package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"my.domain/lb/util"
)

type VirNetSpec struct {
	NICs    []util.NamespacedName `json:"nic_s"` //nic的索引
	VIPPool string                `json:"vip_pool"`
}

type VirNetStatus struct {
	UnuseVIPPool []string `json:"unuse_vip_pool"`
	IsAlloc      bool     `json:"is_alloc"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VirNet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirNetSpec   `json:"spec,omitempty"`
	Status VirNetStatus `json:"status,omitempty"`
}

func (vn *VirNet) Alloc() string {
	if !vn.Status.IsAlloc {
		vn.init()
	}
	if len(vn.Status.UnuseVIPPool) == 0 {
		return ""
	}
	vip := vn.Status.UnuseVIPPool[0]
	vn.Status.UnuseVIPPool = vn.Status.UnuseVIPPool[1:]
	return vip
}

func (vn *VirNet) Free(vip string) {
	vn.Status.UnuseVIPPool = append(vn.Status.UnuseVIPPool, vip)
}

func (vn *VirNet) init() {
	if vn.Status.IsAlloc {
		return
	}
	vn.Status.IsAlloc = true
	vn.Status.UnuseVIPPool = []string{"1.1.1.1", "2.2.2.2", "3.3.3.3", "4.4.4.4"}
}

// +kubebuilder:object:root=true
type VirNetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirNet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirNet{}, &VirNetList{})
}
