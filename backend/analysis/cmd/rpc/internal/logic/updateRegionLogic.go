package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/model"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/common/llm"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"github.com/biter777/countries"
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
		regionConfidence map[string]float64
		regionItem       *model.Region

		mostPossibleRegion     string
		mostPossibleConfidence float64 = 0
	)

	if login, err = customGithub.GetLoginById(l.ctx, id); err != nil {
		return
	}

	if regionConfidence, err = l.getRegionWithConfidenceByPythonScript(login); err != nil {
		return
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
		req            *http.Request
		resp           *http.Response
		role           = "你是一名专业的 GitHub 数据分析师，对不同地区、语言和文化有深刻了解。GitHub 的主要语言为英语，但你可以通过用户使用的非英语语言、地名和位置等信息来推测其国籍或地区。" +
			"\n任务：我将提供某 GitHub 用户的简介、Issue、Pull Request 或 Comment。请分析其内容，判断用户的最可能国籍或地区名。" +
			"\n回复格式：请仅根据以下格式，不要回复其他内容。" +
			"\n{\"region\": \"regionName\", \"confidence\": confidenceValue}" +
			"\n\"region\"：请填入分析出的国籍或地区名（使用英文）。" +
			"\n\"confidence\"：请填入置信度值，表示对该判断的确信程度。" +
			"请务必始终保持该状态。"
		allText          = ""
		jsonStr          string
		body             []byte
		regionConfidence struct {
			Region     string  `json:"region"`
			Confidence float64 `json:"confidence"`
		}
	)

	if text, err := l.getTextFromProfile(id); err != nil {
		return "", 0, err
	} else {
		allText += text
	}

	if text, err := l.getTextFromContribution(id, 1000); err != nil {
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

	if err = json.Unmarshal([]byte(sparkModelResp.Choices[0].Message.Content), &regionConfidence); err != nil {
		return
	}

	region = regionConfidence.Region
	region = strings.ToLower(countries.ByName(region).Alpha2())
	confidence = regionConfidence.Confidence

	return
}

func (l *UpdateRegionLogic) getTextFromContribution(id int64, limitCharacterCount int64) (text string, err error) {
	var (
		req = contribution.SearchByUserIdReq{
			UserId: id,
			Limit:  1000,
			Page:   1,
		}
		resp *contribution.SearchByUserIdResp
	)

	if resp, err = l.svcCtx.ContributionRpcClient.SearchByUserId(l.ctx, &req); err != nil {
		return
	}

	text += "|Contribution Start|"
	for _, theContribution := range resp.Contributions {
		text += theContribution.Content
		if int64(len(text)) > limitCharacterCount {
			break
		}
	}
	strings.ReplaceAll(text, "\n", " ")
	if int64(len(text)) > limitCharacterCount {
		text = text[:limitCharacterCount]
	}
	text += "|Contribution End|"

	return
}

func (l *UpdateRegionLogic) getTextFromProfile(id int64) (text string, err error) {
	var (
		req          = developer.GetDeveloperByIdReq{Id: id}
		resp         *developer.GetDeveloperByIdResp
		theDeveloper *developer.Developer
	)

	if resp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &req); err != nil {
		return
	}

	theDeveloper = resp.Developer

	if theDeveloper == nil {
		err = errors.New("developer not found")
		return
	}

	text += "|Developer Profile Start|" + "|Name:" +
		theDeveloper.Name + "|Bio:" + theDeveloper.Bio + "|TwitterUsername:" + theDeveloper.TwitterUsername +
		"|Company:" + theDeveloper.Company + "|Location:" + theDeveloper.Location + "|Developer Profile End|"

	return
}
