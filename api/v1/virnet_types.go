package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirNetSpec struct {
	NICs    []string `json:"nic_s"` //nic的索引
	VIPPool string
	LIPPool string
}

type VirNetStatus struct {
	UnuseVIPPool string
	UnuseLIPPool string
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VirNet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirNetSpec   `json:"spec,omitempty"`
	Status VirNetStatus `json:"status,omitempty"`
}

func (vn *VirNet) GetVIP() string {
	return ""
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
