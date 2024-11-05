package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

type FetchType int8

const FetchTaskName = "fetch"

const (
	FetchDeveloper = iota
	FetchCreatedRepo
	FetchStarredRepo
	FetchFollow
	FetchFollowing
	FetchFollower
	FetchContributionOfUser
	FetchIssuePROfUser
	FetchCommentOfUser

	FetchRepo
	FetchFork
)

type FetchPayload struct {
	Type FetchType `json:"type"`
	Id   int64     `json:"id"`
}

func NewFetcherTask(fetchType FetchType, id int64) (*asynq.Task, error) {
	payload, err := json.Marshal(FetchPayload{Type: fetchType, Id: id})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(FetchTaskName, payload), nil
}
