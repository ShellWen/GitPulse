// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: relation.proto

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
	Relation_AddCreateRepo_FullMethodName               = "/pb.relation/AddCreateRepo"
	Relation_DelCreateRepo_FullMethodName               = "/pb.relation/DelCreateRepo"
	Relation_GetCreatorId_FullMethodName                = "/pb.relation/GetCreatorId"
	Relation_SearchCreatedRepo_FullMethodName           = "/pb.relation/SearchCreatedRepo"
	Relation_AddFollow_FullMethodName                   = "/pb.relation/AddFollow"
	Relation_DelFollow_FullMethodName                   = "/pb.relation/DelFollow"
	Relation_CheckIfFollow_FullMethodName               = "/pb.relation/CheckIfFollow"
	Relation_SearchFollowedByFollowingId_FullMethodName = "/pb.relation/SearchFollowedByFollowingId"
	Relation_SearchFollowingByFollowedId_FullMethodName = "/pb.relation/SearchFollowingByFollowedId"
	Relation_AddFork_FullMethodName                     = "/pb.relation/AddFork"
	Relation_DelFork_FullMethodName                     = "/pb.relation/DelFork"
	Relation_GetOrigin_FullMethodName                   = "/pb.relation/GetOrigin"
	Relation_SearchFork_FullMethodName                  = "/pb.relation/SearchFork"
	Relation_AddStar_FullMethodName                     = "/pb.relation/AddStar"
	Relation_DelStar_FullMethodName                     = "/pb.relation/DelStar"
	Relation_CheckIfStar_FullMethodName                 = "/pb.relation/CheckIfStar"
	Relation_SearchStaredRepo_FullMethodName            = "/pb.relation/SearchStaredRepo"
	Relation_SearchStaringDev_FullMethodName            = "/pb.relation/SearchStaringDev"
)

// RelationClient is the client API for Relation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationClient interface {
	// -----------------------createRepo-----------------------
	AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error)
	DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error)
	GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error)
	SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error)
	// -----------------------follow-----------------------
	AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error)
	DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error)
	CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error)
	SearchFollowedByFollowingId(ctx context.Context, in *SearchFollowedByFollowingIdReq, opts ...grpc.CallOption) (*SearchFollowByFollowingIdResp, error)
	SearchFollowingByFollowedId(ctx context.Context, in *SearchFollowingByFollowedIdReq, opts ...grpc.CallOption) (*SearchFollowByFollowedIdResp, error)
	// -----------------------fork-----------------------
	AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error)
	DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error)
	GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error)
	SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error)
	// -----------------------star-----------------------
	AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error)
	DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error)
	CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error)
	SearchStaredRepo(ctx context.Context, in *SearchStaredRepoReq, opts ...grpc.CallOption) (*SearchStaredRepoResp, error)
	SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error)
}

type relationClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationClient(cc grpc.ClientConnInterface) RelationClient {
	return &relationClient{cc}
}

func (c *relationClient) AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCreateRepoResp)
	err := c.cc.Invoke(ctx, Relation_AddCreateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelCreateRepoResp)
	err := c.cc.Invoke(ctx, Relation_DelCreateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCreatorIdResp)
	err := c.cc.Invoke(ctx, Relation_GetCreatorId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCreatedRepoResp)
	err := c.cc.Invoke(ctx, Relation_SearchCreatedRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddFollowResp)
	err := c.cc.Invoke(ctx, Relation_AddFollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelFollowResp)
	err := c.cc.Invoke(ctx, Relation_DelFollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckFollowResp)
	err := c.cc.Invoke(ctx, Relation_CheckIfFollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchFollowedByFollowingId(ctx context.Context, in *SearchFollowedByFollowingIdReq, opts ...grpc.CallOption) (*SearchFollowByFollowingIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchFollowByFollowingIdResp)
	err := c.cc.Invoke(ctx, Relation_SearchFollowedByFollowingId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchFollowingByFollowedId(ctx context.Context, in *SearchFollowingByFollowedIdReq, opts ...grpc.CallOption) (*SearchFollowByFollowedIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchFollowByFollowedIdResp)
	err := c.cc.Invoke(ctx, Relation_SearchFollowingByFollowedId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddForkResp)
	err := c.cc.Invoke(ctx, Relation_AddFork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelForkResp)
	err := c.cc.Invoke(ctx, Relation_DelFork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOriginResp)
	err := c.cc.Invoke(ctx, Relation_GetOrigin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchForkResp)
	err := c.cc.Invoke(ctx, Relation_SearchFork_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddStarResp)
	err := c.cc.Invoke(ctx, Relation_AddStar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelStarResp)
	err := c.cc.Invoke(ctx, Relation_DelStar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckIfStarResp)
	err := c.cc.Invoke(ctx, Relation_CheckIfStar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchStaredRepo(ctx context.Context, in *SearchStaredRepoReq, opts ...grpc.CallOption) (*SearchStaredRepoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchStaredRepoResp)
	err := c.cc.Invoke(ctx, Relation_SearchStaredRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchStaringDevResp)
	err := c.cc.Invoke(ctx, Relation_SearchStaringDev_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServer is the server API for Relation service.
// All implementations must embed UnimplementedRelationServer
// for forward compatibility.
type RelationServer interface {
	// -----------------------createRepo-----------------------
	AddCreateRepo(context.Context, *AddCreateRepoReq) (*AddCreateRepoResp, error)
	DelCreateRepo(context.Context, *DelCreateRepoReq) (*DelCreateRepoResp, error)
	GetCreatorId(context.Context, *GetCreatorIdReq) (*GetCreatorIdResp, error)
	SearchCreatedRepo(context.Context, *SearchCreatedRepoReq) (*SearchCreatedRepoResp, error)
	// -----------------------follow-----------------------
	AddFollow(context.Context, *AddFollowReq) (*AddFollowResp, error)
	DelFollow(context.Context, *DelFollowReq) (*DelFollowResp, error)
	CheckIfFollow(context.Context, *CheckIfFollowReq) (*CheckFollowResp, error)
	SearchFollowedByFollowingId(context.Context, *SearchFollowedByFollowingIdReq) (*SearchFollowByFollowingIdResp, error)
	SearchFollowingByFollowedId(context.Context, *SearchFollowingByFollowedIdReq) (*SearchFollowByFollowedIdResp, error)
	// -----------------------fork-----------------------
	AddFork(context.Context, *AddForkReq) (*AddForkResp, error)
	DelFork(context.Context, *DelForkReq) (*DelForkResp, error)
	GetOrigin(context.Context, *GetOriginReq) (*GetOriginResp, error)
	SearchFork(context.Context, *SearchForkReq) (*SearchForkResp, error)
	// -----------------------star-----------------------
	AddStar(context.Context, *AddStarReq) (*AddStarResp, error)
	DelStar(context.Context, *DelStarReq) (*DelStarResp, error)
	CheckIfStar(context.Context, *CheckIfStarReq) (*CheckIfStarResp, error)
	SearchStaredRepo(context.Context, *SearchStaredRepoReq) (*SearchStaredRepoResp, error)
	SearchStaringDev(context.Context, *SearchStaringDevReq) (*SearchStaringDevResp, error)
	mustEmbedUnimplementedRelationServer()
}

// UnimplementedRelationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRelationServer struct{}

func (UnimplementedRelationServer) AddCreateRepo(context.Context, *AddCreateRepoReq) (*AddCreateRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCreateRepo not implemented")
}
func (UnimplementedRelationServer) DelCreateRepo(context.Context, *DelCreateRepoReq) (*DelCreateRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelCreateRepo not implemented")
}
func (UnimplementedRelationServer) GetCreatorId(context.Context, *GetCreatorIdReq) (*GetCreatorIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreatorId not implemented")
}
func (UnimplementedRelationServer) SearchCreatedRepo(context.Context, *SearchCreatedRepoReq) (*SearchCreatedRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCreatedRepo not implemented")
}
func (UnimplementedRelationServer) AddFollow(context.Context, *AddFollowReq) (*AddFollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollow not implemented")
}
func (UnimplementedRelationServer) DelFollow(context.Context, *DelFollowReq) (*DelFollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFollow not implemented")
}
func (UnimplementedRelationServer) CheckIfFollow(context.Context, *CheckIfFollowReq) (*CheckFollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfFollow not implemented")
}
func (UnimplementedRelationServer) SearchFollowedByFollowingId(context.Context, *SearchFollowedByFollowingIdReq) (*SearchFollowByFollowingIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFollowedByFollowingId not implemented")
}
func (UnimplementedRelationServer) SearchFollowingByFollowedId(context.Context, *SearchFollowingByFollowedIdReq) (*SearchFollowByFollowedIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFollowingByFollowedId not implemented")
}
func (UnimplementedRelationServer) AddFork(context.Context, *AddForkReq) (*AddForkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFork not implemented")
}
func (UnimplementedRelationServer) DelFork(context.Context, *DelForkReq) (*DelForkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFork not implemented")
}
func (UnimplementedRelationServer) GetOrigin(context.Context, *GetOriginReq) (*GetOriginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrigin not implemented")
}
func (UnimplementedRelationServer) SearchFork(context.Context, *SearchForkReq) (*SearchForkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFork not implemented")
}
func (UnimplementedRelationServer) AddStar(context.Context, *AddStarReq) (*AddStarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStar not implemented")
}
func (UnimplementedRelationServer) DelStar(context.Context, *DelStarReq) (*DelStarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelStar not implemented")
}
func (UnimplementedRelationServer) CheckIfStar(context.Context, *CheckIfStarReq) (*CheckIfStarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfStar not implemented")
}
func (UnimplementedRelationServer) SearchStaredRepo(context.Context, *SearchStaredRepoReq) (*SearchStaredRepoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchStaredRepo not implemented")
}
func (UnimplementedRelationServer) SearchStaringDev(context.Context, *SearchStaringDevReq) (*SearchStaringDevResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchStaringDev not implemented")
}
func (UnimplementedRelationServer) mustEmbedUnimplementedRelationServer() {}
func (UnimplementedRelationServer) testEmbeddedByValue()                  {}

// UnsafeRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServer will
// result in compilation errors.
type UnsafeRelationServer interface {
	mustEmbedUnimplementedRelationServer()
}

func RegisterRelationServer(s grpc.ServiceRegistrar, srv RelationServer) {
	// If the following call pancis, it indicates UnimplementedRelationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Relation_ServiceDesc, srv)
}

func _Relation_AddCreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCreateRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).AddCreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_AddCreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).AddCreateRepo(ctx, req.(*AddCreateRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_DelCreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCreateRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).DelCreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_DelCreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).DelCreateRepo(ctx, req.(*DelCreateRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetCreatorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreatorIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetCreatorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_GetCreatorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetCreatorId(ctx, req.(*GetCreatorIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchCreatedRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCreatedRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchCreatedRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchCreatedRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchCreatedRepo(ctx, req.(*SearchCreatedRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_AddFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).AddFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_AddFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).AddFollow(ctx, req.(*AddFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_DelFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).DelFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_DelFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).DelFollow(ctx, req.(*DelFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_CheckIfFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).CheckIfFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_CheckIfFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).CheckIfFollow(ctx, req.(*CheckIfFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchFollowedByFollowingId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFollowedByFollowingIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchFollowedByFollowingId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchFollowedByFollowingId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchFollowedByFollowingId(ctx, req.(*SearchFollowedByFollowingIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchFollowingByFollowedId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFollowingByFollowedIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchFollowingByFollowedId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchFollowingByFollowedId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchFollowingByFollowedId(ctx, req.(*SearchFollowingByFollowedIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_AddFork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddForkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).AddFork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_AddFork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).AddFork(ctx, req.(*AddForkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_DelFork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelForkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).DelFork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_DelFork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).DelFork(ctx, req.(*DelForkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_GetOrigin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetOrigin(ctx, req.(*GetOriginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchFork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchForkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchFork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchFork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchFork(ctx, req.(*SearchForkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_AddStar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).AddStar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_AddStar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).AddStar(ctx, req.(*AddStarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_DelStar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelStarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).DelStar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_DelStar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).DelStar(ctx, req.(*DelStarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_CheckIfStar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfStarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).CheckIfStar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_CheckIfStar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).CheckIfStar(ctx, req.(*CheckIfStarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchStaredRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStaredRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchStaredRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchStaredRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchStaredRepo(ctx, req.(*SearchStaredRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_SearchStaringDev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStaringDevReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).SearchStaringDev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_SearchStaringDev_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).SearchStaringDev(ctx, req.(*SearchStaringDevReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Relation_ServiceDesc is the grpc.ServiceDesc for Relation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.relation",
	HandlerType: (*RelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCreateRepo",
			Handler:    _Relation_AddCreateRepo_Handler,
		},
		{
			MethodName: "DelCreateRepo",
			Handler:    _Relation_DelCreateRepo_Handler,
		},
		{
			MethodName: "GetCreatorId",
			Handler:    _Relation_GetCreatorId_Handler,
		},
		{
			MethodName: "SearchCreatedRepo",
			Handler:    _Relation_SearchCreatedRepo_Handler,
		},
		{
			MethodName: "AddFollow",
			Handler:    _Relation_AddFollow_Handler,
		},
		{
			MethodName: "DelFollow",
			Handler:    _Relation_DelFollow_Handler,
		},
		{
			MethodName: "CheckIfFollow",
			Handler:    _Relation_CheckIfFollow_Handler,
		},
		{
			MethodName: "SearchFollowedByFollowingId",
			Handler:    _Relation_SearchFollowedByFollowingId_Handler,
		},
		{
			MethodName: "SearchFollowingByFollowedId",
			Handler:    _Relation_SearchFollowingByFollowedId_Handler,
		},
		{
			MethodName: "AddFork",
			Handler:    _Relation_AddFork_Handler,
		},
		{
			MethodName: "DelFork",
			Handler:    _Relation_DelFork_Handler,
		},
		{
			MethodName: "GetOrigin",
			Handler:    _Relation_GetOrigin_Handler,
		},
		{
			MethodName: "SearchFork",
			Handler:    _Relation_SearchFork_Handler,
		},
		{
			MethodName: "AddStar",
			Handler:    _Relation_AddStar_Handler,
		},
		{
			MethodName: "DelStar",
			Handler:    _Relation_DelStar_Handler,
		},
		{
			MethodName: "CheckIfStar",
			Handler:    _Relation_CheckIfStar_Handler,
		},
		{
			MethodName: "SearchStaredRepo",
			Handler:    _Relation_SearchStaredRepo_Handler,
		},
		{
			MethodName: "SearchStaringDev",
			Handler:    _Relation_SearchStaringDev_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}
