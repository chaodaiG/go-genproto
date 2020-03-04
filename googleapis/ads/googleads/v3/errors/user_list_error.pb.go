// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/user_list_error.proto

package errors

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

// Enum describing possible user list errors.
type UserListErrorEnum_UserListError int32

const (
	// Enum unspecified.
	UserListErrorEnum_UNSPECIFIED UserListErrorEnum_UserListError = 0
	// The received error code is not known in this version.
	UserListErrorEnum_UNKNOWN UserListErrorEnum_UserListError = 1
	// Creating and updating external remarketing user lists is not supported.
	UserListErrorEnum_EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 2
	// Concrete type of user list is required.
	UserListErrorEnum_CONCRETE_TYPE_REQUIRED UserListErrorEnum_UserListError = 3
	// Creating/updating user list conversion types requires specifying the
	// conversion type Id.
	UserListErrorEnum_CONVERSION_TYPE_ID_REQUIRED UserListErrorEnum_UserListError = 4
	// Remarketing user list cannot have duplicate conversion types.
	UserListErrorEnum_DUPLICATE_CONVERSION_TYPES UserListErrorEnum_UserListError = 5
	// Conversion type is invalid/unknown.
	UserListErrorEnum_INVALID_CONVERSION_TYPE UserListErrorEnum_UserListError = 6
	// User list description is empty or invalid.
	UserListErrorEnum_INVALID_DESCRIPTION UserListErrorEnum_UserListError = 7
	// User list name is empty or invalid.
	UserListErrorEnum_INVALID_NAME UserListErrorEnum_UserListError = 8
	// Type of the UserList does not match.
	UserListErrorEnum_INVALID_TYPE UserListErrorEnum_UserListError = 9
	// Embedded logical user lists are not allowed.
	UserListErrorEnum_CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 10
	// User list rule operand is invalid.
	UserListErrorEnum_INVALID_USER_LIST_LOGICAL_RULE_OPERAND UserListErrorEnum_UserListError = 11
	// Name is already being used for another user list for the account.
	UserListErrorEnum_NAME_ALREADY_USED UserListErrorEnum_UserListError = 12
	// Name is required when creating a new conversion type.
	UserListErrorEnum_NEW_CONVERSION_TYPE_NAME_REQUIRED UserListErrorEnum_UserListError = 13
	// The given conversion type name has been used.
	UserListErrorEnum_CONVERSION_TYPE_NAME_ALREADY_USED UserListErrorEnum_UserListError = 14
	// Only an owner account may edit a user list.
	UserListErrorEnum_OWNERSHIP_REQUIRED_FOR_SET UserListErrorEnum_UserListError = 15
	// Creating user list without setting type in oneof user_list field, or
	// creating/updating read-only user list types is not allowed.
	UserListErrorEnum_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 16
	// Rule is invalid.
	UserListErrorEnum_INVALID_RULE UserListErrorEnum_UserListError = 17
	// The specified date range is empty.
	UserListErrorEnum_INVALID_DATE_RANGE UserListErrorEnum_UserListError = 27
	// A UserList which is privacy sensitive or legal rejected cannot be mutated
	// by external users.
	UserListErrorEnum_CAN_NOT_MUTATE_SENSITIVE_USERLIST UserListErrorEnum_UserListError = 28
	// Maximum number of rulebased user lists a customer can have.
	UserListErrorEnum_MAX_NUM_RULEBASED_USERLISTS UserListErrorEnum_UserListError = 29
	// BasicUserList's billable record field cannot be modified once it is set.
	UserListErrorEnum_CANNOT_MODIFY_BILLABLE_RECORD_COUNT UserListErrorEnum_UserListError = 30
	// crm_based_user_list.app_id field must be set when upload_key_type is
	// MOBILE_ADVERTISING_ID.
	UserListErrorEnum_APP_ID_NOT_SET UserListErrorEnum_UserListError = 31
	// Name of the user list is reserved for system generated lists and cannot
	// be used.
	UserListErrorEnum_USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST UserListErrorEnum_UserListError = 32
	// Advertiser needs to be whitelisted to use remarketing lists created from
	// advertiser uploaded data (e.g., Customer Match lists).
	UserListErrorEnum_ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA UserListErrorEnum_UserListError = 33
	// The provided rule_type is not supported for the user list.
	UserListErrorEnum_RULE_TYPE_IS_NOT_SUPPORTED UserListErrorEnum_UserListError = 34
	// Similar user list cannot be used as a logical user list operand.
	UserListErrorEnum_CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 35
	// Logical user list should not have a mix of CRM based user list and other
	// types of lists in its rules.
	UserListErrorEnum_CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS UserListErrorEnum_UserListError = 36
)

var UserListErrorEnum_UserListError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED",
	3:  "CONCRETE_TYPE_REQUIRED",
	4:  "CONVERSION_TYPE_ID_REQUIRED",
	5:  "DUPLICATE_CONVERSION_TYPES",
	6:  "INVALID_CONVERSION_TYPE",
	7:  "INVALID_DESCRIPTION",
	8:  "INVALID_NAME",
	9:  "INVALID_TYPE",
	10: "CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND",
	11: "INVALID_USER_LIST_LOGICAL_RULE_OPERAND",
	12: "NAME_ALREADY_USED",
	13: "NEW_CONVERSION_TYPE_NAME_REQUIRED",
	14: "CONVERSION_TYPE_NAME_ALREADY_USED",
	15: "OWNERSHIP_REQUIRED_FOR_SET",
	16: "USER_LIST_MUTATE_NOT_SUPPORTED",
	17: "INVALID_RULE",
	27: "INVALID_DATE_RANGE",
	28: "CAN_NOT_MUTATE_SENSITIVE_USERLIST",
	29: "MAX_NUM_RULEBASED_USERLISTS",
	30: "CANNOT_MODIFY_BILLABLE_RECORD_COUNT",
	31: "APP_ID_NOT_SET",
	32: "USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST",
	33: "ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA",
	34: "RULE_TYPE_IS_NOT_SUPPORTED",
	35: "CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND",
	36: "CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS",
}

var UserListErrorEnum_UserListError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED":    2,
	"CONCRETE_TYPE_REQUIRED":                                 3,
	"CONVERSION_TYPE_ID_REQUIRED":                            4,
	"DUPLICATE_CONVERSION_TYPES":                             5,
	"INVALID_CONVERSION_TYPE":                                6,
	"INVALID_DESCRIPTION":                                    7,
	"INVALID_NAME":                                           8,
	"INVALID_TYPE":                                           9,
	"CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND":       10,
	"INVALID_USER_LIST_LOGICAL_RULE_OPERAND":                 11,
	"NAME_ALREADY_USED":                                      12,
	"NEW_CONVERSION_TYPE_NAME_REQUIRED":                      13,
	"CONVERSION_TYPE_NAME_ALREADY_USED":                      14,
	"OWNERSHIP_REQUIRED_FOR_SET":                             15,
	"USER_LIST_MUTATE_NOT_SUPPORTED":                         16,
	"INVALID_RULE":                                           17,
	"INVALID_DATE_RANGE":                                     27,
	"CAN_NOT_MUTATE_SENSITIVE_USERLIST":                      28,
	"MAX_NUM_RULEBASED_USERLISTS":                            29,
	"CANNOT_MODIFY_BILLABLE_RECORD_COUNT":                    30,
	"APP_ID_NOT_SET":                                         31,
	"USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST":              32,
	"ADVERTISER_NOT_WHITELISTED_FOR_USING_UPLOADED_DATA":     33,
	"RULE_TYPE_IS_NOT_SUPPORTED":                             34,
	"CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND": 35,
	"CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS": 36,
}

func (x UserListErrorEnum_UserListError) String() string {
	return proto.EnumName(UserListErrorEnum_UserListError_name, int32(x))
}

func (UserListErrorEnum_UserListError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94d23c47a73d2df0, []int{0, 0}
}

// Container for enum describing possible user list errors.
type UserListErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListErrorEnum) Reset()         { *m = UserListErrorEnum{} }
func (m *UserListErrorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListErrorEnum) ProtoMessage()    {}
func (*UserListErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d23c47a73d2df0, []int{0}
}

func (m *UserListErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListErrorEnum.Unmarshal(m, b)
}
func (m *UserListErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListErrorEnum.Marshal(b, m, deterministic)
}
func (m *UserListErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListErrorEnum.Merge(m, src)
}
func (m *UserListErrorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListErrorEnum.Size(m)
}
func (m *UserListErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.UserListErrorEnum_UserListError", UserListErrorEnum_UserListError_name, UserListErrorEnum_UserListError_value)
	proto.RegisterType((*UserListErrorEnum)(nil), "google.ads.googleads.v3.errors.UserListErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/user_list_error.proto", fileDescriptor_94d23c47a73d2df0)
}

var fileDescriptor_94d23c47a73d2df0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xed, 0x6e, 0x23, 0x35,
	0x14, 0xa5, 0x59, 0xd8, 0x80, 0xbb, 0x1f, 0xae, 0x11, 0xbb, 0x52, 0xbb, 0x64, 0xd9, 0x2c, 0x1f,
	0x02, 0x89, 0x09, 0x22, 0x55, 0x91, 0x86, 0x5f, 0xce, 0xf8, 0x36, 0xb1, 0x3a, 0xe3, 0x19, 0x6c,
	0xcf, 0xa4, 0x41, 0x91, 0xac, 0x40, 0xa2, 0x28, 0x52, 0x9b, 0xa9, 0x32, 0x69, 0x1f, 0x88, 0x9f,
	0x3c, 0x0a, 0x0f, 0xc0, 0x43, 0x20, 0x78, 0x07, 0x64, 0x4f, 0x66, 0x9a, 0x44, 0x94, 0xfd, 0x15,
	0xc7, 0xf7, 0x9c, 0xe3, 0x7b, 0xcf, 0xbd, 0x73, 0xd1, 0xe9, 0x3c, 0xcf, 0xe7, 0x57, 0xb3, 0xce,
	0x64, 0x5a, 0x74, 0xca, 0xa3, 0x3d, 0xdd, 0x75, 0x3b, 0xb3, 0xd5, 0x2a, 0x5f, 0x15, 0x9d, 0xdb,
	0x62, 0xb6, 0x32, 0x57, 0x8b, 0x62, 0x6d, 0xdc, 0x85, 0x77, 0xb3, 0xca, 0xd7, 0x39, 0x69, 0x95,
	0x50, 0x6f, 0x32, 0x2d, 0xbc, 0x9a, 0xe5, 0xdd, 0x75, 0xbd, 0x92, 0x75, 0xfc, 0xaa, 0x52, 0xbd,
	0x59, 0x74, 0x26, 0xcb, 0x65, 0xbe, 0x9e, 0xac, 0x17, 0xf9, 0xb2, 0x28, 0xd9, 0xed, 0xbf, 0x9b,
	0xe8, 0x28, 0x2d, 0x66, 0xab, 0x70, 0x51, 0xac, 0xc1, 0x12, 0x60, 0x79, 0x7b, 0xdd, 0xfe, 0xb3,
	0x89, 0x9e, 0xee, 0xdc, 0x92, 0xe7, 0xe8, 0x30, 0x15, 0x2a, 0x81, 0x80, 0x9f, 0x73, 0x60, 0xf8,
	0x3d, 0x72, 0x88, 0x9a, 0xa9, 0xb8, 0x10, 0xf1, 0x50, 0xe0, 0x03, 0xf2, 0x03, 0xea, 0xc2, 0xa5,
	0x06, 0x29, 0x68, 0x68, 0x24, 0x44, 0x54, 0x5e, 0x80, 0xe6, 0xa2, 0x6f, 0x52, 0x05, 0xd2, 0x84,
	0x5c, 0x69, 0x13, 0xa5, 0x9a, 0x6a, 0x30, 0x22, 0xd6, 0x46, 0xa5, 0x49, 0x12, 0x4b, 0x0d, 0x0c,
	0x37, 0xc8, 0x31, 0x7a, 0x11, 0xc4, 0x22, 0x90, 0xa0, 0xc1, 0xe8, 0x51, 0x02, 0x46, 0xc2, 0x4f,
	0x29, 0x97, 0xc0, 0xf0, 0x23, 0xf2, 0x1a, 0x9d, 0x04, 0xb1, 0xc8, 0x40, 0x2a, 0x1e, 0x8b, 0x32,
	0xca, 0xd9, 0x3d, 0xe0, 0x7d, 0xd2, 0x42, 0xc7, 0x2c, 0x4d, 0x42, 0x1e, 0x58, 0xe5, 0x3d, 0xa8,
	0xc2, 0x1f, 0x90, 0x13, 0xf4, 0x92, 0x8b, 0x8c, 0x86, 0x9c, 0xed, 0x47, 0xf1, 0x63, 0xf2, 0x12,
	0x7d, 0x5c, 0x05, 0x19, 0xa8, 0x40, 0xf2, 0x44, 0xf3, 0x58, 0xe0, 0x26, 0xc1, 0xe8, 0x49, 0x15,
	0x10, 0x34, 0x02, 0xfc, 0xe1, 0xf6, 0x8d, 0x23, 0x7f, 0x44, 0x4e, 0xd1, 0x77, 0x01, 0x15, 0xae,
	0x1a, 0xca, 0x98, 0x09, 0xe3, 0x3e, 0x0f, 0x68, 0x58, 0x56, 0x4a, 0xd5, 0xee, 0xff, 0x38, 0x01,
	0x49, 0x05, 0xc3, 0x88, 0x7c, 0x83, 0xbe, 0xac, 0x74, 0xee, 0x8d, 0xa9, 0xb0, 0x32, 0x0d, 0xa1,
	0xc6, 0x1e, 0x92, 0x4f, 0xd0, 0x91, 0x7d, 0xdd, 0xd0, 0x50, 0x02, 0x65, 0x23, 0x4b, 0x60, 0xf8,
	0x09, 0xf9, 0x02, 0xbd, 0x11, 0x30, 0xdc, 0x2f, 0xc7, 0x25, 0x7a, 0xef, 0xcc, 0x53, 0x0b, 0xfb,
	0x4f, 0xc8, 0x8e, 0xda, 0x33, 0x6b, 0x60, 0x3c, 0x14, 0x20, 0xd5, 0x80, 0x27, 0x35, 0xdd, 0x9c,
	0xc7, 0xd2, 0x28, 0xd0, 0xf8, 0x39, 0x69, 0xa3, 0xd6, 0x3b, 0x3a, 0x88, 0xb7, 0xcd, 0xb1, 0x25,
	0xe0, 0x23, 0xf2, 0x02, 0x91, 0xda, 0x59, 0xcb, 0x90, 0x54, 0xf4, 0x01, 0x9f, 0xb8, 0xa4, 0x36,
	0xa6, 0x6d, 0xb4, 0x14, 0x08, 0xc5, 0x35, 0xcf, 0xc0, 0xf9, 0x61, 0x5f, 0xc1, 0xaf, 0x6c, 0xdb,
	0x23, 0x7a, 0x69, 0x44, 0x1a, 0x39, 0xc1, 0x1e, 0x55, 0xc0, 0xea, 0xb8, 0xc2, 0x9f, 0x92, 0xaf,
	0xd0, 0xdb, 0x80, 0x0a, 0x27, 0x13, 0x33, 0x7e, 0x3e, 0x32, 0x3d, 0x1e, 0x86, 0xb4, 0x17, 0x5a,
	0x07, 0x82, 0x58, 0xda, 0x66, 0xa7, 0x42, 0xe3, 0x16, 0x21, 0xe8, 0x19, 0x4d, 0x12, 0x3b, 0x34,
	0x2e, 0x69, 0xd0, 0xf8, 0x35, 0xf9, 0x16, 0x7d, 0x5d, 0x69, 0x95, 0x96, 0x70, 0x65, 0x24, 0x28,
	0x90, 0x59, 0x55, 0xf9, 0x48, 0x69, 0x88, 0x5c, 0xc9, 0xf8, 0x33, 0x72, 0x86, 0xbe, 0xa7, 0x2c,
	0x03, 0xa9, 0xb9, 0xf5, 0xc1, 0xca, 0x0c, 0x07, 0x5c, 0x83, 0x8d, 0x6e, 0xf0, 0xa9, 0x72, 0x83,
	0x9e, 0x84, 0x31, 0x65, 0xe0, 0x2a, 0xa6, 0xf8, 0x8d, 0x75, 0xd6, 0x35, 0xb4, 0x9c, 0x5a, 0xb5,
	0xe7, 0x5a, 0x9b, 0xf8, 0xe8, 0x6c, 0x7b, 0x80, 0xa8, 0x51, 0x3c, 0xe2, 0x21, 0x95, 0x75, 0xa1,
	0x0f, 0x8e, 0xd1, 0xdb, 0x6d, 0x6e, 0xc4, 0x2f, 0x4d, 0x20, 0x23, 0x53, 0x9a, 0xc4, 0xc5, 0x2e,
	0x65, 0xc8, 0xf5, 0xc0, 0xc4, 0x7a, 0xb0, 0xe9, 0xa0, 0xc2, 0x9f, 0xf7, 0xfe, 0x39, 0x40, 0xed,
	0x5f, 0xf3, 0x6b, 0xef, 0xff, 0x77, 0x46, 0x8f, 0xec, 0x7c, 0xfc, 0x89, 0xdd, 0x14, 0xc9, 0xc1,
	0xcf, 0x6c, 0xc3, 0x9a, 0xe7, 0x57, 0x93, 0xe5, 0xdc, 0xcb, 0x57, 0xf3, 0xce, 0x7c, 0xb6, 0x74,
	0x7b, 0xa4, 0xda, 0x57, 0x37, 0x8b, 0xe2, 0xa1, 0xf5, 0xf5, 0x63, 0xf9, 0xf3, 0x5b, 0xe3, 0x51,
	0x9f, 0xd2, 0xdf, 0x1b, 0xad, 0x7e, 0x29, 0x46, 0xa7, 0x85, 0x57, 0x1e, 0xed, 0x29, 0xeb, 0x7a,
	0xee, 0xc9, 0xe2, 0x8f, 0x0a, 0x30, 0xa6, 0xd3, 0x62, 0x5c, 0x03, 0xc6, 0x59, 0x77, 0x5c, 0x02,
	0xfe, 0x6a, 0xb4, 0xcb, 0x5b, 0xdf, 0xa7, 0xd3, 0xc2, 0xf7, 0x6b, 0x88, 0xef, 0x67, 0x5d, 0xdf,
	0x2f, 0x41, 0xbf, 0x3c, 0x76, 0xd9, 0x75, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x4c, 0x0e,
	0x3e, 0x5b, 0x05, 0x00, 0x00,
}