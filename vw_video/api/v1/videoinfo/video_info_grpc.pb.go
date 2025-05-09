// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/videoinfo/video_info.proto

package videoinfo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VideoInfo_GetVideoInfo_FullMethodName            = "/video.v1.videoinfo.VideoInfo/GetVideoInfo"
	VideoInfo_GetVideoList_FullMethodName            = "/video.v1.videoinfo.VideoInfo/GetVideoList"
	VideoInfo_GetVideoFile_FullMethodName            = "/video.v1.videoinfo.VideoInfo/GetVideoFile"
	VideoInfo_GetVideoMpd_FullMethodName             = "/video.v1.videoinfo.VideoInfo/GetVideoMpd"
	VideoInfo_GetVideoSegments_FullMethodName        = "/video.v1.videoinfo.VideoInfo/GetVideoSegments"
	VideoInfo_GetVideoCover_FullMethodName           = "/video.v1.videoinfo.VideoInfo/GetVideoCover"
	VideoInfo_UploadVideoInfo_FullMethodName         = "/video.v1.videoinfo.VideoInfo/UploadVideoInfo"
	VideoInfo_UploadVideoFile_FullMethodName         = "/video.v1.videoinfo.VideoInfo/UploadVideoFile"
	VideoInfo_UploadVideoCover_FullMethodName        = "/video.v1.videoinfo.VideoInfo/UploadVideoCover"
	VideoInfo_GetPublisherIdByVideoId_FullMethodName = "/video.v1.videoinfo.VideoInfo/GetPublisherIdByVideoId"
	VideoInfo_AddVideoCntShared_FullMethodName       = "/video.v1.videoinfo.VideoInfo/AddVideoCntShared"
	VideoInfo_AddVideoCntSharedRevert_FullMethodName = "/video.v1.videoinfo.VideoInfo/AddVideoCntSharedRevert"
)

// VideoInfoClient is the client API for VideoInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoInfoClient interface {
	GetVideoInfo(ctx context.Context, in *GetVideoInfoReq, opts ...grpc.CallOption) (*GetVideoInfoResp, error)
	GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...grpc.CallOption) (*GetVideoListResp, error)
	GetVideoFile(ctx context.Context, in *GetVideoFileReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoFileResp], error)
	GetVideoMpd(ctx context.Context, in *GetVideoMpdReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoMpdResp], error)
	GetVideoSegments(ctx context.Context, in *GetVideoSegmentReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoSegmentResp], error)
	GetVideoCover(ctx context.Context, in *GetVideoCoverReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoCoverResp], error)
	UploadVideoInfo(ctx context.Context, in *UploadVideoInfoReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UploadVideoFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadVideoFileReq, emptypb.Empty], error)
	UploadVideoCover(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadVideoCoverReq, emptypb.Empty], error)
	GetPublisherIdByVideoId(ctx context.Context, in *GetPublisherIdByVideoIdReq, opts ...grpc.CallOption) (*GetPublisherIdByVideoIdResp, error)
	AddVideoCntShared(ctx context.Context, in *AddVideoCntSharedReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddVideoCntSharedRevert(ctx context.Context, in *AddVideoCntSharedReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type videoInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoInfoClient(cc grpc.ClientConnInterface) VideoInfoClient {
	return &videoInfoClient{cc}
}

func (c *videoInfoClient) GetVideoInfo(ctx context.Context, in *GetVideoInfoReq, opts ...grpc.CallOption) (*GetVideoInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVideoInfoResp)
	err := c.cc.Invoke(ctx, VideoInfo_GetVideoInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoInfoClient) GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...grpc.CallOption) (*GetVideoListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVideoListResp)
	err := c.cc.Invoke(ctx, VideoInfo_GetVideoList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoInfoClient) GetVideoFile(ctx context.Context, in *GetVideoFileReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoFileResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[0], VideoInfo_GetVideoFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetVideoFileReq, GetVideoFileResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoFileClient = grpc.ServerStreamingClient[GetVideoFileResp]

func (c *videoInfoClient) GetVideoMpd(ctx context.Context, in *GetVideoMpdReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoMpdResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[1], VideoInfo_GetVideoMpd_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetVideoMpdReq, GetVideoMpdResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoMpdClient = grpc.ServerStreamingClient[GetVideoMpdResp]

func (c *videoInfoClient) GetVideoSegments(ctx context.Context, in *GetVideoSegmentReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoSegmentResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[2], VideoInfo_GetVideoSegments_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetVideoSegmentReq, GetVideoSegmentResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoSegmentsClient = grpc.ServerStreamingClient[GetVideoSegmentResp]

func (c *videoInfoClient) GetVideoCover(ctx context.Context, in *GetVideoCoverReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoCoverResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[3], VideoInfo_GetVideoCover_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetVideoCoverReq, GetVideoCoverResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoCoverClient = grpc.ServerStreamingClient[GetVideoCoverResp]

func (c *videoInfoClient) UploadVideoInfo(ctx context.Context, in *UploadVideoInfoReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VideoInfo_UploadVideoInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoInfoClient) UploadVideoFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadVideoFileReq, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[4], VideoInfo_UploadVideoFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadVideoFileReq, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_UploadVideoFileClient = grpc.ClientStreamingClient[UploadVideoFileReq, emptypb.Empty]

func (c *videoInfoClient) UploadVideoCover(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadVideoCoverReq, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[5], VideoInfo_UploadVideoCover_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadVideoCoverReq, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_UploadVideoCoverClient = grpc.ClientStreamingClient[UploadVideoCoverReq, emptypb.Empty]

func (c *videoInfoClient) GetPublisherIdByVideoId(ctx context.Context, in *GetPublisherIdByVideoIdReq, opts ...grpc.CallOption) (*GetPublisherIdByVideoIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPublisherIdByVideoIdResp)
	err := c.cc.Invoke(ctx, VideoInfo_GetPublisherIdByVideoId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoInfoClient) AddVideoCntShared(ctx context.Context, in *AddVideoCntSharedReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VideoInfo_AddVideoCntShared_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoInfoClient) AddVideoCntSharedRevert(ctx context.Context, in *AddVideoCntSharedReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VideoInfo_AddVideoCntSharedRevert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoInfoServer is the server API for VideoInfo service.
// All implementations must embed UnimplementedVideoInfoServer
// for forward compatibility.
type VideoInfoServer interface {
	GetVideoInfo(context.Context, *GetVideoInfoReq) (*GetVideoInfoResp, error)
	GetVideoList(context.Context, *GetVideoListReq) (*GetVideoListResp, error)
	GetVideoFile(*GetVideoFileReq, grpc.ServerStreamingServer[GetVideoFileResp]) error
	GetVideoMpd(*GetVideoMpdReq, grpc.ServerStreamingServer[GetVideoMpdResp]) error
	GetVideoSegments(*GetVideoSegmentReq, grpc.ServerStreamingServer[GetVideoSegmentResp]) error
	GetVideoCover(*GetVideoCoverReq, grpc.ServerStreamingServer[GetVideoCoverResp]) error
	UploadVideoInfo(context.Context, *UploadVideoInfoReq) (*emptypb.Empty, error)
	UploadVideoFile(grpc.ClientStreamingServer[UploadVideoFileReq, emptypb.Empty]) error
	UploadVideoCover(grpc.ClientStreamingServer[UploadVideoCoverReq, emptypb.Empty]) error
	GetPublisherIdByVideoId(context.Context, *GetPublisherIdByVideoIdReq) (*GetPublisherIdByVideoIdResp, error)
	AddVideoCntShared(context.Context, *AddVideoCntSharedReq) (*emptypb.Empty, error)
	AddVideoCntSharedRevert(context.Context, *AddVideoCntSharedReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedVideoInfoServer()
}

// UnimplementedVideoInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVideoInfoServer struct{}

func (UnimplementedVideoInfoServer) GetVideoInfo(context.Context, *GetVideoInfoReq) (*GetVideoInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfo not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoList(context.Context, *GetVideoListReq) (*GetVideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoList not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoFile(*GetVideoFileReq, grpc.ServerStreamingServer[GetVideoFileResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetVideoFile not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoMpd(*GetVideoMpdReq, grpc.ServerStreamingServer[GetVideoMpdResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetVideoMpd not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoSegments(*GetVideoSegmentReq, grpc.ServerStreamingServer[GetVideoSegmentResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetVideoSegments not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoCover(*GetVideoCoverReq, grpc.ServerStreamingServer[GetVideoCoverResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetVideoCover not implemented")
}
func (UnimplementedVideoInfoServer) UploadVideoInfo(context.Context, *UploadVideoInfoReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideoInfo not implemented")
}
func (UnimplementedVideoInfoServer) UploadVideoFile(grpc.ClientStreamingServer[UploadVideoFileReq, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideoFile not implemented")
}
func (UnimplementedVideoInfoServer) UploadVideoCover(grpc.ClientStreamingServer[UploadVideoCoverReq, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideoCover not implemented")
}
func (UnimplementedVideoInfoServer) GetPublisherIdByVideoId(context.Context, *GetPublisherIdByVideoIdReq) (*GetPublisherIdByVideoIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublisherIdByVideoId not implemented")
}
func (UnimplementedVideoInfoServer) AddVideoCntShared(context.Context, *AddVideoCntSharedReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideoCntShared not implemented")
}
func (UnimplementedVideoInfoServer) AddVideoCntSharedRevert(context.Context, *AddVideoCntSharedReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideoCntSharedRevert not implemented")
}
func (UnimplementedVideoInfoServer) mustEmbedUnimplementedVideoInfoServer() {}
func (UnimplementedVideoInfoServer) testEmbeddedByValue()                   {}

// UnsafeVideoInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoInfoServer will
// result in compilation errors.
type UnsafeVideoInfoServer interface {
	mustEmbedUnimplementedVideoInfoServer()
}

func RegisterVideoInfoServer(s grpc.ServiceRegistrar, srv VideoInfoServer) {
	// If the following call pancis, it indicates UnimplementedVideoInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VideoInfo_ServiceDesc, srv)
}

func _VideoInfo_GetVideoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).GetVideoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_GetVideoInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).GetVideoInfo(ctx, req.(*GetVideoInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoInfo_GetVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).GetVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_GetVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).GetVideoList(ctx, req.(*GetVideoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoInfo_GetVideoFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVideoFileReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetVideoFile(m, &grpc.GenericServerStream[GetVideoFileReq, GetVideoFileResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoFileServer = grpc.ServerStreamingServer[GetVideoFileResp]

func _VideoInfo_GetVideoMpd_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVideoMpdReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetVideoMpd(m, &grpc.GenericServerStream[GetVideoMpdReq, GetVideoMpdResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoMpdServer = grpc.ServerStreamingServer[GetVideoMpdResp]

func _VideoInfo_GetVideoSegments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVideoSegmentReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetVideoSegments(m, &grpc.GenericServerStream[GetVideoSegmentReq, GetVideoSegmentResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoSegmentsServer = grpc.ServerStreamingServer[GetVideoSegmentResp]

func _VideoInfo_GetVideoCover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVideoCoverReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetVideoCover(m, &grpc.GenericServerStream[GetVideoCoverReq, GetVideoCoverResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoCoverServer = grpc.ServerStreamingServer[GetVideoCoverResp]

func _VideoInfo_UploadVideoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).UploadVideoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_UploadVideoInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).UploadVideoInfo(ctx, req.(*UploadVideoInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoInfo_UploadVideoFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoInfoServer).UploadVideoFile(&grpc.GenericServerStream[UploadVideoFileReq, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_UploadVideoFileServer = grpc.ClientStreamingServer[UploadVideoFileReq, emptypb.Empty]

func _VideoInfo_UploadVideoCover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoInfoServer).UploadVideoCover(&grpc.GenericServerStream[UploadVideoCoverReq, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_UploadVideoCoverServer = grpc.ClientStreamingServer[UploadVideoCoverReq, emptypb.Empty]

func _VideoInfo_GetPublisherIdByVideoId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublisherIdByVideoIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).GetPublisherIdByVideoId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_GetPublisherIdByVideoId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).GetPublisherIdByVideoId(ctx, req.(*GetPublisherIdByVideoIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoInfo_AddVideoCntShared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideoCntSharedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).AddVideoCntShared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_AddVideoCntShared_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).AddVideoCntShared(ctx, req.(*AddVideoCntSharedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoInfo_AddVideoCntSharedRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideoCntSharedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoInfoServer).AddVideoCntSharedRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoInfo_AddVideoCntSharedRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoInfoServer).AddVideoCntSharedRevert(ctx, req.(*AddVideoCntSharedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoInfo_ServiceDesc is the grpc.ServiceDesc for VideoInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.v1.videoinfo.VideoInfo",
	HandlerType: (*VideoInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoInfo",
			Handler:    _VideoInfo_GetVideoInfo_Handler,
		},
		{
			MethodName: "GetVideoList",
			Handler:    _VideoInfo_GetVideoList_Handler,
		},
		{
			MethodName: "UploadVideoInfo",
			Handler:    _VideoInfo_UploadVideoInfo_Handler,
		},
		{
			MethodName: "GetPublisherIdByVideoId",
			Handler:    _VideoInfo_GetPublisherIdByVideoId_Handler,
		},
		{
			MethodName: "AddVideoCntShared",
			Handler:    _VideoInfo_AddVideoCntShared_Handler,
		},
		{
			MethodName: "AddVideoCntSharedRevert",
			Handler:    _VideoInfo_AddVideoCntSharedRevert_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetVideoFile",
			Handler:       _VideoInfo_GetVideoFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVideoMpd",
			Handler:       _VideoInfo_GetVideoMpd_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVideoSegments",
			Handler:       _VideoInfo_GetVideoSegments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVideoCover",
			Handler:       _VideoInfo_GetVideoCover_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadVideoFile",
			Handler:       _VideoInfo_UploadVideoFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadVideoCover",
			Handler:       _VideoInfo_UploadVideoCover_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "v1/videoinfo/video_info.proto",
}
