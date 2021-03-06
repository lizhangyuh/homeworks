// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.2
// source: v1/mobile.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MobileClient is the client API for Mobile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileClient interface {
	// Auth
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignInByPwd(ctx context.Context, in *SignInByPwdRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// User
	CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...grpc.CallOption) (*CurrentUserResponse, error)
	// Post
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostReply, error)
	PostIndex(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostReply, error)
	PostShow(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostReply, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostReply, error)
}

type mobileClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileClient(cc grpc.ClientConnInterface) MobileClient {
	return &mobileClient{cc}
}

func (c *mobileClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) SignInByPwd(ctx context.Context, in *SignInByPwdRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/SignInByPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...grpc.CallOption) (*CurrentUserResponse, error) {
	out := new(CurrentUserResponse)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/CurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostReply, error) {
	out := new(CreatePostReply)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) PostIndex(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostReply, error) {
	out := new(ListPostReply)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/PostIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) PostShow(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostReply, error) {
	out := new(GetPostReply)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/PostShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostReply, error) {
	out := new(UpdatePostReply)
	err := c.cc.Invoke(ctx, "/bff.mobile.v1.Mobile/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServer is the server API for Mobile service.
// All implementations must embed UnimplementedMobileServer
// for forward compatibility
type MobileServer interface {
	// Auth
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignInByPwd(context.Context, *SignInByPwdRequest) (*SignInResponse, error)
	// User
	CurrentUser(context.Context, *CurrentUserRequest) (*CurrentUserResponse, error)
	// Post
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error)
	PostIndex(context.Context, *ListPostRequest) (*ListPostReply, error)
	PostShow(context.Context, *GetPostRequest) (*GetPostReply, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error)
	mustEmbedUnimplementedMobileServer()
}

// UnimplementedMobileServer must be embedded to have forward compatible implementations.
type UnimplementedMobileServer struct {
}

func (UnimplementedMobileServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedMobileServer) SignInByPwd(context.Context, *SignInByPwdRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInByPwd not implemented")
}
func (UnimplementedMobileServer) CurrentUser(context.Context, *CurrentUserRequest) (*CurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentUser not implemented")
}
func (UnimplementedMobileServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedMobileServer) PostIndex(context.Context, *ListPostRequest) (*ListPostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostIndex not implemented")
}
func (UnimplementedMobileServer) PostShow(context.Context, *GetPostRequest) (*GetPostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostShow not implemented")
}
func (UnimplementedMobileServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedMobileServer) mustEmbedUnimplementedMobileServer() {}

// UnsafeMobileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileServer will
// result in compilation errors.
type UnsafeMobileServer interface {
	mustEmbedUnimplementedMobileServer()
}

func RegisterMobileServer(s grpc.ServiceRegistrar, srv MobileServer) {
	s.RegisterService(&Mobile_ServiceDesc, srv)
}

func _Mobile_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_SignInByPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInByPwdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).SignInByPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/SignInByPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).SignInByPwd(ctx, req.(*SignInByPwdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_CurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).CurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/CurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).CurrentUser(ctx, req.(*CurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_PostIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).PostIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/PostIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).PostIndex(ctx, req.(*ListPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_PostShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).PostShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/PostShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).PostShow(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.mobile.v1.Mobile/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mobile_ServiceDesc is the grpc.ServiceDesc for Mobile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mobile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.mobile.v1.Mobile",
	HandlerType: (*MobileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Mobile_SignUp_Handler,
		},
		{
			MethodName: "SignInByPwd",
			Handler:    _Mobile_SignInByPwd_Handler,
		},
		{
			MethodName: "CurrentUser",
			Handler:    _Mobile_CurrentUser_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Mobile_CreatePost_Handler,
		},
		{
			MethodName: "PostIndex",
			Handler:    _Mobile_PostIndex_Handler,
		},
		{
			MethodName: "PostShow",
			Handler:    _Mobile_PostShow_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Mobile_UpdatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/mobile.proto",
}
