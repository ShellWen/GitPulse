// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: repo.proto

package pb

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
	Repo_AddRepo_FullMethodName     = "/pb.repo/AddRepo"
	Repo_UpdateRepo_FullMethodName  = "/pb.repo/UpdateRepo"
	Repo_DelRepoById_FullMethodName = "/pb.repo/DelRepoById"
	Repo_GetRepoById_FullMethodName = "/pb.repo/GetRepoById"
	Repo_SearchRepo_FullMethodName  = "/pb.repo/SearchRepo"
)

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoClient interface {
	// -----------------------repo-----------------------
	AddRepo(ctx context.Context, in *AddRepoReq, opts ...grpc.CallOption) (*AddRepoResp, error)
	UpdateRepo(ctx context.Context, in *UpdateRepoReq, opts ...grpc.CallOption) (*UpdateRepoResp, error)
	DelRepoById(ctx context.Context, in *DelRepoByIdReq, opts ...grpc.CallOption) (*DelRepoByIdResp, error)
	GetRepoById(ctx context.Context, in *GetRepoByIdReq, opts ...grpc.CallOption) (*GetRepoByIdResp, error)
	SearchRepo(ctx context.Context, in *SearchRepoReq, opts ...grpc.CallOption) (*SearchRepoResp, error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) AddRepo(ctx context.Context, in *AddRepoReq, opts ...grpc.CallOption) (*AddRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddRepoResp)
	err := c.cc.Invoke(ctx, Repo_AddRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) UpdateRepo(ctx context.Context, in *UpdateRepoReq, opts ...grpc.CallOption) (*UpdateRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRepoResp)
	err := c.cc.Invoke(ctx, Repo_UpdateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) DelRepoById(ctx context.Context, in *DelRepoByIdReq, opts ...grpc.CallOption) (*DelRepoByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelRepoByIdResp)
	err := c.cc.Invoke(ctx, Repo_DelRepoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) GetRepoById(ctx context.Context, in *GetRepoByIdReq, opts ...grpc.CallOption) (*GetRepoByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRepoByIdResp)
	err := c.cc.Invoke(ctx, Repo_GetRepoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) SearchRepo(ctx context.Context, in *SearchRepoReq, opts ...grpc.CallOption) (*SearchRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchRepoResp)
	err := c.cc.Invoke(ctx, Repo_SearchRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServer is the server API for Repo service.
// All implementations must embed UnimplementedRepoServer
// for forward compatibility.
type RepoServer interface {
	// -----------------------repo-----------------------
	AddRepo(context.Context, *AddRepoReq) (*AddRepoResp, error)
	UpdateRepo(context.Context, *UpdateRepoReq) (*UpdateRepoResp, error)
	DelRepoById(context.Context, *DelRepoByIdReq) (*DelRepoByIdResp, error)
	GetRepoById(context.Context, *GetRepoByIdReq) (*GetRepoByIdResp, error)
	SearchRepo(context.Context, *SearchRepoReq) (*SearchRepoResp, error)
	mustEmbedUnimplementedRepoServer()
}

// UnimplementedRepoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRepoServer struct{}

func (UnimplementedRepoServer) AddRepo(context.Context, *AddRepoReq) (*AddRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRepo not implemented")
}
func (UnimplementedRepoServer) UpdateRepo(context.Context, *UpdateRepoReq) (*UpdateRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedRepoServer) DelRepoById(context.Context, *DelRepoByIdReq) (*DelRepoByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRepoById not implemented")
}
func (UnimplementedRepoServer) GetRepoById(context.Context, *GetRepoByIdReq) (*GetRepoByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepoById not implemented")
}
func (UnimplementedRepoServer) SearchRepo(context.Context, *SearchRepoReq) (*SearchRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRepo not implemented")
}
func (UnimplementedRepoServer) mustEmbedUnimplementedRepoServer() {}
func (UnimplementedRepoServer) testEmbeddedByValue()              {}

// UnsafeRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoServer will
// result in compilation errors.
type UnsafeRepoServer interface {
	mustEmbedUnimplementedRepoServer()
}

func RegisterRepoServer(s grpc.ServiceRegistrar, srv RepoServer) {
	// If the following call pancis, it indicates UnimplementedRepoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Repo_ServiceDesc, srv)
}

func _Repo_AddRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).AddRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_AddRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).AddRepo(ctx, req.(*AddRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_UpdateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).UpdateRepo(ctx, req.(*UpdateRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_DelRepoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRepoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).DelRepoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_DelRepoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).DelRepoById(ctx, req.(*DelRepoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_GetRepoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).GetRepoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_GetRepoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).GetRepoById(ctx, req.(*GetRepoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_SearchRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).SearchRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repo_SearchRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).SearchRepo(ctx, req.(*SearchRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Repo_ServiceDesc is the grpc.ServiceDesc for Repo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.repo",
	HandlerType: (*RepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRepo",
			Handler:    _Repo_AddRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _Repo_UpdateRepo_Handler,
		},
		{
			MethodName: "DelRepoById",
			Handler:    _Repo_DelRepoById_Handler,
		},
		{
			MethodName: "GetRepoById",
			Handler:    _Repo_GetRepoById_Handler,
		},
		{
			MethodName: "SearchRepo",
			Handler:    _Repo_SearchRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo.proto",
}
