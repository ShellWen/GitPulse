package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentOfUserLogic {
	return &UpdateCommentOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentOfUserLogic) UpdateCommentOfUser(in *pb.UpdateCommentOfUserReq) (*pb.UpdateCommentOfUserResp, error) {
	err := l.doUpdateCommentOfUser(in)
	if err != nil {
		return &pb.UpdateCommentOfUserResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateCommentOfUserResp{
		Code:    http.StatusOK,
		Message: "Successfully updated comment",
	}, nil
}

func (l *UpdateCommentOfUserLogic) doUpdateCommentOfUser(in *pb.UpdateCommentOfUserReq) error {
	lock, err := l.acquireUpdateCommentOfUserLock(in.UserId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateCommentOfUser(in.UserId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateCommentOfUser, in.UserId)); err != nil {
		return err
	}

	var (
		updateAfter string
		searchLimit int64
	)

	if in.UpdateAfter == "" {
		updateAfter = customGithub.DefaultUpdateAfterTime()
	} else {
		updateAfter = in.UpdateAfter
	}

	if in.SearchLimit == 0 {
		searchLimit = customGithub.DefaultSearchLimit
	} else {
		searchLimit = in.SearchLimit
	}

	err = l.pushUpdateCommentOfUserTask(in.UserId, updateAfter, searchLimit)

	if err != nil {
		logx.Error("Failed to push update comment task: ", err)
		return err
	}

	err = l.blockUntilCommentOfUserUpdated(in.UserId)

	if err != nil {
		logx.Error("Failed to block until comment updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateCommentOfUserLogic) acquireUpdateCommentOfUserLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateCommentOfUser, id),
		Value:       []byte("locked"),
		SessionTTL:  "10s",
		SessionName: uuid.Must(uuid.NewV7()).String(),
	})

	if err != nil {
		logx.Error("Failed to create lock: ", err)
		return nil, err
	}

	_, err = lock.Lock(nil)

	if err != nil {
		logx.Error("Failed to acquire lock: ", err)
		return nil, err
	}

	return lock, nil
}

func (l *UpdateCommentOfUserLogic) pushUpdateCommentOfUserTask(id int64, updateAfter string, searchLimit int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchCommentOfUser, id, updateAfter, searchLimit); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for comment updated")
	return
}

func (l *UpdateCommentOfUserLogic) blockUntilCommentOfUserUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateCommentOfUser, id)); err != nil {
		return
	}

	logx.Info("Comment of user updated: ", result)
	return
}

func (l *UpdateCommentOfUserLogic) checkIfNeedUpdateCommentOfUser(id int64) (bool, error) {
	if commentUpdatedAt, err := l.svcCtx.CommentOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(commentUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
