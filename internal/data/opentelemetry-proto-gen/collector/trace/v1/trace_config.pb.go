// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/trace/v1/trace_config.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfTimedEvents int64 `protobuf:"varint,5,opt,name=max_number_of_timed_events,json=maxNumberOfTimedEvents,proto3" json:"max_number_of_timed_events,omitempty"`
	// The global default max number of attributes per timed event.
	MaxNumberOfAttributesPerTimedEvent int64 `protobuf:"varint,6,opt,name=max_number_of_attributes_per_timed_event,json=maxNumberOfAttributesPerTimedEvent,proto3" json:"max_number_of_attributes_per_timed_event,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributesPerLink int64 `protobuf:"varint,8,opt,name=max_number_of_attributes_per_link,json=maxNumberOfAttributesPerLink,proto3" json:"max_number_of_attributes_per_link,omitempty"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,1,opt,name=constant_sampler,json=constantSampler,proto3,oneof" json:"constant_sampler,omitempty"`
}
type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,2,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof" json:"probability_sampler,omitempty"`
}
type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof" json:"rate_limiting_sampler,omitempty"`
}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler()     {}
func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler()  {}
func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfTimedEvents() int64 {
	if m != nil {
		return m.MaxNumberOfTimedEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerTimedEvent() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerTimedEvent
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerLink() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerLink
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return m.Size()
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability float64 `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{2}
}
func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(m, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return m.Size()
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps int64 `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opentelemetry.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ConstantSampler)(nil), "opentelemetry.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*ProbabilitySampler)(nil), "opentelemetry.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opentelemetry.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/trace/v1/trace_config.proto", fileDescriptor_5936aa8fa6443e6f)
}

var fileDescriptor_5936aa8fa6443e6f = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0xc6, 0x71, 0x59, 0xff, 0xbd, 0x55, 0xdb, 0xcc, 0xa8, 0x53, 0x34, 0x55, 0x59, 0x97, 0xcb,
	0xb8, 0x40, 0x4a, 0x77, 0x98, 0xb4, 0xc3, 0x24, 0x68, 0xcb, 0x76, 0x40, 0x14, 0xa5, 0x68, 0xd3,
	0xb8, 0x44, 0x26, 0x98, 0xc8, 0x5a, 0x62, 0x67, 0x8e, 0x8b, 0xda, 0x0f, 0xb0, 0xfb, 0xbe, 0xc4,
	0x3e, 0xc1, 0xbe, 0xc4, 0x8e, 0x3d, 0xee, 0x38, 0xc1, 0x17, 0x99, 0x62, 0x28, 0x84, 0x3f, 0x45,
	0xea, 0xcd, 0xef, 0xfb, 0xf8, 0xf9, 0x3d, 0x36, 0xbc, 0x31, 0x38, 0x22, 0xa6, 0x5c, 0xd1, 0x90,
	0x46, 0x54, 0xc9, 0x3b, 0x27, 0x96, 0x42, 0x09, 0x47, 0x49, 0xe2, 0x53, 0x67, 0x50, 0x19, 0x2f,
	0x3c, 0x5f, 0xf0, 0x3e, 0x0b, 0xca, 0x5a, 0xc3, 0xc7, 0x73, 0x86, 0x71, 0xb3, 0xac, 0xf7, 0x95,
	0x07, 0x15, 0xfb, 0xc7, 0x26, 0xec, 0xb5, 0xd3, 0xe2, 0x5c, 0x7b, 0x70, 0x07, 0x0c, 0x5f, 0xf0,
	0x44, 0x11, 0xae, 0xbc, 0x84, 0x44, 0x71, 0x48, 0xa5, 0x89, 0x4e, 0x50, 0x71, 0xef, 0xac, 0x54,
	0x5e, 0x07, 0x2a, 0x9f, 0x4f, 0x5c, 0xd7, 0x63, 0xd3, 0xa7, 0x9c, 0x7b, 0xe8, 0xcf, 0xb7, 0xb0,
	0x0f, 0x85, 0x58, 0x8a, 0x2e, 0xe9, 0xb2, 0x90, 0xa9, 0xbb, 0x29, 0x7e, 0x43, 0xe3, 0x4f, 0xd7,
	0xe3, 0x5b, 0x33, 0xe3, 0x2c, 0x01, 0xc7, 0x4b, 0x5d, 0x1c, 0xc0, 0x91, 0x24, 0x8a, 0x7a, 0x21,
	0x8b, 0x98, 0x62, 0x3c, 0x98, 0xc6, 0xe4, 0x75, 0x4c, 0x65, 0x7d, 0x8c, 0x4b, 0x14, 0x6d, 0x4c,
	0x9c, 0xb3, 0x9c, 0x82, 0x5c, 0x6e, 0xe3, 0x77, 0x60, 0x46, 0xe4, 0xd6, 0xe3, 0x37, 0x51, 0x97,
	0x4a, 0x4f, 0xf4, 0x3d, 0xa2, 0x94, 0x64, 0xdd, 0x1b, 0x45, 0x13, 0xf3, 0xd9, 0x09, 0x2a, 0xe6,
	0xdd, 0xa3, 0x88, 0xdc, 0x36, 0xb5, 0x7c, 0xd5, 0xaf, 0x4e, 0x45, 0xfc, 0x1e, 0x5e, 0xce, 0x1b,
	0x15, 0x8b, 0x68, 0xcf, 0xa3, 0x03, 0xca, 0x55, 0x62, 0x6e, 0x6a, 0xeb, 0x8b, 0x8c, 0xb5, 0x9d,
	0xca, 0x97, 0x5a, 0xc5, 0x6d, 0x28, 0x3e, 0x16, 0xea, 0xc5, 0x54, 0x66, 0x51, 0xe6, 0x96, 0x26,
	0xd9, 0x2b, 0x0f, 0xd1, 0xa2, 0x72, 0x86, 0xc5, 0x25, 0x28, 0xcc, 0x53, 0x43, 0xc6, 0xbf, 0x25,
	0xe6, 0xb6, 0x06, 0x18, 0x19, 0x40, 0x23, 0xed, 0xe3, 0x8f, 0xf0, 0x7a, 0xed, 0x21, 0x52, 0xb7,
	0xb9, 0xa3, 0xcd, 0xc7, 0x8f, 0xa5, 0xa7, 0xa4, 0xda, 0x2e, 0x6c, 0x4f, 0xfe, 0x1d, 0xfb, 0x37,
	0x82, 0xc3, 0x85, 0x11, 0xc2, 0x1d, 0xd8, 0xe9, 0x51, 0x9f, 0x25, 0x4c, 0x70, 0x3d, 0x83, 0x07,
	0x67, 0x1f, 0x9e, 0x34, 0x83, 0xd3, 0xfa, 0x62, 0x42, 0x71, 0xa7, 0x3c, 0xfb, 0x02, 0x8c, 0x45,
	0x15, 0x1f, 0x00, 0x54, 0x1b, 0x5f, 0xaa, 0x5f, 0xaf, 0xbd, 0xab, 0x7a, 0xdd, 0xc8, 0xe1, 0x7d,
	0xd8, 0x7d, 0xa8, 0x9b, 0x06, 0xc2, 0xcf, 0x61, 0x7f, 0x52, 0xb6, 0xaa, 0xee, 0x65, 0xb3, 0x6d,
	0x6c, 0xd8, 0x75, 0xc0, 0xcb, 0x83, 0x89, 0x4f, 0xa1, 0xa0, 0xaf, 0xc5, 0x78, 0x90, 0x51, 0xf5,
	0x15, 0x90, 0xbb, 0x4a, 0xb2, 0xdf, 0x40, 0x61, 0xc5, 0xe4, 0x61, 0x03, 0xf2, 0xdf, 0xe3, 0x44,
	0x1b, 0xf3, 0x6e, 0xba, 0xac, 0xfd, 0x42, 0x7f, 0x86, 0x16, 0xba, 0x1f, 0x5a, 0xe8, 0xdf, 0xd0,
	0x42, 0x3f, 0x47, 0x56, 0xee, 0x7e, 0x64, 0xe5, 0xfe, 0x8e, 0xac, 0x1c, 0xbc, 0x62, 0x62, 0xed,
	0xaf, 0x53, 0x33, 0x32, 0xdf, 0x79, 0x2b, 0x95, 0x5a, 0xa8, 0xf3, 0x39, 0x58, 0x34, 0x31, 0xe1,
	0xf8, 0x22, 0x0c, 0xa9, 0xaf, 0x84, 0x74, 0x18, 0x57, 0x54, 0x72, 0x12, 0x3a, 0x3d, 0xa2, 0xc8,
	0xfc, 0xcb, 0x53, 0xd2, 0xf4, 0x52, 0x40, 0x79, 0x66, 0xff, 0xc3, 0x3b, 0xd4, 0xdd, 0xd2, 0xea,
	0xdb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0xfd, 0x53, 0x41, 0xae, 0x04, 0x00, 0x00,
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxNumberOfAttributesPerLink != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributesPerLink))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxNumberOfLinks != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfLinks))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxNumberOfAttributesPerTimedEvent != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributesPerTimedEvent))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxNumberOfTimedEvents != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfTimedEvents))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxNumberOfAttributes != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributes))
		i--
		dAtA[i] = 0x20
	}
	if m.Sampler != nil {
		{
			size := m.Sampler.Size()
			i -= size
			if _, err := m.Sampler.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *TraceConfig_ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ConstantSampler != nil {
		{
			size, err := m.ConstantSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_ProbabilitySampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProbabilitySampler != nil {
		{
			size, err := m.ProbabilitySampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RateLimitingSampler != nil {
		{
			size, err := m.RateLimitingSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ConstantSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decision != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Decision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProbabilitySampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProbabilitySampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SamplingProbability))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitingSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Qps != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Qps))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTraceConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraceConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sampler != nil {
		n += m.Sampler.Size()
	}
	if m.MaxNumberOfAttributes != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributes))
	}
	if m.MaxNumberOfTimedEvents != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfTimedEvents))
	}
	if m.MaxNumberOfAttributesPerTimedEvent != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributesPerTimedEvent))
	}
	if m.MaxNumberOfLinks != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfLinks))
	}
	if m.MaxNumberOfAttributesPerLink != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributesPerLink))
	}
	return n
}

func (m *TraceConfig_ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConstantSampler != nil {
		l = m.ConstantSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProbabilitySampler != nil {
		l = m.ProbabilitySampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RateLimitingSampler != nil {
		l = m.RateLimitingSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Decision != 0 {
		n += 1 + sovTraceConfig(uint64(m.Decision))
	}
	return n
}

func (m *ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		n += 9
	}
	return n
}

func (m *RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Qps != 0 {
		n += 1 + sovTraceConfig(uint64(m.Qps))
	}
	return n
}

func sovTraceConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraceConfig(x uint64) (n int) {
	return sovTraceConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConstantSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ConstantSampler{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProbabilitySampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ProbabilitySampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ProbabilitySampler{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitingSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RateLimitingSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_RateLimitingSampler{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributes", wireType)
			}
			m.MaxNumberOfAttributes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfTimedEvents", wireType)
			}
			m.MaxNumberOfTimedEvents = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfTimedEvents |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributesPerTimedEvent", wireType)
			}
			m.MaxNumberOfAttributesPerTimedEvent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributesPerTimedEvent |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfLinks", wireType)
			}
			m.MaxNumberOfLinks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfLinks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributesPerLink", wireType)
			}
			m.MaxNumberOfAttributesPerLink = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributesPerLink |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConstantSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConstantSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConstantSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decision", wireType)
			}
			m.Decision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decision |= ConstantSampler_ConstantDecision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProbabilitySampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProbabilitySampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProbabilitySampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingProbability", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SamplingProbability = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitingSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitingSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitingSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qps", wireType)
			}
			m.Qps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Qps |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTraceConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTraceConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTraceConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTraceConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTraceConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTraceConfig = fmt.Errorf("proto: unexpected end of group")
)