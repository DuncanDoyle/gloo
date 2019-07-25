// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Optional, if not set, Gloo will try to use the default AWS secret specified by environment variables.
	// If a secret is not provided, the environment must specify both the AWS access key and secret.
	// The environment variables used to indicate the AWS account can be:
	// - for the access key: "AWS_ACCESS_KEY_ID" or "AWS_ACCESS_KEY"
	// - for the secret: "AWS_SECRET_ACCESS_KEY" or "AWS_SECRET_KEY"
	// If set, a [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	// Gloo will create an EC2 API client with this credential. You may choose to use a credential with limited access
	// in conjunction with a list of Roles, specified by their Amazon Resource Number (ARN).
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// Optional, list of Amazon Resource Numbers (ARNs) referring to IAM Roles that should be assumed when the Upstream
	// queries for eligible EC2 instances.
	// If provided, Gloo will create an EC2 API client with the provided roles. If not provided, Gloo will not assume
	// any roles.
	RoleArns []string `protobuf:"bytes,6,rep,name=role_arns,json=roleArns,proto3" json:"role_arns,omitempty"`
	// List of tag filters for selecting instances
	// An instance must match all the filters in order to be selected
	// Filter keys are not case-sensitive
	Filters []*TagFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// If set, will use the EC2 public IP address. Defaults to the private IP address.
	PublicIp bool `protobuf:"varint,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	// If set, will use this port on EC2 instances. Defaults to port 80.
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetRoleArns() []string {
	if m != nil {
		return m.RoleArns
	}
	return nil
}

func (m *UpstreamSpec) GetFilters() []*TagFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *UpstreamSpec) GetPublicIp() bool {
	if m != nil {
		return m.PublicIp
	}
	return false
}

func (m *UpstreamSpec) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type TagFilter struct {
	// Types that are valid to be assigned to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec                 isTagFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter.Unmarshal(m, b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return xxx_messageInfo_TagFilter.Size(m)
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

type isTagFilter_Spec interface {
	isTagFilter_Spec()
	Equal(interface{}) bool
}

type TagFilter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}
type TagFilter_KvPair_ struct {
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof"`
}

func (*TagFilter_Key) isTagFilter_Spec()     {}
func (*TagFilter_KvPair_) isTagFilter_Spec() {}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TagFilter) GetKey() string {
	if x, ok := m.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := m.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TagFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TagFilter_OneofMarshaler, _TagFilter_OneofUnmarshaler, _TagFilter_OneofSizer, []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
}

func _TagFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TagFilter)
	// spec
	switch x := m.Spec.(type) {
	case *TagFilter_Key:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Key)
	case *TagFilter_KvPair_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KvPair); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TagFilter.Spec has unexpected type %T", x)
	}
	return nil
}

func _TagFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TagFilter)
	switch tag {
	case 1: // spec.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Spec = &TagFilter_Key{x}
		return true, err
	case 2: // spec.kv_pair
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TagFilter_KvPair)
		err := b.DecodeMessage(msg)
		m.Spec = &TagFilter_KvPair_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TagFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TagFilter)
	// spec
	switch x := m.Spec.(type) {
	case *TagFilter_Key:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Key)))
		n += len(x.Key)
	case *TagFilter_KvPair_:
		s := proto.Size(x.KvPair)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TagFilter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter_KvPair) Reset()         { *m = TagFilter_KvPair{} }
func (m *TagFilter_KvPair) String() string { return proto.CompactTextString(m) }
func (*TagFilter_KvPair) ProtoMessage()    {}
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1, 0}
}
func (m *TagFilter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter_KvPair.Unmarshal(m, b)
}
func (m *TagFilter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter_KvPair.Marshal(b, m, deterministic)
}
func (m *TagFilter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter_KvPair.Merge(m, src)
}
func (m *TagFilter_KvPair) XXX_Size() int {
	return xxx_messageInfo_TagFilter_KvPair.Size(m)
}
func (m *TagFilter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter_KvPair proto.InternalMessageInfo

func (m *TagFilter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TagFilter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*TagFilter)(nil), "aws_ec2.plugins.gloo.solo.io.TagFilter")
	proto.RegisterType((*TagFilter_KvPair)(nil), "aws_ec2.plugins.gloo.solo.io.TagFilter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto", fileDescriptor_fc1fd6f1173c4563)
}

var fileDescriptor_fc1fd6f1173c4563 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x6d, 0xd8, 0x6d, 0xba, 0x71, 0x41, 0x42, 0x56, 0x85, 0xc2, 0x82, 0x50, 0xd4, 0x0b, 0x39,
	0x80, 0x0d, 0xcb, 0x85, 0x23, 0xed, 0xa1, 0x6a, 0xe1, 0x82, 0x0c, 0x5c, 0xb8, 0x44, 0x5e, 0x6b,
	0x62, 0x4c, 0xd2, 0x1d, 0x6b, 0xec, 0x2c, 0xe2, 0x7f, 0x38, 0xf0, 0x5d, 0xfc, 0x01, 0x7f, 0x80,
	0x12, 0x77, 0x2b, 0x0e, 0x50, 0xf5, 0xe4, 0x79, 0xa3, 0x37, 0xef, 0xcd, 0x93, 0x87, 0xbd, 0xb5,
	0x2e, 0x7e, 0x19, 0xd6, 0xc2, 0xe0, 0xa5, 0x0c, 0xd8, 0xe3, 0x73, 0x87, 0xd2, 0xf6, 0x88, 0xd2,
	0x13, 0x7e, 0x05, 0x13, 0x43, 0x42, 0xda, 0x3b, 0xb9, 0x7d, 0x29, 0x7d, 0x3f, 0x58, 0xb7, 0x09,
	0x52, 0x7f, 0x0b, 0x12, 0xcc, 0x6a, 0x7c, 0x1b, 0x30, 0x2b, 0xe1, 0x09, 0x23, 0xf2, 0xc7, 0xd7,
	0x30, 0xd1, 0xc4, 0x38, 0x2a, 0x46, 0x55, 0xe1, 0x70, 0x79, 0x64, 0xd1, 0xe2, 0x44, 0x94, 0x63,
	0x95, 0x66, 0x96, 0xcf, 0xfe, 0xe1, 0x3f, 0xbd, 0x9d, 0x8b, 0x3b, 0x57, 0x82, 0x36, 0xb1, 0x8f,
	0x7f, 0x67, 0xec, 0xee, 0x27, 0x1f, 0x22, 0x81, 0xbe, 0xfc, 0xe0, 0xc1, 0xf0, 0x07, 0x2c, 0x27,
	0xb0, 0x0e, 0x37, 0x65, 0x56, 0x65, 0x75, 0xa1, 0xae, 0x10, 0x7f, 0xcd, 0x58, 0x00, 0x43, 0x10,
	0x1b, 0x82, 0xb6, 0xbc, 0x53, 0x65, 0xf5, 0xe1, 0xea, 0xa1, 0x30, 0x48, 0xb0, 0xdb, 0x47, 0x28,
	0x08, 0x38, 0x90, 0x01, 0x05, 0xad, 0x2a, 0x12, 0x59, 0x41, 0xcb, 0x1f, 0xb1, 0x82, 0xb0, 0x87,
	0x46, 0xd3, 0x26, 0x94, 0x79, 0x35, 0xab, 0x0b, 0xb5, 0x18, 0x1b, 0x27, 0xb4, 0x09, 0xfc, 0x84,
	0x1d, 0xb4, 0xae, 0x8f, 0x40, 0xa1, 0x9c, 0x55, 0xb3, 0xfa, 0x70, 0xf5, 0x54, 0xdc, 0x94, 0x59,
	0x7c, 0xd4, 0xf6, 0x6c, 0xe2, 0xab, 0xdd, 0xdc, 0xa8, 0xef, 0x87, 0x75, 0xef, 0x4c, 0xe3, 0x7c,
	0x39, 0xaf, 0xb2, 0x7a, 0xa1, 0x16, 0xa9, 0x71, 0xe1, 0x39, 0x67, 0x73, 0x8f, 0x14, 0xcb, 0xfd,
	0x2a, 0xab, 0xef, 0xa9, 0xa9, 0x3e, 0xfe, 0x91, 0xb1, 0xe2, 0x5a, 0x87, 0x73, 0x36, 0xeb, 0xe0,
	0x7b, 0x4a, 0x7b, 0xbe, 0xa7, 0x46, 0xc0, 0x2f, 0xd8, 0x41, 0xb7, 0x6d, 0xbc, 0x76, 0x74, 0x95,
	0x54, 0xdc, 0x72, 0x2b, 0xf1, 0x6e, 0xfb, 0x5e, 0x3b, 0x3a, 0xdf, 0x53, 0x79, 0x37, 0x55, 0xcb,
	0x17, 0x2c, 0x4f, 0x3d, 0x7e, 0xff, 0x2f, 0xa3, 0x64, 0x73, 0xc4, 0xf6, 0xb7, 0xba, 0x1f, 0x60,
	0x32, 0x29, 0x54, 0x02, 0xa7, 0x39, 0x9b, 0x07, 0x0f, 0xe6, 0xf4, 0xec, 0xe7, 0xaf, 0x27, 0xd9,
	0xe7, 0x37, 0xb7, 0x3b, 0x27, 0xdf, 0xd9, 0xff, 0x9c, 0xd4, 0x3a, 0x9f, 0x7e, 0xfa, 0xd5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x5b, 0x17, 0xa7, 0x99, 0x02, 0x00, 0x00,
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
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if len(this.RoleArns) != len(that1.RoleArns) {
		return false
	}
	for i := range this.RoleArns {
		if this.RoleArns[i] != that1.RoleArns[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if this.PublicIp != that1.PublicIp {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
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
	if that1.Spec == nil {
		if this.Spec != nil {
			return false
		}
	} else if this.Spec == nil {
		return false
	} else if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_Key)
	if !ok {
		that2, ok := that.(TagFilter_Key)
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
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *TagFilter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair_)
	if !ok {
		that2, ok := that.(TagFilter_KvPair_)
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
	if !this.KvPair.Equal(that1.KvPair) {
		return false
	}
	return true
}
func (this *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}