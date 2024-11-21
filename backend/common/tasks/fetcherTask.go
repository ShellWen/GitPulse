package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"strconv"
	"time"
)

type FetchType int8

const (
	FetchMaxRetry   = 3
	FetchRetryDelay = time.Second * 5
)

const (
	FetcherTaskName  = "fetch"
	FetcherTaskQueue = "fetch"
)

const (
	FetchDeveloper = iota
	FetchCreatedRepo
	FetchStarredRepo

	FetchFollowing
	FetchFollower

	FetchIssuePROfUser
	FetchCommentOfUser
	FetchReviewOfUser

	FetchRepo
	FetchFork
)

const (
	FetchCreatedRepoCompletedDataId = -iota - 1

	FetchStarredRepoCompletedDataId
	FetchStarringDeveloperCompletedDataId

	FetchFollowingCompletedDataId
	FetchFollowerCompletedDataId

	FetchIssuePROfUserCompletedDataId
	FetchCommentOfUserCompletedDataId
	FetchReviewOfUserCompletedDataId

	FetchForkCompletedDataId
)

type FetchPayload struct {
	Type        FetchType `json:"type"`
	Id          int64     `json:"id"`
	UpdateAfter string    `json:"updateAfter"`
	SearchLimit int64     `json:"searchLimit"`
}

func getNewFetcherTaskKey(fetchType FetchType, id int64) string {
	return FetcherTaskName + separator + strconv.Itoa(int(fetchType)) + separator + strconv.Itoa(int(id))
}

func NewFetcherTask(fetchType FetchType, id int64, updateAfter string, searchLimit int64) (*asynq.Task, string, error) {
	payload, err := json.Marshal(FetchPayload{
		Type:        fetchType,
		Id:          id,
		UpdateAfter: updateAfter,
		SearchLimit: searchLimit,
	})
	if err != nil {
		return nil, "", err
	}
	return asynq.NewTask(FetcherTaskName, payload), getNewFetcherTaskKey(fetchType, id), nil
}
