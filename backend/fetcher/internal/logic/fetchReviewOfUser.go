package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/common/tasks"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"strconv"
	"time"
)

const reviewFetcherTopic = "review"

type reviewWithRepoId struct {
	review *github.PullRequestReview
	repoId int64
}

func FetchReviewOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, searchLimit int64) (number int, err error) {
	if number, err = doFetchReviewOfUser(ctx, svcContext, userId, createAfter, searchLimit); err != nil {
		return
	}
	return
}

func doFetchReviewOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, searchLimit int64) (successPush int, err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allReview    []*reviewWithRepoId
		allRepo      map[int64]*github.Repository
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching review of user: ", githubUser.GetLogin())
	if allReview, allRepo, err = getAllGithubReviewByLogin(ctx, svcContext, githubClient, githubUser.GetLogin(), createAfter, searchLimit); err != nil {
		return
	}
	logx.Info("Finish fetching review of user: ", githubUser.GetLogin()+", total review: "+strconv.Itoa(len(allReview)))

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryReview); err != nil {
		return
	}

	logx.Info("Start pushing review of user: ", githubUser.GetLogin())
	for _, githubReview := range allReview {
		if err = pushContribution(ctx, svcContext, buildReview(ctx, svcContext, githubReview, userId)); err != nil {
			continue
		}
		successPush++
	}

	for _, repo := range allRepo {
		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, repo); err != nil {
			continue
		}
	}

	if err = pushFetchReviewOfUserCompleted(ctx, svcContext, userId); err != nil {
		return
	}

	logx.Info("Successfully pushed review of user: ", githubUser.GetLogin()+", total review: "+strconv.Itoa(successPush))
	return
}

func getAllGithubReviewByLogin(ctx context.Context, svcContext *svc.ServiceContext, githubClient *github.Client, login string, createAfter string, searchLimit int64) (allReviewWithRepoId []*reviewWithRepoId, repos map[int64]*github.Repository, err error) {
	var allIssue []*github.Issue
	allIssue, err = getAllGithubIssuePRByLogin(ctx, githubClient, login, RoleReviewer, createAfter, searchLimit)

	opt := &github.ListOptions{
		PerPage: 100,
	}

	var prResp *github.Response
	repos = make(map[int64]*github.Repository)

	for _, issue := range allIssue {
		if len(allReviewWithRepoId) >= int(searchLimit) {
			break
		}

		if issue.IsPullRequest() {
			var prReviews []*github.PullRequestReview

			var repo *github.Repository
			if repo, err = getRepoByUrl(ctx, githubClient, issue.GetRepositoryURL()); err != nil {
				return
			}
			repos[repo.GetID()] = repo

			if prReviews, prResp, err = githubClient.PullRequests.ListReviews(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issue.GetNumber(), opt); err != nil {
				if prResp == nil || prResp.StatusCode != http.StatusNotFound {
					return
				}
			}

			for _, review := range prReviews {
				if review.GetUser().GetLogin() == login {
					allReviewWithRepoId = append(allReviewWithRepoId, &reviewWithRepoId{review: review, repoId: repo.GetID()})
					logx.Info("Found review of user: ", login, " in repo: ", repo.GetFullName())
				}
			}
		}
	}

	if len(allReviewWithRepoId) > int(searchLimit) {
		allReviewWithRepoId = allReviewWithRepoId[:int(searchLimit)]
	}

	logx.Info("Finish fetching review of user: ", login+", total review: "+strconv.Itoa(len(allReviewWithRepoId)))
	return
}

func buildReview(ctx context.Context, svcContext *svc.ServiceContext, githubReviewWithRepoId *reviewWithRepoId, userId int64) (newReview *model.Contribution) {
	newReview = &model.Contribution{
		DataCreatedAt:  time.Now(),
		DataUpdatedAt:  time.Now(),
		UserId:         userId,
		RepoId:         githubReviewWithRepoId.repoId,
		Category:       model.CategoryReview,
		Content:        githubReviewWithRepoId.review.GetBody(),
		CreatedAt:      githubReviewWithRepoId.review.GetSubmittedAt().Time,
		UpdatedAt:      githubReviewWithRepoId.review.GetSubmittedAt().Time,
		ContributionId: githubReviewWithRepoId.review.GetID(),
	}
	return
}

func pushFetchReviewOfUserCompleted(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	if err = pushContribution(ctx, svcContext, &model.Contribution{
		DataId: tasks.FetchReviewOfUserCompletedDataId,
		UserId: userId,
	}); err != nil {
		return
	}
	return
}
