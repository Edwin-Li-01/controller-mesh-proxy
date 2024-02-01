// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package proto

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using FaultInjectConfigResp within kubernetes types, where deepcopy-gen is used.
func (in *FaultInjectConfigResp) DeepCopyInto(out *FaultInjectConfigResp) {
	p := proto.Clone(in).(*FaultInjectConfigResp)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FaultInjectConfigResp. Required by controller-gen.
func (in *FaultInjectConfigResp) DeepCopy() *FaultInjectConfigResp {
	if in == nil {
		return nil
	}
	out := new(FaultInjectConfigResp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FaultInjectConfigResp. Required by controller-gen.
func (in *FaultInjectConfigResp) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FaultInjection within kubernetes types, where deepcopy-gen is used.
func (in *FaultInjection) DeepCopyInto(out *FaultInjection) {
	p := proto.Clone(in).(*FaultInjection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FaultInjection. Required by controller-gen.
func (in *FaultInjection) DeepCopy() *FaultInjection {
	if in == nil {
		return nil
	}
	out := new(FaultInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FaultInjection. Required by controller-gen.
func (in *FaultInjection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection) DeepCopyInto(out *HTTPFaultInjection) {
	p := proto.Clone(in).(*HTTPFaultInjection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopy() *HTTPFaultInjection {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Delay within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Delay) DeepCopyInto(out *HTTPFaultInjection_Delay) {
	p := proto.Clone(in).(*HTTPFaultInjection_Delay)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopy() *HTTPFaultInjection_Delay {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Delay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Abort within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Abort) DeepCopyInto(out *HTTPFaultInjection_Abort) {
	p := proto.Clone(in).(*HTTPFaultInjection_Abort)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopy() *HTTPFaultInjection_Abort {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Abort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EffectiveTimeRange within kubernetes types, where deepcopy-gen is used.
func (in *EffectiveTimeRange) DeepCopyInto(out *EffectiveTimeRange) {
	p := proto.Clone(in).(*EffectiveTimeRange)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveTimeRange. Required by controller-gen.
func (in *EffectiveTimeRange) DeepCopy() *EffectiveTimeRange {
	if in == nil {
		return nil
	}
	out := new(EffectiveTimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveTimeRange. Required by controller-gen.
func (in *EffectiveTimeRange) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Match within kubernetes types, where deepcopy-gen is used.
func (in *Match) DeepCopyInto(out *Match) {
	p := proto.Clone(in).(*Match)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match. Required by controller-gen.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Match. Required by controller-gen.
func (in *Match) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HttpMatch within kubernetes types, where deepcopy-gen is used.
func (in *HttpMatch) DeepCopyInto(out *HttpMatch) {
	p := proto.Clone(in).(*HttpMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpMatch. Required by controller-gen.
func (in *HttpMatch) DeepCopy() *HttpMatch {
	if in == nil {
		return nil
	}
	out := new(HttpMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HttpMatch. Required by controller-gen.
func (in *HttpMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MatchContent within kubernetes types, where deepcopy-gen is used.
func (in *MatchContent) DeepCopyInto(out *MatchContent) {
	p := proto.Clone(in).(*MatchContent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchContent. Required by controller-gen.
func (in *MatchContent) DeepCopy() *MatchContent {
	if in == nil {
		return nil
	}
	out := new(MatchContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MatchContent. Required by controller-gen.
func (in *MatchContent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HttpHeader within kubernetes types, where deepcopy-gen is used.
func (in *HttpHeader) DeepCopyInto(out *HttpHeader) {
	p := proto.Clone(in).(*HttpHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpHeader. Required by controller-gen.
func (in *HttpHeader) DeepCopy() *HttpHeader {
	if in == nil {
		return nil
	}
	out := new(HttpHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HttpHeader. Required by controller-gen.
func (in *HttpHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ResourceMatch within kubernetes types, where deepcopy-gen is used.
func (in *ResourceMatch) DeepCopyInto(out *ResourceMatch) {
	p := proto.Clone(in).(*ResourceMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMatch. Required by controller-gen.
func (in *ResourceMatch) DeepCopy() *ResourceMatch {
	if in == nil {
		return nil
	}
	out := new(ResourceMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMatch. Required by controller-gen.
func (in *ResourceMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}