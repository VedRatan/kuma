//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	"github.com/kumahq/kuma/api/common/v1alpha1/tls"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conf) DeepCopyInto(out *Conf) {
	*out = *in
	if in.TlsVersion != nil {
		in, out := &in.TlsVersion, &out.TlsVersion
		*out = new(tls.Version)
		(*in).DeepCopyInto(*out)
	}
	if in.TlsCiphers != nil {
		in, out := &in.TlsCiphers, &out.TlsCiphers
		*out = new([]tls.TlsCipher)
		if **in != nil {
			in, out := *in, *out
			*out = make([]tls.TlsCipher, len(*in))
			copy(*out, *in)
		}
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(Mode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conf.
func (in *Conf) DeepCopy() *Conf {
	if in == nil {
		return nil
	}
	out := new(Conf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *From) DeepCopyInto(out *From) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	in.Default.DeepCopyInto(&out.Default)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new From.
func (in *From) DeepCopy() *From {
	if in == nil {
		return nil
	}
	out := new(From)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshTLS) DeepCopyInto(out *MeshTLS) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.TargetRef)
		(*in).DeepCopyInto(*out)
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new([]From)
		if **in != nil {
			in, out := *in, *out
			*out = make([]From, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new([]Rule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Rule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshTLS.
func (in *MeshTLS) DeepCopy() *MeshTLS {
	if in == nil {
		return nil
	}
	out := new(MeshTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}
