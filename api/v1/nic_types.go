package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NICSpec struct {
	LIP string `json:"lip"`
}

type NICStatus struct {
	LIP string         `json:"lip"`
	VIP map[string]int `json:"vip"` //{vip:load}
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
