package customGithub

import (
	"context"
	"github.com/google/go-github/v66/github"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

const (
	DefaultSearchLimit int64         = 150
	DataExpiredTime    time.Duration = 24 * time.Hour
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

func GetLoginById(ctx context.Context, id int64) (login string, err error) {
	var githubClient *github.Client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_API_TOKEN"))
	var githubUser *github.User
	if githubUser, _, err = githubClient.Users.GetByID(ctx, id); err != nil {
		logx.Error("Unexpected error when fetching user ", login, " from github: ", err)
		return
	}
	login = githubUser.GetLogin()
	logx.Info("Successfully get login ", login, " of id", id)
	return
}

func DefaultUpdateAfterTime() string {
	return time.Unix(time.Now().Unix()-int64(7*24*time.Hour.Seconds()), 0).Format("2006-01-02")
}

func CheckIfDataExpired(lastUpdate time.Time) bool {
	return time.Now().Sub(lastUpdate) > DataExpiredTime
}
