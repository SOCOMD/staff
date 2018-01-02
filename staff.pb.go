// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staff.proto

/*
Package staff is a generated protocol buffer package.

It is generated from these files:
	staff.proto

It has these top-level messages:
	GetUserMessage
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

type GetUserMessage struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserMessage) Reset()                    { *m = GetUserMessage{} }
func (m *GetUserMessage) String() string            { return proto.CompactTextString(m) }
func (*GetUserMessage) ProtoMessage()               {}
func (*GetUserMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetUserMessage) GetId() string {
	if m != nil {
		return m.Id
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
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func init() {
	proto.RegisterType((*GetUserMessage)(nil), "GetUserMessage")
	proto.RegisterType((*User)(nil), "User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Members service

type MembersClient interface {
	GetUser(ctx context.Context, in *GetUserMessage, opts ...grpc.CallOption) (*User, error)
}

type membersClient struct {
	cc *grpc.ClientConn
}

func NewMembersClient(cc *grpc.ClientConn) MembersClient {
	return &membersClient{cc}
}

func (c *membersClient) GetUser(ctx context.Context, in *GetUserMessage, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/members/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Members service

type MembersServer interface {
	GetUser(context.Context, *GetUserMessage) (*User, error)
}

func RegisterMembersServer(s *grpc.Server, srv MembersServer) {
	s.RegisterService(&_Members_serviceDesc, srv)
}

func _Members_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/members/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServer).GetUser(ctx, req.(*GetUserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Members_serviceDesc = grpc.ServiceDesc{
	ServiceName: "members",
	HandlerType: (*MembersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Members_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}

func init() { proto.RegisterFile("staff.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x49, 0xda, 0xfc, 0x4d, 0xa5, 0x95, 0x45, 0x64, 0x28, 0x1e, 0x62, 0x4f, 0x39, 0x48,
	0x0e, 0xfa, 0x10, 0x9e, 0xf4, 0x10, 0xf0, 0x01, 0x36, 0x99, 0x69, 0x59, 0x69, 0x12, 0xc9, 0x4e,
	0x7d, 0x2e, 0x1f, 0x51, 0x76, 0x37, 0x55, 0xda, 0x5b, 0xbe, 0xef, 0x1b, 0xc8, 0x30, 0x0b, 0x2b,
	0x2b, 0x7a, 0xbf, 0xaf, 0xbf, 0xa6, 0x51, 0xc6, 0x5d, 0x09, 0xeb, 0x57, 0x96, 0x0f, 0xcb, 0xd3,
	0x1b, 0x5b, 0xab, 0x0f, 0xac, 0xd6, 0x10, 0x1b, 0xc2, 0xa8, 0x8c, 0xaa, 0xa2, 0x89, 0x0d, 0xed,
	0x7e, 0x62, 0x58, 0xba, 0x7e, 0x1d, 0xd4, 0x3d, 0xa4, 0x62, 0xdf, 0x75, 0xcf, 0x18, 0x7b, 0x37,
	0x53, 0xf0, 0xd4, 0x1a, 0xc2, 0xc5, 0xd9, 0x3b, 0x0a, 0xfe, 0x74, 0x32, 0x84, 0xcb, 0xb3, 0x77,
	0xa4, 0x1e, 0xa0, 0x10, 0xdb, 0x4d, 0xac, 0x85, 0x09, 0x13, 0x9f, 0xfe, 0x85, 0xaa, 0x60, 0x23,
	0xf6, 0xa8, 0xad, 0x74, 0xe3, 0x30, 0x70, 0xe7, 0x66, 0x52, 0x3f, 0x73, 0xad, 0xd5, 0x1d, 0x24,
	0xdc, 0x6b, 0x73, 0xc4, 0xcc, 0xf7, 0x00, 0x6a, 0x0b, 0xf9, 0xe7, 0x68, 0x06, 0xd2, 0xc2, 0x98,
	0xfb, 0xf0, 0xc7, 0xea, 0x16, 0x16, 0x34, 0xb6, 0x58, 0x78, 0xed, 0x3e, 0xdd, 0x8e, 0x07, 0x1e,
	0x88, 0x27, 0x84, 0xb0, 0x63, 0x20, 0xe7, 0x75, 0x27, 0xe6, 0x9b, 0x71, 0x55, 0x46, 0x55, 0xde,
	0xcc, 0xe4, 0xfe, 0xa9, 0xa9, 0x37, 0x03, 0xde, 0x94, 0x51, 0x95, 0x34, 0x01, 0x9e, 0x9f, 0x20,
	0xeb, 0xb9, 0x6f, 0x79, 0xb2, 0xea, 0x11, 0xb2, 0xf9, 0xbe, 0x6a, 0x53, 0x5f, 0x5e, 0x7a, 0x9b,
	0xd4, 0x8e, 0xda, 0xd4, 0xbf, 0xc4, 0xcb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x98, 0x63,
	0x29, 0x98, 0x01, 0x00, 0x00,
}
