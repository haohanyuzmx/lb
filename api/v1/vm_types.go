/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VMSpec struct {
	VirServers []VSTypes `json:"vir_servers"`
	NICs       []string  `json:"nic_s"`
}
type VSTypes struct {
	Identity string `json:"identity"` //master|backup
	VS       string `json:"vs"`       //vs的索引
}

type VMStatus struct {
	HealthCheck int64 `json:"health_check"` //时间戳
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMSpec   `json:"spec,omitempty"`
	Status VMStatus `json:"status,omitempty"`
}

func (vm *VM) AddMaster(master string) {
	vm.Spec.VirServers = append(vm.Spec.VirServers, VSTypes{
		Identity: "master",
		VS:       master,
	})

}
func (vm *VM) AddBackup(backup string) {
	vm.Spec.VirServers = append(vm.Spec.VirServers, VSTypes{
		Identity: "backup",
		VS:       backup,
	})

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
