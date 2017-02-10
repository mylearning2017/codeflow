// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration controlling usage of a service.
type Usage struct {
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements" json:"requirements,omitempty"`
	// A list of usage rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules" json:"rules,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// True, if the method allows unregistered calls; false otherwise.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls" json:"allow_unregistered_calls,omitempty"`
}

func (m *UsageRule) Reset()                    { *m = UsageRule{} }
func (m *UsageRule) String() string            { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()               {}
func (*UsageRule) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto", fileDescriptor15)
}

var fileDescriptor15 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x6a, 0x4b, 0x3b, 0x8a, 0x87, 0x05, 0x65, 0xe9, 0x49, 0x16, 0x04, 0x41, 0x48,
	0x40, 0x2f, 0x5e, 0x6d, 0x0f, 0xd2, 0xdb, 0xb2, 0x50, 0xf0, 0x56, 0x62, 0x1c, 0x43, 0x20, 0xcd,
	0xd4, 0x4c, 0x56, 0xdf, 0xc7, 0x27, 0x35, 0x9b, 0x95, 0x5a, 0xaf, 0xbd, 0x04, 0xf2, 0x7f, 0x3f,
	0xdf, 0xcc, 0xc0, 0xd2, 0x10, 0x19, 0x87, 0xc2, 0x90, 0x53, 0xde, 0x08, 0x0a, 0x46, 0x1a, 0xf4,
	0xbb, 0x40, 0x91, 0xe4, 0x80, 0xd4, 0xce, 0xb2, 0x4c, 0x8f, 0x64, 0x0c, 0x9f, 0x56, 0xa3, 0x26,
	0xff, 0x6e, 0x8d, 0xec, 0x58, 0x19, 0x14, 0xb9, 0x58, 0xc2, 0xaf, 0x24, 0xb5, 0xe6, 0xab, 0x63,
	0x85, 0xca, 0x7b, 0x8a, 0x2a, 0x5a, 0xf2, 0x3c, 0x68, 0xeb, 0x17, 0x18, 0xaf, 0xfb, 0x29, 0x65,
	0x0d, 0xe7, 0x01, 0x3f, 0x3a, 0x1b, 0x70, 0x8b, 0x3e, 0x72, 0x55, 0x5c, 0x9f, 0xdc, 0xce, 0xda,
	0x7f, 0x59, 0x79, 0x07, 0xe3, 0xd0, 0x39, 0xe4, 0x6a, 0x92, 0xe0, 0xd9, 0xfd, 0xa5, 0xf8, 0xdb,
	0x49, 0x64, 0x4b, 0x9b, 0x68, 0x3b, 0x74, 0x6a, 0x05, 0xb3, 0x7d, 0x56, 0xce, 0x61, 0xca, 0xe8,
	0x50, 0x47, 0x0a, 0xc9, 0x5c, 0x24, 0xf3, 0xfe, 0x5f, 0x3e, 0x42, 0xa5, 0x9c, 0xa3, 0xaf, 0x4d,
	0xe7, 0x03, 0x1a, 0xcb, 0x11, 0x03, 0xbe, 0x6d, 0x74, 0xca, 0xb8, 0x1a, 0xa5, 0xee, 0xb4, 0xbd,
	0xca, 0x7c, 0x7d, 0x80, 0x97, 0x3d, 0x5d, 0xdc, 0xc0, 0x85, 0xa6, 0xed, 0xc1, 0x16, 0x0b, 0xc8,
	0x23, 0x9b, 0xfe, 0xb4, 0xa6, 0xf8, 0x1e, 0x9d, 0x3e, 0x3f, 0x35, 0xab, 0xd7, 0x49, 0x3e, 0xf5,
	0xe1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x72, 0x2d, 0x47, 0x30, 0x88, 0x01, 0x00, 0x00,
}
