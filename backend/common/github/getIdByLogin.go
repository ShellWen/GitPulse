package customGithub

import (
	"context"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func GetIdByLogin(ctx context.Context, login string) (id int64, err error) {
	var githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
	var githubUser *github.User
	if githubUser, _, err = githubClient.Users.Get(ctx, login); err != nil {
		logx.Error("Unexpected error when fetching user ", login, " from github: ", err)
		return
	}
	id = githubUser.GetID()
	logx.Info("Successfully get id ", id, " of login", login)
	return
}
