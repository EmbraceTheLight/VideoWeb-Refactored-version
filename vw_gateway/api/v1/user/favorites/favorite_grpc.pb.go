// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/user/favorites/favorite.proto

package favorv1

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
	Favorite_CreateFavorites_FullMethodName = "/gateway.api.v1.user.favorites.Favorite/CreateFavorites"
	Favorite_ModifyFavorites_FullMethodName = "/gateway.api.v1.user.favorites.Favorite/ModifyFavorites"
	Favorite_DeleteFavorites_FullMethodName = "/gateway.api.v1.user.favorites.Favorite/DeleteFavorites"
)

// FavoriteClient is the client API for Favorite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteClient interface {
	CreateFavorites(ctx context.Context, in *CreateFavoritesReq, opts ...grpc.CallOption) (*CreateFavoritesResp, error)
	ModifyFavorites(ctx context.Context, in *ModifyFavoritesReq, opts ...grpc.CallOption) (*ModifyFavoritesResp, error)
	DeleteFavorites(ctx context.Context, in *DeleteFavoritesReq, opts ...grpc.CallOption) (*DeleteFavoritesResp, error)
}

type favoriteClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteClient(cc grpc.ClientConnInterface) FavoriteClient {
	return &favoriteClient{cc}
}

func (c *favoriteClient) CreateFavorites(ctx context.Context, in *CreateFavoritesReq, opts ...grpc.CallOption) (*CreateFavoritesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFavoritesResp)
	err := c.cc.Invoke(ctx, Favorite_CreateFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) ModifyFavorites(ctx context.Context, in *ModifyFavoritesReq, opts ...grpc.CallOption) (*ModifyFavoritesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyFavoritesResp)
	err := c.cc.Invoke(ctx, Favorite_ModifyFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) DeleteFavorites(ctx context.Context, in *DeleteFavoritesReq, opts ...grpc.CallOption) (*DeleteFavoritesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFavoritesResp)
	err := c.cc.Invoke(ctx, Favorite_DeleteFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServer is the server API for Favorite service.
// All implementations must embed UnimplementedFavoriteServer
// for forward compatibility.
type FavoriteServer interface {
	CreateFavorites(context.Context, *CreateFavoritesReq) (*CreateFavoritesResp, error)
	ModifyFavorites(context.Context, *ModifyFavoritesReq) (*ModifyFavoritesResp, error)
	DeleteFavorites(context.Context, *DeleteFavoritesReq) (*DeleteFavoritesResp, error)
	mustEmbedUnimplementedFavoriteServer()
}

// UnimplementedFavoriteServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFavoriteServer struct{}

func (UnimplementedFavoriteServer) CreateFavorites(context.Context, *CreateFavoritesReq) (*CreateFavoritesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFavorites not implemented")
}
func (UnimplementedFavoriteServer) ModifyFavorites(context.Context, *ModifyFavoritesReq) (*ModifyFavoritesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyFavorites not implemented")
}
func (UnimplementedFavoriteServer) DeleteFavorites(context.Context, *DeleteFavoritesReq) (*DeleteFavoritesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorites not implemented")
}
func (UnimplementedFavoriteServer) mustEmbedUnimplementedFavoriteServer() {}
func (UnimplementedFavoriteServer) testEmbeddedByValue()                  {}

// UnsafeFavoriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServer will
// result in compilation errors.
type UnsafeFavoriteServer interface {
	mustEmbedUnimplementedFavoriteServer()
}

func RegisterFavoriteServer(s grpc.ServiceRegistrar, srv FavoriteServer) {
	// If the following call pancis, it indicates UnimplementedFavoriteServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Favorite_ServiceDesc, srv)
}

func _Favorite_CreateFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFavoritesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).CreateFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_CreateFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).CreateFavorites(ctx, req.(*CreateFavoritesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_ModifyFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyFavoritesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).ModifyFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_ModifyFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).ModifyFavorites(ctx, req.(*ModifyFavoritesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_DeleteFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFavoritesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).DeleteFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_DeleteFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).DeleteFavorites(ctx, req.(*DeleteFavoritesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Favorite_ServiceDesc is the grpc.ServiceDesc for Favorite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Favorite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.api.v1.user.favorites.Favorite",
	HandlerType: (*FavoriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFavorites",
			Handler:    _Favorite_CreateFavorites_Handler,
		},
		{
			MethodName: "ModifyFavorites",
			Handler:    _Favorite_ModifyFavorites_Handler,
		},
		{
			MethodName: "DeleteFavorites",
			Handler:    _Favorite_DeleteFavorites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user/favorites/favorite.proto",
}
