// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamestate.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Vector3 struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float64  `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3) Reset()         { *m = Vector3{} }
func (m *Vector3) String() string { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()    {}
func (*Vector3) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{0}
}

func (m *Vector3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3.Unmarshal(m, b)
}
func (m *Vector3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3.Marshal(b, m, deterministic)
}
func (m *Vector3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3.Merge(m, src)
}
func (m *Vector3) XXX_Size() int {
	return xxx_messageInfo_Vector3.Size(m)
}
func (m *Vector3) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3 proto.InternalMessageInfo

func (m *Vector3) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

type SyncPos struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos                  *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPos) Reset()         { *m = SyncPos{} }
func (m *SyncPos) String() string { return proto.CompactTextString(m) }
func (*SyncPos) ProtoMessage()    {}
func (*SyncPos) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{1}
}

func (m *SyncPos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPos.Unmarshal(m, b)
}
func (m *SyncPos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPos.Marshal(b, m, deterministic)
}
func (m *SyncPos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPos.Merge(m, src)
}
func (m *SyncPos) XXX_Size() int {
	return xxx_messageInfo_SyncPos.Size(m)
}
func (m *SyncPos) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPos.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPos proto.InternalMessageInfo

func (m *SyncPos) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SyncPos) GetPos() *Vector3 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SyncPos) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SyncPlayerRequest struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos                  *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPlayerRequest) Reset()         { *m = SyncPlayerRequest{} }
func (m *SyncPlayerRequest) String() string { return proto.CompactTextString(m) }
func (*SyncPlayerRequest) ProtoMessage()    {}
func (*SyncPlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{2}
}

func (m *SyncPlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayerRequest.Unmarshal(m, b)
}
func (m *SyncPlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayerRequest.Marshal(b, m, deterministic)
}
func (m *SyncPlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayerRequest.Merge(m, src)
}
func (m *SyncPlayerRequest) XXX_Size() int {
	return xxx_messageInfo_SyncPlayerRequest.Size(m)
}
func (m *SyncPlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayerRequest proto.InternalMessageInfo

func (m *SyncPlayerRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SyncPlayerRequest) GetPos() *Vector3 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SyncPlayerRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SyncPlayerResponse struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPlayerResponse) Reset()         { *m = SyncPlayerResponse{} }
func (m *SyncPlayerResponse) String() string { return proto.CompactTextString(m) }
func (*SyncPlayerResponse) ProtoMessage()    {}
func (*SyncPlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{3}
}

func (m *SyncPlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayerResponse.Unmarshal(m, b)
}
func (m *SyncPlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayerResponse.Marshal(b, m, deterministic)
}
func (m *SyncPlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayerResponse.Merge(m, src)
}
func (m *SyncPlayerResponse) XXX_Size() int {
	return xxx_messageInfo_SyncPlayerResponse.Size(m)
}
func (m *SyncPlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayerResponse proto.InternalMessageInfo

func (m *SyncPlayerResponse) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type RcvPlayerRequest struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcvPlayerRequest) Reset()         { *m = RcvPlayerRequest{} }
func (m *RcvPlayerRequest) String() string { return proto.CompactTextString(m) }
func (*RcvPlayerRequest) ProtoMessage()    {}
func (*RcvPlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{4}
}

func (m *RcvPlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcvPlayerRequest.Unmarshal(m, b)
}
func (m *RcvPlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcvPlayerRequest.Marshal(b, m, deterministic)
}
func (m *RcvPlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcvPlayerRequest.Merge(m, src)
}
func (m *RcvPlayerRequest) XXX_Size() int {
	return xxx_messageInfo_RcvPlayerRequest.Size(m)
}
func (m *RcvPlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RcvPlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RcvPlayerRequest proto.InternalMessageInfo

func (m *RcvPlayerRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type RcvPlayerResponse struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos                  *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcvPlayerResponse) Reset()         { *m = RcvPlayerResponse{} }
func (m *RcvPlayerResponse) String() string { return proto.CompactTextString(m) }
func (*RcvPlayerResponse) ProtoMessage()    {}
func (*RcvPlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4287d9bb1ae0ce0d, []int{5}
}

func (m *RcvPlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcvPlayerResponse.Unmarshal(m, b)
}
func (m *RcvPlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcvPlayerResponse.Marshal(b, m, deterministic)
}
func (m *RcvPlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcvPlayerResponse.Merge(m, src)
}
func (m *RcvPlayerResponse) XXX_Size() int {
	return xxx_messageInfo_RcvPlayerResponse.Size(m)
}
func (m *RcvPlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RcvPlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RcvPlayerResponse proto.InternalMessageInfo

func (m *RcvPlayerResponse) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *RcvPlayerResponse) GetPos() *Vector3 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func init() {
	proto.RegisterType((*Vector3)(nil), "pb.Vector3")
	proto.RegisterType((*SyncPos)(nil), "pb.SyncPos")
	proto.RegisterType((*SyncPlayerRequest)(nil), "pb.SyncPlayerRequest")
	proto.RegisterType((*SyncPlayerResponse)(nil), "pb.SyncPlayerResponse")
	proto.RegisterType((*RcvPlayerRequest)(nil), "pb.RcvPlayerRequest")
	proto.RegisterType((*RcvPlayerResponse)(nil), "pb.RcvPlayerResponse")
}

func init() {
	proto.RegisterFile("gamestate.proto", fileDescriptor_4287d9bb1ae0ce0d)
}

var fileDescriptor_4287d9bb1ae0ce0d = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x4b, 0xc4, 0x30,
	0x18, 0x36, 0x29, 0x78, 0xf4, 0x3d, 0xbf, 0xee, 0x45, 0xa5, 0x1c, 0x0a, 0x12, 0x1c, 0x9c, 0x3a,
	0x5c, 0x37, 0xc1, 0x45, 0x6e, 0xb9, 0x4d, 0x72, 0x72, 0xb3, 0x6d, 0x09, 0x22, 0xd8, 0x26, 0x36,
	0x51, 0x9a, 0xfe, 0x00, 0x7f, 0xb7, 0x24, 0x55, 0x5b, 0xa2, 0xe8, 0x72, 0xe3, 0xf3, 0xf0, 0x7c,
	0xe5, 0x03, 0x0e, 0x1f, 0xf3, 0x4a, 0x68, 0x93, 0x1b, 0x91, 0xaa, 0x46, 0x1a, 0x89, 0x54, 0x15,
	0x2c, 0x83, 0xc9, 0x46, 0x94, 0x46, 0x36, 0x19, 0xee, 0x01, 0x69, 0x13, 0x72, 0x41, 0xae, 0x08,
	0x27, 0xad, 0x43, 0x36, 0xa1, 0x3d, 0xb2, 0x0e, 0x75, 0x49, 0xd4, 0xa3, 0x8e, 0x6d, 0x60, 0xb2,
	0xb6, 0x75, 0x79, 0x27, 0x35, 0x1e, 0x00, 0x5d, 0x2d, 0xbd, 0x6b, 0x9f, 0xd3, 0xd5, 0x12, 0xcf,
	0x21, 0x52, 0x52, 0x7b, 0xe3, 0x74, 0x31, 0x4d, 0x55, 0x91, 0x7e, 0xc6, 0x73, 0xc7, 0xe3, 0x19,
	0xc4, 0xf7, 0x4f, 0x7e, 0x45, 0xa5, 0x7c, 0x5e, 0xc4, 0x07, 0x82, 0x3d, 0xc0, 0xcc, 0xe7, 0x3e,
	0xe7, 0x56, 0x34, 0x5c, 0xbc, 0xbc, 0x0a, 0x6d, 0xb6, 0xdb, 0x70, 0x09, 0x38, 0x6e, 0xd0, 0x4a,
	0xd6, 0x5a, 0x84, 0x15, 0x8c, 0xc1, 0x11, 0x2f, 0xdf, 0xfe, 0x9c, 0xc1, 0x6e, 0x61, 0x36, 0xd2,
	0xfc, 0x1e, 0xf4, 0xcf, 0xd6, 0xc5, 0x3b, 0x81, 0xd8, 0xcd, 0x59, 0xbb, 0x47, 0xc1, 0x6b, 0x88,
	0xbf, 0x13, 0xf1, 0xd8, 0x89, 0xc3, 0x11, 0xf3, 0x93, 0x80, 0xed, 0x6b, 0xd9, 0x0e, 0xde, 0x00,
	0x0c, 0xe7, 0x42, 0x2f, 0xfb, 0x71, 0x93, 0xf3, 0xd3, 0x90, 0xfe, 0xb2, 0x17, 0xbb, 0xfe, 0x43,
	0x64, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xc3, 0xab, 0x12, 0x23, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncStateClient is the client API for SyncState service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncStateClient interface {
	RcvPlayer(ctx context.Context, in *RcvPlayerRequest, opts ...grpc.CallOption) (*RcvPlayerResponse, error)
	SyncPlayer(ctx context.Context, in *SyncPlayerRequest, opts ...grpc.CallOption) (*SyncPlayerResponse, error)
}

type syncStateClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncStateClient(cc grpc.ClientConnInterface) SyncStateClient {
	return &syncStateClient{cc}
}

func (c *syncStateClient) RcvPlayer(ctx context.Context, in *RcvPlayerRequest, opts ...grpc.CallOption) (*RcvPlayerResponse, error) {
	out := new(RcvPlayerResponse)
	err := c.cc.Invoke(ctx, "/pb.SyncState/RcvPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncStateClient) SyncPlayer(ctx context.Context, in *SyncPlayerRequest, opts ...grpc.CallOption) (*SyncPlayerResponse, error) {
	out := new(SyncPlayerResponse)
	err := c.cc.Invoke(ctx, "/pb.SyncState/SyncPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncStateServer is the server API for SyncState service.
type SyncStateServer interface {
	RcvPlayer(context.Context, *RcvPlayerRequest) (*RcvPlayerResponse, error)
	SyncPlayer(context.Context, *SyncPlayerRequest) (*SyncPlayerResponse, error)
}

// UnimplementedSyncStateServer can be embedded to have forward compatible implementations.
type UnimplementedSyncStateServer struct {
}

func (*UnimplementedSyncStateServer) RcvPlayer(ctx context.Context, req *RcvPlayerRequest) (*RcvPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RcvPlayer not implemented")
}
func (*UnimplementedSyncStateServer) SyncPlayer(ctx context.Context, req *SyncPlayerRequest) (*SyncPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPlayer not implemented")
}

func RegisterSyncStateServer(s *grpc.Server, srv SyncStateServer) {
	s.RegisterService(&_SyncState_serviceDesc, srv)
}

func _SyncState_RcvPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RcvPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncStateServer).RcvPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SyncState/RcvPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncStateServer).RcvPlayer(ctx, req.(*RcvPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncState_SyncPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncStateServer).SyncPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SyncState/SyncPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncStateServer).SyncPlayer(ctx, req.(*SyncPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncState_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SyncState",
	HandlerType: (*SyncStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RcvPlayer",
			Handler:    _SyncState_RcvPlayer_Handler,
		},
		{
			MethodName: "SyncPlayer",
			Handler:    _SyncState_SyncPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamestate.proto",
}
