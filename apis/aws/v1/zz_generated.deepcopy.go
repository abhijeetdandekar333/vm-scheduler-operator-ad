//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartad) DeepCopyInto(out *AWSVMSchedulerStartad) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartad.
func (in *AWSVMSchedulerStartad) DeepCopy() *AWSVMSchedulerStartad {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartad)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStartad) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartadList) DeepCopyInto(out *AWSVMSchedulerStartadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSVMSchedulerStartad, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartadList.
func (in *AWSVMSchedulerStartadList) DeepCopy() *AWSVMSchedulerStartadList {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStartadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartadSpec) DeepCopyInto(out *AWSVMSchedulerStartadSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartadSpec.
func (in *AWSVMSchedulerStartadSpec) DeepCopy() *AWSVMSchedulerStartadSpec {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartadStatus) DeepCopyInto(out *AWSVMSchedulerStartadStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartadStatus.
func (in *AWSVMSchedulerStartadStatus) DeepCopy() *AWSVMSchedulerStartadStatus {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStopad) DeepCopyInto(out *AWSVMSchedulerStopad) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStopad.
func (in *AWSVMSchedulerStopad) DeepCopy() *AWSVMSchedulerStopad {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStopad)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStopad) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStopadList) DeepCopyInto(out *AWSVMSchedulerStopadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSVMSchedulerStopad, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStopadList.
func (in *AWSVMSchedulerStopadList) DeepCopy() *AWSVMSchedulerStopadList {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStopadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStopadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStopadSpec) DeepCopyInto(out *AWSVMSchedulerStopadSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStopadSpec.
func (in *AWSVMSchedulerStopadSpec) DeepCopy() *AWSVMSchedulerStopadSpec {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStopadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStopadStatus) DeepCopyInto(out *AWSVMSchedulerStopadStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStopadStatus.
func (in *AWSVMSchedulerStopadStatus) DeepCopy() *AWSVMSchedulerStopadStatus {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStopadStatus)
	in.DeepCopyInto(out)
	return out
}
