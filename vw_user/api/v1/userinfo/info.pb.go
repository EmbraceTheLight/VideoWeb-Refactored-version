// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: v1/userinfo/info.proto

package infov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName   string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email      string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Signature  string                 `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Shells     int32                  `protobuf:"varint,4,opt,name=shells,proto3" json:"shells,omitempty"`
	CntFans    int32                  `protobuf:"varint,5,opt,name=cnt_fans,json=cntFans,proto3" json:"cnt_fans,omitempty"`
	CntFollows int32                  `protobuf:"varint,6,opt,name=cnt_follows,json=cntFollows,proto3" json:"cnt_follows,omitempty"`
	CntVideos  int32                  `protobuf:"varint,7,opt,name=cnt_videos,json=cntVideos,proto3" json:"cnt_videos,omitempty"`
	AvatarPath string                 `protobuf:"bytes,8,opt,name=avatar_path,json=avatarPath,proto3" json:"avatar_path,omitempty"`
	Gender     int32                  `protobuf:"varint,9,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UserInfo) GetShells() int32 {
	if x != nil {
		return x.Shells
	}
	return 0
}

func (x *UserInfo) GetCntFans() int32 {
	if x != nil {
		return x.CntFans
	}
	return 0
}

func (x *UserInfo) GetCntFollows() int32 {
	if x != nil {
		return x.CntFollows
	}
	return 0
}

func (x *UserInfo) GetCntVideos() int32 {
	if x != nil {
		return x.CntVideos
	}
	return 0
}

func (x *UserInfo) GetAvatarPath() string {
	if x != nil {
		return x.AvatarPath
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

type UserinfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserinfoReq) Reset() {
	*x = UserinfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserinfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserinfoReq) ProtoMessage() {}

func (x *UserinfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserinfoReq.ProtoReflect.Descriptor instead.
func (*UserinfoReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserinfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserinfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *UserinfoResp) Reset() {
	*x = UserinfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserinfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserinfoResp) ProtoMessage() {}

func (x *UserinfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserinfoResp.ProtoReflect.Descriptor instead.
func (*UserinfoResp) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{2}
}

func (x *UserinfoResp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type ModifyEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	InputCode string `protobuf:"bytes,3,opt,name=input_code,json=inputCode,proto3" json:"input_code,omitempty"` // the verify code that inputted by userbiz
}

func (x *ModifyEmailReq) Reset() {
	*x = ModifyEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyEmailReq) ProtoMessage() {}

func (x *ModifyEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyEmailReq.ProtoReflect.Descriptor instead.
func (*ModifyEmailReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyEmailReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModifyEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifyEmailReq) GetInputCode() string {
	if x != nil {
		return x.InputCode
	}
	return ""
}

type ModifyEmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ModifyEmailResp) Reset() {
	*x = ModifyEmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyEmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyEmailResp) ProtoMessage() {}

func (x *ModifyEmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyEmailResp.ProtoReflect.Descriptor instead.
func (*ModifyEmailResp) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{4}
}

func (x *ModifyEmailResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModifyEmailResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ModifyPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ModifyPasswordReq) Reset() {
	*x = ModifyPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPasswordReq) ProtoMessage() {}

func (x *ModifyPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPasswordReq.ProtoReflect.Descriptor instead.
func (*ModifyPasswordReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyPasswordReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModifyPasswordReq) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ModifyPasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ModifySignatureReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ModifySignatureReq) Reset() {
	*x = ModifySignatureReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySignatureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySignatureReq) ProtoMessage() {}

func (x *ModifySignatureReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySignatureReq.ProtoReflect.Descriptor instead.
func (*ModifySignatureReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{6}
}

func (x *ModifySignatureReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModifySignatureReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type ModifySignatureResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewSignature string `protobuf:"bytes,1,opt,name=new_signature,json=newSignature,proto3" json:"new_signature,omitempty"`
}

func (x *ModifySignatureResp) Reset() {
	*x = ModifySignatureResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySignatureResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySignatureResp) ProtoMessage() {}

func (x *ModifySignatureResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySignatureResp.ProtoReflect.Descriptor instead.
func (*ModifySignatureResp) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{7}
}

func (x *ModifySignatureResp) GetNewSignature() string {
	if x != nil {
		return x.NewSignature
	}
	return ""
}

type ForgetPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	InputCode   string `protobuf:"bytes,3,opt,name=input_code,json=inputCode,proto3" json:"input_code,omitempty"`
	NewPassword string `protobuf:"bytes,4,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ForgetPasswordReq) Reset() {
	*x = ForgetPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordReq) ProtoMessage() {}

func (x *ForgetPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordReq.ProtoReflect.Descriptor instead.
func (*ForgetPasswordReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{8}
}

func (x *ForgetPasswordReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ForgetPasswordReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ForgetPasswordReq) GetInputCode() string {
	if x != nil {
		return x.InputCode
	}
	return ""
}

func (x *ForgetPasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ModifyUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewUsername string `protobuf:"bytes,2,opt,name=new_username,json=newUsername,proto3" json:"new_username,omitempty"`
}

func (x *ModifyUsernameReq) Reset() {
	*x = ModifyUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyUsernameReq) ProtoMessage() {}

func (x *ModifyUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyUsernameReq.ProtoReflect.Descriptor instead.
func (*ModifyUsernameReq) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{9}
}

func (x *ModifyUsernameReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModifyUsernameReq) GetNewUsername() string {
	if x != nil {
		return x.NewUsername
	}
	return ""
}

type ModifyUsernameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewUsername string `protobuf:"bytes,1,opt,name=new_username,json=newUsername,proto3" json:"new_username,omitempty"`
}

func (x *ModifyUsernameResp) Reset() {
	*x = ModifyUsernameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_userinfo_info_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyUsernameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyUsernameResp) ProtoMessage() {}

func (x *ModifyUsernameResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_userinfo_info_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyUsernameResp.ProtoReflect.Descriptor instead.
func (*ModifyUsernameResp) Descriptor() ([]byte, []int) {
	return file_v1_userinfo_info_proto_rawDescGZIP(), []int{10}
}

func (x *ModifyUsernameResp) GetNewUsername() string {
	if x != nil {
		return x.NewUsername
	}
	return ""
}

var File_v1_userinfo_info_proto protoreflect.FileDescriptor

var file_v1_userinfo_info_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6e,
	0x74, 0x5f, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6e,
	0x74, 0x46, 0x61, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6e, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6e, 0x74, 0x5f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6e, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x36,
	0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x26, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x40, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x72, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x12, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4f, 0x0a, 0x11, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xef, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x55,
	0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x76, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x3b, 0x69, 0x6e, 0x66, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_userinfo_info_proto_rawDescOnce sync.Once
	file_v1_userinfo_info_proto_rawDescData = file_v1_userinfo_info_proto_rawDesc
)

func file_v1_userinfo_info_proto_rawDescGZIP() []byte {
	file_v1_userinfo_info_proto_rawDescOnce.Do(func() {
		file_v1_userinfo_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_userinfo_info_proto_rawDescData)
	})
	return file_v1_userinfo_info_proto_rawDescData
}

var file_v1_userinfo_info_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_userinfo_info_proto_goTypes = []any{
	(*UserInfo)(nil),              // 0: userbiz.v1.info.UserInfo
	(*UserinfoReq)(nil),           // 1: userbiz.v1.info.userinfoReq
	(*UserinfoResp)(nil),          // 2: userbiz.v1.info.userinfoResp
	(*ModifyEmailReq)(nil),        // 3: userbiz.v1.info.modifyEmailReq
	(*ModifyEmailResp)(nil),       // 4: userbiz.v1.info.modifyEmailResp
	(*ModifyPasswordReq)(nil),     // 5: userbiz.v1.info.modifyPasswordReq
	(*ModifySignatureReq)(nil),    // 6: userbiz.v1.info.modifySignatureReq
	(*ModifySignatureResp)(nil),   // 7: userbiz.v1.info.modifySignatureResp
	(*ForgetPasswordReq)(nil),     // 8: userbiz.v1.info.forgetPasswordReq
	(*ModifyUsernameReq)(nil),     // 9: userbiz.v1.info.modifyUsernameReq
	(*ModifyUsernameResp)(nil),    // 10: userbiz.v1.info.modifyUsernameResp
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_v1_userinfo_info_proto_depIdxs = []int32{
	11, // 0: userbiz.v1.info.UserInfo.birthday:type_name -> google.protobuf.Timestamp
	0,  // 1: userbiz.v1.info.userinfoResp.user_info:type_name -> userbiz.v1.info.UserInfo
	1,  // 2: userbiz.v1.info.Userinfo.GetUserinfo:input_type -> userbiz.v1.info.userinfoReq
	3,  // 3: userbiz.v1.info.Userinfo.ModifyEmail:input_type -> userbiz.v1.info.modifyEmailReq
	5,  // 4: userbiz.v1.info.Userinfo.ModifyPassword:input_type -> userbiz.v1.info.modifyPasswordReq
	6,  // 5: userbiz.v1.info.Userinfo.ModifyUserSignature:input_type -> userbiz.v1.info.modifySignatureReq
	8,  // 6: userbiz.v1.info.Userinfo.ForgetPassword:input_type -> userbiz.v1.info.forgetPasswordReq
	9,  // 7: userbiz.v1.info.Userinfo.ModifyUsername:input_type -> userbiz.v1.info.modifyUsernameReq
	2,  // 8: userbiz.v1.info.Userinfo.GetUserinfo:output_type -> userbiz.v1.info.userinfoResp
	4,  // 9: userbiz.v1.info.Userinfo.ModifyEmail:output_type -> userbiz.v1.info.modifyEmailResp
	12, // 10: userbiz.v1.info.Userinfo.ModifyPassword:output_type -> google.protobuf.Empty
	7,  // 11: userbiz.v1.info.Userinfo.ModifyUserSignature:output_type -> userbiz.v1.info.modifySignatureResp
	12, // 12: userbiz.v1.info.Userinfo.ForgetPassword:output_type -> google.protobuf.Empty
	10, // 13: userbiz.v1.info.Userinfo.ModifyUsername:output_type -> userbiz.v1.info.modifyUsernameResp
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_v1_userinfo_info_proto_init() }
func file_v1_userinfo_info_proto_init() {
	if File_v1_userinfo_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_userinfo_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfo); i {
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
		file_v1_userinfo_info_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserinfoReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserinfoResp); i {
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
		file_v1_userinfo_info_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ModifyEmailReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ModifyEmailResp); i {
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
		file_v1_userinfo_info_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ModifyPasswordReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySignatureReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySignatureResp); i {
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
		file_v1_userinfo_info_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ForgetPasswordReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ModifyUsernameReq); i {
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
		file_v1_userinfo_info_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ModifyUsernameResp); i {
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
			RawDescriptor: file_v1_userinfo_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_userinfo_info_proto_goTypes,
		DependencyIndexes: file_v1_userinfo_info_proto_depIdxs,
		MessageInfos:      file_v1_userinfo_info_proto_msgTypes,
	}.Build()
	File_v1_userinfo_info_proto = out.File
	file_v1_userinfo_info_proto_rawDesc = nil
	file_v1_userinfo_info_proto_goTypes = nil
	file_v1_userinfo_info_proto_depIdxs = nil
}
