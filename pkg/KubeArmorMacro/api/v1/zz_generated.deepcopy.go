//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2021 Authors of KubeArmor
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorMacro) DeepCopyInto(out *KubeArmorMacro) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorMacro.
func (in *KubeArmorMacro) DeepCopy() *KubeArmorMacro {
	if in == nil {
		return nil
	}
	out := new(KubeArmorMacro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorMacro) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorMacroList) DeepCopyInto(out *KubeArmorMacroList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeArmorMacro, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorMacroList.
func (in *KubeArmorMacroList) DeepCopy() *KubeArmorMacroList {
	if in == nil {
		return nil
	}
	out := new(KubeArmorMacroList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorMacroList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorMacroSpec) DeepCopyInto(out *KubeArmorMacroSpec) {
	*out = *in
	if in.Macros != nil {
		in, out := &in.Macros, &out.Macros
		*out = make([]MacrosType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorMacroSpec.
func (in *KubeArmorMacroSpec) DeepCopy() *KubeArmorMacroSpec {
	if in == nil {
		return nil
	}
	out := new(KubeArmorMacroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorMacroStatus) DeepCopyInto(out *KubeArmorMacroStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorMacroStatus.
func (in *KubeArmorMacroStatus) DeepCopy() *KubeArmorMacroStatus {
	if in == nil {
		return nil
	}
	out := new(KubeArmorMacroStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacrosType) DeepCopyInto(out *MacrosType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacrosType.
func (in *MacrosType) DeepCopy() *MacrosType {
	if in == nil {
		return nil
	}
	out := new(MacrosType)
	in.DeepCopyInto(out)
	return out
}
