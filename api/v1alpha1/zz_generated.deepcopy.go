//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)


// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressController.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}


// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	out.ApplicationName = in.ApplicationName
	out.Namespace = in.Namespace

	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(Repository)
		(*in).DeepCopyInto(*out)
	}

	if in.DockerConfig != nil {
		in, out := &in.DockerConfig, &out.DockerConfig
		*out = new(DockerConfig)
		(*in).DeepCopyInto(*out)
	}

	if in.Acr != nil {
		in, out := &in.Acr, &out.Acr
		*out = new(Acr)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSSLCertificate.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.Owner = in.Owner
	out.Name = in.Name
	out.BranchName = in.BranchName
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSSLCertificate.
func (in *DockerConfig) DeepCopy() *DockerConfig {
	if in == nil {
		return nil
	}
	out := new(DockerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfig) DeepCopyInto(out *DockerConfig) {
	*out = *in
	out.Dockerfile = in.Dockerfile
	out.BuildContext = in.BuildContext
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSSLCertificate.
func (in *Acr) DeepCopy() *Acr {
	if in == nil {
		return nil
	}
	out := new(Acr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Acr) DeepCopyInto(out *Acr) {
	*out = *in
	out.Id = in.Id
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}