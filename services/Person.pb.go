// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.0
// source: Person.proto

package services

import (
	context "context"
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

// 定义list请求
type PersonListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *PersonListReq) Reset() {
	*x = PersonListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonListReq) ProtoMessage() {}

func (x *PersonListReq) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonListReq.ProtoReflect.Descriptor instead.
func (*PersonListReq) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{0}
}

func (x *PersonListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 定义list响应
type PersonListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error int32             `protobuf:"varint,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Msg   string            `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data  []*PersonRespData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PersonListResp) Reset() {
	*x = PersonListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonListResp) ProtoMessage() {}

func (x *PersonListResp) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonListResp.ProtoReflect.Descriptor instead.
func (*PersonListResp) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonListResp) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *PersonListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PersonListResp) GetData() []*PersonRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PersonRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Login int32  `protobuf:"varint,3,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *PersonRespData) Reset() {
	*x = PersonRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRespData) ProtoMessage() {}

func (x *PersonRespData) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRespData.ProtoReflect.Descriptor instead.
func (*PersonRespData) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{2}
}

func (x *PersonRespData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PersonRespData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonRespData) GetLogin() int32 {
	if x != nil {
		return x.Login
	}
	return 0
}

type PersonInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *PersonInfoReq) Reset() {
	*x = PersonInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfoReq) ProtoMessage() {}

func (x *PersonInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfoReq.ProtoReflect.Descriptor instead.
func (*PersonInfoReq) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{3}
}

func (x *PersonInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PersonInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error int32           `protobuf:"varint,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Msg   string          `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data  *PersonRespData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PersonInfoResp) Reset() {
	*x = PersonInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfoResp) ProtoMessage() {}

func (x *PersonInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfoResp.ProtoReflect.Descriptor instead.
func (*PersonInfoResp) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{4}
}

func (x *PersonInfoResp) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *PersonInfoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PersonInfoResp) GetData() *PersonRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Person_proto protoreflect.FileDescriptor

var file_Person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a,
	0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x1f, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x9b, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Person_proto_rawDescOnce sync.Once
	file_Person_proto_rawDescData = file_Person_proto_rawDesc
)

func file_Person_proto_rawDescGZIP() []byte {
	file_Person_proto_rawDescOnce.Do(func() {
		file_Person_proto_rawDescData = protoimpl.X.CompressGZIP(file_Person_proto_rawDescData)
	})
	return file_Person_proto_rawDescData
}

var file_Person_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Person_proto_goTypes = []interface{}{
	(*PersonListReq)(nil),  // 0: services.PersonListReq
	(*PersonListResp)(nil), // 1: services.PersonListResp
	(*PersonRespData)(nil), // 2: services.PersonRespData
	(*PersonInfoReq)(nil),  // 3: services.PersonInfoReq
	(*PersonInfoResp)(nil), // 4: services.PersonInfoResp
}
var file_Person_proto_depIdxs = []int32{
	2, // 0: services.PersonListResp.data:type_name -> services.PersonRespData
	2, // 1: services.PersonInfoResp.data:type_name -> services.PersonRespData
	0, // 2: services.PersonService.GetPersonList:input_type -> services.PersonListReq
	3, // 3: services.PersonService.GetPersonInfo:input_type -> services.PersonInfoReq
	1, // 4: services.PersonService.GetPersonList:output_type -> services.PersonListResp
	4, // 5: services.PersonService.GetPersonInfo:output_type -> services.PersonInfoResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Person_proto_init() }
func file_Person_proto_init() {
	if File_Person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonListReq); i {
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
		file_Person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonListResp); i {
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
		file_Person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRespData); i {
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
		file_Person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfoReq); i {
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
		file_Person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfoResp); i {
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
			RawDescriptor: file_Person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Person_proto_goTypes,
		DependencyIndexes: file_Person_proto_depIdxs,
		MessageInfos:      file_Person_proto_msgTypes,
	}.Build()
	File_Person_proto = out.File
	file_Person_proto_rawDesc = nil
	file_Person_proto_goTypes = nil
	file_Person_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonServiceClient interface {
	GetPersonList(ctx context.Context, in *PersonListReq, opts ...grpc.CallOption) (*PersonListResp, error)
	GetPersonInfo(ctx context.Context, in *PersonInfoReq, opts ...grpc.CallOption) (*PersonInfoResp, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) GetPersonList(ctx context.Context, in *PersonListReq, opts ...grpc.CallOption) (*PersonListResp, error) {
	out := new(PersonListResp)
	err := c.cc.Invoke(ctx, "/services.PersonService/GetPersonList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPersonInfo(ctx context.Context, in *PersonInfoReq, opts ...grpc.CallOption) (*PersonInfoResp, error) {
	out := new(PersonInfoResp)
	err := c.cc.Invoke(ctx, "/services.PersonService/GetPersonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
type PersonServiceServer interface {
	GetPersonList(context.Context, *PersonListReq) (*PersonListResp, error)
	GetPersonInfo(context.Context, *PersonInfoReq) (*PersonInfoResp, error)
}

// UnimplementedPersonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (*UnimplementedPersonServiceServer) GetPersonList(context.Context, *PersonListReq) (*PersonListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonList not implemented")
}
func (*UnimplementedPersonServiceServer) GetPersonInfo(context.Context, *PersonInfoReq) (*PersonInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonInfo not implemented")
}

func RegisterPersonServiceServer(s *grpc.Server, srv PersonServiceServer) {
	s.RegisterService(&_PersonService_serviceDesc, srv)
}

func _PersonService_GetPersonList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPersonList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PersonService/GetPersonList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPersonList(ctx, req.(*PersonListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPersonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPersonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PersonService/GetPersonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPersonInfo(ctx, req.(*PersonInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonList",
			Handler:    _PersonService_GetPersonList_Handler,
		},
		{
			MethodName: "GetPersonInfo",
			Handler:    _PersonService_GetPersonInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Person.proto",
}
