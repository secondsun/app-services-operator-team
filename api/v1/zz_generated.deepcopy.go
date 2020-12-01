// +build !ignore_autogenerated


// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapServerSpec) DeepCopyInto(out *BootstrapServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapServerSpec.
func (in *BootstrapServerSpec) DeepCopy() *BootstrapServerSpec {
	if in == nil {
		return nil
	}
	out := new(BootstrapServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsSpec) DeepCopyInto(out *CredentialsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsSpec.
func (in *CredentialsSpec) DeepCopy() *CredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConnection) DeepCopyInto(out *ManagedKafkaConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConnection.
func (in *ManagedKafkaConnection) DeepCopy() *ManagedKafkaConnection {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConnectionList) DeepCopyInto(out *ManagedKafkaConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedKafkaConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConnectionList.
func (in *ManagedKafkaConnectionList) DeepCopy() *ManagedKafkaConnectionList {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedKafkaConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConnectionSpec) DeepCopyInto(out *ManagedKafkaConnectionSpec) {
	*out = *in
	out.BootstrapServer = in.BootstrapServer
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConnectionSpec.
func (in *ManagedKafkaConnectionSpec) DeepCopy() *ManagedKafkaConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKafkaConnectionStatus) DeepCopyInto(out *ManagedKafkaConnectionStatus) {
	*out = *in
	out.Secret = in.Secret
	out.Deployment = in.Deployment
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKafkaConnectionStatus.
func (in *ManagedKafkaConnectionStatus) DeepCopy() *ManagedKafkaConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedKafkaConnectionStatus)
	in.DeepCopyInto(out)
	return out
}
