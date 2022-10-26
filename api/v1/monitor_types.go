package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MonitorSpec struct {
	Type     string `json:"type"`
	Interval int    `json:"intercal"`
	Timeout  int    `json:"timeout"`
	Method   string `json:"method"`
	URL      string `json:"url"`
}

type MonitorStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitorSpec   `json:"spec,omitempty"`
	Status MonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}
