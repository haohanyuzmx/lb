package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirtualServerSpec struct {
	Protocol    string `json:"protocol"`
	VIPVirNet   string `json:"vip_vir_net"`
	Port        int    `json:"port"`
	ServerPools string `json:"server_pools"`
}

type VirtualServerStatus struct {
	VIP string `json:"vip"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VirtualServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServerSpec   `json:"spec,omitempty"`
	Status VirtualServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type VirtualServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualServer{}, &VirtualServerList{})
}
