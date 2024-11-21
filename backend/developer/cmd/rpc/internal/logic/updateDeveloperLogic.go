package logic

import (
	"context"
	"errors"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/pb"
	"github.com/ShellWen/GitPulse/developer/model"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeveloperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeveloperLogic {
	return &UpdateDeveloperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeveloperLogic) UpdateDeveloper(in *pb.UpdateDeveloperReq) (*pb.UpdateDeveloperResp, error) {
	err := l.doUpdateDeveloper(in)
	if err != nil {
		return &pb.UpdateDeveloperResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateDeveloperResp{
		Code:    http.StatusOK,
		Message: "Successfully updated developer",
	}, nil
}

func (l *UpdateDeveloperLogic) doUpdateDeveloper(in *pb.UpdateDeveloperReq) error {
	lock, err := l.acquireUpdateDeveloperLock()
	if err != nil {
		return err
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateDeveloper(in.Id)
	if err != nil {
		return err
	}

	if !needUpdate {
		return nil
	}

	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, locks.GetNewLockKey(locks.UpdateDeveloper, in.Id)); err != nil {
		return err
	}

	err = l.pushUpdateDeveloperTask(in.Id)

	if err != nil {
		logx.Error("Failed to push update developer task: ", err)
		return err
	}

	err = l.blockUntilDeveloperUpdated(in.Id)

	if err != nil {
		logx.Error("Failed to block until developer updated: ", err)
		return err
	}

	return nil
}

func (l *UpdateDeveloperLogic) acquireUpdateDeveloperLock() (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         strconv.Itoa(tasks.FetchDeveloper),
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

func (l *UpdateDeveloperLogic) pushUpdateDeveloperTask(id int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = tasks.NewFetcherTask(tasks.FetchDeveloper, id, "", 0); err != nil {
		return
	}

	if _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.TaskID(taskId), asynq.Queue(tasks.FetcherTaskQueue), asynq.MaxRetry(tasks.FetchMaxRetry), asynq.MaxRetry(tasks.FetchMaxRetry)); err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			return
		}
	}

	logx.Info("Successfully pushed task ", task.Payload(), " to fetcher, waiting for developer updated")
	return
}

func (l *UpdateDeveloperLogic) blockUntilDeveloperUpdated(id int64) (err error) {
	var (
		node   redis.ClosableNode
		result string
	)

	if node, err = redis.CreateBlockingNode(l.svcCtx.RedisClient); err != nil {
		return
	}
	defer node.Close()

	if result, err = l.svcCtx.RedisClient.BlpopWithTimeoutCtx(l.ctx, node, time.Duration(l.svcCtx.Config.Timeout)*time.Millisecond, locks.GetNewLockKey(locks.UpdateDeveloper, id)); err != nil {
		return
	}

	logx.Info("Developer updated: ", result)
	return
}

func (l *UpdateDeveloperLogic) checkIfNeedUpdateDeveloper(id int64) (bool, error) {
	if developer, err := l.svcCtx.DeveloperModel.FindOneById(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			logx.Error("FindOneById error: ", err)
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(developer.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
