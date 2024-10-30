package developer

import (
	"context"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperLogic {
	return &GetDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeveloperLogic) GetDeveloper(req *types.GetDeveloperReq) (resp *types.GetDeveloperResp, err error) {
	var rpcResp *developer.GetDeveloperByUsernameResp
	rpcResp, err = l.svcCtx.DeveloperRpc.GetDeveloperByUsername(l.ctx, &developer.GetDeveloperByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		return
	}

	resp = &types.GetDeveloperResp{
		Data: struct {
			Id        int64  `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"login"`
			AvatarUrl string `json:"avatar_url"`
			Company   string `json:"company"`
			Location  string `json:"location"`
			Bio       string `json:"bio"`
			Blog      string `json:"blog"`
			Email     string `json:"email"`
			CreateAt  string `json:"created_at"`
			UpdateAt  string `json:"updated_at"`
		}{
			rpcResp.Developer.Id,
			rpcResp.Developer.Name,
			rpcResp.Developer.Username,
			rpcResp.Developer.AvatarUrl,
			rpcResp.Developer.Company,
			rpcResp.Developer.Location,
			rpcResp.Developer.Bio,
			rpcResp.Developer.Blog,
			rpcResp.Developer.Email,
			time.Unix(rpcResp.Developer.CreateAt, 0).Format(time.RFC3339),
			time.Unix(rpcResp.Developer.UpdateAt, 0).Format(time.RFC3339),
		},
	}
	return
}
