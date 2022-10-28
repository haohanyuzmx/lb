package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type ServerPoolSpec struct {
	Name      string               `json:"name"`
	Algorithm string               `json:"algorithm"`
	Members   []PoolMember         `json:"members"`
	Monitor   types.NamespacedName `json:"members"`
}

type PoolMember struct {
	ServerAddr  string `json:"server_address"`
	ServerPort  int    `json:"server_port"`
	Weight      int    `json:"weight"`
	MonitorPort int    `json:"monitor_port"`
}

type ServerPoolStatus struct {
	Members []MemberStatus
}

type MemberStatus struct {
	Conn int `json:"conn"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ServerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerPoolSpec   `json:"spec,omitempty"`
	Status ServerPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ServerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServerPool{}, &ServerPoolList{})
}
