package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"net/http"
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

func (l *GetDeveloperLogic) GetDeveloper(req *types.GetDeveloperReq) (resp *types.Developer, err error) {
	resp, err = l.doGetDeveloper(req)
	return
}

func (l *GetDeveloperLogic) doGetDeveloper(req *types.GetDeveloperReq) (resp *types.Developer, err error) {
	var (
		id      int64
		rpcResp *developer.GetDeveloperByIdResp
	)

	resp = &types.Developer{}

	if id, err = customGithub.GetIdByLogin(l.ctx, req.Login); err != nil {
		logx.Error("Failed to get id by login ", req.Login, err)
		return
	}

	if _, err = l.rpcUpdateDeveloperById(id); err != nil {
		logx.Error("Failed to update developer by id ", id, err)
		return
	}

	if rpcResp, err = l.rpcGetDeveloperById(id); err != nil {
		logx.Error("Failed to get developer by id ", id, err)
		return
	}

	resp = &types.Developer{
		Id:        rpcResp.Developer.Id,
		Name:      rpcResp.Developer.Name,
		Login:     rpcResp.Developer.Login,
		AvatarUrl: rpcResp.Developer.AvatarUrl,
		Company:   rpcResp.Developer.Company,
		Location:  rpcResp.Developer.Location,
		Bio:       rpcResp.Developer.Bio,
		Blog:      rpcResp.Developer.Blog,
		Email:     rpcResp.Developer.Email,
		Followers: rpcResp.Developer.Followers,
		Following: rpcResp.Developer.Following,
		Stars:     rpcResp.Developer.Stars,
		Repos:     rpcResp.Developer.Repos,
		Gists:     rpcResp.Developer.Gists,
		CreatedAt: time.Unix(rpcResp.Developer.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(rpcResp.Developer.UpdatedAt, 0).Format(time.RFC3339),
	}

	logx.Info("Successfully get developer by login ", req.Login)
	return
}

func (l *GetDeveloperLogic) rpcUpdateDeveloperById(id int64) (resp *developer.UpdateDeveloperResp, err error) {
	if resp, err = l.svcCtx.DeveloperRpcClient.UpdateDeveloper(l.ctx, &developer.UpdateDeveloperReq{
		Id: id,
	}); err != nil {
		return
	} else if resp.Code != http.StatusOK {
		err = errors.New(resp.Message)
		return
	}

	logx.Info("Successfully update developer by id ", id)
	return
}

func (l *GetDeveloperLogic) rpcGetDeveloperById(id int64) (resp *developer.GetDeveloperByIdResp, err error) {
	if resp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: id,
	}); err != nil {
		return
	} else if resp.Code != http.StatusOK {
		err = errors.New(resp.Message)
		return
	}

	logx.Info("Successfully get developer by id ", id)
	return
}
