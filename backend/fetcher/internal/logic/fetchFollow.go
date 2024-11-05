package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

func FetchFollow(ctx context.Context, svcContext *svc.ServiceContext, developerId int64) (err error) {
	if err = FetchFollower(ctx, svcContext, developerId); err != nil {
		return
	}

	if err = FetchFollowing(ctx, svcContext, developerId); err != nil {
		return
	}

	return
}

func buildFollow(ctx context.Context, svcContext *svc.ServiceContext, followerId int64, followingId int64) (follow *model.Follow) {
	return &model.Follow{
		FollowerId:  followerId,
		FollowingId: followingId,
	}
}

func pushFollow(ctx context.Context, svcContext *svc.ServiceContext, follow *model.Follow) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(follow); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling follow: " + err.Error()))
		return
	}

	if err = svcContext.KqFollowPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when pushing follow: " + err.Error()))
		return
	}

	return
}

func updateFollowFetchTimeOfDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	developerZrpcClient := svcContext.DeveloperRpcClient
	var resp *developer.GetDeveloperByIdResp
	var theDeveloper *developer.Developer

	if resp, err = developerZrpcClient.GetDeveloperById(ctx, &developer.GetDeveloperByIdReq{Id: userId}); err != nil {
		return
	}

	switch resp.Code {
	case http.StatusOK:
		theDeveloper = resp.Developer
	case http.StatusNotFound:
		err = errors.New("Developer not found")
		return
	default:
		err = errors.New("Unexpected error when getting developer: " + resp.Message)
		return
	}

	theDeveloper.LastFetchFollowAt = time.Now().Unix()
	if _, err = developerZrpcClient.UpdateDeveloper(ctx, &developer.UpdateDeveloperReq{
		Id:                      userId,
		Name:                    theDeveloper.Name,
		Login:                   theDeveloper.Login,
		AvatarUrl:               theDeveloper.AvatarUrl,
		Company:                 theDeveloper.Company,
		Location:                theDeveloper.Location,
		Bio:                     theDeveloper.Bio,
		Blog:                    theDeveloper.Blog,
		Email:                   theDeveloper.Email,
		CreatedAt:               theDeveloper.CreatedAt,
		UpdatedAt:               theDeveloper.UpdatedAt,
		TwitterUsername:         theDeveloper.TwitterUsername,
		Repos:                   theDeveloper.Repos,
		Following:               theDeveloper.Following,
		Followers:               theDeveloper.Followers,
		Gists:                   theDeveloper.Gists,
		Stars:                   theDeveloper.Stars,
		LastFetchContributionAt: theDeveloper.LastFetchContributionAt,
		LastFetchFollowAt:       theDeveloper.LastFetchFollowAt,
		LastFetchStarAt:         theDeveloper.LastFetchStarAt,
		LastFetchCreateRepoAt:   theDeveloper.LastFetchCreateRepoAt,
	}); err != nil {
		return
	}

	return
}
