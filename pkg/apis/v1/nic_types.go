package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"my.domain/lb/pkg/common"
)

type NICSpec struct {
	// +optional
	Name string `json:"name"`
	// +optional
	LIP    string                `json:"lip"`
	VM     common.NamespacedName `json:"vm"`
	Master common.NamespacedName `json:"master"`
}

type NICStatus struct {
	LIP  []string `json:"lip"`
	VIP  []string `json:"vip"`
	Load int      `json:"load"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type NIC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NICSpec   `json:"spec,omitempty"`
	Status NICStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type NICList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NIC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NIC{}, &NICList{})
}
