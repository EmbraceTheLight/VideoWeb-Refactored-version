// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: user/v1/identity/captcha.proto

package v1id

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetImageCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetImageCaptchaRequest) Reset() {
	*x = GetImageCaptchaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_identity_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageCaptchaRequest) ProtoMessage() {}

func (x *GetImageCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_identity_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageCaptchaRequest.ProtoReflect.Descriptor instead.
func (*GetImageCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_identity_captcha_proto_rawDescGZIP(), []int{0}
}

type GetImageCaptchaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32                              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg           string                             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	CaptchaResult *GetImageCaptchaResp_CaptchaResult `protobuf:"bytes,3,opt,name=captcha_result,json=captchaResult,proto3" json:"captcha_result,omitempty"`
}

func (x *GetImageCaptchaResp) Reset() {
	*x = GetImageCaptchaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_identity_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageCaptchaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageCaptchaResp) ProtoMessage() {}

func (x *GetImageCaptchaResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_identity_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageCaptchaResp.ProtoReflect.Descriptor instead.
func (*GetImageCaptchaResp) Descriptor() ([]byte, []int) {
	return file_user_v1_identity_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *GetImageCaptchaResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetImageCaptchaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetImageCaptchaResp) GetCaptchaResult() *GetImageCaptchaResp_CaptchaResult {
	if x != nil {
		return x.CaptchaResult
	}
	return nil
}

type GetCodeCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // user email
}

func (x *GetCodeCaptchaRequest) Reset() {
	*x = GetCodeCaptchaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_identity_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeCaptchaRequest) ProtoMessage() {}

func (x *GetCodeCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_identity_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeCaptchaRequest.ProtoReflect.Descriptor instead.
func (*GetCodeCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_identity_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *GetCodeCaptchaRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetCodeCaptchaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetCodeCaptchaResp) Reset() {
	*x = GetCodeCaptchaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_identity_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeCaptchaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeCaptchaResp) ProtoMessage() {}

func (x *GetCodeCaptchaResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_identity_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeCaptchaResp.ProtoReflect.Descriptor instead.
func (*GetCodeCaptchaResp) Descriptor() ([]byte, []int) {
	return file_user_v1_identity_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *GetCodeCaptchaResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetCodeCaptchaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCodeCaptchaResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetImageCaptchaResp_CaptchaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	B64Log string `protobuf:"bytes,2,opt,name=b64log,proto3" json:"b64log,omitempty"`
	Answer string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *GetImageCaptchaResp_CaptchaResult) Reset() {
	*x = GetImageCaptchaResp_CaptchaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_identity_captcha_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageCaptchaResp_CaptchaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageCaptchaResp_CaptchaResult) ProtoMessage() {}

func (x *GetImageCaptchaResp_CaptchaResult) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_identity_captcha_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageCaptchaResp_CaptchaResult.ProtoReflect.Descriptor instead.
func (*GetImageCaptchaResp_CaptchaResult) Descriptor() ([]byte, []int) {
	return file_user_v1_identity_captcha_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetImageCaptchaResp_CaptchaResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetImageCaptchaResp_CaptchaResult) GetB64Log() string {
	if x != nil {
		return x.B64Log
	}
	return ""
}

func (x *GetImageCaptchaResp_CaptchaResult) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_user_v1_identity_captcha_proto protoreflect.FileDescriptor

var file_user_v1_identity_captcha_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xef, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x54, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x4f, 0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x36, 0x34, 0x6c, 0x6f, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x36, 0x34, 0x6c, 0x6f, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0xba, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x58,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x20, 0x5a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3b, 0x76, 0x31, 0x69,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_identity_captcha_proto_rawDescOnce sync.Once
	file_user_v1_identity_captcha_proto_rawDescData = file_user_v1_identity_captcha_proto_rawDesc
)

func file_user_v1_identity_captcha_proto_rawDescGZIP() []byte {
	file_user_v1_identity_captcha_proto_rawDescOnce.Do(func() {
		file_user_v1_identity_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_identity_captcha_proto_rawDescData)
	})
	return file_user_v1_identity_captcha_proto_rawDescData
}

var file_user_v1_identity_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_v1_identity_captcha_proto_goTypes = []any{
	(*GetImageCaptchaRequest)(nil),            // 0: user.v1.id.GetImageCaptchaRequest
	(*GetImageCaptchaResp)(nil),               // 1: user.v1.id.GetImageCaptchaResp
	(*GetCodeCaptchaRequest)(nil),             // 2: user.v1.id.GetCodeCaptchaRequest
	(*GetCodeCaptchaResp)(nil),                // 3: user.v1.id.GetCodeCaptchaResp
	(*GetImageCaptchaResp_CaptchaResult)(nil), // 4: user.v1.id.GetImageCaptchaResp.CaptchaResult
}
var file_user_v1_identity_captcha_proto_depIdxs = []int32{
	4, // 0: user.v1.id.GetImageCaptchaResp.captcha_result:type_name -> user.v1.id.GetImageCaptchaResp.CaptchaResult
	0, // 1: user.v1.id.Captcha.GetImageCaptcha:input_type -> user.v1.id.GetImageCaptchaRequest
	2, // 2: user.v1.id.Captcha.GetCodeCaptcha:input_type -> user.v1.id.GetCodeCaptchaRequest
	1, // 3: user.v1.id.Captcha.GetImageCaptcha:output_type -> user.v1.id.GetImageCaptchaResp
	3, // 4: user.v1.id.Captcha.GetCodeCaptcha:output_type -> user.v1.id.GetCodeCaptchaResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_v1_identity_captcha_proto_init() }
func file_user_v1_identity_captcha_proto_init() {
	if File_user_v1_identity_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_identity_captcha_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetImageCaptchaRequest); i {
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
		file_user_v1_identity_captcha_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetImageCaptchaResp); i {
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
		file_user_v1_identity_captcha_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetCodeCaptchaRequest); i {
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
		file_user_v1_identity_captcha_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCodeCaptchaResp); i {
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
		file_user_v1_identity_captcha_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetImageCaptchaResp_CaptchaResult); i {
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
			RawDescriptor: file_user_v1_identity_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_identity_captcha_proto_goTypes,
		DependencyIndexes: file_user_v1_identity_captcha_proto_depIdxs,
		MessageInfos:      file_user_v1_identity_captcha_proto_msgTypes,
	}.Build()
	File_user_v1_identity_captcha_proto = out.File
	file_user_v1_identity_captcha_proto_rawDesc = nil
	file_user_v1_identity_captcha_proto_goTypes = nil
	file_user_v1_identity_captcha_proto_depIdxs = nil
}
