// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: v1/video_comment/comment.proto

package comment

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

type CommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId   int64                  `protobuf:"zigzag64,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	PublisherId int64                  `protobuf:"zigzag64,2,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	RootId      int64                  `protobuf:"zigzag64,3,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	Content     string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpvoteCount uint64                 `protobuf:"varint,6,opt,name=upvote_count,json=upvoteCount,proto3" json:"upvote_count,omitempty"`
	IpAddress   string                 `protobuf:"bytes,7,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	CntReplies  uint32                 `protobuf:"varint,8,opt,name=cnt_replies,json=cntReplies,proto3" json:"cnt_replies,omitempty"`
}

func (x *CommentInfo) Reset() {
	*x = CommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentInfo) ProtoMessage() {}

func (x *CommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentInfo.ProtoReflect.Descriptor instead.
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentInfo) GetPublisherId() int64 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *CommentInfo) GetRootId() int64 {
	if x != nil {
		return x.RootId
	}
	return 0
}

func (x *CommentInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CommentInfo) GetUpvoteCount() uint64 {
	if x != nil {
		return x.UpvoteCount
	}
	return 0
}

func (x *CommentInfo) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *CommentInfo) GetCntReplies() uint32 {
	if x != nil {
		return x.CntReplies
	}
	return 0
}

type PublishCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId     int64  `protobuf:"zigzag64,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	ParentId    int64  `protobuf:"zigzag64,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PublisherId int64  `protobuf:"zigzag64,3,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PublishCommentReq) Reset() {
	*x = PublishCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCommentReq) ProtoMessage() {}

func (x *PublishCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCommentReq.ProtoReflect.Descriptor instead.
func (*PublishCommentReq) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{1}
}

func (x *PublishCommentReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *PublishCommentReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *PublishCommentReq) GetPublisherId() int64 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *PublishCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PublishCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"zigzag64,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *PublishCommentResp) Reset() {
	*x = PublishCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCommentResp) ProtoMessage() {}

func (x *PublishCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCommentResp.ProtoReflect.Descriptor instead.
func (*PublishCommentResp) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{2}
}

func (x *PublishCommentResp) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type GetCommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId  int64  `protobuf:"zigzag64,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	PageNum  int32  `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy   string `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Order    string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetCommentListReq) Reset() {
	*x = GetCommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentListReq) ProtoMessage() {}

func (x *GetCommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentListReq.ProtoReflect.Descriptor instead.
func (*GetCommentListReq) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentListReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *GetCommentListReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetCommentListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetCommentListReq) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetCommentListReq) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetCommentListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*CommentInfo `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	TotalCount int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetCommentListResp) Reset() {
	*x = GetCommentListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentListResp) ProtoMessage() {}

func (x *GetCommentListResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentListResp.ProtoReflect.Descriptor instead.
func (*GetCommentListResp) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentListResp) GetComments() []*CommentInfo {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetCommentListResp) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetCommentRepliesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64  `protobuf:"zigzag64,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	PageNum   int32  `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize  int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy    string `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Order     string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetCommentRepliesReq) Reset() {
	*x = GetCommentRepliesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRepliesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRepliesReq) ProtoMessage() {}

func (x *GetCommentRepliesReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRepliesReq.ProtoReflect.Descriptor instead.
func (*GetCommentRepliesReq) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentRepliesReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *GetCommentRepliesReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetCommentRepliesReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetCommentRepliesReq) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetCommentRepliesReq) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetCommentRepliesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replies []*CommentInfo `protobuf:"bytes,1,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *GetCommentRepliesResp) Reset() {
	*x = GetCommentRepliesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRepliesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRepliesResp) ProtoMessage() {}

func (x *GetCommentRepliesResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRepliesResp.ProtoReflect.Descriptor instead.
func (*GetCommentRepliesResp) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommentRepliesResp) GetReplies() []*CommentInfo {
	if x != nil {
		return x.Replies
	}
	return nil
}

type UpvoteCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"zigzag64,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId    int64 `protobuf:"zigzag64,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// If is_upvote is true, the user has upvoted the comment before, so we need to cancel the upvote.
	// Otherwise, the user has not upvoted the comment before, so we need to add the upvote.
	IsUpvoted bool `protobuf:"varint,3,opt,name=is_upvoted,json=isUpvoted,proto3" json:"is_upvoted,omitempty"`
}

func (x *UpvoteCommentReq) Reset() {
	*x = UpvoteCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteCommentReq) ProtoMessage() {}

func (x *UpvoteCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpvoteCommentReq.ProtoReflect.Descriptor instead.
func (*UpvoteCommentReq) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{7}
}

func (x *UpvoteCommentReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *UpvoteCommentReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpvoteCommentReq) GetIsUpvoted() bool {
	if x != nil {
		return x.IsUpvoted
	}
	return false
}

type UpvoteCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpvoteCount uint64 `protobuf:"varint,1,opt,name=upvote_count,json=upvoteCount,proto3" json:"upvote_count,omitempty"`
}

func (x *UpvoteCommentResp) Reset() {
	*x = UpvoteCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteCommentResp) ProtoMessage() {}

func (x *UpvoteCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpvoteCommentResp.ProtoReflect.Descriptor instead.
func (*UpvoteCommentResp) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{8}
}

func (x *UpvoteCommentResp) GetUpvoteCount() uint64 {
	if x != nil {
		return x.UpvoteCount
	}
	return 0
}

type CheckIfUserUpvotedCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"zigzag64,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId    int64 `protobuf:"zigzag64,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckIfUserUpvotedCommentReq) Reset() {
	*x = CheckIfUserUpvotedCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfUserUpvotedCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfUserUpvotedCommentReq) ProtoMessage() {}

func (x *CheckIfUserUpvotedCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfUserUpvotedCommentReq.ProtoReflect.Descriptor instead.
func (*CheckIfUserUpvotedCommentReq) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{9}
}

func (x *CheckIfUserUpvotedCommentReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CheckIfUserUpvotedCommentReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckIfUserUpvotedCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUpvoted bool `protobuf:"varint,1,opt,name=is_upvoted,json=isUpvoted,proto3" json:"is_upvoted,omitempty"`
}

func (x *CheckIfUserUpvotedCommentResp) Reset() {
	*x = CheckIfUserUpvotedCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_video_comment_comment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfUserUpvotedCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfUserUpvotedCommentResp) ProtoMessage() {}

func (x *CheckIfUserUpvotedCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_video_comment_comment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfUserUpvotedCommentResp.ProtoReflect.Descriptor instead.
func (*CheckIfUserUpvotedCommentResp) Descriptor() ([]byte, []int) {
	return file_v1_video_comment_comment_proto_rawDescGZIP(), []int{10}
}

func (x *CheckIfUserUpvotedCommentResp) GetIsUpvoted() bool {
	if x != nil {
		return x.IsUpvoted
	}
	return false
}

var File_v1_video_comment_comment_proto protoreflect.FileDescriptor

var file_v1_video_comment_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x33, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x72, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x52, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x10, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x22,
	0x36, 0x0a, 0x11, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x1c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3e, 0x0a, 0x1d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x32,
	0x96, 0x04, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x61, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_video_comment_comment_proto_rawDescOnce sync.Once
	file_v1_video_comment_comment_proto_rawDescData = file_v1_video_comment_comment_proto_rawDesc
)

func file_v1_video_comment_comment_proto_rawDescGZIP() []byte {
	file_v1_video_comment_comment_proto_rawDescOnce.Do(func() {
		file_v1_video_comment_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_video_comment_comment_proto_rawDescData)
	})
	return file_v1_video_comment_comment_proto_rawDescData
}

var file_v1_video_comment_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_video_comment_comment_proto_goTypes = []any{
	(*CommentInfo)(nil),                   // 0: comment.v1.comment.CommentInfo
	(*PublishCommentReq)(nil),             // 1: comment.v1.comment.PublishCommentReq
	(*PublishCommentResp)(nil),            // 2: comment.v1.comment.PublishCommentResp
	(*GetCommentListReq)(nil),             // 3: comment.v1.comment.GetCommentListReq
	(*GetCommentListResp)(nil),            // 4: comment.v1.comment.GetCommentListResp
	(*GetCommentRepliesReq)(nil),          // 5: comment.v1.comment.GetCommentRepliesReq
	(*GetCommentRepliesResp)(nil),         // 6: comment.v1.comment.GetCommentRepliesResp
	(*UpvoteCommentReq)(nil),              // 7: comment.v1.comment.UpvoteCommentReq
	(*UpvoteCommentResp)(nil),             // 8: comment.v1.comment.UpvoteCommentResp
	(*CheckIfUserUpvotedCommentReq)(nil),  // 9: comment.v1.comment.CheckIfUserUpvotedCommentReq
	(*CheckIfUserUpvotedCommentResp)(nil), // 10: comment.v1.comment.CheckIfUserUpvotedCommentResp
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_v1_video_comment_comment_proto_depIdxs = []int32{
	11, // 0: comment.v1.comment.CommentInfo.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: comment.v1.comment.GetCommentListResp.comments:type_name -> comment.v1.comment.CommentInfo
	0,  // 2: comment.v1.comment.GetCommentRepliesResp.replies:type_name -> comment.v1.comment.CommentInfo
	1,  // 3: comment.v1.comment.VideoComment.PublishComment:input_type -> comment.v1.comment.PublishCommentReq
	3,  // 4: comment.v1.comment.VideoComment.GetCommentList:input_type -> comment.v1.comment.GetCommentListReq
	5,  // 5: comment.v1.comment.VideoComment.GetCommentReplies:input_type -> comment.v1.comment.GetCommentRepliesReq
	7,  // 6: comment.v1.comment.VideoComment.UpvoteComment:input_type -> comment.v1.comment.UpvoteCommentReq
	9,  // 7: comment.v1.comment.VideoComment.CheckIfUserUpvotedComment:input_type -> comment.v1.comment.CheckIfUserUpvotedCommentReq
	2,  // 8: comment.v1.comment.VideoComment.PublishComment:output_type -> comment.v1.comment.PublishCommentResp
	4,  // 9: comment.v1.comment.VideoComment.GetCommentList:output_type -> comment.v1.comment.GetCommentListResp
	6,  // 10: comment.v1.comment.VideoComment.GetCommentReplies:output_type -> comment.v1.comment.GetCommentRepliesResp
	12, // 11: comment.v1.comment.VideoComment.UpvoteComment:output_type -> google.protobuf.Empty
	10, // 12: comment.v1.comment.VideoComment.CheckIfUserUpvotedComment:output_type -> comment.v1.comment.CheckIfUserUpvotedCommentResp
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_v1_video_comment_comment_proto_init() }
func file_v1_video_comment_comment_proto_init() {
	if File_v1_video_comment_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_video_comment_comment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommentInfo); i {
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
		file_v1_video_comment_comment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PublishCommentReq); i {
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
		file_v1_video_comment_comment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PublishCommentResp); i {
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
		file_v1_video_comment_comment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentListReq); i {
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
		file_v1_video_comment_comment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentListResp); i {
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
		file_v1_video_comment_comment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentRepliesReq); i {
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
		file_v1_video_comment_comment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentRepliesResp); i {
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
		file_v1_video_comment_comment_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpvoteCommentReq); i {
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
		file_v1_video_comment_comment_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpvoteCommentResp); i {
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
		file_v1_video_comment_comment_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CheckIfUserUpvotedCommentReq); i {
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
		file_v1_video_comment_comment_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CheckIfUserUpvotedCommentResp); i {
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
			RawDescriptor: file_v1_video_comment_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_video_comment_comment_proto_goTypes,
		DependencyIndexes: file_v1_video_comment_comment_proto_depIdxs,
		MessageInfos:      file_v1_video_comment_comment_proto_msgTypes,
	}.Build()
	File_v1_video_comment_comment_proto = out.File
	file_v1_video_comment_comment_proto_rawDesc = nil
	file_v1_video_comment_comment_proto_goTypes = nil
	file_v1_video_comment_comment_proto_depIdxs = nil
}
