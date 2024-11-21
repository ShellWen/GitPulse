package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIssuePROfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIssuePROfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIssuePROfUserLogic {
	return &UpdateIssuePROfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateIssuePROfUserLogic) UpdateIssuePROfUser(in *pb.UpdateIssuePROfUserReq) (*pb.UpdateIssuePROfUserResp, error) {
	err := l.doUpdateIssuePROfUser(in)
	if err != nil {
		return &pb.UpdateIssuePROfUserResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateIssuePROfUserResp{
		Code:    http.StatusOK,
		Message: "Successfully updated issue pr",
	}, nil
}

func (l *UpdateIssuePROfUserLogic) doUpdateIssuePROfUser(in *pb.UpdateIssuePROfUserReq) error {
	lock, err := l.acquireUpdateIssuePROfUserLock(in.UserId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateIssuePROfUser(in.UserId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateIssuePROfUser, in.UserId)); err != nil {
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

	err = l.pushUpdateIssuePROfUserTask(in.UserId, updateAfter, searchLimit)

	if err != nil {
		logx.Error("Failed to push update issue pr task: ", err)
		return err
	}

	err = l.blockUntilIssuePROfUserUpdated(in.UserId)

	if err != nil {
		logx.Error("Failed to block until issue pr updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateIssuePROfUserLogic) acquireUpdateIssuePROfUserLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateIssuePROfUser, id),
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

func (l *UpdateIssuePROfUserLogic) pushUpdateIssuePROfUserTask(id int64, updateAfter string, searchLimit int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchIssuePROfUser, id, updateAfter, searchLimit); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId), asynq.Queue(tasks.FetcherTaskQueue), asynq.MaxRetry(tasks.FetchMaxRetry)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for issue pr updated")
	return
}

func (l *UpdateIssuePROfUserLogic) blockUntilIssuePROfUserUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateIssuePROfUser, id)); err != nil {
		return
	}

	logx.Info("IssuePR of user updated: ", result)
	return
}

func (l *UpdateIssuePROfUserLogic) checkIfNeedUpdateIssuePROfUser(id int64) (bool, error) {
	if issuePRUpdatedAt, err := l.svcCtx.IssuePrOfUserUpdatedAtModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(issuePRUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
