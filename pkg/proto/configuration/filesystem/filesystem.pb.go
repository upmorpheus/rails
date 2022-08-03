// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pkg/proto/configuration/filesystem/filesystem.proto

package filesystem

import (
	blockdevice "github.com/buildbarn/bb-storage/pkg/proto/configuration/blockdevice"
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

type FilePoolConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Backend:
	//
	//	*FilePoolConfiguration_InMemory
	//	*FilePoolConfiguration_DirectoryPath
	//	*FilePoolConfiguration_BlockDevice
	Backend isFilePoolConfiguration_Backend `protobuf_oneof:"backend"`
}

func (x *FilePoolConfiguration) Reset() {
	*x = FilePoolConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePoolConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePoolConfiguration) ProtoMessage() {}

func (x *FilePoolConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePoolConfiguration.ProtoReflect.Descriptor instead.
func (*FilePoolConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescGZIP(), []int{0}
}

func (m *FilePoolConfiguration) GetBackend() isFilePoolConfiguration_Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (x *FilePoolConfiguration) GetInMemory() *emptypb.Empty {
	if x, ok := x.GetBackend().(*FilePoolConfiguration_InMemory); ok {
		return x.InMemory
	}
	return nil
}

func (x *FilePoolConfiguration) GetDirectoryPath() string {
	if x, ok := x.GetBackend().(*FilePoolConfiguration_DirectoryPath); ok {
		return x.DirectoryPath
	}
	return ""
}

func (x *FilePoolConfiguration) GetBlockDevice() *blockdevice.Configuration {
	if x, ok := x.GetBackend().(*FilePoolConfiguration_BlockDevice); ok {
		return x.BlockDevice
	}
	return nil
}

type isFilePoolConfiguration_Backend interface {
	isFilePoolConfiguration_Backend()
}

type FilePoolConfiguration_InMemory struct {
	InMemory *emptypb.Empty `protobuf:"bytes,1,opt,name=in_memory,json=inMemory,proto3,oneof"`
}

type FilePoolConfiguration_DirectoryPath struct {
	DirectoryPath string `protobuf:"bytes,2,opt,name=directory_path,json=directoryPath,proto3,oneof"`
}

type FilePoolConfiguration_BlockDevice struct {
	BlockDevice *blockdevice.Configuration `protobuf:"bytes,3,opt,name=block_device,json=blockDevice,proto3,oneof"`
}

func (*FilePoolConfiguration_InMemory) isFilePoolConfiguration_Backend() {}

func (*FilePoolConfiguration_DirectoryPath) isFilePoolConfiguration_Backend() {}

func (*FilePoolConfiguration_BlockDevice) isFilePoolConfiguration_Backend() {}

var File_pkg_proto_configuration_filesystem_filesystem_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_filesystem_filesystem_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x27,
	0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x57, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x42, 0x4d, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescData = file_pkg_proto_configuration_filesystem_filesystem_proto_rawDesc
)

func file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescData)
	})
	return file_pkg_proto_configuration_filesystem_filesystem_proto_rawDescData
}

var file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_configuration_filesystem_filesystem_proto_goTypes = []interface{}{
	(*FilePoolConfiguration)(nil),     // 0: buildbarn.configuration.filesystem.FilePoolConfiguration
	(*emptypb.Empty)(nil),             // 1: google.protobuf.Empty
	(*blockdevice.Configuration)(nil), // 2: buildbarn.configuration.blockdevice.Configuration
}
var file_pkg_proto_configuration_filesystem_filesystem_proto_depIdxs = []int32{
	1, // 0: buildbarn.configuration.filesystem.FilePoolConfiguration.in_memory:type_name -> google.protobuf.Empty
	2, // 1: buildbarn.configuration.filesystem.FilePoolConfiguration.block_device:type_name -> buildbarn.configuration.blockdevice.Configuration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_filesystem_filesystem_proto_init() }
func file_pkg_proto_configuration_filesystem_filesystem_proto_init() {
	if File_pkg_proto_configuration_filesystem_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePoolConfiguration); i {
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
	file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilePoolConfiguration_InMemory)(nil),
		(*FilePoolConfiguration_DirectoryPath)(nil),
		(*FilePoolConfiguration_BlockDevice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_filesystem_filesystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_filesystem_filesystem_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_filesystem_filesystem_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_filesystem_filesystem_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_filesystem_filesystem_proto = out.File
	file_pkg_proto_configuration_filesystem_filesystem_proto_rawDesc = nil
	file_pkg_proto_configuration_filesystem_filesystem_proto_goTypes = nil
	file_pkg_proto_configuration_filesystem_filesystem_proto_depIdxs = nil
}
