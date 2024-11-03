package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"os"
)

const commentFetcherTopic = "comment"

type commentWithRepoId struct {
	isIssueComment bool
	issueComment   *github.IssueComment
	prComment      *github.PullRequestComment
	repoId         int64
}

func FetchCommentOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	err = fetchWithRetry(ctx, svcContext, userId, commentFetcherTopic, doFetchCommentOfUser)
	return
}

func doFetchCommentOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64) (err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allComment   []*commentWithRepoId
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	if allComment, err = getAllGithubCommentByLogin(ctx, githubClient, githubUser.GetLogin()); err != nil {
		return
	}

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryComment); err != nil {
		return
	}

	for _, githubComment := range allComment {
		if err = pushContribution(ctx, svcContext, buildComment(ctx, svcContext, githubComment, userId)); err != nil {
			return
		}
	}

	return
}

func getAllGithubCommentByLogin(ctx context.Context, githubClient *github.Client, login string) (allCommentWithRepoId []*commentWithRepoId, err error) {
	var allIssue []*github.Issue
	allIssue, err = getAllGithubIssuePRByLogin(ctx, githubClient, login, RoleCommenter)

	issueOpt := &github.IssueListCommentsOptions{ListOptions: github.ListOptions{PerPage: 100}}
	prOpt := &github.PullRequestListCommentsOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for _, issue := range allIssue {
		issueComments, _, err := githubClient.Issues.ListComments(ctx, *issue.Repository.Owner.Login, *issue.Repository.Name, *issue.Number, issueOpt)

		if err != nil {
			return nil, err
		}

		for _, comment := range issueComments {
			if comment.User.GetLogin() == login {
				allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: true, issueComment: comment, repoId: *issue.Repository.ID})
			}
		}

		if issue.IsPullRequest() {
			prComments, _, err := githubClient.PullRequests.ListComments(ctx, *issue.Repository.Owner.Login, *issue.Repository.Name, *issue.Number, prOpt)
			if err != nil {
				return nil, err
			}

			for _, comment := range prComments {
				if comment.User.GetLogin() == login {
					allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: false, prComment: comment, repoId: *issue.Repository.ID})
				}
			}
		}
	}

	return
}

func buildComment(ctx context.Context, svcContext *svc.ServiceContext, githubCommentWithRepoId *commentWithRepoId, userId int64) (newComment *model.Contribution) {
	if githubCommentWithRepoId.isIssueComment {
		newComment = &model.Contribution{
			UserId:         userId,
			RepoId:         githubCommentWithRepoId.repoId,
			Category:       model.CategoryComment,
			Content:        githubCommentWithRepoId.issueComment.GetBody(),
			CreatedAt:      githubCommentWithRepoId.issueComment.GetCreatedAt().Time,
			UpdatedAt:      githubCommentWithRepoId.issueComment.GetUpdatedAt().Time,
			ContributionId: githubCommentWithRepoId.issueComment.GetID(),
		}
	} else {
		newComment = &model.Contribution{
			UserId:         userId,
			RepoId:         githubCommentWithRepoId.repoId,
			Category:       model.CategoryReview,
			Content:        githubCommentWithRepoId.prComment.GetBody(),
			CreatedAt:      githubCommentWithRepoId.prComment.GetCreatedAt().Time,
			UpdatedAt:      githubCommentWithRepoId.prComment.GetUpdatedAt().Time,
			ContributionId: githubCommentWithRepoId.prComment.GetID(),
		}
	}
	return
}
