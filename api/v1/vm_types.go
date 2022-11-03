package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"my.domain/lb/common"
)

type VMSpec struct {
	// +optional
	Master []common.NamespacedName `json:"master"`
	// +optional
	Backup []common.NamespacedName `json:"backup"`
	NICs   []common.NamespacedName `json:"nic_s"`
}

type VMStatus struct {
	HealthCheck int64  `json:"health_check"`
	Hostname    string `json:"hostname"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMSpec   `json:"spec,omitempty"`
	Status VMStatus `json:"status,omitempty"`
}

func (vm *VM) AddMaster(master types.NamespacedName) {
	n := common.NamespacedName{}
	n.FromTypes(master)
	vm.Spec.Master = append(vm.Spec.Master, n)

}
func (vm *VM) AddBackup(backup types.NamespacedName) {
	n := common.NamespacedName{}
	n.FromTypes(backup)
	vm.Spec.Backup = append(vm.Spec.Backup, n)
}

// +kubebuilder:object:root=true
type VMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VM `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VM{}, &VMList{})
}
