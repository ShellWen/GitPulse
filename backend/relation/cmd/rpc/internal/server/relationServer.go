// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: relation.proto

package server

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/logic"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
)

type RelationServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRelationServer
}

func NewRelationServer(svcCtx *svc.ServiceContext) *RelationServer {
	return &RelationServer{
		svcCtx: svcCtx,
	}
}

// -----------------------createRepo-----------------------
func (s *RelationServer) AddCreateRepo(ctx context.Context, in *pb.AddCreateRepoReq) (*pb.AddCreateRepoResp, error) {
	l := logic.NewAddCreateRepoLogic(ctx, s.svcCtx)
	return l.AddCreateRepo(in)
}

func (s *RelationServer) DelCreateRepo(ctx context.Context, in *pb.DelCreateRepoReq) (*pb.DelCreateRepoResp, error) {
	l := logic.NewDelCreateRepoLogic(ctx, s.svcCtx)
	return l.DelCreateRepo(in)
}

func (s *RelationServer) GetCreatorId(ctx context.Context, in *pb.GetCreatorIdReq) (*pb.GetCreatorIdResp, error) {
	l := logic.NewGetCreatorIdLogic(ctx, s.svcCtx)
	return l.GetCreatorId(in)
}

func (s *RelationServer) SearchCreatedRepo(ctx context.Context, in *pb.SearchCreatedRepoReq) (*pb.SearchCreatedRepoResp, error) {
	l := logic.NewSearchCreatedRepoLogic(ctx, s.svcCtx)
	return l.SearchCreatedRepo(in)
}

// -----------------------follow-----------------------
func (s *RelationServer) AddFollow(ctx context.Context, in *pb.AddFollowReq) (*pb.AddFollowResp, error) {
	l := logic.NewAddFollowLogic(ctx, s.svcCtx)
	return l.AddFollow(in)
}

func (s *RelationServer) DelFollow(ctx context.Context, in *pb.DelFollowReq) (*pb.DelFollowResp, error) {
	l := logic.NewDelFollowLogic(ctx, s.svcCtx)
	return l.DelFollow(in)
}

func (s *RelationServer) CheckIfFollow(ctx context.Context, in *pb.CheckIfFollowReq) (*pb.CheckFollowResp, error) {
	l := logic.NewCheckIfFollowLogic(ctx, s.svcCtx)
	return l.CheckIfFollow(in)
}

func (s *RelationServer) SearchFollowedByFollowingId(ctx context.Context, in *pb.SearchFollowedByFollowingIdReq) (*pb.SearchFollowByFollowingIdResp, error) {
	l := logic.NewSearchFollowedByFollowingIdLogic(ctx, s.svcCtx)
	return l.SearchFollowedByFollowingId(in)
}

func (s *RelationServer) SearchFollowingByFollowedId(ctx context.Context, in *pb.SearchFollowingByFollowedIdReq) (*pb.SearchFollowByFollowedIdResp, error) {
	l := logic.NewSearchFollowingByFollowedIdLogic(ctx, s.svcCtx)
	return l.SearchFollowingByFollowedId(in)
}

// -----------------------fork-----------------------
func (s *RelationServer) AddFork(ctx context.Context, in *pb.AddForkReq) (*pb.AddForkResp, error) {
	l := logic.NewAddForkLogic(ctx, s.svcCtx)
	return l.AddFork(in)
}

func (s *RelationServer) DelFork(ctx context.Context, in *pb.DelForkReq) (*pb.DelForkResp, error) {
	l := logic.NewDelForkLogic(ctx, s.svcCtx)
	return l.DelFork(in)
}

func (s *RelationServer) GetOrigin(ctx context.Context, in *pb.GetOriginReq) (*pb.GetOriginResp, error) {
	l := logic.NewGetOriginLogic(ctx, s.svcCtx)
	return l.GetOrigin(in)
}

func (s *RelationServer) SearchFork(ctx context.Context, in *pb.SearchForkReq) (*pb.SearchForkResp, error) {
	l := logic.NewSearchForkLogic(ctx, s.svcCtx)
	return l.SearchFork(in)
}

// -----------------------star-----------------------
func (s *RelationServer) AddStar(ctx context.Context, in *pb.AddStarReq) (*pb.AddStarResp, error) {
	l := logic.NewAddStarLogic(ctx, s.svcCtx)
	return l.AddStar(in)
}

func (s *RelationServer) DelStar(ctx context.Context, in *pb.DelStarReq) (*pb.DelStarResp, error) {
	l := logic.NewDelStarLogic(ctx, s.svcCtx)
	return l.DelStar(in)
}

func (s *RelationServer) CheckIfStar(ctx context.Context, in *pb.CheckIfStarReq) (*pb.CheckIfStarResp, error) {
	l := logic.NewCheckIfStarLogic(ctx, s.svcCtx)
	return l.CheckIfStar(in)
}

func (s *RelationServer) SearchStaredRepo(ctx context.Context, in *pb.SearchStaredRepoReq) (*pb.SearchStaredRepoResp, error) {
	l := logic.NewSearchStaredRepoLogic(ctx, s.svcCtx)
	return l.SearchStaredRepo(in)
}

func (s *RelationServer) SearchStaringDev(ctx context.Context, in *pb.SearchStaringDevReq) (*pb.SearchStaringDevResp, error) {
	l := logic.NewSearchStaringDevLogic(ctx, s.svcCtx)
	return l.SearchStaringDev(in)
}
