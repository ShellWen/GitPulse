package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/repo/cmd/rpc/repo"
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
	ScoreReview          = 4
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
		allContributionResp                 *contribution.SearchByUserIdResp
		allContributions                    []*contribution.Contribution
		allContributionsCategorizedByRepoId = make(map[int64][]*contribution.Contribution)
		pulsePoint                          float64
		pulsePointItem                      *model.PulsePoint
	)

	if allContributionResp, err = contributionZrpcClient.SearchByUserId(l.ctx, &contribution.SearchByUserIdReq{
		UserId: id,
		Limit:  1000,
		Page:   1,
	}); err != nil {
		return
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
