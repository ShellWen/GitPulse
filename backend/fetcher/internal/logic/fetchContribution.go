package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

const (
	RoleAuthor    = "author"
	RoleCommenter = "commenter"
)

func FetchContributionOfUser(ctx context.Context, svcContext *svc.ServiceContext, repoId int64, createAfter string, issueSearchLimit int64, commentSearchLimit int64) (err error) {
	if err = FetchIssuePROfUser(ctx, svcContext, repoId, createAfter, issueSearchLimit); err != nil {
		return
	}

	if err = FetchCommentOfUser(ctx, svcContext, repoId, createAfter, commentSearchLimit); err != nil {
		return
	}

	return
}

func buildContribution(ctx context.Context, svcContext *svc.ServiceContext, githubContribution *github.ContributorStats, repoId int64) (newContribution *model.Contribution) {
	newContribution = &model.Contribution{
		RepoId: repoId,
	}
	return
}

func delAllOldContributionInCategory(ctx context.Context, svcContext *svc.ServiceContext, userId int64, category string) (err error) {
	contributionZrpcClient := svcContext.ContributionRpcClient

	if delAllOldContributionResp, err := contributionZrpcClient.DelAllContributionInCategoryByUserId(ctx, &contribution.DelAllContributionInCategoryByUserIdReq{Category: category, UserId: userId}); err != nil {
		logx.Error("Unexpected error when deleting all old " + category + " contributions: " + err.Error())
		return err
	} else if delAllOldContributionResp.Code != http.StatusOK {
		errMsg := "Unexpected error when deleting all old " + category + " contributions: " + delAllOldContributionResp.Message
		err = errors.New(errMsg)
		logx.Error(errMsg)
		return err
	}

	logx.Info("Successfully delete all old " + category + " contributions")
	return nil
}

func pushContribution(ctx context.Context, svcContext *svc.ServiceContext, newContribution *model.Contribution) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(newContribution); err != nil {
		return
	}

	if err = svcContext.KqContributionPusher.Push(ctx, jsonStr); err != nil {
		return
	}

	logx.Info("Successfully pushed a contribution, size: ", len(jsonStr))
	return
}

func updateContributionFetchTimeOfDeveloper(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
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

	theDeveloper.LastFetchContributionAt = time.Now().Unix()

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
