package developer

import (
	"context"
	"github.com/ShellWen/GitPulse/common/message"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/api/internal/types"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"net/http"
	"os"
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
	resp, err = l.doGetDeveloper(req)
	return
}

func (l *GetDeveloperLogic) doGetDeveloper(req *types.GetDeveloperReq) (resp *types.GetDeveloperResp, err error) {
	var (
		id      int64
		rpcResp *developer.GetDeveloperByIdResp
	)

	if id, err = l.getIdByLogin(req.Login); err != nil {
		return
	}

	if rpcResp, err = l.rpcGetDeveloperById(id); err != nil {
		return
	}

	switch rpcResp.Code {
	case http.StatusOK:
		logx.Info("Found in local cache")
		if time.Now().Unix()-rpcResp.Developer.GetDataUpdatedAt() < int64(time.Hour.Seconds()*24) {
			break
		}
		logx.Info("Local cache expired, fetching from github")
		fallthrough
	case http.StatusNotFound:
		logx.Info("Not found in local cache, fetching from github")
		if err = l.updateDeveloperWithBlock(id); err != nil {
			return
		}

		if rpcResp, err = l.rpcGetDeveloperById(id); err != nil {
			return
		}
		fallthrough
	default:
		if rpcResp.Code != http.StatusOK {
			err = jsonx.UnmarshalFromString(rpcResp.Message, &resp)
			return
		}
	}

	resp = &types.GetDeveloperResp{
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
	return
}

func (l *GetDeveloperLogic) getIdByLogin(login string) (id int64, err error) {
	var githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
	var githubUser *github.User
	if githubUser, _, err = githubClient.Users.Get(l.ctx, login); err != nil {
		logx.Error("Unexpected error when fetching user ", login, " from github: ", err)
		return
	}
	id = githubUser.GetID()
	logx.Info("Successfully get id ", id, " of login", login)
	return
}

func (l *GetDeveloperLogic) rpcGetDeveloperById(id int64) (resp *developer.GetDeveloperByIdResp, err error) {
	var rpcResp *developer.GetDeveloperByIdResp
	if rpcResp, err = l.svcCtx.DeveloperRpc.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: id,
	}); err != nil {
		resp = &developer.GetDeveloperByIdResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	logx.Info("Successfully get developer by id ", id)
	resp = rpcResp
	return
}

func (l *GetDeveloperLogic) updateDeveloperWithBlock(id int64) (err error) {
	var (
		task    *message.FetcherTask
		jsonStr string
	)

	task = &message.FetcherTask{Type: message.FetchDeveloper, Id: id}
	if jsonStr, err = jsonx.MarshalToString(task); err != nil {
		return
	}

	if err = l.svcCtx.KqFetcherTaskPusher.Push(l.ctx, jsonStr); err != nil {
		return
	}
	logx.Info("Successfully pushed task ", jsonStr, " to fetcher, waiting for developer updated")

	if _, err = l.svcCtx.DeveloperRpc.BlockUntilDeveloperUpdated(l.ctx, &developer.BlockUntilDeveloperUpdatedReq{Id: id}); err != nil {
		return
	}
	logx.Info("Developer successfully updated")

	return
}
