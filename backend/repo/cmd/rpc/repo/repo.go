// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: repo.proto

package repo

import (
	"context"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRepoReq                = pb.AddRepoReq
	AddRepoResp               = pb.AddRepoResp
	BlockUntilRepoUpdatedReq  = pb.BlockUntilRepoUpdatedReq
	BlockUntilRepoUpdatedResp = pb.BlockUntilRepoUpdatedResp
	DelRepoByIdReq            = pb.DelRepoByIdReq
	DelRepoByIdResp           = pb.DelRepoByIdResp
	GetRepoByIdReq            = pb.GetRepoByIdReq
	GetRepoByIdResp           = pb.GetRepoByIdResp
	Repo                      = pb.Repo
	UpdateRepoReq             = pb.UpdateRepoReq
	UpdateRepoResp            = pb.UpdateRepoResp

	RepoZrpcClient interface {
		// -----------------------repo-----------------------
		AddRepo(ctx context.Context, in *AddRepoReq, opts ...grpc.CallOption) (*AddRepoResp, error)
		UpdateRepo(ctx context.Context, in *UpdateRepoReq, opts ...grpc.CallOption) (*UpdateRepoResp, error)
		DelRepoById(ctx context.Context, in *DelRepoByIdReq, opts ...grpc.CallOption) (*DelRepoByIdResp, error)
		GetRepoById(ctx context.Context, in *GetRepoByIdReq, opts ...grpc.CallOption) (*GetRepoByIdResp, error)
		BlockUntilRepoUpdated(ctx context.Context, in *BlockUntilRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilRepoUpdatedResp, error)
	}

	defaultRepoZrpcClient struct {
		cli zrpc.Client
	}
)

func NewRepoZrpcClient(cli zrpc.Client) RepoZrpcClient {
	return &defaultRepoZrpcClient{
		cli: cli,
	}
}

// -----------------------repo-----------------------
func (m *defaultRepoZrpcClient) AddRepo(ctx context.Context, in *AddRepoReq, opts ...grpc.CallOption) (*AddRepoResp, error) {
	client := pb.NewRepoClient(m.cli.Conn())
	return client.AddRepo(ctx, in, opts...)
}

func (m *defaultRepoZrpcClient) UpdateRepo(ctx context.Context, in *UpdateRepoReq, opts ...grpc.CallOption) (*UpdateRepoResp, error) {
	client := pb.NewRepoClient(m.cli.Conn())
	return client.UpdateRepo(ctx, in, opts...)
}

func (m *defaultRepoZrpcClient) DelRepoById(ctx context.Context, in *DelRepoByIdReq, opts ...grpc.CallOption) (*DelRepoByIdResp, error) {
	client := pb.NewRepoClient(m.cli.Conn())
	return client.DelRepoById(ctx, in, opts...)
}

func (m *defaultRepoZrpcClient) GetRepoById(ctx context.Context, in *GetRepoByIdReq, opts ...grpc.CallOption) (*GetRepoByIdResp, error) {
	client := pb.NewRepoClient(m.cli.Conn())
	return client.GetRepoById(ctx, in, opts...)
}

func (m *defaultRepoZrpcClient) BlockUntilRepoUpdated(ctx context.Context, in *BlockUntilRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilRepoUpdatedResp, error) {
	client := pb.NewRepoClient(m.cli.Conn())
	return client.BlockUntilRepoUpdated(ctx, in, opts...)
}
