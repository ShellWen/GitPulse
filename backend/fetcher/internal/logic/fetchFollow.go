package logic

import (
	"context"
	"errors"
	"github.com/ShellWen/GitPulse/fetcher/internal/svc"
	"github.com/ShellWen/GitPulse/relation/model"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func buildFollow(ctx context.Context, svcContext *svc.ServiceContext, followerId int64, followingId int64) (follow *model.Follow) {
	return &model.Follow{
		FollowerId:  followerId,
		FollowingId: followingId,
	}
}

func pushFollow(ctx context.Context, svcContext *svc.ServiceContext, follow *model.Follow) (err error) {
	var jsonStr string

	if jsonStr, err = jsonx.MarshalToString(follow); err != nil {
		logx.Error(errors.New("Unexpected error when marshalling follow: " + err.Error()))
		return
	}

	if err = svcContext.KqFollowPusher.Push(ctx, jsonStr); err != nil {
		logx.Error(errors.New("Unexpected error when pushing follow: " + err.Error()))
		return
	}

	return
}
