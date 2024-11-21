package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/repo/model"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRepoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRepoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRepoLogic {
	return &UpdateRepoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRepoLogic) UpdateRepo(in *pb.UpdateRepoReq) (*pb.UpdateRepoResp, error) {
	err := l.doUpdateRepo(in)
	if err != nil {
		return &pb.UpdateRepoResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateRepoResp{
		Code:    http.StatusOK,
		Message: "Successfully updated repo",
	}, nil
}

func (l *UpdateRepoLogic) doUpdateRepo(in *pb.UpdateRepoReq) error {
	lock, err := l.acquireUpdateRepoLock(in.Id)
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateRepo(in.Id)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateRepo, in.Id)); err != nil {
		return err
	}

	err = l.pushUpdateRepoTask(in.Id)
	if err != nil {
		logx.Error("Failed to push update repo task: ", err)
		return err
	}

	err = l.blockUntilRepoUpdated(in.Id)
	if err != nil {
		logx.Error("Failed to block until repo updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateRepoLogic) acquireUpdateRepoLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateRepo, id),
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
		logx.Error("Failed to accquire lock: ", err)
		return nil, err
	}

	return lock, nil
}

func (l *UpdateRepoLogic) pushUpdateRepoTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchRepo, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId), asynq.Queue(tasks.FetcherTaskQueue), asynq.MaxRetry(tasks.FetchMaxRetry)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for repo updated")
	return
}

func (l *UpdateRepoLogic) blockUntilRepoUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateRepo, id)); err != nil {
		return
	}

	logx.Info("Repo updated: ", result)
	return
}

func (l *UpdateRepoLogic) checkIfNeedUpdateRepo(id int64) (bool, error) {
	if repo, err := l.svcCtx.RepoModel.FindOneById(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			logx.Error("FindOneById error: ", err)
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(repo.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
