package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/relation/model"
	"net/http"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStaringDevLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStaringDevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStaringDevLogic {
	return &SearchStaringDevLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStaringDevLogic) SearchStaringDev(in *pb.SearchStaringDevReq) (resp *pb.SearchStaringDevResp, err error) {
	var stars *[]*model.Star
	var staringDevIds *[]int64

	if stars, err = l.svcCtx.StarModel.SearchStaringDeveloper(l.ctx, in.RepoId, in.Page, in.Limit); err != nil {
		resp = &pb.SearchStaringDevResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else if staringDevIds = l.buildStaringDevIds(stars); len(*staringDevIds) == 0 {
		resp = &pb.SearchStaringDevResp{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	} else {
		resp = &pb.SearchStaringDevResp{
			Code:         http.StatusOK,
			Message:      http.StatusText(http.StatusOK),
			DeveloperIds: *staringDevIds,
		}
	}

	err = nil
	return
}

func (l *SearchStaringDevLogic) buildStaringDevIds(stars *[]*model.Star) (staringDevIds *[]int64) {
	staringDevIds = new([]int64)
	for _, star := range *stars {
		*staringDevIds = append(*staringDevIds, star.DeveloperId)
	}

	return
}
