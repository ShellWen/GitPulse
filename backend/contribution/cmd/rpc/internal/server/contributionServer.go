// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: contribution.proto

package server

import (
	"context"

	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/logic"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"
)

type ContributionServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedContributionServer
}

func NewContributionServer(svcCtx *svc.ServiceContext) *ContributionServer {
	return &ContributionServer{
		svcCtx: svcCtx,
	}
}

// -----------------------contribution-----------------------
func (s *ContributionServer) AddContribution(ctx context.Context, in *pb.AddContributionReq) (*pb.AddContributionResp, error) {
	l := logic.NewAddContributionLogic(ctx, s.svcCtx)
	return l.AddContribution(in)
}

func (s *ContributionServer) UpdateContribution(ctx context.Context, in *pb.UpdateContributionReq) (*pb.UpdateContributionResp, error) {
	l := logic.NewUpdateContributionLogic(ctx, s.svcCtx)
	return l.UpdateContribution(in)
}

func (s *ContributionServer) DelContribution(ctx context.Context, in *pb.DelContributionReq) (*pb.DelContributionResp, error) {
	l := logic.NewDelContributionLogic(ctx, s.svcCtx)
	return l.DelContribution(in)
}

func (s *ContributionServer) DelAllContributionInCategoryByUserId(ctx context.Context, in *pb.DelAllContributionInCategoryByUserIdReq) (*pb.DelAllContributionInCategoryByUserIdResp, error) {
	l := logic.NewDelAllContributionInCategoryByUserIdLogic(ctx, s.svcCtx)
	return l.DelAllContributionInCategoryByUserId(in)
}

func (s *ContributionServer) GetContribution(ctx context.Context, in *pb.GetContributionReq) (*pb.GetContributionResp, error) {
	l := logic.NewGetContributionLogic(ctx, s.svcCtx)
	return l.GetContribution(in)
}

func (s *ContributionServer) SearchByCategory(ctx context.Context, in *pb.SearchByCategoryReq) (*pb.SearchByCategoryResp, error) {
	l := logic.NewSearchByCategoryLogic(ctx, s.svcCtx)
	return l.SearchByCategory(in)
}

func (s *ContributionServer) SearchByUserId(ctx context.Context, in *pb.SearchByUserIdReq) (*pb.SearchByUserIdResp, error) {
	l := logic.NewSearchByUserIdLogic(ctx, s.svcCtx)
	return l.SearchByUserId(in)
}

func (s *ContributionServer) SearchByRepoId(ctx context.Context, in *pb.SearchByRepoIdReq) (*pb.SearchByRepoIdResp, error) {
	l := logic.NewSearchByRepoIdLogic(ctx, s.svcCtx)
	return l.SearchByRepoId(in)
}

func (s *ContributionServer) BlockUntilIssuePrOfUserUpdated(ctx context.Context, in *pb.BlockUntilIssuePrOfUserUpdatedReq) (*pb.BlockUntilIssuePrOfUserUpdatedResp, error) {
	l := logic.NewBlockUntilIssuePrOfUserUpdatedLogic(ctx, s.svcCtx)
	return l.BlockUntilIssuePrOfUserUpdated(in)
}

func (s *ContributionServer) BlockUntilCommentReviewOfUserUpdated(ctx context.Context, in *pb.BlockUntilCommentReviewOfUserUpdatedReq) (*pb.BlockUntilCommentReviewOfUserUpdatedResp, error) {
	l := logic.NewBlockUntilCommentReviewOfUserUpdatedLogic(ctx, s.svcCtx)
	return l.BlockUntilCommentReviewOfUserUpdated(in)
}

func (s *ContributionServer) BlockUntilAllUpdated(ctx context.Context, in *pb.BlockUntilAllUpdatedReq) (*pb.BlockUntilAllUpdatedResp, error) {
	l := logic.NewBlockUntilAllUpdatedLogic(ctx, s.svcCtx)
	return l.BlockUntilAllUpdated(in)
}

func (s *ContributionServer) UnblockContribution(ctx context.Context, in *pb.UnblockContributionReq) (*pb.UnblockContributionResp, error) {
	l := logic.NewUnblockContributionLogic(ctx, s.svcCtx)
	return l.UnblockContribution(in)
}
