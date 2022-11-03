package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"my.domain/lb/util"
)

type VirtualServerSpec struct {
	Enabled bool `json:"enabled"`
	// +optional
	Name               string              `json:"name"`
	Protocol           string              `json:"protocol"`
	VirtualNetwork     util.NamespacedName `json:"virtual_network"`
	Port               int                 `json:"port"`
	DefaultServerPool  util.NamespacedName `json:"default_server_pool"`
	ApplicationProfile util.NamespacedName `json:"application_profile"`
}

type VirtualServerStatus struct {
	VIP    string              `json:"vip"`
	Master util.NamespacedName `json:"master"`
	Backup util.NamespacedName `json:"backup"`
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
