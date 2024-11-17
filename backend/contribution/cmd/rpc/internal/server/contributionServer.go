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

func (s *ContributionServer) UpdateContributionOfUser(ctx context.Context, in *pb.UpdateContributionOfUserReq) (*pb.UpdateContributionOfUserResp, error) {
	l := logic.NewUpdateContributionOfUserLogic(ctx, s.svcCtx)
	return l.UpdateContributionOfUser(in)
}

func (s *ContributionServer) UpdateIssuePROfUser(ctx context.Context, in *pb.UpdateIssuePROfUserReq) (*pb.UpdateIssuePROfUserResp, error) {
	l := logic.NewUpdateIssuePROfUserLogic(ctx, s.svcCtx)
	return l.UpdateIssuePROfUser(in)
}

func (s *ContributionServer) UpdateCommentOfUser(ctx context.Context, in *pb.UpdateCommentOfUserReq) (*pb.UpdateCommentOfUserResp, error) {
	l := logic.NewUpdateCommentOfUserLogic(ctx, s.svcCtx)
	return l.UpdateCommentOfUser(in)
}

func (s *ContributionServer) UpdateReviewOfUser(ctx context.Context, in *pb.UpdateReviewOfUserReq) (*pb.UpdateReviewOfUserResp, error) {
	l := logic.NewUpdateReviewOfUserLogic(ctx, s.svcCtx)
	return l.UpdateReviewOfUser(in)
}

func (s *ContributionServer) GetIssuePROfUserUpdatedAt(ctx context.Context, in *pb.GetIssuePROfUserUpdatedAtReq) (*pb.GetIssuePROfUserUpdatedAtResp, error) {
	l := logic.NewGetIssuePROfUserUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetIssuePROfUserUpdatedAt(in)
}

func (s *ContributionServer) GetCommentOfUserUpdatedAt(ctx context.Context, in *pb.GetCommentOfUserUpdatedAtReq) (*pb.GetCommentOfUserUpdatedAtResp, error) {
	l := logic.NewGetCommentOfUserUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetCommentOfUserUpdatedAt(in)
}

func (s *ContributionServer) GetReviewOfUserUpdatedAt(ctx context.Context, in *pb.GetReviewOfUserUpdatedAtReq) (*pb.GetReviewOfUserUpdatedAtResp, error) {
	l := logic.NewGetReviewOfUserUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetReviewOfUserUpdatedAt(in)
}
