// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/video/videoinfo/videoinfo.proto

package videoinfo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VideoInfo_GetMpd_FullMethodName        = "/gateway.api.v1.video.videoinfo.VideoInfo/GetMpd"
	VideoInfo_GetSegments_FullMethodName   = "/gateway.api.v1.video.videoinfo.VideoInfo/GetSegments"
	VideoInfo_GetVideoCover_FullMethodName = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoCover"
	VideoInfo_GetVideoInfo_FullMethodName  = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoInfo"
	VideoInfo_GetVideoList_FullMethodName  = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoList"
)

// VideoInfoClient is the client API for VideoInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoInfoClient interface {
	GetMpd(ctx context.Context, in *ProvideMpdReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProvideMpdResp], error)
	GetSegments(ctx context.Context, in *ProvideSegmentsReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProvideSegmentsResp], error)
	GetVideoCover(ctx context.Context, in *GetVideoCoverReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoCoverResp], error)
	GetVideoInfo(ctx context.Context, in *GetVideoInfoReq, opts ...grpc.CallOption) (*GetVideoInfoResp, error)
	GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...grpc.CallOption) (*GetVideoListResp, error)
}

type videoInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoInfoClient(cc grpc.ClientConnInterface) VideoInfoClient {
	return &videoInfoClient{cc}
}

func (c *videoInfoClient) GetMpd(ctx context.Context, in *ProvideMpdReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProvideMpdResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[0], VideoInfo_GetMpd_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ProvideMpdReq, ProvideMpdResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetMpdClient = grpc.ServerStreamingClient[ProvideMpdResp]

func (c *videoInfoClient) GetSegments(ctx context.Context, in *ProvideSegmentsReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProvideSegmentsResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[1], VideoInfo_GetSegments_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ProvideSegmentsReq, ProvideSegmentsResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetSegmentsClient = grpc.ServerStreamingClient[ProvideSegmentsResp]

func (c *videoInfoClient) GetVideoCover(ctx context.Context, in *GetVideoCoverReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetVideoCoverResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoInfo_ServiceDesc.Streams[2], VideoInfo_GetVideoCover_FullMethodName, cOpts...)
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

// VideoInfoServer is the server API for VideoInfo service.
// All implementations must embed UnimplementedVideoInfoServer
// for forward compatibility.
type VideoInfoServer interface {
	GetMpd(*ProvideMpdReq, grpc.ServerStreamingServer[ProvideMpdResp]) error
	GetSegments(*ProvideSegmentsReq, grpc.ServerStreamingServer[ProvideSegmentsResp]) error
	GetVideoCover(*GetVideoCoverReq, grpc.ServerStreamingServer[GetVideoCoverResp]) error
	GetVideoInfo(context.Context, *GetVideoInfoReq) (*GetVideoInfoResp, error)
	GetVideoList(context.Context, *GetVideoListReq) (*GetVideoListResp, error)
	mustEmbedUnimplementedVideoInfoServer()
}

// UnimplementedVideoInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVideoInfoServer struct{}

func (UnimplementedVideoInfoServer) GetMpd(*ProvideMpdReq, grpc.ServerStreamingServer[ProvideMpdResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetMpd not implemented")
}
func (UnimplementedVideoInfoServer) GetSegments(*ProvideSegmentsReq, grpc.ServerStreamingServer[ProvideSegmentsResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetSegments not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoCover(*GetVideoCoverReq, grpc.ServerStreamingServer[GetVideoCoverResp]) error {
	return status.Errorf(codes.Unimplemented, "method GetVideoCover not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoInfo(context.Context, *GetVideoInfoReq) (*GetVideoInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfo not implemented")
}
func (UnimplementedVideoInfoServer) GetVideoList(context.Context, *GetVideoListReq) (*GetVideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoList not implemented")
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

func _VideoInfo_GetMpd_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProvideMpdReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetMpd(m, &grpc.GenericServerStream[ProvideMpdReq, ProvideMpdResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetMpdServer = grpc.ServerStreamingServer[ProvideMpdResp]

func _VideoInfo_GetSegments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProvideSegmentsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetSegments(m, &grpc.GenericServerStream[ProvideSegmentsReq, ProvideSegmentsResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetSegmentsServer = grpc.ServerStreamingServer[ProvideSegmentsResp]

func _VideoInfo_GetVideoCover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetVideoCoverReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoInfoServer).GetVideoCover(m, &grpc.GenericServerStream[GetVideoCoverReq, GetVideoCoverResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoInfo_GetVideoCoverServer = grpc.ServerStreamingServer[GetVideoCoverResp]

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

// VideoInfo_ServiceDesc is the grpc.ServiceDesc for VideoInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.api.v1.video.videoinfo.VideoInfo",
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMpd",
			Handler:       _VideoInfo_GetMpd_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSegments",
			Handler:       _VideoInfo_GetSegments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVideoCover",
			Handler:       _VideoInfo_GetVideoCover_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/video/videoinfo/videoinfo.proto",
}
