package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"my.domain/lb/pkg/common"
)

type VMSpec struct {
	Memery   string                  `json:"memery"`
	CPU      string                  `json:"cpu"`
	MgtIP    string                  `json:"managment_ip"`
	Hostname string                  `json:"hostname"`
	NICs     []common.NamespacedName `json:"nic_s"`
}

type VMStatus struct {
	HealthCheck int64 `json:"health_check"`
	// +optional
	MasterLBs []common.NamespacedName `json:"master_lb"`
	// +optional
	BackupLBs []common.NamespacedName `json:"backup_lb"`
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
	vm.Status.MasterLBs = append(vm.Status.MasterLBs, n)

}
func (vm *VM) AddBackup(backup types.NamespacedName) {
	n := common.NamespacedName{}
	n.FromTypes(backup)
	vm.Status.BackupLBs = append(vm.Status.BackupLBs, n)
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
