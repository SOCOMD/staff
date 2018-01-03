// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staff.proto

/*
Package staff is a generated protocol buffer package.

It is generated from these files:
	staff.proto

It has these top-level messages:
	GetAuthStatusRequest
	GetAuthStatusResult
	GetUserRequest
	User
*/
package staff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetUserRequestSearchType int32

const (
	GetUserRequest_ID      GetUserRequestSearchType = 0
	GetUserRequest_TSDBID  GetUserRequestSearchType = 1
	GetUserRequest_TSUUID  GetUserRequestSearchType = 2
	GetUserRequest_EMAIL   GetUserRequestSearchType = 3
	GetUserRequest_STEAMID GetUserRequestSearchType = 4
	GetUserRequest_TOKEN   GetUserRequestSearchType = 5
)

var GetUserRequestSearchType_name = map[int32]string{
	0: "ID",
	1: "TSDBID",
	2: "TSUUID",
	3: "EMAIL",
	4: "STEAMID",
	5: "TOKEN",
}
var GetUserRequestSearchType_value = map[string]int32{
	"ID":      0,
	"TSDBID":  1,
	"TSUUID":  2,
	"EMAIL":   3,
	"STEAMID": 4,
	"TOKEN":   5,
}

func (x GetUserRequestSearchType) String() string {
	return proto.EnumName(GetUserRequestSearchType_name, int32(x))
}
func (GetUserRequestSearchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type GetAuthStatusRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetAuthStatusRequest) Reset()                    { *m = GetAuthStatusRequest{} }
func (m *GetAuthStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAuthStatusRequest) ProtoMessage()               {}
func (*GetAuthStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetAuthStatusRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetAuthStatusResult struct {
}

func (m *GetAuthStatusResult) Reset()                    { *m = GetAuthStatusResult{} }
func (m *GetAuthStatusResult) String() string            { return proto.CompactTextString(m) }
func (*GetAuthStatusResult) ProtoMessage()               {}
func (*GetAuthStatusResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetUserRequest struct {
	Search string                   `protobuf:"bytes,1,opt,name=search" json:"search,omitempty"`
	Type   GetUserRequestSearchType `protobuf:"varint,2,opt,name=type,enum=staff.GetUserRequestSearchType" json:"type,omitempty"`
	Token  string                   `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *GetUserRequest) GetType() GetUserRequestSearchType {
	if m != nil {
		return m.Type
	}
	return GetUserRequest_ID
}

func (m *GetUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type User struct {
	Id              string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TsName          string `protobuf:"bytes,2,opt,name=tsName" json:"tsName,omitempty"`
	Tsdbid          string `protobuf:"bytes,3,opt,name=tsdbid" json:"tsdbid,omitempty"`
	Tsuuid          string `protobuf:"bytes,4,opt,name=tsuuid" json:"tsuuid,omitempty"`
	Tscreated       string `protobuf:"bytes,5,opt,name=tscreated" json:"tscreated,omitempty"`
	Tslastconnected string `protobuf:"bytes,6,opt,name=tslastconnected" json:"tslastconnected,omitempty"`
	Email           string `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Joindate        string `protobuf:"bytes,8,opt,name=joindate" json:"joindate,omitempty"`
	Dob             string `protobuf:"bytes,9,opt,name=dob" json:"dob,omitempty"`
	Gender          string `protobuf:"bytes,10,opt,name=gender" json:"gender,omitempty"`
	Active          bool   `protobuf:"varint,11,opt,name=active" json:"active,omitempty"`
	Admin           int32  `protobuf:"varint,12,opt,name=admin" json:"admin,omitempty"`
	Steamid         string `protobuf:"bytes,13,opt,name=steamid" json:"steamid,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetTsName() string {
	if m != nil {
		return m.TsName
	}
	return ""
}

func (m *User) GetTsdbid() string {
	if m != nil {
		return m.Tsdbid
	}
	return ""
}

func (m *User) GetTsuuid() string {
	if m != nil {
		return m.Tsuuid
	}
	return ""
}

func (m *User) GetTscreated() string {
	if m != nil {
		return m.Tscreated
	}
	return ""
}

func (m *User) GetTslastconnected() string {
	if m != nil {
		return m.Tslastconnected
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetJoindate() string {
	if m != nil {
		return m.Joindate
	}
	return ""
}

func (m *User) GetDob() string {
	if m != nil {
		return m.Dob
	}
	return ""
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *User) GetAdmin() int32 {
	if m != nil {
		return m.Admin
	}
	return 0
}

func (m *User) GetSteamid() string {
	if m != nil {
		return m.Steamid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAuthStatusRequest)(nil), "staff.GetAuthStatusRequest")
	proto.RegisterType((*GetAuthStatusResult)(nil), "staff.GetAuthStatusResult")
	proto.RegisterType((*GetUserRequest)(nil), "staff.GetUserRequest")
	proto.RegisterType((*User)(nil), "staff.User")
	proto.RegisterEnum("staff.GetUserRequestSearchType", GetUserRequestSearchType_name, GetUserRequestSearchType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Staff service

type StaffClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	AuthStatus(ctx context.Context, in *GetAuthStatusRequest, opts ...grpc.CallOption) (*GetAuthStatusResult, error)
}

type staffClient struct {
	cc *grpc.ClientConn
}

func NewStaffClient(cc *grpc.ClientConn) StaffClient {
	return &staffClient{cc}
}

func (c *staffClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/staff.staff/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) AuthStatus(ctx context.Context, in *GetAuthStatusRequest, opts ...grpc.CallOption) (*GetAuthStatusResult, error) {
	out := new(GetAuthStatusResult)
	err := grpc.Invoke(ctx, "/staff.staff/AuthStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Staff service

type StaffServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	AuthStatus(context.Context, *GetAuthStatusRequest) (*GetAuthStatusResult, error)
}

func RegisterStaffServer(s *grpc.Server, srv StaffServer) {
	s.RegisterService(&_Staff_serviceDesc, srv)
}

func _Staff_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.staff/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_AuthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).AuthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.staff/AuthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).AuthStatus(ctx, req.(*GetAuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Staff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "staff.staff",
	HandlerType: (*StaffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Staff_GetUser_Handler,
		},
		{
			MethodName: "AuthStatus",
			Handler:    _Staff_AuthStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}

func init() { proto.RegisterFile("staff.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xb1, 0x13, 0x3b, 0xf1, 0x0d, 0x04, 0x6b, 0x68, 0xd1, 0x28, 0xb0, 0xb0, 0xb2, 0xf2,
	0x02, 0xb2, 0x28, 0xbc, 0x40, 0x90, 0xa3, 0xca, 0x82, 0xb6, 0x92, 0x93, 0x3c, 0xc0, 0xc4, 0x73,
	0x4b, 0x0d, 0x89, 0x1d, 0x3c, 0xd7, 0x48, 0x5d, 0xf0, 0x78, 0x6c, 0x78, 0x2a, 0x34, 0x3f, 0xc1,
	0x50, 0xba, 0x9b, 0xef, 0xcc, 0xb9, 0xe3, 0xe3, 0x99, 0x03, 0x13, 0x45, 0xe2, 0xf6, 0x76, 0x71,
	0x6c, 0x1b, 0x6a, 0x58, 0x60, 0x60, 0xfe, 0x06, 0xce, 0x2e, 0x91, 0x96, 0x1d, 0xdd, 0xad, 0x49,
	0x50, 0xa7, 0x0a, 0xfc, 0xd6, 0xa1, 0x22, 0x76, 0x06, 0x01, 0x35, 0x5f, 0xb1, 0xe6, 0x5e, 0xe2,
	0xa5, 0x51, 0x61, 0x61, 0x7e, 0x0e, 0x2f, 0x1e, 0xb8, 0x55, 0xb7, 0xa7, 0xf9, 0x4f, 0x0f, 0xa6,
	0x97, 0x48, 0x5b, 0x85, 0xed, 0x69, 0xfe, 0x25, 0x84, 0x0a, 0x45, 0x5b, 0xde, 0xb9, 0x03, 0x1c,
	0xb1, 0xf7, 0x30, 0xa4, 0xfb, 0x23, 0x72, 0x3f, 0xf1, 0xd2, 0xe9, 0x45, 0xb2, 0xb0, 0x91, 0xfe,
	0x1d, 0x5e, 0x58, 0xef, 0xe6, 0xfe, 0x88, 0x85, 0x71, 0xf7, 0x69, 0x06, 0x7f, 0xa7, 0xb9, 0x01,
	0xe8, 0x9d, 0x2c, 0x04, 0x3f, 0xcf, 0xe2, 0x27, 0x0c, 0x20, 0xdc, 0xac, 0xb3, 0x0f, 0x79, 0x16,
	0x7b, 0x76, 0xbd, 0xdd, 0xe6, 0x59, 0xec, 0xb3, 0x08, 0x82, 0xd5, 0xd5, 0x32, 0xff, 0x14, 0x0f,
	0xd8, 0x04, 0x46, 0xeb, 0xcd, 0x6a, 0x79, 0x95, 0x67, 0xf1, 0x50, 0xeb, 0x9b, 0x9b, 0x8f, 0xab,
	0xeb, 0x38, 0x98, 0xff, 0xf2, 0x61, 0xa8, 0x73, 0xb0, 0x29, 0xf8, 0x95, 0x74, 0xc9, 0xfd, 0x4a,
	0xea, 0xbf, 0x21, 0x75, 0x2d, 0x0e, 0x36, 0x77, 0x54, 0x38, 0xb2, 0xba, 0xdc, 0x55, 0xd2, 0x05,
	0x73, 0x64, 0xf5, 0xae, 0xab, 0x24, 0x1f, 0x9e, 0x74, 0x4d, 0xec, 0x35, 0x44, 0xa4, 0xca, 0x16,
	0x05, 0xa1, 0xe4, 0x81, 0xd9, 0xea, 0x05, 0x96, 0xc2, 0x73, 0x52, 0x7b, 0xa1, 0xa8, 0x6c, 0xea,
	0x1a, 0x4b, 0xed, 0x09, 0x8d, 0xe7, 0xa1, 0xac, 0xef, 0x03, 0x0f, 0xa2, 0xda, 0xf3, 0x91, 0xbd,
	0x0f, 0x03, 0x6c, 0x06, 0xe3, 0x2f, 0x4d, 0x55, 0x4b, 0x41, 0xc8, 0xc7, 0x66, 0xe3, 0x0f, 0xb3,
	0x18, 0x06, 0xb2, 0xd9, 0xf1, 0xc8, 0xc8, 0x7a, 0xa9, 0x33, 0x7e, 0xc6, 0x5a, 0x62, 0xcb, 0xc1,
	0x66, 0xb4, 0xa4, 0x75, 0x51, 0x52, 0xf5, 0x1d, 0xf9, 0x24, 0xf1, 0xd2, 0x71, 0xe1, 0x48, 0x7f,
	0x53, 0xc8, 0x43, 0x55, 0xf3, 0xa7, 0x89, 0x97, 0x06, 0x85, 0x05, 0xc6, 0x61, 0xa4, 0x08, 0xc5,
	0xa1, 0x92, 0xfc, 0x99, 0x39, 0xe6, 0x84, 0x17, 0x3f, 0xc0, 0x56, 0x8c, 0xbd, 0x85, 0x91, 0x7b,
	0x5f, 0x76, 0xfe, 0xe8, 0x7b, 0xcf, 0x26, 0x4e, 0x36, 0x9e, 0x15, 0x40, 0x5f, 0x30, 0xf6, 0xaa,
	0x9f, 0xf8, 0xaf, 0xa4, 0xb3, 0xd9, 0xe3, 0x9b, 0xba, 0x93, 0xbb, 0xd0, 0xd4, 0xfc, 0xdd, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xb7, 0x77, 0x77, 0xf5, 0x02, 0x00, 0x00,
}
