package logic

import (
	"context"
	"github.com/ShellWen/GitPulse/contribution/model"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"strconv"
	"time"
)

const commentFetcherTopic = "comment"

type commentWithRepoId struct {
	isIssueComment bool
	issueComment   *github.IssueComment
	prComment      *github.PullRequestComment
	repoId         int64
}

func FetchCommentOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, searchLimit int64) (successPush int, err error) {
	if successPush, err = doFetchCommentOfUser(ctx, svcContext, userId, createAfter, searchLimit); err != nil {
		return
	}
	return
}

func doFetchCommentOfUser(ctx context.Context, svcContext *svc.ServiceContext, userId int64, createAfter string, searchLimit int64) (successPush int, err error) {
	var (
		githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
		githubUser   *github.User
		allComment   []*commentWithRepoId
		allRepo      map[int64]*github.Repository
	)

	if githubUser, _, err = getGithubUserById(ctx, githubClient, userId); err != nil {
		return
	}

	logx.Info("Start fetching comment of user: ", githubUser.GetLogin())
	if allComment, allRepo, err = getAllGithubCommentByLogin(ctx, svcContext, githubClient, githubUser.GetLogin(), createAfter, searchLimit); err != nil {
		return
	}
	logx.Info("Finish fetching comment of user: ", githubUser.GetLogin()+", total comment: "+strconv.Itoa(len(allComment)))

	if err = delAllOldContributionInCategory(ctx, svcContext, userId, model.CategoryComment); err != nil {
		return
	}

	logx.Info("Start pushing comment of user: ", githubUser.GetLogin())
	for _, githubComment := range allComment {
		if err = pushContribution(ctx, svcContext, buildComment(ctx, svcContext, githubComment, userId)); err != nil {
			continue
		}
		successPush++
	}

	for _, repo := range allRepo {
		if err = buildAndPushRepoByGithubRepo(ctx, svcContext, githubClient, repo); err != nil {
			continue
		}
	}

	logx.Info("Successfully pushed comment of user: ", githubUser.GetLogin()+", total comment: "+strconv.Itoa(successPush))
	return
}

func getAllGithubCommentByLogin(ctx context.Context, svcContext *svc.ServiceContext, githubClient *github.Client, login string, createAfter string, searchLimit int64) (allCommentWithRepoId []*commentWithRepoId, repos map[int64]*github.Repository, err error) {
	var allIssue []*github.Issue
	allIssue, err = getAllGithubIssuePRByLogin(ctx, githubClient, login, RoleCommenter, createAfter, searchLimit)

	updated := "updated"
	desc := "desc"
	createAfterTime, _ := time.Parse("2006-01-02", createAfter)

	issueOpt := &github.IssueListCommentsOptions{
		Sort:        &updated,
		Direction:   &desc,
		Since:       &createAfterTime,
		ListOptions: github.ListOptions{PerPage: int(searchLimit)},
	}
	prOpt := &github.PullRequestListCommentsOptions{
		Sort:        updated,
		Direction:   desc,
		Since:       createAfterTime,
		ListOptions: github.ListOptions{PerPage: int(searchLimit)},
	}

	var issueResp *github.Response
	var prResp *github.Response
	repos = make(map[int64]*github.Repository)

	for _, issue := range allIssue {
		if len(allCommentWithRepoId) >= int(searchLimit) {
			break
		}

		var repo *github.Repository
		if repo, err = getRepoByUrl(ctx, githubClient, issue.GetRepositoryURL()); err != nil {
			return
		}
		repos[repo.GetID()] = repo

		if issue.IsPullRequest() {
			var prComments []*github.PullRequestComment
			if prComments, prResp, err = githubClient.PullRequests.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), 0, prOpt); err != nil {
				if prResp == nil || prResp.StatusCode != http.StatusNotFound {
					return
				}
			}

			for _, comment := range prComments {
				if comment.User.GetLogin() == login {
					allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: false, prComment: comment, repoId: repo.GetID()})
				}
			}
		}
	}

	for _, issue := range allIssue {
		if len(allCommentWithRepoId) >= int(searchLimit) {
			break
		}

		var repo *github.Repository
		if repo, err = getRepoByUrl(ctx, githubClient, issue.GetRepositoryURL()); err != nil {
			return
		}
		repos[repo.GetID()] = repo

		var issueComments []*github.IssueComment
		if issueComments, issueResp, err = githubClient.Issues.ListComments(ctx, repo.GetOwner().GetLogin(), repo.GetName(), issue.GetNumber(), issueOpt); err != nil {
			if issueResp == nil || issueResp.StatusCode != http.StatusNotFound {
				return
			}
		}

		for _, comment := range issueComments {
			if comment.User.GetLogin() == login {
				allCommentWithRepoId = append(allCommentWithRepoId, &commentWithRepoId{isIssueComment: true, issueComment: comment, repoId: repo.GetID()})
			}
		}
	}

	if len(allCommentWithRepoId) > int(searchLimit) {
		allCommentWithRepoId = allCommentWithRepoId[:int(searchLimit)]
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
			Category:       model.CategoryComment,
			Content:        githubCommentWithRepoId.prComment.GetBody(),
			CreatedAt:      githubCommentWithRepoId.prComment.GetCreatedAt().Time,
			UpdatedAt:      githubCommentWithRepoId.prComment.GetUpdatedAt().Time,
			ContributionId: githubCommentWithRepoId.prComment.GetID(),
		}
	}
	return
}
