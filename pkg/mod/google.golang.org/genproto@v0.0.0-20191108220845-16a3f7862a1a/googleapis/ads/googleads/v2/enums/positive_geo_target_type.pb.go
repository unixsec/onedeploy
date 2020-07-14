// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/positive_geo_target_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The possible positive geo target types.
type PositiveGeoTargetTypeEnum_PositiveGeoTargetType int32

const (
	// Not specified.
	PositiveGeoTargetTypeEnum_UNSPECIFIED PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 0
	// The value is unknown in this version.
	PositiveGeoTargetTypeEnum_UNKNOWN PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 1
	// Specifies that an ad is triggered if the user is in,
	// or shows interest in, advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 5
	// Specifies that an ad is triggered if the user
	// searches for advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_SEARCH_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 6
	// Specifies that an ad is triggered if the user is in
	// or regularly in advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 7
)

var PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "PRESENCE_OR_INTEREST",
	6: "SEARCH_INTEREST",
	7: "PRESENCE",
}

var PositiveGeoTargetTypeEnum_PositiveGeoTargetType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PRESENCE_OR_INTEREST": 5,
	"SEARCH_INTEREST":      6,
	"PRESENCE":             7,
}

func (x PositiveGeoTargetTypeEnum_PositiveGeoTargetType) String() string {
	return proto.EnumName(PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name, int32(x))
}

func (PositiveGeoTargetTypeEnum_PositiveGeoTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e39be6a5f5414c2d, []int{0, 0}
}

// Container for enum describing possible positive geo target types.
type PositiveGeoTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositiveGeoTargetTypeEnum) Reset()         { *m = PositiveGeoTargetTypeEnum{} }
func (m *PositiveGeoTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PositiveGeoTargetTypeEnum) ProtoMessage()    {}
func (*PositiveGeoTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39be6a5f5414c2d, []int{0}
}

func (m *PositiveGeoTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Unmarshal(m, b)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositiveGeoTargetTypeEnum.Merge(m, src)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Size(m)
}
func (m *PositiveGeoTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PositiveGeoTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PositiveGeoTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.PositiveGeoTargetTypeEnum_PositiveGeoTargetType", PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name, PositiveGeoTargetTypeEnum_PositiveGeoTargetType_value)
	proto.RegisterType((*PositiveGeoTargetTypeEnum)(nil), "google.ads.googleads.v2.enums.PositiveGeoTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/positive_geo_target_type.proto", fileDescriptor_e39be6a5f5414c2d)
}

var fileDescriptor_e39be6a5f5414c2d = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x13, 0x37, 0xc9, 0x84, 0x95, 0xaa, 0xa0, 0xc3, 0x1d, 0xb6, 0x07, 0x48, 0xa1, 0xde,
	0xa2, 0x97, 0x6e, 0xc6, 0x39, 0x84, 0xae, 0xb4, 0xdd, 0x04, 0x29, 0x94, 0x6a, 0x43, 0x28, 0x6c,
	0x49, 0x69, 0xb2, 0xc1, 0x9e, 0xc2, 0x77, 0xf0, 0xe8, 0xa3, 0xf8, 0x28, 0x5e, 0x7c, 0x05, 0x69,
	0xb2, 0xce, 0xcb, 0xf4, 0x12, 0x3e, 0xf2, 0xfd, 0xe1, 0xf7, 0x7d, 0xe0, 0x96, 0x72, 0x4e, 0x17,
	0xc4, 0x4e, 0x33, 0x61, 0x6b, 0x58, 0xa1, 0xb5, 0x63, 0x13, 0xb6, 0x5a, 0x0a, 0xbb, 0xe0, 0x22,
	0x97, 0xf9, 0x9a, 0x24, 0x94, 0xf0, 0x44, 0xa6, 0x25, 0x25, 0x32, 0x91, 0x9b, 0x82, 0xc0, 0xa2,
	0xe4, 0x92, 0x5b, 0x3d, 0x6d, 0x81, 0x69, 0x26, 0xe0, 0xce, 0x0d, 0xd7, 0x0e, 0x54, 0xee, 0xee,
	0x55, 0x1d, 0x5e, 0xe4, 0x76, 0xca, 0x18, 0x97, 0xa9, 0xcc, 0x39, 0x13, 0xda, 0x3c, 0x78, 0x33,
	0xc0, 0xa5, 0xbf, 0xcd, 0x1f, 0x13, 0x1e, 0xa9, 0xf4, 0x68, 0x53, 0x10, 0xcc, 0x56, 0xcb, 0x41,
	0x09, 0xce, 0xf7, 0x92, 0x56, 0x07, 0xb4, 0x67, 0x5e, 0xe8, 0xe3, 0xd1, 0xe4, 0x7e, 0x82, 0xef,
	0xcc, 0x03, 0xab, 0x0d, 0x5a, 0x33, 0xef, 0xd1, 0x9b, 0x3e, 0x79, 0xa6, 0x61, 0x5d, 0x80, 0x33,
	0x3f, 0xc0, 0x21, 0xf6, 0x46, 0x38, 0x99, 0x06, 0xc9, 0xc4, 0x8b, 0x70, 0x80, 0xc3, 0xc8, 0x3c,
	0xb2, 0x4e, 0x41, 0x27, 0xc4, 0x6e, 0x30, 0x7a, 0xf8, 0xfd, 0x6c, 0x5a, 0x27, 0xe0, 0xb8, 0x96,
	0x9b, 0xad, 0xe1, 0xb7, 0x01, 0xfa, 0xaf, 0x7c, 0x09, 0xff, 0x6d, 0x35, 0xec, 0xee, 0xbd, 0xcb,
	0xaf, 0x3a, 0xf9, 0xc6, 0xf3, 0x70, 0x6b, 0xa6, 0x7c, 0x91, 0x32, 0x0a, 0x79, 0x49, 0x6d, 0x4a,
	0x98, 0x6a, 0x5c, 0x0f, 0x5c, 0xe4, 0xe2, 0x8f, 0xbd, 0x6f, 0xd4, 0xfb, 0xde, 0x38, 0x1c, 0xbb,
	0xee, 0x47, 0xa3, 0x37, 0xd6, 0x51, 0x6e, 0x26, 0xa0, 0x86, 0x15, 0x9a, 0x3b, 0xb0, 0x1a, 0x48,
	0x7c, 0xd6, 0x7c, 0xec, 0x66, 0x22, 0xde, 0xf1, 0xf1, 0xdc, 0x89, 0x15, 0xff, 0xd5, 0xe8, 0xeb,
	0x4f, 0x84, 0xdc, 0x4c, 0x20, 0xb4, 0x53, 0x20, 0x34, 0x77, 0x10, 0x52, 0x9a, 0x97, 0xa6, 0x3a,
	0xec, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xac, 0x07, 0x53, 0x94, 0x07, 0x02, 0x00, 0x00,
}