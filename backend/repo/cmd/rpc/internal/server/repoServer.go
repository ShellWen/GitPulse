// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: repo.proto

package server

import (
	"context"

	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/logic"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"
)

type RepoServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRepoServer
}

func NewRepoServer(svcCtx *svc.ServiceContext) *RepoServer {
	return &RepoServer{
		svcCtx: svcCtx,
	}
}

// -----------------------repo-----------------------
func (s *RepoServer) AddRepo(ctx context.Context, in *pb.AddRepoReq) (*pb.AddRepoResp, error) {
	l := logic.NewAddRepoLogic(ctx, s.svcCtx)
	return l.AddRepo(in)
}

func (s *RepoServer) UpdateRepo(ctx context.Context, in *pb.UpdateRepoReq) (*pb.UpdateRepoResp, error) {
	l := logic.NewUpdateRepoLogic(ctx, s.svcCtx)
	return l.UpdateRepo(in)
}

func (s *RepoServer) DelRepoById(ctx context.Context, in *pb.DelRepoByIdReq) (*pb.DelRepoByIdResp, error) {
	l := logic.NewDelRepoByIdLogic(ctx, s.svcCtx)
	return l.DelRepoById(in)
}

func (s *RepoServer) GetRepoById(ctx context.Context, in *pb.GetRepoByIdReq) (*pb.GetRepoByIdResp, error) {
	l := logic.NewGetRepoByIdLogic(ctx, s.svcCtx)
	return l.GetRepoById(in)
}
