// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0
// source: user/v1/identity/identity.proto

package v1id

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

const OperationIdentityLogin = "/user.v1.id.Identity/Login"
const OperationIdentityLogout = "/user.v1.id.Identity/Logout"
const OperationIdentityRegister = "/user.v1.id.Identity/Register"

type IdentityHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginResp, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResp, error)
	Register(context.Context, *RegisterRequest) (*RegisterResp, error)
}

func RegisterIdentityHTTPServer(s *http.Server, srv IdentityHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/login", _Identity_Login0_HTTP_Handler(srv))
	r.POST("/api/v1/user/register", _Identity_Register0_HTTP_Handler(srv))
	r.DELETE("/api/v1/user/logout", _Identity_Logout0_HTTP_Handler(srv))
}

func _Identity_Login0_HTTP_Handler(srv IdentityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationIdentityLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _Identity_Register0_HTTP_Handler(srv IdentityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationIdentityRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResp)
		return ctx.Result(200, reply)
	}
}

func _Identity_Logout0_HTTP_Handler(srv IdentityHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationIdentityLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResp)
		return ctx.Result(200, reply)
	}
}

type IdentityHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResp, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutResp, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResp, err error)
}

type IdentityHTTPClientImpl struct {
	cc *http.Client
}

func NewIdentityHTTPClient(client *http.Client) IdentityHTTPClient {
	return &IdentityHTTPClientImpl{client}
}

func (c *IdentityHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/api/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationIdentityLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *IdentityHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutResp, error) {
	var out LogoutResp
	pattern := "/api/v1/user/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationIdentityLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *IdentityHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResp, error) {
	var out RegisterResp
	pattern := "/api/v1/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationIdentityRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}