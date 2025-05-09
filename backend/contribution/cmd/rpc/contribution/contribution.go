// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: contribution.proto

package contribution

import (
	"context"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddContributionReq                       = pb.AddContributionReq
	AddContributionResp                      = pb.AddContributionResp
	Contribution                             = pb.Contribution
	DelAllContributionInCategoryByUserIdReq  = pb.DelAllContributionInCategoryByUserIdReq
	DelAllContributionInCategoryByUserIdResp = pb.DelAllContributionInCategoryByUserIdResp
	DelContributionReq                       = pb.DelContributionReq
	DelContributionResp                      = pb.DelContributionResp
	GetCommentOfUserUpdatedAtReq             = pb.GetCommentOfUserUpdatedAtReq
	GetCommentOfUserUpdatedAtResp            = pb.GetCommentOfUserUpdatedAtResp
	GetContributionReq                       = pb.GetContributionReq
	GetContributionResp                      = pb.GetContributionResp
	GetIssuePROfUserUpdatedAtReq             = pb.GetIssuePROfUserUpdatedAtReq
	GetIssuePROfUserUpdatedAtResp            = pb.GetIssuePROfUserUpdatedAtResp
	GetReviewOfUserUpdatedAtReq              = pb.GetReviewOfUserUpdatedAtReq
	GetReviewOfUserUpdatedAtResp             = pb.GetReviewOfUserUpdatedAtResp
	SearchByCategoryReq                      = pb.SearchByCategoryReq
	SearchByCategoryResp                     = pb.SearchByCategoryResp
	SearchByRepoIdReq                        = pb.SearchByRepoIdReq
	SearchByRepoIdResp                       = pb.SearchByRepoIdResp
	SearchByUserIdReq                        = pb.SearchByUserIdReq
	SearchByUserIdResp                       = pb.SearchByUserIdResp
	UpdateCommentOfUserReq                   = pb.UpdateCommentOfUserReq
	UpdateCommentOfUserResp                  = pb.UpdateCommentOfUserResp
	UpdateContributionOfUserReq              = pb.UpdateContributionOfUserReq
	UpdateContributionOfUserResp             = pb.UpdateContributionOfUserResp
	UpdateIssuePROfUserReq                   = pb.UpdateIssuePROfUserReq
	UpdateIssuePROfUserResp                  = pb.UpdateIssuePROfUserResp
	UpdateReviewOfUserReq                    = pb.UpdateReviewOfUserReq
	UpdateReviewOfUserResp                   = pb.UpdateReviewOfUserResp

	ContributionZrpcClient interface {
		// -----------------------contribution-----------------------
		AddContribution(ctx context.Context, in *AddContributionReq, opts ...grpc.CallOption) (*AddContributionResp, error)
		DelContribution(ctx context.Context, in *DelContributionReq, opts ...grpc.CallOption) (*DelContributionResp, error)
		DelAllContributionInCategoryByUserId(ctx context.Context, in *DelAllContributionInCategoryByUserIdReq, opts ...grpc.CallOption) (*DelAllContributionInCategoryByUserIdResp, error)
		GetContribution(ctx context.Context, in *GetContributionReq, opts ...grpc.CallOption) (*GetContributionResp, error)
		SearchByCategory(ctx context.Context, in *SearchByCategoryReq, opts ...grpc.CallOption) (*SearchByCategoryResp, error)
		SearchByUserId(ctx context.Context, in *SearchByUserIdReq, opts ...grpc.CallOption) (*SearchByUserIdResp, error)
		SearchByRepoId(ctx context.Context, in *SearchByRepoIdReq, opts ...grpc.CallOption) (*SearchByRepoIdResp, error)
		UpdateContributionOfUser(ctx context.Context, in *UpdateContributionOfUserReq, opts ...grpc.CallOption) (*UpdateContributionOfUserResp, error)
		UpdateIssuePROfUser(ctx context.Context, in *UpdateIssuePROfUserReq, opts ...grpc.CallOption) (*UpdateIssuePROfUserResp, error)
		UpdateCommentOfUser(ctx context.Context, in *UpdateCommentOfUserReq, opts ...grpc.CallOption) (*UpdateCommentOfUserResp, error)
		UpdateReviewOfUser(ctx context.Context, in *UpdateReviewOfUserReq, opts ...grpc.CallOption) (*UpdateReviewOfUserResp, error)
		GetIssuePROfUserUpdatedAt(ctx context.Context, in *GetIssuePROfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetIssuePROfUserUpdatedAtResp, error)
		GetCommentOfUserUpdatedAt(ctx context.Context, in *GetCommentOfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetCommentOfUserUpdatedAtResp, error)
		GetReviewOfUserUpdatedAt(ctx context.Context, in *GetReviewOfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetReviewOfUserUpdatedAtResp, error)
	}

	defaultContributionZrpcClient struct {
		cli zrpc.Client
	}
)

func NewContributionZrpcClient(cli zrpc.Client) ContributionZrpcClient {
	return &defaultContributionZrpcClient{
		cli: cli,
	}
}

// -----------------------contribution-----------------------
func (m *defaultContributionZrpcClient) AddContribution(ctx context.Context, in *AddContributionReq, opts ...grpc.CallOption) (*AddContributionResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.AddContribution(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) DelContribution(ctx context.Context, in *DelContributionReq, opts ...grpc.CallOption) (*DelContributionResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.DelContribution(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) DelAllContributionInCategoryByUserId(ctx context.Context, in *DelAllContributionInCategoryByUserIdReq, opts ...grpc.CallOption) (*DelAllContributionInCategoryByUserIdResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.DelAllContributionInCategoryByUserId(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) GetContribution(ctx context.Context, in *GetContributionReq, opts ...grpc.CallOption) (*GetContributionResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.GetContribution(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) SearchByCategory(ctx context.Context, in *SearchByCategoryReq, opts ...grpc.CallOption) (*SearchByCategoryResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.SearchByCategory(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) SearchByUserId(ctx context.Context, in *SearchByUserIdReq, opts ...grpc.CallOption) (*SearchByUserIdResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.SearchByUserId(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) SearchByRepoId(ctx context.Context, in *SearchByRepoIdReq, opts ...grpc.CallOption) (*SearchByRepoIdResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.SearchByRepoId(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) UpdateContributionOfUser(ctx context.Context, in *UpdateContributionOfUserReq, opts ...grpc.CallOption) (*UpdateContributionOfUserResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.UpdateContributionOfUser(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) UpdateIssuePROfUser(ctx context.Context, in *UpdateIssuePROfUserReq, opts ...grpc.CallOption) (*UpdateIssuePROfUserResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.UpdateIssuePROfUser(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) UpdateCommentOfUser(ctx context.Context, in *UpdateCommentOfUserReq, opts ...grpc.CallOption) (*UpdateCommentOfUserResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.UpdateCommentOfUser(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) UpdateReviewOfUser(ctx context.Context, in *UpdateReviewOfUserReq, opts ...grpc.CallOption) (*UpdateReviewOfUserResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.UpdateReviewOfUser(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) GetIssuePROfUserUpdatedAt(ctx context.Context, in *GetIssuePROfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetIssuePROfUserUpdatedAtResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.GetIssuePROfUserUpdatedAt(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) GetCommentOfUserUpdatedAt(ctx context.Context, in *GetCommentOfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetCommentOfUserUpdatedAtResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.GetCommentOfUserUpdatedAt(ctx, in, opts...)
}

func (m *defaultContributionZrpcClient) GetReviewOfUserUpdatedAt(ctx context.Context, in *GetReviewOfUserUpdatedAtReq, opts ...grpc.CallOption) (*GetReviewOfUserUpdatedAtResp, error) {
	client := pb.NewContributionClient(m.cli.Conn())
	return client.GetReviewOfUserUpdatedAt(ctx, in, opts...)
}
