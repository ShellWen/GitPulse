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

type UpdateFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFollowerLogic {
	return &UpdateFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFollowerLogic) UpdateFollower(in *pb.UpdateFollowerReq) (*pb.UpdateFollowerResp, error) {
	err := l.doUpdateFollower(in)
	if err != nil {
		return &pb.UpdateFollowerResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateFollowerResp{
		Code:    http.StatusOK,
		Message: "Successfully updated follower",
	}, nil
}

func (l *UpdateFollowerLogic) doUpdateFollower(in *pb.UpdateFollowerReq) error {
	lock, err := l.acquireUpdateFollowerLock(in.DeveloperId)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateFollower(in.DeveloperId)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateFollower, in.DeveloperId)); err != nil {
		return err
	}

	err = l.pushUpdateFollowerTask(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to push update follower task: ", err)
		return err
	}

	err = l.blockUntilFollowerUpdated(in.DeveloperId)

	if err != nil {
		logx.Error("Failed to block until follower updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateFollowerLogic) acquireUpdateFollowerLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateFollower, id),
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

func (l *UpdateFollowerLogic) pushUpdateFollowerTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchFollower, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for follower updated")
	return
}

func (l *UpdateFollowerLogic) blockUntilFollowerUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateFollower, id)); err != nil {
		return
	}

	logx.Info("Follower updated: ", result)
	return
}

func (l *UpdateFollowerLogic) checkIfNeedUpdateFollower(id int64) (bool, error) {
	if followerUpdatedAt, err := l.svcCtx.FollowerUpdatedAtModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(followerUpdatedAt.UpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
