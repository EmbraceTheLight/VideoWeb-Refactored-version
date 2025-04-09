// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0
// source: v1/video/videoinfo/videoinfo.proto

package videoinfo

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationVideoInfoDownloadVideo = "/gateway.api.v1.video.videoinfo.VideoInfo/DownloadVideo"
const OperationVideoInfoGetMpd = "/gateway.api.v1.video.videoinfo.VideoInfo/GetMpd"
const OperationVideoInfoGetSegments = "/gateway.api.v1.video.videoinfo.VideoInfo/GetSegments"
const OperationVideoInfoGetVideoCover = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoCover"
const OperationVideoInfoGetVideoInfo = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoInfo"
const OperationVideoInfoGetVideoList = "/gateway.api.v1.video.videoinfo.VideoInfo/GetVideoList"
const OperationVideoInfoUploadVideoCover = "/gateway.api.v1.video.videoinfo.VideoInfo/UploadVideoCover"
const OperationVideoInfoUploadVideoFile = "/gateway.api.v1.video.videoinfo.VideoInfo/UploadVideoFile"
const OperationVideoInfoUploadVideoInfo = "/gateway.api.v1.video.videoinfo.VideoInfo/UploadVideoInfo"

type VideoInfoHTTPServer interface {
	DownloadVideo(context.Context, *DownloadVideoReq) (*FileResp, error)
	GetMpd(context.Context, *ProvideMpdReq) (*FileResp, error)
	GetSegments(context.Context, *ProvideSegmentsReq) (*FileResp, error)
	GetVideoCover(context.Context, *GetVideoCoverReq) (*FileResp, error)
	GetVideoInfo(context.Context, *GetVideoInfoReq) (*GetVideoInfoResp, error)
	GetVideoList(context.Context, *GetVideoListReq) (*GetVideoListResp, error)
	UploadVideoCover(context.Context, *UploadVideoCoverReq) (*UploadVideoCoverResp, error)
	UploadVideoFile(context.Context, *UploadVideoFileReq) (*UploadVideoFileResp, error)
	// UploadVideoInfo Request Order: 1. UploadVideoInfo 2. UploadVideoFile 3. UploadVideoCover
	// 1(UploadVideoInfo) Will Create a new directory for the video,
	// 2(UploadVideoFile) and 3(UploadVideoCover) will upload the video file and cover file to the directory.
	// The order of 2 and 3 is not important.
	UploadVideoInfo(context.Context, *UploadVideoInfoReq) (*UploadVideoInfoResp, error)
}

func RegisterVideoInfoHTTPServer(s *http.Server, srv VideoInfoHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/{user_id}/video/{video_id}/mpd", _VideoInfo_GetMpd0_HTTP_Handler(srv))
	r.GET("/api/v1/{user_id}/video/{video_id}/segments", _VideoInfo_GetSegments0_HTTP_Handler(srv))
	r.GET("/api/v1/{user_id}/video/{video_id}/cover", _VideoInfo_GetVideoCover0_HTTP_Handler(srv))
	r.GET("/api/v1/{user_id}/video/{video_id}/file", _VideoInfo_DownloadVideo0_HTTP_Handler(srv))
	r.POST("/api/v1/{user_id}/video/metaInfo", _VideoInfo_UploadVideoInfo0_HTTP_Handler(srv))
	r.POST("/api/v1/{user_id}/video/{video_id}/file", _VideoInfo_UploadVideoFile0_HTTP_Handler(srv))
	r.POST("/api/v1/{user_id}/video/{video_id}/cover", _VideoInfo_UploadVideoCover0_HTTP_Handler(srv))
	r.GET("/api/v1/video/{video_id}/info", _VideoInfo_GetVideoInfo0_HTTP_Handler(srv))
	r.GET("/api/v1/video/list", _VideoInfo_GetVideoList0_HTTP_Handler(srv))
}

func _VideoInfo_GetMpd0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProvideMpdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoGetMpd)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMpd(ctx, req.(*ProvideMpdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_GetSegments0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProvideSegmentsReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoGetSegments)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSegments(ctx, req.(*ProvideSegmentsReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_GetVideoCover0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoCoverReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoGetVideoCover)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoCover(ctx, req.(*GetVideoCoverReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_DownloadVideo0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadVideoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoDownloadVideo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadVideo(ctx, req.(*DownloadVideoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_UploadVideoInfo0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadVideoInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoUploadVideoInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadVideoInfo(ctx, req.(*UploadVideoInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadVideoInfoResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_UploadVideoFile0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadVideoFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoUploadVideoFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadVideoFile(ctx, req.(*UploadVideoFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadVideoFileResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_UploadVideoCover0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadVideoCoverReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoUploadVideoCover)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadVideoCover(ctx, req.(*UploadVideoCoverReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadVideoCoverResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_GetVideoInfo0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoInfoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoGetVideoInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoInfo(ctx, req.(*GetVideoInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVideoInfoResp)
		return ctx.Result(200, reply)
	}
}

func _VideoInfo_GetVideoList0_HTTP_Handler(srv VideoInfoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVideoListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVideoInfoGetVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVideoList(ctx, req.(*GetVideoListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVideoListResp)
		return ctx.Result(200, reply)
	}
}

type VideoInfoHTTPClient interface {
	DownloadVideo(ctx context.Context, req *DownloadVideoReq, opts ...http.CallOption) (rsp *FileResp, err error)
	GetMpd(ctx context.Context, req *ProvideMpdReq, opts ...http.CallOption) (rsp *FileResp, err error)
	GetSegments(ctx context.Context, req *ProvideSegmentsReq, opts ...http.CallOption) (rsp *FileResp, err error)
	GetVideoCover(ctx context.Context, req *GetVideoCoverReq, opts ...http.CallOption) (rsp *FileResp, err error)
	GetVideoInfo(ctx context.Context, req *GetVideoInfoReq, opts ...http.CallOption) (rsp *GetVideoInfoResp, err error)
	GetVideoList(ctx context.Context, req *GetVideoListReq, opts ...http.CallOption) (rsp *GetVideoListResp, err error)
	UploadVideoCover(ctx context.Context, req *UploadVideoCoverReq, opts ...http.CallOption) (rsp *UploadVideoCoverResp, err error)
	UploadVideoFile(ctx context.Context, req *UploadVideoFileReq, opts ...http.CallOption) (rsp *UploadVideoFileResp, err error)
	UploadVideoInfo(ctx context.Context, req *UploadVideoInfoReq, opts ...http.CallOption) (rsp *UploadVideoInfoResp, err error)
}

type VideoInfoHTTPClientImpl struct {
	cc *http.Client
}

func NewVideoInfoHTTPClient(client *http.Client) VideoInfoHTTPClient {
	return &VideoInfoHTTPClientImpl{client}
}

func (c *VideoInfoHTTPClientImpl) DownloadVideo(ctx context.Context, in *DownloadVideoReq, opts ...http.CallOption) (*FileResp, error) {
	var out FileResp
	pattern := "/api/v1/{user_id}/video/{video_id}/file"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoDownloadVideo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) GetMpd(ctx context.Context, in *ProvideMpdReq, opts ...http.CallOption) (*FileResp, error) {
	var out FileResp
	pattern := "/api/v1/{user_id}/video/{video_id}/mpd"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoGetMpd))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) GetSegments(ctx context.Context, in *ProvideSegmentsReq, opts ...http.CallOption) (*FileResp, error) {
	var out FileResp
	pattern := "/api/v1/{user_id}/video/{video_id}/segments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoGetSegments))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) GetVideoCover(ctx context.Context, in *GetVideoCoverReq, opts ...http.CallOption) (*FileResp, error) {
	var out FileResp
	pattern := "/api/v1/{user_id}/video/{video_id}/cover"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoGetVideoCover))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) GetVideoInfo(ctx context.Context, in *GetVideoInfoReq, opts ...http.CallOption) (*GetVideoInfoResp, error) {
	var out GetVideoInfoResp
	pattern := "/api/v1/video/{video_id}/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoGetVideoInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) GetVideoList(ctx context.Context, in *GetVideoListReq, opts ...http.CallOption) (*GetVideoListResp, error) {
	var out GetVideoListResp
	pattern := "/api/v1/video/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVideoInfoGetVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) UploadVideoCover(ctx context.Context, in *UploadVideoCoverReq, opts ...http.CallOption) (*UploadVideoCoverResp, error) {
	var out UploadVideoCoverResp
	pattern := "/api/v1/{user_id}/video/{video_id}/cover"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVideoInfoUploadVideoCover))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) UploadVideoFile(ctx context.Context, in *UploadVideoFileReq, opts ...http.CallOption) (*UploadVideoFileResp, error) {
	var out UploadVideoFileResp
	pattern := "/api/v1/{user_id}/video/{video_id}/file"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVideoInfoUploadVideoFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *VideoInfoHTTPClientImpl) UploadVideoInfo(ctx context.Context, in *UploadVideoInfoReq, opts ...http.CallOption) (*UploadVideoInfoResp, error) {
	var out UploadVideoInfoResp
	pattern := "/api/v1/{user_id}/video/metaInfo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVideoInfoUploadVideoInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
