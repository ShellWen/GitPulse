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
		role           = "你是一名专业的GitHub数据分析师。同时，你对世界不同地区、语言、文化有着非常深刻的了解。" +
			"GitHub是一个代码协作平台，该平台的主语言以英文为主，可以通过多种方式来猜测用户可能的区域，" +
			"例如使用的非英语语言、文化地区元素、位置、地名等等。请主要以语言、位置、地名来判别用户国籍或地区名。" +
			"接下来，我将给你GitHub上某位用户的简介、Issue、PullRequest或Comment。" +
			"你的任务是对GitHub上某位用户的简介、Issue、PullRequest或Comment进行分析，分析该用户可能的所在国籍或地区名。" +
			"你只需要以{\"region\": regionName, \"confidence\": confidenceValue}格式进行回复，将regionName替换为" +
			"你所判别的国籍或地区名，confidenceValue替换为置信度, 表示用户最可能的所在国籍（或地区名）及置信度。" +
			"Region使用英文表示。请不要回复其他任何无关文字。请不要回复Unknown" +
			"请一直保持该状态，不要停止。"
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
