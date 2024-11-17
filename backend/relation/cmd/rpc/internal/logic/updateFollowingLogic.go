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

type UpdateFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFollowingLogic {
	return &UpdateFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFollowingLogic) UpdateFollowing(in *pb.UpdateFollowingReq) (*pb.UpdateFollowingResp, error) {
	err := l.doUpdateFollowing(in)
	if err != nil {
		return &pb.UpdateFollowingResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateFollowingResp{
		Code:    http.StatusOK,
		Message: "Successfully updated following",
	}, nil
}

func (l *UpdateFollowingLogic) doUpdateFollowing(in *pb.UpdateFollowingReq) error {
	lock, err := l.acquireUpdateFollowingLock(in.DeveloperId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateFollowing(in.DeveloperId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateFollowing, in.DeveloperId)); err != nil {
		return err
	}

	err = l.pushUpdateFollowingTask(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to push update following task: ", err)
		return err
	}

	err = l.blockUntilFollowingUpdated(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to block until following updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateFollowingLogic) acquireUpdateFollowingLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateFollowing, id),
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

func (l *UpdateFollowingLogic) pushUpdateFollowingTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchFollowing, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for following updated")
	return
}

func (l *UpdateFollowingLogic) blockUntilFollowingUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateFollowing, id)); err != nil {
		return
	}

	logx.Info("Following updated: ", result)
	return
}

func (l *UpdateFollowingLogic) checkIfNeedUpdateFollowing(id int64) (bool, error) {
	if followingUpdatedAt, err := l.svcCtx.FollowingUpdatedAtModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(followingUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
