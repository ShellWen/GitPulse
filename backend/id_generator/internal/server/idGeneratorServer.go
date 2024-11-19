// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: idGenerator.proto

package server

import (
	"context"

	"github.com/ShellWen/GitPulse/id_generator/internal/logic"
	"github.com/ShellWen/GitPulse/id_generator/internal/svc"
	"github.com/ShellWen/GitPulse/id_generator/pb"
)

type IdGeneratorServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedIdGeneratorServer
}

func NewIdGeneratorServer(svcCtx *svc.ServiceContext) *IdGeneratorServer {
	return &IdGeneratorServer{
		svcCtx: svcCtx,
	}
}

func (s *IdGeneratorServer) GetId(ctx context.Context, in *pb.GetIdReq) (*pb.GetIdResp, error) {
	l := logic.NewGetIdLogic(ctx, s.svcCtx)
	return l.GetId(in)
}