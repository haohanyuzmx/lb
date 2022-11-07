package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApplicationProfileSpec struct {
	SessionPersistence bool   `json:"session_persistence"`
	AccessControl      string `json:"access_control"`
	TrafficControl     []int  `json:"traffic_control"`
}

type ApplicationProfileStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ApplicationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationProfileSpec   `json:"spec,omitempty"`
	Status ApplicationProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ApplicationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationProfile{}, &ApplicationProfileList{})
}
