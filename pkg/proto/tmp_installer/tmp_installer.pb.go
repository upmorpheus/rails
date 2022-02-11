// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/proto/tmp_installer/tmp_installer.proto

package tmp_installer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InstallTemporaryDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemporaryDirectory string `protobuf:"bytes,1,opt,name=temporary_directory,json=temporaryDirectory,proto3" json:"temporary_directory,omitempty"`
}

func (x *InstallTemporaryDirectoryRequest) Reset() {
	*x = InstallTemporaryDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallTemporaryDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallTemporaryDirectoryRequest) ProtoMessage() {}

func (x *InstallTemporaryDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallTemporaryDirectoryRequest.ProtoReflect.Descriptor instead.
func (*InstallTemporaryDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescGZIP(), []int{0}
}

func (x *InstallTemporaryDirectoryRequest) GetTemporaryDirectory() string {
	if x != nil {
		return x.TemporaryDirectory
	}
	return ""
}

var File_pkg_proto_tmp_installer_tmp_installer_proto protoreflect.FileDescriptor

var file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6d, 0x70, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x32, 0xcf, 0x01, 0x0a, 0x1b, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6e, 0x0a, 0x19, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2e, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescOnce sync.Once
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData = file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc
)

func file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescGZIP() []byte {
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescOnce.Do(func() {
		file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData)
	})
	return file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData
}

var file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes = []interface{}{
	(*InstallTemporaryDirectoryRequest)(nil), // 0: buildbarn.tmp_installer.InstallTemporaryDirectoryRequest
	(*emptypb.Empty)(nil),                    // 1: google.protobuf.Empty
}
var file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs = []int32{
	1, // 0: buildbarn.tmp_installer.TemporaryDirectoryInstaller.CheckReadiness:input_type -> google.protobuf.Empty
	0, // 1: buildbarn.tmp_installer.TemporaryDirectoryInstaller.InstallTemporaryDirectory:input_type -> buildbarn.tmp_installer.InstallTemporaryDirectoryRequest
	1, // 2: buildbarn.tmp_installer.TemporaryDirectoryInstaller.CheckReadiness:output_type -> google.protobuf.Empty
	1, // 3: buildbarn.tmp_installer.TemporaryDirectoryInstaller.InstallTemporaryDirectory:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_tmp_installer_tmp_installer_proto_init() }
func file_pkg_proto_tmp_installer_tmp_installer_proto_init() {
	if File_pkg_proto_tmp_installer_tmp_installer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallTemporaryDirectoryRequest); i {
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
			RawDescriptor: file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes,
		DependencyIndexes: file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs,
		MessageInfos:      file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes,
	}.Build()
	File_pkg_proto_tmp_installer_tmp_installer_proto = out.File
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc = nil
	file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes = nil
	file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemporaryDirectoryInstallerClient is the client API for TemporaryDirectoryInstaller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemporaryDirectoryInstallerClient interface {
	CheckReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InstallTemporaryDirectory(ctx context.Context, in *InstallTemporaryDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type temporaryDirectoryInstallerClient struct {
	cc grpc.ClientConnInterface
}

func NewTemporaryDirectoryInstallerClient(cc grpc.ClientConnInterface) TemporaryDirectoryInstallerClient {
	return &temporaryDirectoryInstallerClient{cc}
}

func (c *temporaryDirectoryInstallerClient) CheckReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/CheckReadiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temporaryDirectoryInstallerClient) InstallTemporaryDirectory(ctx context.Context, in *InstallTemporaryDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/InstallTemporaryDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemporaryDirectoryInstallerServer is the server API for TemporaryDirectoryInstaller service.
type TemporaryDirectoryInstallerServer interface {
	CheckReadiness(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	InstallTemporaryDirectory(context.Context, *InstallTemporaryDirectoryRequest) (*emptypb.Empty, error)
}

// UnimplementedTemporaryDirectoryInstallerServer can be embedded to have forward compatible implementations.
type UnimplementedTemporaryDirectoryInstallerServer struct {
}

func (*UnimplementedTemporaryDirectoryInstallerServer) CheckReadiness(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReadiness not implemented")
}
func (*UnimplementedTemporaryDirectoryInstallerServer) InstallTemporaryDirectory(context.Context, *InstallTemporaryDirectoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallTemporaryDirectory not implemented")
}

func RegisterTemporaryDirectoryInstallerServer(s grpc.ServiceRegistrar, srv TemporaryDirectoryInstallerServer) {
	s.RegisterService(&_TemporaryDirectoryInstaller_serviceDesc, srv)
}

func _TemporaryDirectoryInstaller_CheckReadiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporaryDirectoryInstallerServer).CheckReadiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/CheckReadiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporaryDirectoryInstallerServer).CheckReadiness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemporaryDirectoryInstaller_InstallTemporaryDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallTemporaryDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporaryDirectoryInstallerServer).InstallTemporaryDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/InstallTemporaryDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporaryDirectoryInstallerServer).InstallTemporaryDirectory(ctx, req.(*InstallTemporaryDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemporaryDirectoryInstaller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.tmp_installer.TemporaryDirectoryInstaller",
	HandlerType: (*TemporaryDirectoryInstallerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckReadiness",
			Handler:    _TemporaryDirectoryInstaller_CheckReadiness_Handler,
		},
		{
			MethodName: "InstallTemporaryDirectory",
			Handler:    _TemporaryDirectoryInstaller_InstallTemporaryDirectory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/tmp_installer/tmp_installer.proto",
}
