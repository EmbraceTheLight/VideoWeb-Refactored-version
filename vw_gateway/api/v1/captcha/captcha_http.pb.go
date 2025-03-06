// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0
// source: v1/captcha/captcha.proto

package captv1

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

const OperationCaptchaGetCodeCaptcha = "/gateway.v1.captcha.Captcha/GetCodeCaptcha"
const OperationCaptchaGetImageCaptcha = "/gateway.v1.captcha.Captcha/GetImageCaptcha"

type CaptchaHTTPServer interface {
	GetCodeCaptcha(context.Context, *GetCodeCaptchaReq) (*GetCodeCaptchaResp, error)
	GetImageCaptcha(context.Context, *GetImageCaptchaRequest) (*GetImageCaptchaResp, error)
}

func RegisterCaptchaHTTPServer(s *http.Server, srv CaptchaHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/captcha/graphic_captcha", _Captcha_GetImageCaptcha0_HTTP_Handler(srv))
	r.GET("/api/v1/captcha/code_captcha", _Captcha_GetCodeCaptcha0_HTTP_Handler(srv))
}

func _Captcha_GetImageCaptcha0_HTTP_Handler(srv CaptchaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetImageCaptchaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaptchaGetImageCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetImageCaptcha(ctx, req.(*GetImageCaptchaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetImageCaptchaResp)
		return ctx.Result(200, reply)
	}
}

func _Captcha_GetCodeCaptcha0_HTTP_Handler(srv CaptchaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCodeCaptchaReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaptchaGetCodeCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCodeCaptcha(ctx, req.(*GetCodeCaptchaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCodeCaptchaResp)
		return ctx.Result(200, reply)
	}
}

type CaptchaHTTPClient interface {
	GetCodeCaptcha(ctx context.Context, req *GetCodeCaptchaReq, opts ...http.CallOption) (rsp *GetCodeCaptchaResp, err error)
	GetImageCaptcha(ctx context.Context, req *GetImageCaptchaRequest, opts ...http.CallOption) (rsp *GetImageCaptchaResp, err error)
}

type CaptchaHTTPClientImpl struct {
	cc *http.Client
}

func NewCaptchaHTTPClient(client *http.Client) CaptchaHTTPClient {
	return &CaptchaHTTPClientImpl{client}
}

func (c *CaptchaHTTPClientImpl) GetCodeCaptcha(ctx context.Context, in *GetCodeCaptchaReq, opts ...http.CallOption) (*GetCodeCaptchaResp, error) {
	var out GetCodeCaptchaResp
	pattern := "/api/v1/captcha/code_captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCaptchaGetCodeCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CaptchaHTTPClientImpl) GetImageCaptcha(ctx context.Context, in *GetImageCaptchaRequest, opts ...http.CallOption) (*GetImageCaptchaResp, error) {
	var out GetImageCaptchaResp
	pattern := "/api/v1/captcha/graphic_captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCaptchaGetImageCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
