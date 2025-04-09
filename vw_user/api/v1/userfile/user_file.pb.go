// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: v1/userfile/user_file.proto

package filev1

import (
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

type UploadAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadAvatarReq_FileName
	//	*UploadAvatarReq_FileContent
	Data isUploadAvatarReq_Data `protobuf_oneof:"data"`
}

func (x *UploadAvatarReq) Reset() {
	*x = UploadAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userfile_user_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarReq) ProtoMessage() {}

func (x *UploadAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userfile_user_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarReq.ProtoReflect.Descriptor instead.
func (*UploadAvatarReq) Descriptor() ([]byte, []int) {
	return file_v1_userfile_user_file_proto_rawDescGZIP(), []int{0}
}

func (m *UploadAvatarReq) GetData() isUploadAvatarReq_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadAvatarReq) GetFileName() string {
	if x, ok := x.GetData().(*UploadAvatarReq_FileName); ok {
		return x.FileName
	}
	return ""
}

func (x *UploadAvatarReq) GetFileContent() []byte {
	if x, ok := x.GetData().(*UploadAvatarReq_FileContent); ok {
		return x.FileContent
	}
	return nil
}

type isUploadAvatarReq_Data interface {
	isUploadAvatarReq_Data()
}

type UploadAvatarReq_FileName struct {
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3,oneof"`
}

type UploadAvatarReq_FileContent struct {
	FileContent []byte `protobuf:"bytes,2,opt,name=file_content,json=fileContent,proto3,oneof"`
}

func (*UploadAvatarReq_FileName) isUploadAvatarReq_Data() {}

func (*UploadAvatarReq_FileContent) isUploadAvatarReq_Data() {}

type UploadAvatarResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *UploadAvatarResp) Reset() {
	*x = UploadAvatarResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userfile_user_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarResp) ProtoMessage() {}

func (x *UploadAvatarResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userfile_user_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarResp.ProtoReflect.Descriptor instead.
func (*UploadAvatarResp) Descriptor() ([]byte, []int) {
	return file_v1_userfile_user_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAvatarResp) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type UpdateAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UpdateAvatarReq_MetaData
	//	*UpdateAvatarReq_FileContent
	Data isUpdateAvatarReq_Data `protobuf_oneof:"data"`
}

func (x *UpdateAvatarReq) Reset() {
	*x = UpdateAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userfile_user_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvatarReq) ProtoMessage() {}

func (x *UpdateAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userfile_user_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvatarReq.ProtoReflect.Descriptor instead.
func (*UpdateAvatarReq) Descriptor() ([]byte, []int) {
	return file_v1_userfile_user_file_proto_rawDescGZIP(), []int{2}
}

func (m *UpdateAvatarReq) GetData() isUpdateAvatarReq_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateAvatarReq) GetMetaData() *UpdateAvatarReq_FileMetadata {
	if x, ok := x.GetData().(*UpdateAvatarReq_MetaData); ok {
		return x.MetaData
	}
	return nil
}

func (x *UpdateAvatarReq) GetFileContent() []byte {
	if x, ok := x.GetData().(*UpdateAvatarReq_FileContent); ok {
		return x.FileContent
	}
	return nil
}

type isUpdateAvatarReq_Data interface {
	isUpdateAvatarReq_Data()
}

type UpdateAvatarReq_MetaData struct {
	MetaData *UpdateAvatarReq_FileMetadata `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3,oneof"`
}

type UpdateAvatarReq_FileContent struct {
	FileContent []byte `protobuf:"bytes,2,opt,name=file_content,json=fileContent,proto3,oneof"`
}

func (*UpdateAvatarReq_MetaData) isUpdateAvatarReq_Data() {}

func (*UpdateAvatarReq_FileContent) isUpdateAvatarReq_Data() {}

type UpdateAvatarReq_FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileExtension string `protobuf:"bytes,2,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
}

func (x *UpdateAvatarReq_FileMetadata) Reset() {
	*x = UpdateAvatarReq_FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userfile_user_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvatarReq_FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvatarReq_FileMetadata) ProtoMessage() {}

func (x *UpdateAvatarReq_FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userfile_user_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvatarReq_FileMetadata.ProtoReflect.Descriptor instead.
func (*UpdateAvatarReq_FileMetadata) Descriptor() ([]byte, []int) {
	return file_v1_userfile_user_file_proto_rawDescGZIP(), []int{2, 0}
}

func (x *UpdateAvatarReq_FileMetadata) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateAvatarReq_FileMetadata) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

var File_v1_userfile_user_file_proto protoreflect.FileDescriptor

var file_v1_userfile_user_file_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x49, 0x0a, 0x09,
	0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x4e, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66,
	0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xab, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x76, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x3b, 0x66, 0x69,
	0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_userfile_user_file_proto_rawDescOnce sync.Once
	file_v1_userfile_user_file_proto_rawDescData = file_v1_userfile_user_file_proto_rawDesc
)

func file_v1_userfile_user_file_proto_rawDescGZIP() []byte {
	file_v1_userfile_user_file_proto_rawDescOnce.Do(func() {
		file_v1_userfile_user_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_userfile_user_file_proto_rawDescData)
	})
	return file_v1_userfile_user_file_proto_rawDescData
}

var file_v1_userfile_user_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_userfile_user_file_proto_goTypes = []any{
	(*UploadAvatarReq)(nil),              // 0: user.v1.file.UploadAvatarReq
	(*UploadAvatarResp)(nil),             // 1: user.v1.file.UploadAvatarResp
	(*UpdateAvatarReq)(nil),              // 2: user.v1.file.UpdateAvatarReq
	(*UpdateAvatarReq_FileMetadata)(nil), // 3: user.v1.file.UpdateAvatarReq.FileMetadata
	(*emptypb.Empty)(nil),                // 4: google.protobuf.Empty
}
var file_v1_userfile_user_file_proto_depIdxs = []int32{
	3, // 0: user.v1.file.UpdateAvatarReq.meta_data:type_name -> user.v1.file.UpdateAvatarReq.FileMetadata
	0, // 1: user.v1.file.FileService.UploadAvatar:input_type -> user.v1.file.UploadAvatarReq
	2, // 2: user.v1.file.FileService.UpdateAvatar:input_type -> user.v1.file.UpdateAvatarReq
	1, // 3: user.v1.file.FileService.UploadAvatar:output_type -> user.v1.file.UploadAvatarResp
	4, // 4: user.v1.file.FileService.UpdateAvatar:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_userfile_user_file_proto_init() }
func file_v1_userfile_user_file_proto_init() {
	if File_v1_userfile_user_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_userfile_user_file_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UploadAvatarReq); i {
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
		file_v1_userfile_user_file_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadAvatarResp); i {
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
		file_v1_userfile_user_file_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAvatarReq); i {
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
		file_v1_userfile_user_file_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAvatarReq_FileMetadata); i {
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
	file_v1_userfile_user_file_proto_msgTypes[0].OneofWrappers = []any{
		(*UploadAvatarReq_FileName)(nil),
		(*UploadAvatarReq_FileContent)(nil),
	}
	file_v1_userfile_user_file_proto_msgTypes[2].OneofWrappers = []any{
		(*UpdateAvatarReq_MetaData)(nil),
		(*UpdateAvatarReq_FileContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_userfile_user_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_userfile_user_file_proto_goTypes,
		DependencyIndexes: file_v1_userfile_user_file_proto_depIdxs,
		MessageInfos:      file_v1_userfile_user_file_proto_msgTypes,
	}.Build()
	File_v1_userfile_user_file_proto = out.File
	file_v1_userfile_user_file_proto_rawDesc = nil
	file_v1_userfile_user_file_proto_goTypes = nil
	file_v1_userfile_user_file_proto_depIdxs = nil
}
