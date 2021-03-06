// POST PROCESSED USING by build_cfg.sh
// 80513108 4392 cfg.proto
// Code generated by protoc-gen-go.
// source: cfg.proto
// DO NOT EDIT!

/*
Package istio_mixer_v1_config is a generated protocol buffer package.

It is generated from these files:
	cfg.proto

It has these top-level messages:
	ServiceConfig
	AspectRule
	Aspect
	Adapter
	GlobalConfig
	ClientConfig
*/
package istio_mixer_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configures a set of services
// following example configures metrics collection and ratelimit for
// all services
// # service config
// subject: "namespace:ns1"
// revision: "1011"
// rules:
// - selector: target_name == "*"
//  aspects:
//  - kind: metrics
//    params:
//      metrics:   # defines metric collection across the board.
//      - name: response_time_by_status_code
//        value: metric.response_time     # certain attributes are metrics
//        metric_kind: DELTA
//        labels:
//        - key: response.status_code
//  - kind: ratelimiter
//    params:
//      limits:  # imposes 2 limits, 100/s per source and destination
//      - limit: "100/s"
//        labels:
//          - key: src.service_id
//          - key: target.service_id
//       - limit: "1000/s"  # every destination service gets 1000/s
//        labels:
//          - key: target.service_id
type ServiceConfig struct {
	// subject is unique for a config type
	// 2 config with the same subject will overwrite each other
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	// revision of this config. This is assigned by the server
	Revision string        `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Rules    []*AspectRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *ServiceConfig) Reset()                    { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()               {}
func (*ServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceConfig) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ServiceConfig) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ServiceConfig) GetRules() []*AspectRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// AspectRules are intent based
type AspectRule struct {
	// selector is an attributes based predicate.
	// attr1 == "20" && attr2 == "30"
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// The following aspects apply when the selector predicate evaluates to True
	Aspects []*Aspect `protobuf:"bytes,2,rep,name=aspects" json:"aspects,omitempty"`
	// Nested aspect Rule is evaluated if selector predicate evaluates to True
	Rules []*AspectRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *AspectRule) Reset()                    { *m = AspectRule{} }
func (m *AspectRule) String() string            { return proto.CompactTextString(m) }
func (*AspectRule) ProtoMessage()               {}
func (*AspectRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AspectRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *AspectRule) GetAspects() []*Aspect {
	if m != nil {
		return m.Aspects
	}
	return nil
}

func (m *AspectRule) GetRules() []*AspectRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Aspect is intent based. It specifies the intent "kind"
// following example specifies that the user would like to collect
// response_time with 3 labels (src_consumer_id, target_response_status_code,
// target_service_name)
//
// The Input section tells if target_service_name is not available it can be
// computed using the given expression
//
//      kind: istio/metrics
//      params:
//        metrics:
//        - name: response_time     # What to call this metric outbound.
//          value: metric_response_time  # from wellknown vocabulary
//          metric_kind: DELTA
//          labels:
//          - key: src_consumer_id
//          - key: target_response_status_code
//          - key: target_service_name
//      Inputs:
//           Attr.target_service_name: target_service_name || target_service_id
type Aspect struct {
	Kind    string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Adapter string `protobuf:"bytes,2,opt,name=adapter" json:"adapter,omitempty"`
	// maps from isio Attribute space to adapter.Input proto defined
	// by the aspect
	Inputs map[string]string `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Struct representation of a proto defined by the aspect
	Params interface{} `protobuf:"bytes,4,opt,name=params" json:"params,omitempty"`
}

func (m *Aspect) Reset()                    { *m = Aspect{} }
func (m *Aspect) String() string            { return proto.CompactTextString(m) }
func (*Aspect) ProtoMessage()               {}
func (*Aspect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Aspect) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Aspect) GetAdapter() string {
	if m != nil {
		return m.Adapter
	}
	return ""
}

func (m *Aspect) GetInputs() map[string]string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Aspect) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

// Adapter config defines specifics of adapter implementations
// We define an adapter that provides "metrics" aspect
// Kind: istio/metrics
// Name: metrics-statsd
// Impl: “istio.io/adapters/statsd”
// Args:
//    Host: statd.svc.cluster
//    Port: 8125
type Adapter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Impl string `protobuf:"bytes,3,opt,name=impl" json:"impl,omitempty"`
	// Struct representation of a proto defined by the implementation
	Params interface{} `protobuf:"bytes,4,opt,name=params" json:"params,omitempty"`
}

func (m *Adapter) Reset()                    { *m = Adapter{} }
func (m *Adapter) String() string            { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()               {}
func (*Adapter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Adapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Adapter) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Adapter) GetImpl() string {
	if m != nil {
		return m.Impl
	}
	return ""
}

func (m *Adapter) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

// GlobalConfig defines configuration elements that are available
// for the rest of the config
// It is used to configure adapters and make them available in AspectRules
type GlobalConfig struct {
	Revision string     `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	Adapters []*Adapter `protobuf:"bytes,2,rep,name=adapters" json:"adapters,omitempty"`
}

func (m *GlobalConfig) Reset()                    { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string            { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()               {}
func (*GlobalConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GlobalConfig) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *GlobalConfig) GetAdapters() []*Adapter {
	if m != nil {
		return m.Adapters
	}
	return nil
}

// ClientConfig defines configuration from a client perspective.
// ServiceA can define rules about what happens when it is acting as client
// to other services
type ClientConfig struct {
	Subject  string        `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	Revision string        `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Rules    []*AspectRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientConfig) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ClientConfig) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ClientConfig) GetRules() []*AspectRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "istio.mixer.v1.config.ServiceConfig")
	proto.RegisterType((*AspectRule)(nil), "istio.mixer.v1.config.AspectRule")
	proto.RegisterType((*Aspect)(nil), "istio.mixer.v1.config.Aspect")
	proto.RegisterType((*Adapter)(nil), "istio.mixer.v1.config.Adapter")
	proto.RegisterType((*GlobalConfig)(nil), "istio.mixer.v1.config.GlobalConfig")
	proto.RegisterType((*ClientConfig)(nil), "istio.mixer.v1.config.ClientConfig")
}

func init() { proto.RegisterFile("cfg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x8f, 0x4d, 0x77, 0xa7, 0x8b, 0x84, 0x2c, 0x10, 0x51, 0x05, 0xa8, 0xe4, 0x54,
	0x2e, 0xae, 0x58, 0x0e, 0x0b, 0xbd, 0x55, 0x15, 0x42, 0x5c, 0xd3, 0x27, 0x70, 0xd2, 0x49, 0x64,
	0xea, 0xc4, 0x91, 0xed, 0x44, 0xf4, 0x00, 0x0f, 0xc2, 0xfb, 0xf1, 0x1e, 0x28, 0x8e, 0x13, 0x2a,
	0x04, 0x95, 0xe0, 0xb2, 0xb7, 0x99, 0xc9, 0xcf, 0xf9, 0xff, 0xe7, 0x03, 0x6e, 0xf2, 0xa2, 0xa4,
	0x8d, 0x92, 0x46, 0x92, 0xa7, 0x5c, 0x1b, 0x2e, 0x69, 0xc5, 0xbf, 0xa0, 0xa2, 0xdd, 0x1b, 0x9a,
	0xcb, 0xba, 0xe0, 0xe5, 0xe2, 0x79, 0x29, 0x65, 0x29, 0x70, 0x6d, 0xa1, 0xac, 0x2d, 0xd6, 0xda,
	0xa8, 0x36, 0x37, 0xc3, 0xa3, 0xe4, 0x1b, 0x3c, 0xda, 0xa3, 0xea, 0x78, 0x8e, 0x3b, 0x8b, 0x93,
	0x18, 0x66, 0xba, 0xcd, 0x3e, 0x63, 0x6e, 0x62, 0x6f, 0xe9, 0xad, 0x6e, 0xd2, 0x31, 0x25, 0x0b,
	0xb8, 0x56, 0xd8, 0x71, 0xcd, 0x65, 0x1d, 0xfb, 0xf6, 0xd3, 0x94, 0x93, 0x7b, 0xb8, 0x52, 0xad,
	0x40, 0x1d, 0x07, 0xcb, 0x60, 0x35, 0xbf, 0x7b, 0x45, 0xff, 0xe8, 0x85, 0x6e, 0x75, 0x83, 0xb9,
	0x49, 0x5b, 0x81, 0xe9, 0xc0, 0x27, 0xdf, 0x3d, 0x80, 0x5f, 0xd5, 0x5e, 0x43, 0xa3, 0xc0, 0xdc,
	0x48, 0xe5, 0xe4, 0xa7, 0x9c, 0xdc, 0xc3, 0x8c, 0x59, 0x52, 0xc7, 0xbe, 0x55, 0x79, 0x71, 0x59,
	0x65, 0xa4, 0xff, 0xdf, 0xdc, 0x0f, 0x0f, 0xa2, 0xa1, 0x4a, 0x08, 0x84, 0x47, 0x5e, 0x1f, 0x9c,
	0x29, 0x1b, 0xf7, 0xa3, 0x62, 0x07, 0xd6, 0x18, 0x54, 0x6e, 0x1e, 0x63, 0x4a, 0xb6, 0x10, 0xf1,
	0xba, 0x69, 0xcd, 0x28, 0xf9, 0xfa, 0xa2, 0x24, 0xfd, 0x64, 0xd9, 0x0f, 0xb5, 0x51, 0xa7, 0xd4,
	0x3d, 0x24, 0x6b, 0x88, 0x1a, 0xa6, 0x58, 0xa5, 0xe3, 0x70, 0xe9, 0xad, 0xe6, 0x77, 0xcf, 0xe8,
	0xb0, 0x47, 0x3a, 0xee, 0x91, 0xee, 0xed, 0x1e, 0x53, 0x87, 0x2d, 0xde, 0xc3, 0xfc, 0xec, 0x3f,
	0xe4, 0x31, 0x04, 0x47, 0x3c, 0x39, 0xbf, 0x7d, 0x48, 0x9e, 0xc0, 0x55, 0xc7, 0x44, 0x8b, 0xce,
	0xec, 0x90, 0x6c, 0xfc, 0x77, 0x5e, 0xd2, 0xc1, 0x6c, 0xeb, 0x9c, 0x13, 0x08, 0x6b, 0x56, 0xe1,
	0xd8, 0x67, 0x1f, 0x4f, 0xbd, 0xfb, 0x67, 0xbd, 0x13, 0x08, 0x79, 0xd5, 0x88, 0x38, 0x18, 0x6a,
	0x7d, 0xfc, 0xcf, 0x96, 0x93, 0x02, 0x6e, 0x3f, 0x0a, 0x99, 0x31, 0xe1, 0x6e, 0xef, 0xfc, 0xc2,
	0xbc, 0xdf, 0x2e, 0x6c, 0x03, 0xd7, 0x6e, 0xba, 0xe3, 0xfa, 0x5f, 0xfe, 0x6d, 0xa8, 0x03, 0x96,
	0x4e, 0x7c, 0xf2, 0x15, 0x6e, 0x77, 0x82, 0x63, 0x6d, 0x1e, 0xe4, 0xc6, 0xb3, 0xc8, 0xf6, 0xff,
	0xf6, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xe9, 0x9b, 0x87, 0xac, 0x03, 0x00, 0x00,
}
