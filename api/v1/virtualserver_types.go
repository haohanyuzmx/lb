package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"my.domain/lb/common"
)

type VirtualServerSpec struct {
	Enabled bool `json:"enabled"`
	// +optional
	Name               string                `json:"name"`
	Protocol           string                `json:"protocol"`
	VirtualNetwork     common.NamespacedName `json:"virtual_network"`
	Port               int                   `json:"port"`
	DefaultServerPool  common.NamespacedName `json:"default_server_pool"`
	ApplicationProfile common.NamespacedName `json:"application_profile"`
}

type VirtualServerStatus struct {
	VIP       string                `json:"vip"`
	NowVirNet common.NamespacedName `json:"now_vir_net"`
	Master    common.NamespacedName `json:"master"`
	Backup    common.NamespacedName `json:"backup"`
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
