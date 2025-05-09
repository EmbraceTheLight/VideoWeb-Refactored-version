// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0
// source: v1/user/userfile/file.proto

package filev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFileServiceUpdateAvatar = "/gateway.api.v1.user.file.FileService/UpdateAvatar"
const OperationFileServiceUploadAvatar = "/gateway.api.v1.user.file.FileService/UploadAvatar"

type FileServiceHTTPServer interface {
	UpdateAvatar(context.Context, *UpdateAvatarReq) (*UpdateAvatarResp, error)
	UploadAvatar(context.Context, *emptypb.Empty) (*UploadAvatarResp, error)
}

func RegisterFileServiceHTTPServer(s *http.Server, srv FileServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/profile/avatar", _FileService_UploadAvatar0_HTTP_Handler(srv))
	r.PUT("/api/v1/{user_id}/profile/avatar", _FileService_UpdateAvatar0_HTTP_Handler(srv))
}

func _FileService_UploadAvatar0_HTTP_Handler(srv FileServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileServiceUploadAvatar)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadAvatar(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadAvatarResp)
		return ctx.Result(200, reply)
	}
}

func _FileService_UpdateAvatar0_HTTP_Handler(srv FileServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAvatarReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileServiceUpdateAvatar)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAvatar(ctx, req.(*UpdateAvatarReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAvatarResp)
		return ctx.Result(200, reply)
	}
}

type FileServiceHTTPClient interface {
	UpdateAvatar(ctx context.Context, req *UpdateAvatarReq, opts ...http.CallOption) (rsp *UpdateAvatarResp, err error)
	UploadAvatar(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UploadAvatarResp, err error)
}

type FileServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewFileServiceHTTPClient(client *http.Client) FileServiceHTTPClient {
	return &FileServiceHTTPClientImpl{client}
}

func (c *FileServiceHTTPClientImpl) UpdateAvatar(ctx context.Context, in *UpdateAvatarReq, opts ...http.CallOption) (*UpdateAvatarResp, error) {
	var out UpdateAvatarResp
	pattern := "/api/v1/{user_id}/profile/avatar"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileServiceUpdateAvatar))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileServiceHTTPClientImpl) UploadAvatar(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UploadAvatarResp, error) {
	var out UploadAvatarResp
	pattern := "/api/v1/user/profile/avatar"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileServiceUploadAvatar))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
