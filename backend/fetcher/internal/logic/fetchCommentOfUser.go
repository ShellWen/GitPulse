package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"net/http"
	"os"
	"time"
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
	var issueResp *github.Response
	var prResp *github.Response
	for _, issue := range allIssue {
		var repo *github.Repository
		if repo, err = getRepoByUrl(ctx, githubClient, issue.GetRepositoryURL()); err != nil {
			return
		}

		for {
			var issueComments []*github.IssueComment
			if issueComments, issueResp, err = githubClient.Issues.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issue.GetNumber(), issueOpt); err != nil {
				if issueResp != nil && issueResp.StatusCode == http.StatusNotFound {
					err = nil
					break
				}
				return
			}

			for _, comment := range issueComments {
				if comment.User.GetLogin() == login {
					allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: true, issueComment: comment, repoId: repo.GetID()})
				}
			}
			if issueResp.NextPage == 0 {
				break
			}
			issueOpt.Page = issueResp.NextPage
		}

		if issue.IsPullRequest() {
			for {
				var prComments []*github.PullRequestComment
				if prComments, prResp, err = githubClient.PullRequests.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), 0, prOpt); err != nil {
					if prResp != nil && prResp.StatusCode == http.StatusNotFound {
						err = nil
						break
					}
					return
				}

				for _, comment := range prComments {
					if comment.User.GetLogin() == login {
						allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: false, prComment: comment, repoId: repo.GetID()})
					}
				}
				if prResp.NextPage == 0 {
					break
				}
				prOpt.Page = prResp.NextPage
			}
		}
	}

	return
}

func buildComment(ctx context.Context, svcContext *svc.ServiceContext, githubCommentWithRepoId *commentWithRepoId, userId int64) (newComment *model.Contribution) {
	if githubCommentWithRepoId.isIssueComment {
		newComment = &model.Contribution{
			DataCreatedAt:  time.Now(),
			DataUpdatedAt:  time.Now(),
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
			DataCreatedAt:  time.Now(),
			DataUpdatedAt:  time.Now(),
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
