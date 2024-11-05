package consumer

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/fetcher/internal/logic"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/relation"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

// the recent month YYYY-MM-DD
var createdAfterTime string = time.Unix(time.Now().Unix()-int64(180*24*time.Hour.Seconds()), 0).Format("2006-01-02")
var issueSearchLimit int64 = 50
var commentSearchLimit int64 = 50

type FetcherTaskConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewFetcherTaskConsumer(ctx context.Context, svc *svc.ServiceContext) *FetcherTaskConsumer {
	return &FetcherTaskConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *FetcherTaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks.FetchTaskName, c.Consume)

	return mux
}

func (c *FetcherTaskConsumer) Consume(ctx context.Context, task *asynq.Task) (err error) {
	logx.Info("consume message: ", task.Type(), task.Payload())

	msg := tasks.FetchPayload{}
	if err = jsonx.Unmarshal(task.Payload(), &msg); err != nil {
		return
	}

	switch msg.Type {
	case tasks.FetchDeveloper:
		defer c.svc.DeveloperRpcClient.UnblockDeveloper(c.ctx, &developer.UnblockDeveloperReq{FetchType: tasks.FetchDeveloper, Id: msg.Id})
		err = logic.FetchDeveloper(c.ctx, c.svc, msg.Id)
	case tasks.FetchRepo:
		defer c.svc.RepoRpcClient.UnblockRepo(c.ctx, &repo.UnblockRepoReq{FetchType: tasks.FetchRepo, Id: msg.Id})
		err = logic.FetchRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchCreatedRepo:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchCreatedRepo, Id: msg.Id})
		err = logic.FetchCreatedRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchStarredRepo:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchStarredRepo, Id: msg.Id})
		err = logic.FetchStarredRepo(c.ctx, c.svc, msg.Id)
	case tasks.FetchFollow:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchFollow, Id: msg.Id})
		err = logic.FetchFollow(c.ctx, c.svc, msg.Id)
	case tasks.FetchFollower:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchFollower, Id: msg.Id})
		err = logic.FetchFollower(c.ctx, c.svc, msg.Id)
	case tasks.FetchFollowing:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchFollowing, Id: msg.Id})
		err = logic.FetchFollowing(c.ctx, c.svc, msg.Id)
	case tasks.FetchFork:
		defer c.svc.RelationRpcClient.UnblockRelation(c.ctx, &relation.UnblockRelationReq{FetchType: tasks.FetchFork, Id: msg.Id})
		err = logic.FetchFork(c.ctx, c.svc, msg.Id)
	case tasks.FetchContributionOfUser:
		defer c.svc.ContributionRpcClient.UnblockContribution(c.ctx, &contribution.UnblockContributionReq{FetchType: tasks.FetchContributionOfUser, Id: msg.Id})
		err = logic.FetchContributionOfUser(c.ctx, c.svc, msg.Id, createdAfterTime, issueSearchLimit, commentSearchLimit)
	case tasks.FetchIssuePROfUser:
		defer c.svc.ContributionRpcClient.UnblockContribution(c.ctx, &contribution.UnblockContributionReq{FetchType: tasks.FetchIssuePROfUser, Id: msg.Id})
		err = logic.FetchIssuePROfUser(c.ctx, c.svc, msg.Id, createdAfterTime, issueSearchLimit)
	case tasks.FetchCommentOfUser:
		defer c.svc.ContributionRpcClient.UnblockContribution(c.ctx, &contribution.UnblockContributionReq{FetchType: tasks.FetchCommentOfUser, Id: msg.Id})
		err = logic.FetchCommentOfUser(c.ctx, c.svc, msg.Id, createdAfterTime, commentSearchLimit)
	default:
		err = errors.New("unexpected message type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	// wait to ensure the message is consumed
	time.Sleep(time.Second)

	return
}
