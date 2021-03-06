// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/kubernetes/kubernetes.proto

package kubernetes // import "github.com/solo-io/supergloo/pkg/api/external/gloo/v1/plugins/kubernetes"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import plugins "github.com/solo-io/supergloo/pkg/api/external/gloo/v1/plugins"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for Kubernetes Upstreams
// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support mulitple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The port where the Service is listening.
	ServicePort uint32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Allows finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//     An optional Service Spec describing the service listening at this address
	ServiceSpec          *plugins.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_kubernetes_75f8d4515c6916ea, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (dst *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(dst, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *plugins.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterMapType((map[string]string)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec.SelectorEntry")
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("plugins/kubernetes/kubernetes.proto", fileDescriptor_kubernetes_75f8d4515c6916ea)
}

var fileDescriptor_kubernetes_75f8d4515c6916ea = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6b, 0x45, 0xb7, 0x2d, 0xd4, 0xa5, 0x87, 0xd0, 0x83, 0xa6, 0x7a, 0x09, 0x88,
	0xbb, 0x58, 0x2f, 0x62, 0x6f, 0xa2, 0xe0, 0x41, 0x8a, 0xb4, 0x88, 0xe0, 0x45, 0xb6, 0x61, 0x88,
	0xa1, 0x49, 0x66, 0xd9, 0xdd, 0x14, 0xfb, 0x46, 0xbe, 0x94, 0x17, 0x9f, 0x44, 0x36, 0x69, 0xd3,
	0x15, 0xa4, 0xde, 0x66, 0x7f, 0xbe, 0xf9, 0x67, 0xf6, 0x1f, 0x72, 0x26, 0xd3, 0x22, 0x4e, 0x72,
	0xcd, 0x17, 0xc5, 0x1c, 0x54, 0x0e, 0x06, 0xdc, 0x92, 0x49, 0x85, 0x06, 0xe9, 0x89, 0xab, 0x54,
	0x3c, 0x8b, 0x53, 0x44, 0xa6, 0x31, 0x45, 0x96, 0xe0, 0xa0, 0x1f, 0x63, 0x8c, 0x25, 0xcb, 0x6d,
	0x55, 0xb5, 0x0d, 0x1e, 0xe3, 0xc4, 0xbc, 0x17, 0x73, 0x16, 0x61, 0xc6, 0x2d, 0x79, 0x91, 0x20,
	0xd7, 0x85, 0x04, 0x65, 0x7b, 0xb9, 0x90, 0x09, 0x87, 0x0f, 0x03, 0x2a, 0x17, 0x29, 0x2f, 0x95,
	0xe5, 0x25, 0xdf, 0x6c, 0xa3, 0x41, 0x2d, 0x93, 0x08, 0xde, 0xb4, 0x84, 0xa8, 0x72, 0x3b, 0xfd,
	0x6a, 0x90, 0xce, 0xb3, 0xd4, 0x46, 0x81, 0xc8, 0x66, 0x12, 0x22, 0x3a, 0x24, 0x9d, 0x0d, 0x96,
	0x8b, 0x0c, 0x7c, 0x2f, 0xf0, 0xc2, 0xc3, 0x69, 0x7b, 0xad, 0x4d, 0x44, 0x06, 0xf4, 0x9c, 0x1c,
	0xb9, 0x88, 0x96, 0x22, 0x02, 0xbf, 0x51, 0x72, 0x3d, 0x87, 0x2b, 0x75, 0xd7, 0x4f, 0xa2, 0x32,
	0x7e, 0x33, 0xf0, 0xc2, 0x6e, 0xed, 0xf7, 0x84, 0xca, 0xd0, 0x17, 0x72, 0xa0, 0x21, 0x85, 0xc8,
	0xa0, 0xf2, 0xf7, 0x82, 0x66, 0xd8, 0x1e, 0x8d, 0xd9, 0x3f, 0xd9, 0x30, 0x77, 0x67, 0x36, 0x5b,
	0x77, 0xdf, 0xe7, 0x46, 0xad, 0xa6, 0xb5, 0x19, 0xbd, 0xdb, 0xce, 0xb6, 0x5f, 0xf6, 0x5b, 0x81,
	0x17, 0xb6, 0x47, 0xc3, 0xbf, 0x1d, 0x67, 0x15, 0x69, 0x0d, 0xeb, 0xf5, 0xec, 0x63, 0x30, 0x26,
	0xdd, 0x5f, 0x03, 0x68, 0x8f, 0x34, 0x17, 0xb0, 0x5a, 0x27, 0x63, 0x4b, 0xda, 0x27, 0xad, 0xa5,
	0x48, 0x8b, 0x4d, 0x0a, 0xd5, 0xe3, 0xa6, 0x71, 0xed, 0xdd, 0x4e, 0x3e, 0xbf, 0x8f, 0xbd, 0xd7,
	0x87, 0x9d, 0x37, 0x93, 0x8b, 0x78, 0xf7, 0xdd, 0xb6, 0x61, 0xcc, 0xf7, 0xcb, 0xb3, 0x5d, 0xfd,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xcf, 0x13, 0xcc, 0x62, 0x02, 0x00, 0x00,
}
