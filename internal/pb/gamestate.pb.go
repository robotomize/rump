// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: gamestate.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Vector3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{0}
}

func (x *Vector3) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type SyncPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *SyncPositionRequest) Reset() {
	*x = SyncPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPositionRequest) ProtoMessage() {}

func (x *SyncPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPositionRequest.ProtoReflect.Descriptor instead.
func (*SyncPositionRequest) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{1}
}

func (x *SyncPositionRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *SyncPositionRequest) GetPos() *Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

type SyncPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *SyncPositionResponse) Reset() {
	*x = SyncPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPositionResponse) ProtoMessage() {}

func (x *SyncPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPositionResponse.ProtoReflect.Descriptor instead.
func (*SyncPositionResponse) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{2}
}

func (x *SyncPositionResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *SyncPositionResponse) GetPos() *Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

type RcvPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RcvPositionRequest) Reset() {
	*x = RcvPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RcvPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RcvPositionRequest) ProtoMessage() {}

func (x *RcvPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RcvPositionRequest.ProtoReflect.Descriptor instead.
func (*RcvPositionRequest) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{3}
}

func (x *RcvPositionRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type RcvPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pos *Vector3 `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *RcvPositionResponse) Reset() {
	*x = RcvPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RcvPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RcvPositionResponse) ProtoMessage() {}

func (x *RcvPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RcvPositionResponse.ProtoReflect.Descriptor instead.
func (*RcvPositionResponse) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{4}
}

func (x *RcvPositionResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RcvPositionResponse) GetPos() *Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_gamestate_proto protoreflect.FileDescriptor

var file_gamestate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x33, 0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01,
	0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x22, 0x44, 0x0a, 0x13, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x70, 0x6f, 0x73,
	0x22, 0x45, 0x0a, 0x14, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x63, 0x76, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x44, 0x0a,
	0x13, 0x52, 0x63, 0x76, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03,
	0x70, 0x6f, 0x73, 0x32, 0x92, 0x01, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x52, 0x63, 0x76, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x63, 0x76, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x63,
	0x76, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gamestate_proto_rawDescOnce sync.Once
	file_gamestate_proto_rawDescData = file_gamestate_proto_rawDesc
)

func file_gamestate_proto_rawDescGZIP() []byte {
	file_gamestate_proto_rawDescOnce.Do(func() {
		file_gamestate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gamestate_proto_rawDescData)
	})
	return file_gamestate_proto_rawDescData
}

var file_gamestate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gamestate_proto_goTypes = []interface{}{
	(*Vector3)(nil),              // 0: pb.Vector3
	(*SyncPositionRequest)(nil),  // 1: pb.SyncPositionRequest
	(*SyncPositionResponse)(nil), // 2: pb.SyncPositionResponse
	(*RcvPositionRequest)(nil),   // 3: pb.RcvPositionRequest
	(*RcvPositionResponse)(nil),  // 4: pb.RcvPositionResponse
}
var file_gamestate_proto_depIdxs = []int32{
	0, // 0: pb.SyncPositionRequest.pos:type_name -> pb.Vector3
	0, // 1: pb.SyncPositionResponse.pos:type_name -> pb.Vector3
	0, // 2: pb.RcvPositionResponse.pos:type_name -> pb.Vector3
	3, // 3: pb.SyncState.RcvPosition:input_type -> pb.RcvPositionRequest
	1, // 4: pb.SyncState.SyncPosition:input_type -> pb.SyncPositionRequest
	4, // 5: pb.SyncState.RcvPosition:output_type -> pb.RcvPositionResponse
	2, // 6: pb.SyncState.SyncPosition:output_type -> pb.SyncPositionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gamestate_proto_init() }
func file_gamestate_proto_init() {
	if File_gamestate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gamestate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPositionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPositionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RcvPositionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RcvPositionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gamestate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gamestate_proto_goTypes,
		DependencyIndexes: file_gamestate_proto_depIdxs,
		MessageInfos:      file_gamestate_proto_msgTypes,
	}.Build()
	File_gamestate_proto = out.File
	file_gamestate_proto_rawDesc = nil
	file_gamestate_proto_goTypes = nil
	file_gamestate_proto_depIdxs = nil
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
	RcvPosition(ctx context.Context, in *RcvPositionRequest, opts ...grpc.CallOption) (*RcvPositionResponse, error)
	SyncPosition(ctx context.Context, in *SyncPositionRequest, opts ...grpc.CallOption) (*SyncPositionResponse, error)
}

type syncStateClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncStateClient(cc grpc.ClientConnInterface) SyncStateClient {
	return &syncStateClient{cc}
}

func (c *syncStateClient) RcvPosition(ctx context.Context, in *RcvPositionRequest, opts ...grpc.CallOption) (*RcvPositionResponse, error) {
	out := new(RcvPositionResponse)
	err := c.cc.Invoke(ctx, "/pb.SyncState/RcvPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncStateClient) SyncPosition(ctx context.Context, in *SyncPositionRequest, opts ...grpc.CallOption) (*SyncPositionResponse, error) {
	out := new(SyncPositionResponse)
	err := c.cc.Invoke(ctx, "/pb.SyncState/SyncPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncStateServer is the server API for SyncState service.
type SyncStateServer interface {
	RcvPosition(context.Context, *RcvPositionRequest) (*RcvPositionResponse, error)
	SyncPosition(context.Context, *SyncPositionRequest) (*SyncPositionResponse, error)
}

// UnimplementedSyncStateServer can be embedded to have forward compatible implementations.
type UnimplementedSyncStateServer struct {
}

func (*UnimplementedSyncStateServer) RcvPosition(context.Context, *RcvPositionRequest) (*RcvPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RcvPosition not implemented")
}
func (*UnimplementedSyncStateServer) SyncPosition(context.Context, *SyncPositionRequest) (*SyncPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPosition not implemented")
}

func RegisterSyncStateServer(s *grpc.Server, srv SyncStateServer) {
	s.RegisterService(&_SyncState_serviceDesc, srv)
}

func _SyncState_RcvPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RcvPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncStateServer).RcvPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SyncState/RcvPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncStateServer).RcvPosition(ctx, req.(*RcvPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncState_SyncPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncStateServer).SyncPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SyncState/SyncPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncStateServer).SyncPosition(ctx, req.(*SyncPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncState_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SyncState",
	HandlerType: (*SyncStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RcvPosition",
			Handler:    _SyncState_RcvPosition_Handler,
		},
		{
			MethodName: "SyncPosition",
			Handler:    _SyncState_SyncPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamestate.proto",
}
