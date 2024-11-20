package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/analysis"
	"github.com/ShellWen/GitPulse/analysis/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/contribution"
	"github.com/ShellWen/GitPulse/developer/cmd/rpc/developer"
	"net/http"
	"strings"
)

func getTextFromContribution(ctx context.Context, svcCtx *svc.ServiceContext, id int64, limitCharacterCount int64) (text string, err error) {
	var (
		updateReq  = contribution.UpdateContributionOfUserReq{UserId: id}
		updateResp *contribution.UpdateContributionOfUserResp

		getReq = contribution.SearchByUserIdReq{
			UserId: id,
			Limit:  1000,
			Page:   1,
		}
		getResp *contribution.SearchByUserIdResp
	)

	if updateResp, err = svcCtx.ContributionRpcClient.UpdateContributionOfUser(ctx, &updateReq); err != nil {
		return
	} else if updateResp.Code != http.StatusOK {
		return "", errors.New(updateResp.Message)
	}

	if getResp, err = svcCtx.ContributionRpcClient.SearchByUserId(ctx, &getReq); err != nil {
		return
	} else if getResp.Code == http.StatusInternalServerError {
		return "", errors.New(getResp.Message)
	}

	text += "|Contribution Start|"
	for _, theContribution := range getResp.Contributions {
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

func getTextFromProfile(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (text string, err error) {
	var (
		updateReq    = developer.UpdateDeveloperReq{Id: id}
		updateResp   *developer.UpdateDeveloperResp
		getReq       = developer.GetDeveloperByIdReq{Id: id}
		getResp      *developer.GetDeveloperByIdResp
		theDeveloper *developer.Developer
	)

	if updateResp, err = svcCtx.DeveloperRpcClient.UpdateDeveloper(ctx, &updateReq); err != nil || updateResp.Code != http.StatusOK {
		return
	}
	if getResp, err = svcCtx.DeveloperRpcClient.GetDeveloperById(ctx, &getReq); err != nil || getResp.Code != http.StatusOK {
		return
	}

	theDeveloper = getResp.Developer

	if theDeveloper == nil {
		err = errors.New("developer not found")
		return
	}

	text += "|Developer Profile Start|" + "|Name:" +
		theDeveloper.Name + "|Bio:" + theDeveloper.Bio + "|TwitterUsername:" + theDeveloper.TwitterUsername +
		"|Company:" + theDeveloper.Company + "|Location:" + theDeveloper.Location + "|Developer Profile End|"

	return
}

func getLanguageAsText(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (text string, err error) {
	var (
		updateLogic   = NewUpdateLanguageLogic(ctx, svcCtx)
		updateReq     = analysis.UpdateAnalysisReq{DeveloperId: id}
		updateResp    *analysis.UpdateAnalysisResp
		getLogic      = NewGetLanguagesLogic(ctx, svcCtx)
		getReq        = analysis.GetAnalysisReq{DeveloperId: id}
		getResp       *analysis.GetLanguagesResp
		languageUsage *analysis.Languages
	)

	if updateResp, err = updateLogic.UpdateLanguage(&updateReq); err != nil || updateResp.Code != http.StatusOK {
		return
	}
	if getResp, err = getLogic.GetLanguages(&getReq); err != nil || getResp.Code != http.StatusOK {
		return
	}

	languageUsage = getResp.Languages

	if languageUsage == nil {
		err = errors.New("language usage not found")
		return
	}

	text += "|Language Usage Start|" + languageUsage.Languages + "|Language Usage End|"
	return
}
