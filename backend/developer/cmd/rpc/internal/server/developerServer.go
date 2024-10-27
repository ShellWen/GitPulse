// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: developer.proto

package server

import (
	"context"

	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/logic"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"
)

type DeveloperServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedDeveloperServer
}

func NewDeveloperServer(svcCtx *svc.ServiceContext) *DeveloperServer {
	return &DeveloperServer{
		svcCtx: svcCtx,
	}
}

// -----------------------developer-----------------------
func (s *DeveloperServer) AddDeveloper(ctx context.Context, in *pb.AddDeveloperReq) (*pb.AddDeveloperResp, error) {
	l := logic.NewAddDeveloperLogic(ctx, s.svcCtx)
	return l.AddDeveloper(in)
}

func (s *DeveloperServer) UpdateDeveloper(ctx context.Context, in *pb.UpdateDeveloperReq) (*pb.UpdateDeveloperResp, error) {
	l := logic.NewUpdateDeveloperLogic(ctx, s.svcCtx)
	return l.UpdateDeveloper(in)
}

func (s *DeveloperServer) DelDeveloperById(ctx context.Context, in *pb.DelDeveloperByIdReq) (*pb.DelDeveloperByIdResp, error) {
	l := logic.NewDelDeveloperByIdLogic(ctx, s.svcCtx)
	return l.DelDeveloperById(in)
}

func (s *DeveloperServer) DelDeveloperByUsername(ctx context.Context, in *pb.DelDeveloperByUsernameReq) (*pb.DelDeveloperByUsernameResp, error) {
	l := logic.NewDelDeveloperByUsernameLogic(ctx, s.svcCtx)
	return l.DelDeveloperByUsername(in)
}

func (s *DeveloperServer) GetDeveloperById(ctx context.Context, in *pb.GetDeveloperByIdReq) (*pb.GetDeveloperByIdResp, error) {
	l := logic.NewGetDeveloperByIdLogic(ctx, s.svcCtx)
	return l.GetDeveloperById(in)
}

func (s *DeveloperServer) GetDeveloperByUsername(ctx context.Context, in *pb.GetDeveloperByUsernameReq) (*pb.GetDeveloperByUsernameResp, error) {
	l := logic.NewGetDeveloperByUsernameLogic(ctx, s.svcCtx)
	return l.GetDeveloperByUsername(in)
}
