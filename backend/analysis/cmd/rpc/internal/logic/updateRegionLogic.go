package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/llm"
	locks "github.com/ShellWen/GitPulse/common/lock"
	"github.com/biter777/countries"
	"github.com/hashicorp/consul/api"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"time"

	"net/http"
	"os/exec"
	"strings"

	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const LLModelConfidenceWeight = 0.2

type UpdateRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRegionLogic {
	return &UpdateRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRegionLogic) UpdateRegion(in *pb.UpdateAnalysisReq) (resp *pb.UpdateAnalysisResp, err error) {
	if err = l.doUpdateRegion(in.DeveloperId); err != nil {
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

func (l *UpdateRegionLogic) doUpdateRegion(id int64) (err error) {
	var (
		login            string
		regionConfidence = make(map[string]float64)
		regionItem       *model.Region

		mostPossibleRegion     string
		mostPossibleConfidence float64 = 0
	)

	lock, err := l.acquireUpdateRegionLock(id)
	if err != nil {
		return
	}
	defer lock.Unlock()

	needUpdate, err := l.checkIfNeedUpdateRegion(id)
	if err != nil {
		return
	}

	if !needUpdate {
		return
	}

	if login, err = customGithub.GetLoginById(l.ctx, id); err != nil {
		return
	}

	if regionConfidence, err = l.getRegionWithConfidenceByPythonScript(login); err != nil {
		logx.Error(err)
		err = nil
	}

	if regionConfidence == nil {
		regionConfidence = make(map[string]float64)
	}

	if region, confidence, err := l.getRegionWithConfidenceByLLModel(id, login); err != nil {
		return err
	} else {
		regionConfidence[region] += confidence * LLModelConfidenceWeight
	}

	for region, confidence := range regionConfidence {
		if confidence > mostPossibleConfidence {
			mostPossibleRegion = region
			mostPossibleConfidence = confidence
		}
	}

	if regionItem, err = l.svcCtx.RegionModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			if _, err = l.svcCtx.RegionModel.Insert(l.ctx, &model.Region{
				DataCreatedAt: time.Now(),
				DataUpdatedAt: time.Now(),
				DeveloperId:   id,
				Region:        "",
				Confidence:    0,
			}); err != nil {
				return
			}
			if regionItem, err = l.svcCtx.RegionModel.FindOneByDeveloperId(l.ctx, id); err != nil {
				return
			}
		} else {
			return
		}
	}
	regionItem.Region = mostPossibleRegion
	regionItem.Confidence = mostPossibleConfidence
	if err = l.svcCtx.RegionModel.Update(l.ctx, regionItem); err != nil {
		return
	}

	return
}

func (l *UpdateRegionLogic) getRegionWithConfidenceByPythonScript(login string) (regionConfidence map[string]float64, err error) {
	var (
		cmd *exec.Cmd
		out []byte
	)

	cmd = exec.Command("venv/bin/python", "script/guess_region/main.py", login)
	if out, err = cmd.CombinedOutput(); err != nil {
		err = errors.New(err.Error() + " : " + string(out))
		return
	}

	if err = json.Unmarshal(out, &regionConfidence); err != nil {
		return
	}

	return
}

func (l *UpdateRegionLogic) getRegionWithConfidenceByLLModel(id int64, login string) (region string, confidence float64, err error) {
	var (
		httpClient     = &http.Client{}
		sparkModelData llm.SparkModelData
		sparkModelResp llm.SparkModelResp
		respStr        string
		req            *http.Request
		resp           *http.Response
		role           = "你是一名专业的 GitHub 数据分析师，对不同地区、语言和文化有深刻了解。GitHub 的主要语言为英语，但你可以通过用户使用的非英语语言、地名和位置等信息来推测其国籍或地区。" +
			"\n任务：我将提供某 GitHub 用户的简介、Issue、Pull Request 或 Comment。请分析其内容，判断用户的最可能国籍或地区名。" +
			"\n回复格式：请仅根据以下格式，不要回复其他内容。" +
			"\n{\"region\": \"regionName\", \"confidence\": confidenceValue}" +
			"\n\"region\"：请填入分析出的国籍或地区名（使用英文）。" +
			"\n\"confidence\"：请填入置信度值，表示对该判断的确信程度。该数值应在 0 到 1 之间，且小数点后最多保留 2 位。" +
			"请务必始终保持该状态。"
		allText          = ""
		jsonStr          string
		body             []byte
		regionConfidence struct {
			Region     string  `json:"region"`
			Confidence float64 `json:"confidence"`
		}
	)

	if text, err := getTextFromProfile(l.ctx, l.svcCtx, id); err != nil {
		return "", 0, err
	} else {
		allText += text
	}

	if text, err := getTextFromContribution(l.ctx, l.svcCtx, id, 1000); err != nil {
		return "", 0, err
	} else {
		allText += text
	}

	sparkModelData.MaxTokens = l.svcCtx.Config.SparkModelConf.MaxTokens
	sparkModelData.TopK = l.svcCtx.Config.SparkModelConf.TopK
	sparkModelData.Temperature = l.svcCtx.Config.SparkModelConf.Temperature
	sparkModelData.Messages[0].Role = "system"
	sparkModelData.Messages[0].Content = role
	sparkModelData.Messages[1].Role = "user"
	sparkModelData.Messages[1].Content = allText
	sparkModelData.Model = l.svcCtx.Config.SparkModelConf.Model

	if jsonStr, err = jsonx.MarshalToString(sparkModelData); err != nil {
		return
	}

	if req, err = http.NewRequest("POST", l.svcCtx.Config.SparkModelConf.Url, strings.NewReader(jsonStr)); err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+l.svcCtx.Config.SparkModelConf.APIPassword)
	req.Close = true

	if resp, err = httpClient.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &sparkModelResp); err != nil {
		return
	}

	if len(sparkModelResp.Choices) == 0 {
		return
	}

	respStr = sparkModelResp.Choices[0].Message.Content

	respStr = l.extractJson(respStr)

	if err = json.Unmarshal([]byte(respStr), &regionConfidence); err != nil {
		return
	}

	region = regionConfidence.Region
	region = strings.ToLower(countries.ByName(region).Alpha2())
	confidence = regionConfidence.Confidence
	confidence = l.restrictConfidence(confidence)

	return
}

func (l *UpdateRegionLogic) extractJson(text string) (newText string) {
	splitted := strings.Split(text, "{")

	if len(splitted) < 2 {
		return
	}

	text = "{" + splitted[1]

	splitted = strings.Split(text, "}")

	if len(splitted) < 2 {
		return
	}

	text = splitted[0] + "}"

	newText = text

	return
}

func (l *UpdateRegionLogic) restrictConfidence(confidence float64) (restrictedConfidence float64) {
	if confidence < 0 {
		restrictedConfidence = 0
	} else if confidence > 1 {
		restrictedConfidence = 1
	} else {
		restrictedConfidence = confidence
	}
	return
}

func (l *UpdateRegionLogic) acquireUpdateRegionLock(id int64) (*api.Lock, error) {
	lock, err := l.svcCtx.ConsulClient.LockOpts(&api.LockOptions{
		Key:         locks.GetNewLockKey(locks.UpdateRegion, id),
		Value:       []byte("locked"),
		SessionTTL:  "10s",
		SessionName: "update_region",
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

func (l *UpdateRegionLogic) checkIfNeedUpdateRegion(id int64) (bool, error) {
	if region, err := l.svcCtx.RegionModel.FindOneByDeveloperId(l.ctx, id); err != nil {
		switch {
		case errors.Is(err, model.ErrNotFound):
			return true, nil
		default:
			return false, err
		}
	} else {
		if customGithub.CheckIfDataExpired(region.DataUpdatedAt) {
			return true, nil
		} else {
			return false, nil
		}
	}
}
