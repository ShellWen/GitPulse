package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"os"
)

const reviewFetcherTopic = "review"

type reviewWithRepoId struct {
	review *github.PullRequestReview
	repoId int64
}

func FetchReviewOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, userId, reviewFetcherTopic, doFetchReviewOfUser)
	return
}

func doFetchReviewOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allReview    []*reviewWithRepoId
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	if allReview, err = getAllGithubReviewByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryReview); err != nil {
		return
	}

	for _, githubReview := range allReview {
		if err = pushContribution(ctx, svcContext, buildReview(ctx, svcContext, githubReview, userId)); err != nil {
			return
		}
	}

	return
}

func getAllGithubReviewByLogin(ctx context.Context, githubClient *github.Client, login string) (allReviewWithRepoId []*reviewWithRepoId, err error) {
	var allIssue []*github.Issue
	allIssue, err = getAllGithubIssuePRByLogin(ctx, githubClient, login, RoleReviewer)

	opt := &github.ListOptions{PerPage: 100}
	for _, issue := range allIssue {
		reviews, _, err := githubClient.PullRequests.ListReviews(ctx, *issue.Repository.Owner.Login, *issue.Repository.Name, *issue.Number, opt)
		if err != nil {
			return nil, err
		}

		for _, review := range reviews {
			if review.User.GetLogin() == login {
				allReviewWithRepoId = append(allReviewWithRepoId, &reviewWithRepoId{review: review, repoId: *issue.Repository.ID})
			}
		}
	}

	return
}

func buildReview(ctx context.Context, svcContext *svc.ServiceContext, githubReviewWithRepoId *reviewWithRepoId, userId int64) (newReview *model.Contribution) {
	newReview = &model.Contribution{
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
