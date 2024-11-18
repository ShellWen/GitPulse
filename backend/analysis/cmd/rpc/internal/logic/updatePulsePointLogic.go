package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	contributionModel "github.com/ShellWen/GitPulse/contribution/model"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ScoreComment         = 1
	ScoreOpenIssue       = 2
	ScoreOpenPullRequest = 3
	ScoreReview          = 5
	ScoreMerge           = 5
)

type UpdatePulsePointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePulsePointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePulsePointLogic {
	return &UpdatePulsePointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePulsePointLogic) UpdatePulsePoint(in *pb.UpdateAnalysisReq) (resp *pb.UpdateAnalysisResp, err error) {
	if err = l.doUpdatePulsePoint(in.DeveloperId); err != nil {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = &pb.UpdateAnalysisResp{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		}
	}
	return
}

func (l *UpdatePulsePointLogic) doUpdatePulsePoint(id int64) (err error) {
	var (
		contributionZrpcClient              = l.svcCtx.ContributionRpcClient
		repositoryZrpcClient                = l.svcCtx.RepoRpcClient
		updateAllContributionResp           *contribution.UpdateContributionOfUserResp
		allContributionResp                 *contribution.SearchByUserIdResp
		allContributions                    []*contribution.Contribution
		allContributionsCategorizedByRepoId = make(map[int64][]*contribution.Contribution)
		pulsePoint                          float64
		pulsePointItem                      *model.PulsePoint
	)

	lock, err := l.acquireUpdatePulsePointLock(id)
	if err != nil {
		return
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdatePulsePoint(id)
	if err != nil {
		return
	}

	if !needUpdate {
		return
	}

	if updateAllContributionResp, err = contributionZrpcClient.UpdateContributionOfUser(l.ctx, &contribution.UpdateContributionOfUserReq{
		UserId: id,
	}); err != nil {
		return
	} else if updateAllContributionResp.Code != http.StatusOK {
		return errors.New(updateAllContributionResp.Message)
	}

	if allContributionResp, err = contributionZrpcClient.SearchByUserId(l.ctx, &contribution.SearchByUserIdReq{
		UserId: id,
		Limit:  1000,
		Page:   1,
	}); err != nil {
		return
	} else if allContributionResp.Code == http.StatusInternalServerError {
		return errors.New(allContributionResp.Message)
	}

	allContributions = allContributionResp.Contributions
	for _, theContribution := range allContributions {
		allContributionsCategorizedByRepoId[theContribution.RepoId] = append(allContributionsCategorizedByRepoId[theContribution.RepoId], theContribution)
	}

	const currentDeveloperId = uint32(0)
	virtualDeveloperId := uint32(1)

	for repoId, contributions := range allContributionsCategorizedByRepoId {
		var (
			repoResp                   *repo.GetRepoByIdResp
			theRepo                    *repo.Repo
			currentDeveloperPulsePoint float64
			virtualDeveloperPulsePoint float64
			mean                       float64
		)

		for _, theContribution := range contributions {
			switch theContribution.Category {
			case contributionModel.CategoryComment:
				currentDeveloperPulsePoint += ScoreComment
			case contributionModel.CategoryOpenIssue:
				currentDeveloperPulsePoint += ScoreOpenIssue
			case contributionModel.CategoryOpenPullRequest:
				currentDeveloperPulsePoint += ScoreOpenPullRequest
			case contributionModel.CategoryReview:
				currentDeveloperPulsePoint += ScoreReview
			case contributionModel.CategoryMerge:
				currentDeveloperPulsePoint += ScoreMerge
			}
		}
		currentDeveloperPulsePoint = math.Sqrt(currentDeveloperPulsePoint)

		if repoResp, err = repositoryZrpcClient.GetRepoById(l.ctx, &repo.GetRepoByIdReq{Id: repoId}); err != nil {
			return
		}
		if repoResp.Code != http.StatusOK {
			logx.Error("Unexpected error: " + repoResp.Message)
			continue
		}

		theRepo = repoResp.Repo
		virtualDeveloperPulsePoint = float64(theRepo.CommentCount*ScoreComment + theRepo.IssueCount*ScoreOpenIssue + theRepo.OpenPrCount*ScoreOpenPullRequest + theRepo.ReviewCount*ScoreReview + theRepo.MergedPrCount*ScoreMerge)
		virtualDeveloperPulsePoint = math.Sqrt(virtualDeveloperPulsePoint)

		mean = currentDeveloperPulsePoint * virtualDeveloperPulsePoint / (currentDeveloperPulsePoint + virtualDeveloperPulsePoint)
		pulsePoint += mean

		virtualDeveloperId++
	}

	if pulsePointItem, err = l.svcCtx.PulsePointModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			if _, err = l.svcCtx.PulsePointModel.Insert(l.ctx, &model.PulsePoint{
				DataCreatedAt: time.Now(),
				DataUpdatedAt: time.Now(),
				DeveloperId:   id,
				PulsePoint:    0,
			}); err != nil {
				return
			}
			if pulsePointItem, err = l.svcCtx.PulsePointModel.FindOneByDeveloperId(l.ctx, id); err != nil {
				return
			}
		} else {
			return
		}
	}
	pulsePointItem.PulsePoint = pulsePoint
	if err = l.svcCtx.PulsePointModel.Update(l.ctx, pulsePointItem); err != nil {
		return
	}

	// maintain a rank of pulse point
	if _, err := l.svcCtx.RedisClient.ZaddFloatCtx(l.ctx, "pulse_point_rank", pulsePointItem.PulsePoint, strconv.FormatInt(id, 10)); err != nil {
		return err
	}

	return
}

func (l *UpdatePulsePointLogic) acquireUpdatePulsePointLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdatePulsePoint, id),
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

func (l *UpdatePulsePointLogic) checkIfNeedUpdatePulsePoint(id int64) (bool, error) {
	if pulsePoint, err := l.svcCtx.PulsePointModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(pulsePoint.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
