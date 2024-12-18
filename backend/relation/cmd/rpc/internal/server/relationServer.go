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

func (s *RelationServer) DelAllCreatedRepo(ctx context.Context, in *pb.DelAllCreatedRepoReq) (*pb.DelAllCreatedRepoResp, error) {
	l := logic.NewDelAllCreatedRepoLogic(ctx, s.svcCtx)
	return l.DelAllCreatedRepo(in)
}

func (s *RelationServer) GetCreatorId(ctx context.Context, in *pb.GetCreatorIdReq) (*pb.GetCreatorIdResp, error) {
	l := logic.NewGetCreatorIdLogic(ctx, s.svcCtx)
	return l.GetCreatorId(in)
}

func (s *RelationServer) SearchCreatedRepo(ctx context.Context, in *pb.SearchCreatedRepoReq) (*pb.SearchCreatedRepoResp, error) {
	l := logic.NewSearchCreatedRepoLogic(ctx, s.svcCtx)
	return l.SearchCreatedRepo(in)
}

func (s *RelationServer) UpdateCreateRepo(ctx context.Context, in *pb.UpdateCreateRepoReq) (*pb.UpdateCreateRepoResp, error) {
	l := logic.NewUpdateCreateRepoLogic(ctx, s.svcCtx)
	return l.UpdateCreateRepo(in)
}

func (s *RelationServer) GetCreatedRepoUpdatedAt(ctx context.Context, in *pb.GetCreatedRepoUpdatedAtReq) (*pb.GetCreatedRepoUpdatedAtResp, error) {
	l := logic.NewGetCreatedRepoUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetCreatedRepoUpdatedAt(in)
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

func (s *RelationServer) DelAllFollower(ctx context.Context, in *pb.DelAllFollowerReq) (*pb.DelAllFollowerResp, error) {
	l := logic.NewDelAllFollowerLogic(ctx, s.svcCtx)
	return l.DelAllFollower(in)
}

func (s *RelationServer) DelAllFollowing(ctx context.Context, in *pb.DelAllFollowingReq) (*pb.DelAllFollowingResp, error) {
	l := logic.NewDelAllFollowingLogic(ctx, s.svcCtx)
	return l.DelAllFollowing(in)
}

func (s *RelationServer) CheckIfFollow(ctx context.Context, in *pb.CheckIfFollowReq) (*pb.CheckFollowResp, error) {
	l := logic.NewCheckIfFollowLogic(ctx, s.svcCtx)
	return l.CheckIfFollow(in)
}

func (s *RelationServer) SearchFollowingByDeveloperId(ctx context.Context, in *pb.SearchFollowingByDeveloperIdReq) (*pb.SearchFollowingByDeveloperIdResp, error) {
	l := logic.NewSearchFollowingByDeveloperIdLogic(ctx, s.svcCtx)
	return l.SearchFollowingByDeveloperId(in)
}

func (s *RelationServer) SearchFollowerByDeveloperId(ctx context.Context, in *pb.SearchFollowerByDeveloperIdReq) (*pb.SearchFollowerByDeveloperIdResp, error) {
	l := logic.NewSearchFollowerByDeveloperIdLogic(ctx, s.svcCtx)
	return l.SearchFollowerByDeveloperId(in)
}

func (s *RelationServer) UpdateFollowing(ctx context.Context, in *pb.UpdateFollowingReq) (*pb.UpdateFollowingResp, error) {
	l := logic.NewUpdateFollowingLogic(ctx, s.svcCtx)
	return l.UpdateFollowing(in)
}

func (s *RelationServer) UpdateFollower(ctx context.Context, in *pb.UpdateFollowerReq) (*pb.UpdateFollowerResp, error) {
	l := logic.NewUpdateFollowerLogic(ctx, s.svcCtx)
	return l.UpdateFollower(in)
}

func (s *RelationServer) GetFollowingUpdatedAt(ctx context.Context, in *pb.GetFollowingUpdatedAtReq) (*pb.GetFollowingUpdatedAtResp, error) {
	l := logic.NewGetFollowingUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetFollowingUpdatedAt(in)
}

func (s *RelationServer) GetFollowerUpdatedAt(ctx context.Context, in *pb.GetFollowerUpdatedAtReq) (*pb.GetFollowerUpdatedAtResp, error) {
	l := logic.NewGetFollowerUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetFollowerUpdatedAt(in)
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

func (s *RelationServer) DelAllFork(ctx context.Context, in *pb.DelAllForkReq) (*pb.DelAllForkResp, error) {
	l := logic.NewDelAllForkLogic(ctx, s.svcCtx)
	return l.DelAllFork(in)
}

func (s *RelationServer) GetOrigin(ctx context.Context, in *pb.GetOriginReq) (*pb.GetOriginResp, error) {
	l := logic.NewGetOriginLogic(ctx, s.svcCtx)
	return l.GetOrigin(in)
}

func (s *RelationServer) SearchFork(ctx context.Context, in *pb.SearchForkReq) (*pb.SearchForkResp, error) {
	l := logic.NewSearchForkLogic(ctx, s.svcCtx)
	return l.SearchFork(in)
}

func (s *RelationServer) UpdateFork(ctx context.Context, in *pb.UpdateForkReq) (*pb.UpdateForkResp, error) {
	l := logic.NewUpdateForkLogic(ctx, s.svcCtx)
	return l.UpdateFork(in)
}

func (s *RelationServer) GetForkUpdatedAt(ctx context.Context, in *pb.GetForkUpdatedAtReq) (*pb.GetForkUpdatedAtResp, error) {
	l := logic.NewGetForkUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetForkUpdatedAt(in)
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

func (s *RelationServer) DelAllStarredRepo(ctx context.Context, in *pb.DelAllStarredRepoReq) (*pb.DelAllStarredRepoResp, error) {
	l := logic.NewDelAllStarredRepoLogic(ctx, s.svcCtx)
	return l.DelAllStarredRepo(in)
}

func (s *RelationServer) DelAllStaringDev(ctx context.Context, in *pb.DelAllStaringDevReq) (*pb.DelAllStaringDevResp, error) {
	l := logic.NewDelAllStaringDevLogic(ctx, s.svcCtx)
	return l.DelAllStaringDev(in)
}

func (s *RelationServer) CheckIfStar(ctx context.Context, in *pb.CheckIfStarReq) (*pb.CheckIfStarResp, error) {
	l := logic.NewCheckIfStarLogic(ctx, s.svcCtx)
	return l.CheckIfStar(in)
}

func (s *RelationServer) SearchStarredRepo(ctx context.Context, in *pb.SearchStarredRepoReq) (*pb.SearchStarredRepoResp, error) {
	l := logic.NewSearchStarredRepoLogic(ctx, s.svcCtx)
	return l.SearchStarredRepo(in)
}

func (s *RelationServer) SearchStaringDev(ctx context.Context, in *pb.SearchStaringDevReq) (*pb.SearchStaringDevResp, error) {
	l := logic.NewSearchStaringDevLogic(ctx, s.svcCtx)
	return l.SearchStaringDev(in)
}

func (s *RelationServer) UpdateStarredRepo(ctx context.Context, in *pb.UpdateStarredRepoReq) (*pb.UpdateStarredRepoResp, error) {
	l := logic.NewUpdateStarredRepoLogic(ctx, s.svcCtx)
	return l.UpdateStarredRepo(in)
}

func (s *RelationServer) GetStarredRepoUpdatedAt(ctx context.Context, in *pb.GetStarredRepoUpdatedAtReq) (*pb.GetStarredRepoUpdatedAtResp, error) {
	l := logic.NewGetStarredRepoUpdatedAtLogic(ctx, s.svcCtx)
	return l.GetStarredRepoUpdatedAt(in)
}
