package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/ShellWen/GitPulse/rank/internal/svc"
	"github.com/ShellWen/GitPulse/rank/internal/types"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPulsePointRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPulsePointRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPulsePointRankLogic {
	return &GetPulsePointRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPulsePointRankLogic) GetPulsePointRank(req *types.GetPulsePointRankReq) (resp []types.DeveloperWithPulsePoint, err error) {
	var (
		rank []redis.FloatPair
	)

	resp = make([]types.DeveloperWithPulsePoint, 0, len(rank))

	if rank, err = l.svcCtx.RedisClient.ZrevrangeWithScoresByFloatCtx(l.ctx, "pulse_point_rank", 0, 100); err != nil {
		logx.Error(err)
		return
	}

	for _, pair := range rank {
		if req.Limit != 0 && int64(len(resp)) >= req.Limit {
			break
		}

		var (
			developerId       int64
			pulsePointScore   = pair.Score
			getDeveloperResp  *developer.GetDeveloperByIdResp
			region            string
			languages         = make(map[string]float64)
			getPulsePointResp *analysis.GetPulsePointResp
			typeDeveloper     types.Developer
			typePulsePoint    types.PulsePoint
			result            types.DeveloperWithPulsePoint
		)

		if developerId, err = strconv.ParseInt(pair.Key, 10, 64); err != nil {
			logx.Error(err)
			return
		}

		// filter
		if req.Region != "" {
			if region, _, err = l.getRegionById(developerId); err != nil {
				logx.Error(err)
				continue
			}
			if region != req.Region {
				continue
			}
		}

		if req.Language != "" {
			if languages, err = l.getLanguagesById(developerId); err != nil {
				logx.Error(err)
				continue
			}
			if _, ok := languages[req.Language]; !ok {
				continue
			}
		}

		if getDeveloperResp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
			Id: developerId,
		}); err != nil {
			logx.Error(err)
			continue
		}
		if getDeveloperResp.Code != http.StatusOK {
			logx.Error(err)
			continue
		}

		typeDeveloper = types.Developer{
			Id:        getDeveloperResp.Developer.Id,
			Name:      getDeveloperResp.Developer.Name,
			Login:     getDeveloperResp.Developer.Login,
			AvatarUrl: getDeveloperResp.Developer.AvatarUrl,
			Company:   getDeveloperResp.Developer.Company,
			Location:  getDeveloperResp.Developer.Location,
			Bio:       getDeveloperResp.Developer.Bio,
			Blog:      getDeveloperResp.Developer.Blog,
			Email:     getDeveloperResp.Developer.Email,
			CreatedAt: time.Unix(getDeveloperResp.Developer.CreatedAt, 0).Format(time.RFC3339),
			UpdatedAt: time.Unix(getDeveloperResp.Developer.UpdatedAt, 0).Format(time.RFC3339),
			Following: getDeveloperResp.Developer.Following,
			Followers: getDeveloperResp.Developer.Followers,
			Gists:     getDeveloperResp.Developer.Gists,
			Stars:     getDeveloperResp.Developer.Stars,
			Repos:     getDeveloperResp.Developer.Repos,
		}

		if getPulsePointResp, err = l.svcCtx.AnalysisRpcClient.GetPulsePoint(l.ctx, &analysis.GetAnalysisReq{
			DeveloperId: developerId,
		}); err != nil {
			logx.Error(err)
			continue
		}
		if getPulsePointResp.Code != http.StatusOK {
			logx.Error(err)
			continue
		}

		typePulsePoint = types.PulsePoint{
			Id:         developerId,
			PulsePoint: pulsePointScore,
			UpdatedAt:  time.Unix(getPulsePointResp.PulsePoint.DataUpdatedAt, 0).Format(time.RFC3339),
		}

		result = types.DeveloperWithPulsePoint{
			Developer:  typeDeveloper,
			PulsePoint: typePulsePoint,
		}

		resp = append(resp, result)
	}

	return resp, nil
}

func (l *GetPulsePointRankLogic) getRegionById(id int64) (region string, confidence float64, err error) {
	var getRegionResp *analysis.GetRegionResp

	if getRegionResp, err = l.svcCtx.AnalysisRpcClient.GetRegion(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	}); err != nil {
		logx.Error(err)
		return
	}
	if getRegionResp.Code != http.StatusOK {
		err = errors.New(getRegionResp.Message)
		return
	}

	region = getRegionResp.Region.Region
	confidence = getRegionResp.Region.Confidence

	return
}

func (l *GetPulsePointRankLogic) getLanguagesById(developerId int64) (languages map[string]float64, err error) {
	var (
		getLanguagesResp *analysis.GetLanguagesResp
	)

	if getLanguagesResp, err = l.svcCtx.AnalysisRpcClient.GetLanguages(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: developerId,
	}); err != nil {
		logx.Error(err)
		return
	}
	if getLanguagesResp.Code != http.StatusOK {
		err = errors.New(getLanguagesResp.Message)
		return
	}

	if err = json.Unmarshal([]byte(getLanguagesResp.Languages.Languages), &languages); err != nil {
		logx.Error(err)
		return
	}

	return
}
