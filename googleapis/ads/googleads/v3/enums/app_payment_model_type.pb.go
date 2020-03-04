// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/app_payment_model_type.proto

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

// Enum describing possible app payment models.
type AppPaymentModelTypeEnum_AppPaymentModelType int32

const (
	// Not specified.
	AppPaymentModelTypeEnum_UNSPECIFIED AppPaymentModelTypeEnum_AppPaymentModelType = 0
	// Used for return value only. Represents value unknown in this version.
	AppPaymentModelTypeEnum_UNKNOWN AppPaymentModelTypeEnum_AppPaymentModelType = 1
	// Represents paid-for apps.
	AppPaymentModelTypeEnum_PAID AppPaymentModelTypeEnum_AppPaymentModelType = 30
)

var AppPaymentModelTypeEnum_AppPaymentModelType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	30: "PAID",
}

var AppPaymentModelTypeEnum_AppPaymentModelType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PAID":        30,
}

func (x AppPaymentModelTypeEnum_AppPaymentModelType) String() string {
	return proto.EnumName(AppPaymentModelTypeEnum_AppPaymentModelType_name, int32(x))
}

func (AppPaymentModelTypeEnum_AppPaymentModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3e62b72aff6b996, []int{0, 0}
}

// Represents a criterion for targeting paid apps.
type AppPaymentModelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPaymentModelTypeEnum) Reset()         { *m = AppPaymentModelTypeEnum{} }
func (m *AppPaymentModelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AppPaymentModelTypeEnum) ProtoMessage()    {}
func (*AppPaymentModelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e62b72aff6b996, []int{0}
}

func (m *AppPaymentModelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Unmarshal(m, b)
}
func (m *AppPaymentModelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Marshal(b, m, deterministic)
}
func (m *AppPaymentModelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPaymentModelTypeEnum.Merge(m, src)
}
func (m *AppPaymentModelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Size(m)
}
func (m *AppPaymentModelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPaymentModelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPaymentModelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AppPaymentModelTypeEnum_AppPaymentModelType", AppPaymentModelTypeEnum_AppPaymentModelType_name, AppPaymentModelTypeEnum_AppPaymentModelType_value)
	proto.RegisterType((*AppPaymentModelTypeEnum)(nil), "google.ads.googleads.v3.enums.AppPaymentModelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/app_payment_model_type.proto", fileDescriptor_d3e62b72aff6b996)
}

var fileDescriptor_d3e62b72aff6b996 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xdd, 0x14, 0x95, 0xec, 0xe0, 0xa8, 0x07, 0x45, 0x9c, 0xb0, 0x3d, 0x40, 0x72, 0xc8,
	0x2d, 0xe2, 0x21, 0x73, 0x73, 0x14, 0xb1, 0x16, 0x74, 0x55, 0xa4, 0x50, 0xa2, 0x09, 0xa1, 0xd0,
	0x26, 0x61, 0xe9, 0x06, 0x7d, 0x1d, 0x8f, 0x3e, 0x8a, 0x8f, 0x22, 0xf8, 0x0e, 0xd2, 0xc4, 0xf6,
	0x34, 0xbd, 0x84, 0x8f, 0x7c, 0xff, 0xdf, 0x97, 0x2f, 0x7f, 0x40, 0xa4, 0xd6, 0xb2, 0x10, 0x88,
	0x71, 0x8b, 0xbc, 0x6c, 0xd4, 0x06, 0x23, 0xa1, 0xd6, 0xa5, 0x45, 0xcc, 0x98, 0xcc, 0xb0, 0xba,
	0x14, 0xaa, 0xca, 0x4a, 0xcd, 0x45, 0x91, 0x55, 0xb5, 0x11, 0xd0, 0xac, 0x74, 0xa5, 0x83, 0x91,
	0x07, 0x20, 0xe3, 0x16, 0x76, 0x2c, 0xdc, 0x60, 0xe8, 0xd8, 0xb3, 0xf3, 0x36, 0xda, 0xe4, 0x88,
	0x29, 0xa5, 0x2b, 0x56, 0xe5, 0x5a, 0x59, 0x0f, 0x4f, 0x9e, 0xc1, 0x09, 0x35, 0x26, 0xf6, 0xd9,
	0x77, 0x4d, 0xf4, 0x63, 0x6d, 0xc4, 0x5c, 0xad, 0xcb, 0xc9, 0x15, 0x38, 0xde, 0x62, 0x05, 0x47,
	0x60, 0xb0, 0x8c, 0x1e, 0xe2, 0xf9, 0x75, 0x78, 0x13, 0xce, 0x67, 0xc3, 0x9d, 0x60, 0x00, 0x0e,
	0x96, 0xd1, 0x6d, 0x74, 0xff, 0x14, 0x0d, 0x7b, 0xc1, 0x21, 0xd8, 0x8b, 0x69, 0x38, 0x1b, 0x5e,
	0x4c, 0xbf, 0x7b, 0x60, 0xfc, 0xa6, 0x4b, 0xf8, 0x6f, 0xbb, 0xe9, 0xe9, 0x96, 0x27, 0xe2, 0xa6,
	0x59, 0xdc, 0x7b, 0x99, 0xfe, 0xa2, 0x52, 0x17, 0x4c, 0x49, 0xa8, 0x57, 0x12, 0x49, 0xa1, 0x5c,
	0xef, 0x76, 0x49, 0x26, 0xb7, 0x7f, 0xec, 0xec, 0xd2, 0x9d, 0xef, 0xfd, 0xdd, 0x05, 0xa5, 0x1f,
	0xfd, 0xd1, 0xc2, 0x47, 0x51, 0x6e, 0xa1, 0x97, 0x8d, 0x4a, 0x30, 0x6c, 0x7e, 0x6a, 0x3f, 0x5b,
	0x3f, 0xa5, 0xdc, 0xa6, 0x9d, 0x9f, 0x26, 0x38, 0x75, 0xfe, 0x57, 0x7f, 0xec, 0x2f, 0x09, 0xa1,
	0xdc, 0x12, 0xd2, 0x4d, 0x10, 0x92, 0x60, 0x42, 0xdc, 0xcc, 0xeb, 0xbe, 0x2b, 0x86, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x46, 0x27, 0xe2, 0xcb, 0x01, 0x00, 0x00,
}