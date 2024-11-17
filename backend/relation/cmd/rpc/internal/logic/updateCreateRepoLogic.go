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

type UpdateCreateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCreateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCreateRepoLogic {
	return &UpdateCreateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCreateRepoLogic) UpdateCreateRepo(in *pb.UpdateCreateRepoReq) (*pb.UpdateCreateRepoResp, error) {
	err := l.doUpdateCreateRepo(in)
	if err != nil {
		return &pb.UpdateCreateRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateCreateRepoResp{
		Code:    http.StatusOK,
		Message: "Successfully updated createRepo",
	}, nil
}

func (l *UpdateCreateRepoLogic) doUpdateCreateRepo(in *pb.UpdateCreateRepoReq) error {
	lock, err := l.acquireUpdateCreateRepoLock(in.DeveloperId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateCreateRepo(in.DeveloperId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateCreatedRepo, in.DeveloperId)); err != nil {
		return err
	}

	err = l.pushUpdateCreateRepoTask(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to push update createRepo task: ", err)
		return err
	}

	err = l.blockUntilCreateRepoUpdated(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to block until createRepo updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateCreateRepoLogic) acquireUpdateCreateRepoLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateCreatedRepo, id),
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

func (l *UpdateCreateRepoLogic) pushUpdateCreateRepoTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchCreatedRepo, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for createRepo updated")
	return
}

func (l *UpdateCreateRepoLogic) blockUntilCreateRepoUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateCreatedRepo, id)); err != nil {
		return
	}

	logx.Info("CreateRepo updated: ", result)
	return
}

func (l *UpdateCreateRepoLogic) checkIfNeedUpdateCreateRepo(id int64) (bool, error) {
	if createRepoUpdatedAt, err := l.svcCtx.CreatedRepoUpdatedAtModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(createRepoUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
