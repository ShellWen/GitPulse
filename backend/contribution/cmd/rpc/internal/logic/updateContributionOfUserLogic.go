package logic

import (
	"context"
	customGithub "github.com/ShellWen/GitPulse/common/github"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/internal/svc"
	"github.com/ShellWen/GitPulse/contribution/cmd/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UpdateContributionOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateContributionOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContributionOfUserLogic {
	return &UpdateContributionOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateContributionOfUserLogic) UpdateContributionOfUser(in *pb.UpdateContributionOfUserReq) (*pb.UpdateContributionOfUserResp, error) {
	err := l.doUpdateContributionOfUser(in.UserId, in.UpdateAfter, in.SearchLimit)
	if err != nil {
		return &pb.UpdateContributionOfUserResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdateContributionOfUserResp{
		Code:    http.StatusOK,
		Message: "Successfully updated contribution of user",
	}, nil
}

func (l *UpdateContributionOfUserLogic) doUpdateContributionOfUser(id int64, updateAfter string, searchLimit int64) error {
	if updateAfter == "" {
		updateAfter = customGithub.DefaultUpdateAfterTime()
	}

	if searchLimit == 0 {
		searchLimit = customGithub.DefaultSearchLimit
	}

	updateReviewLogic := NewUpdateReviewOfUserLogic(l.ctx, l.svcCtx)
	err := updateReviewLogic.doUpdateReviewOfUser(&pb.UpdateReviewOfUserReq{
		UserId:      id,
		UpdateAfter: updateAfter,
		SearchLimit: searchLimit,
	})
	if err != nil {
		logx.Error("Failed to update review of user: ", err)
		return err
	}

	updateIssuePROfUserLogic := NewUpdateIssuePROfUserLogic(l.ctx, l.svcCtx)
	err = updateIssuePROfUserLogic.doUpdateIssuePROfUser(&pb.UpdateIssuePROfUserReq{
		UserId:      id,
		UpdateAfter: updateAfter,
		SearchLimit: searchLimit,
	})
	if err != nil {
		logx.Error("Failed to update issue pr of user: ", err)
		return err
	}

	updateCommentOfUserLogic := NewUpdateCommentOfUserLogic(l.ctx, l.svcCtx)
	err = updateCommentOfUserLogic.doUpdateCommentOfUser(&pb.UpdateCommentOfUserReq{
		UserId:      id,
		UpdateAfter: updateAfter,
		SearchLimit: searchLimit,
	})
	if err != nil {
		logx.Error("Failed to update comment of user: ", err)
		return err
	}

	logx.Info("Successfully updated contribution of user")
	return nil
}
