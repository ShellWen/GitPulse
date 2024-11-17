package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const (
	RoleAuthor    = "author"
	RoleCommenter = "commenter"
	RoleReviewer  = "reviewed-by"
)

/*
func FetchContributionOfUser(ctx context.Context, svcContext *svc.ServiceContext, repoId int64, createAfter string, searchLimit int64) (err error) {
	successPush := 0

	if successPush, err = FetchReviewOfUser(ctx, svcContext, repoId, createAfter, searchLimit); err != nil {
		return
	}

	if successPush >= int(searchLimit) {
		return
	}
	searchLimit -= int64(successPush)

	if successPush, err = FetchIssuePROfUser(ctx, svcContext, repoId, createAfter, searchLimit); err != nil {
		return
	}

	if successPush >= int(searchLimit) {
		return
	}
	searchLimit -= int64(successPush)

	if _, err = FetchCommentOfUser(ctx, svcContext, repoId, createAfter, searchLimit); err != nil {
		return
	}

	return
}
*/

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
