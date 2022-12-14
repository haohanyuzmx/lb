//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"my.domain/lb/pkg/common"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationProfile) DeepCopyInto(out *ApplicationProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationProfile.
func (in *ApplicationProfile) DeepCopy() *ApplicationProfile {
	if in == nil {
		return nil
	}
	out := new(ApplicationProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationProfileList) DeepCopyInto(out *ApplicationProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationProfileList.
func (in *ApplicationProfileList) DeepCopy() *ApplicationProfileList {
	if in == nil {
		return nil
	}
	out := new(ApplicationProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationProfileSpec) DeepCopyInto(out *ApplicationProfileSpec) {
	*out = *in
	if in.TrafficControl != nil {
		in, out := &in.TrafficControl, &out.TrafficControl
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationProfileSpec.
func (in *ApplicationProfileSpec) DeepCopy() *ApplicationProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationProfileStatus) DeepCopyInto(out *ApplicationProfileStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationProfileStatus.
func (in *ApplicationProfileStatus) DeepCopy() *ApplicationProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitor) DeepCopyInto(out *Monitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitor.
func (in *Monitor) DeepCopy() *Monitor {
	if in == nil {
		return nil
	}
	out := new(Monitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Monitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorList) DeepCopyInto(out *MonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Monitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorList.
func (in *MonitorList) DeepCopy() *MonitorList {
	if in == nil {
		return nil
	}
	out := new(MonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpec) DeepCopyInto(out *MonitorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpec.
func (in *MonitorSpec) DeepCopy() *MonitorSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorStatus) DeepCopyInto(out *MonitorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorStatus.
func (in *MonitorStatus) DeepCopy() *MonitorStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NIC) DeepCopyInto(out *NIC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NIC.
func (in *NIC) DeepCopy() *NIC {
	if in == nil {
		return nil
	}
	out := new(NIC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NIC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICList) DeepCopyInto(out *NICList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NIC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICList.
func (in *NICList) DeepCopy() *NICList {
	if in == nil {
		return nil
	}
	out := new(NICList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NICList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICSpec) DeepCopyInto(out *NICSpec) {
	*out = *in
	out.VM = in.VM
	out.Master = in.Master
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICSpec.
func (in *NICSpec) DeepCopy() *NICSpec {
	if in == nil {
		return nil
	}
	out := new(NICSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICStatus) DeepCopyInto(out *NICStatus) {
	*out = *in
	if in.LIP != nil {
		in, out := &in.LIP, &out.LIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VIP != nil {
		in, out := &in.VIP, &out.VIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICStatus.
func (in *NICStatus) DeepCopy() *NICStatus {
	if in == nil {
		return nil
	}
	out := new(NICStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolMember) DeepCopyInto(out *PoolMember) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolMember.
func (in *PoolMember) DeepCopy() *PoolMember {
	if in == nil {
		return nil
	}
	out := new(PoolMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPool) DeepCopyInto(out *ServerPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPool.
func (in *ServerPool) DeepCopy() *ServerPool {
	if in == nil {
		return nil
	}
	out := new(ServerPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPoolList) DeepCopyInto(out *ServerPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServerPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPoolList.
func (in *ServerPoolList) DeepCopy() *ServerPoolList {
	if in == nil {
		return nil
	}
	out := new(ServerPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPoolSpec) DeepCopyInto(out *ServerPoolSpec) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]PoolMember, len(*in))
		copy(*out, *in)
	}
	out.Monitor = in.Monitor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPoolSpec.
func (in *ServerPoolSpec) DeepCopy() *ServerPoolSpec {
	if in == nil {
		return nil
	}
	out := new(ServerPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPoolStatus) DeepCopyInto(out *ServerPoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPoolStatus.
func (in *ServerPoolStatus) DeepCopy() *ServerPoolStatus {
	if in == nil {
		return nil
	}
	out := new(ServerPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VM) DeepCopyInto(out *VM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VM.
func (in *VM) DeepCopy() *VM {
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMList) DeepCopyInto(out *VMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMList.
func (in *VMList) DeepCopy() *VMList {
	if in == nil {
		return nil
	}
	out := new(VMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSpec) DeepCopyInto(out *VMSpec) {
	*out = *in
	if in.NICs != nil {
		in, out := &in.NICs, &out.NICs
		*out = make([]common.NamespacedName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSpec.
func (in *VMSpec) DeepCopy() *VMSpec {
	if in == nil {
		return nil
	}
	out := new(VMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStatus) DeepCopyInto(out *VMStatus) {
	*out = *in
	if in.MasterLBs != nil {
		in, out := &in.MasterLBs, &out.MasterLBs
		*out = make([]common.NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.BackupLBs != nil {
		in, out := &in.BackupLBs, &out.BackupLBs
		*out = make([]common.NamespacedName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStatus.
func (in *VMStatus) DeepCopy() *VMStatus {
	if in == nil {
		return nil
	}
	out := new(VMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirNet) DeepCopyInto(out *VirNet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirNet.
func (in *VirNet) DeepCopy() *VirNet {
	if in == nil {
		return nil
	}
	out := new(VirNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirNet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirNetList) DeepCopyInto(out *VirNetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirNet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirNetList.
func (in *VirNetList) DeepCopy() *VirNetList {
	if in == nil {
		return nil
	}
	out := new(VirNetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirNetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirNetSpec) DeepCopyInto(out *VirNetSpec) {
	*out = *in
	if in.NICs != nil {
		in, out := &in.NICs, &out.NICs
		*out = make([]common.NamespacedName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirNetSpec.
func (in *VirNetSpec) DeepCopy() *VirNetSpec {
	if in == nil {
		return nil
	}
	out := new(VirNetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirNetStatus) DeepCopyInto(out *VirNetStatus) {
	*out = *in
	if in.UnuseVIPPool != nil {
		in, out := &in.UnuseVIPPool, &out.UnuseVIPPool
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirNetStatus.
func (in *VirNetStatus) DeepCopy() *VirNetStatus {
	if in == nil {
		return nil
	}
	out := new(VirNetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServer) DeepCopyInto(out *VirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServer.
func (in *VirtualServer) DeepCopy() *VirtualServer {
	if in == nil {
		return nil
	}
	out := new(VirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerList) DeepCopyInto(out *VirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerList.
func (in *VirtualServerList) DeepCopy() *VirtualServerList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerSpec) DeepCopyInto(out *VirtualServerSpec) {
	*out = *in
	out.VirtualNetwork = in.VirtualNetwork
	out.DefaultServerPool = in.DefaultServerPool
	out.ApplicationProfile = in.ApplicationProfile
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerSpec.
func (in *VirtualServerSpec) DeepCopy() *VirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStatus) DeepCopyInto(out *VirtualServerStatus) {
	*out = *in
	out.NowVirNet = in.NowVirNet
	out.Master = in.Master
	out.Backup = in.Backup
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStatus.
func (in *VirtualServerStatus) DeepCopy() *VirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStatus)
	in.DeepCopyInto(out)
	return out
}
