// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pkg/proto/outputpathpersistency/outputpathpersistency.proto

package outputpathpersistency

import (
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RootDirectory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialCreationTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=initial_creation_time,json=initialCreationTime,proto3" json:"initial_creation_time,omitempty"`
	Contents            *Directory             `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *RootDirectory) Reset() {
	*x = RootDirectory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootDirectory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootDirectory) ProtoMessage() {}

func (x *RootDirectory) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootDirectory.ProtoReflect.Descriptor instead.
func (*RootDirectory) Descriptor() ([]byte, []int) {
	return file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescGZIP(), []int{0}
}

func (x *RootDirectory) GetInitialCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.InitialCreationTime
	}
	return nil
}

func (x *RootDirectory) GetContents() *Directory {
	if x != nil {
		return x.Contents
	}
	return nil
}

type Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files       []*v2.FileNode    `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Directories []*DirectoryNode  `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty"`
	Symlinks    []*v2.SymlinkNode `protobuf:"bytes,3,rep,name=symlinks,proto3" json:"symlinks,omitempty"`
}

func (x *Directory) Reset() {
	*x = Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescGZIP(), []int{1}
}

func (x *Directory) GetFiles() []*v2.FileNode {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Directory) GetDirectories() []*DirectoryNode {
	if x != nil {
		return x.Directories
	}
	return nil
}

func (x *Directory) GetSymlinks() []*v2.SymlinkNode {
	if x != nil {
		return x.Symlinks
	}
	return nil
}

type DirectoryNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FileRegion *FileRegion `protobuf:"bytes,2,opt,name=file_region,json=fileRegion,proto3" json:"file_region,omitempty"`
}

func (x *DirectoryNode) Reset() {
	*x = DirectoryNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryNode) ProtoMessage() {}

func (x *DirectoryNode) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryNode.ProtoReflect.Descriptor instead.
func (*DirectoryNode) Descriptor() ([]byte, []int) {
	return file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirectoryNode) GetFileRegion() *FileRegion {
	if x != nil {
		return x.FileRegion
	}
	return nil
}

type FileRegion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffsetBytes int64 `protobuf:"varint,1,opt,name=offset_bytes,json=offsetBytes,proto3" json:"offset_bytes,omitempty"`
	SizeBytes   int32 `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
}

func (x *FileRegion) Reset() {
	*x = FileRegion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRegion) ProtoMessage() {}

func (x *FileRegion) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRegion.ProtoReflect.Descriptor instead.
func (*FileRegion) Descriptor() ([]byte, []int) {
	return file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescGZIP(), []int{3}
}

func (x *FileRegion) GetOffsetBytes() int64 {
	if x != nil {
		return x.OffsetBytes
	}
	return 0
}

func (x *FileRegion) GetSizeBytes() int32 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

var File_pkg_proto_outputpathpersistency_outputpathpersistency_proto protoreflect.FileDescriptor

var file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x70, 0x61, 0x74, 0x68, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70, 0x61, 0x74, 0x68, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70,
	0x61, 0x74, 0x68, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x36,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4e, 0x0a, 0x15, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70, 0x61,
	0x74, 0x68, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xe8, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x3f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x50, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70, 0x61, 0x74, 0x68, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x79, 0x6d, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a,
	0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x6d, 0x6c, 0x69, 0x6e, 0x6b, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x08, 0x73, 0x79, 0x6d, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x71, 0x0a, 0x0d,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70, 0x61, 0x74, 0x68, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22,
	0x4e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42,
	0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x70, 0x61, 0x74, 0x68,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescOnce sync.Once
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescData = file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDesc
)

func file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescGZIP() []byte {
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescOnce.Do(func() {
		file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescData)
	})
	return file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDescData
}

var file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_goTypes = []interface{}{
	(*RootDirectory)(nil),         // 0: buildbarn.outputpathpersistency.RootDirectory
	(*Directory)(nil),             // 1: buildbarn.outputpathpersistency.Directory
	(*DirectoryNode)(nil),         // 2: buildbarn.outputpathpersistency.DirectoryNode
	(*FileRegion)(nil),            // 3: buildbarn.outputpathpersistency.FileRegion
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*v2.FileNode)(nil),           // 5: build.bazel.remote.execution.v2.FileNode
	(*v2.SymlinkNode)(nil),        // 6: build.bazel.remote.execution.v2.SymlinkNode
}
var file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_depIdxs = []int32{
	4, // 0: buildbarn.outputpathpersistency.RootDirectory.initial_creation_time:type_name -> google.protobuf.Timestamp
	1, // 1: buildbarn.outputpathpersistency.RootDirectory.contents:type_name -> buildbarn.outputpathpersistency.Directory
	5, // 2: buildbarn.outputpathpersistency.Directory.files:type_name -> build.bazel.remote.execution.v2.FileNode
	2, // 3: buildbarn.outputpathpersistency.Directory.directories:type_name -> buildbarn.outputpathpersistency.DirectoryNode
	6, // 4: buildbarn.outputpathpersistency.Directory.symlinks:type_name -> build.bazel.remote.execution.v2.SymlinkNode
	3, // 5: buildbarn.outputpathpersistency.DirectoryNode.file_region:type_name -> buildbarn.outputpathpersistency.FileRegion
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_init() }
func file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_init() {
	if File_pkg_proto_outputpathpersistency_outputpathpersistency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootDirectory); i {
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
		file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directory); i {
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
		file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryNode); i {
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
		file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRegion); i {
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
			RawDescriptor: file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_goTypes,
		DependencyIndexes: file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_depIdxs,
		MessageInfos:      file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_msgTypes,
	}.Build()
	File_pkg_proto_outputpathpersistency_outputpathpersistency_proto = out.File
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_rawDesc = nil
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_goTypes = nil
	file_pkg_proto_outputpathpersistency_outputpathpersistency_proto_depIdxs = nil
}
