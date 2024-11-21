package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateForkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateForkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateForkLogic {
	return &UpdateForkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateForkLogic) UpdateFork(in *pb.UpdateForkReq) (*pb.UpdateForkResp, error) {
	err := l.doUpdateFork(in)
	if err != nil {
		return &pb.UpdateForkResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateForkResp{
		Code:    http.StatusOK,
		Message: "Successfully updated fork",
	}, nil
}

func (l *UpdateForkLogic) doUpdateFork(in *pb.UpdateForkReq) error {
	lock, err := l.acquireUpdateForkLock(in.OriginalRepoId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateFork(in.OriginalRepoId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateFork, in.OriginalRepoId)); err != nil {
		return err
	}

	err = l.pushUpdateForkTask(in.OriginalRepoId)

	if err != nil {
		logx.Error("Failed to push update fork task: ", err)
		return err
	}

	err = l.blockUntilForkUpdated(in.OriginalRepoId)

	if err != nil {
		logx.Error("Failed to block until fork updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateForkLogic) acquireUpdateForkLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateFork, id),
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

func (l *UpdateForkLogic) pushUpdateForkTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchFork, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId), asynq.Queue(tasks.FetcherTaskQueue), asynq.MaxRetry(tasks.FetchMaxRetry)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for fork updated")
	return
}

func (l *UpdateForkLogic) blockUntilForkUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateFork, id)); err != nil {
		return
	}

	logx.Info("Fork updated: ", result)
	return
}

func (l *UpdateForkLogic) checkIfNeedUpdateFork(id int64) (bool, error) {
	if forkUpdatedAt, err := l.svcCtx.ForkUpdatedAtModel.FindOneByRepoId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(forkUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
