package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"strconv"
	"time"
)

type APIType int8

const (
	APITaskExpireTime = time.Minute * 10
	APIMaxRetry       = 10
	APIRetryDelay     = time.Second * 10
)

const (
	APITaskName  = "api"
	APITaskQueue = "api"
)

const (
	APIGetDeveloper APIType = iota
	APIGetLanguage
	APIGetPulsePoint
	APIGetRegion
	APIGetSummary
)

type APIPayload struct {
	Type   APIType `json:"type"`
	Id     int64   `json:"id"`
	TaskId string  `json:"taskId"`
}

func GetNewAPITaskKey(fetchType APIType, id int64, reqId string) string {
	return APITaskName + separator + strconv.Itoa(int(fetchType)) + separator + strconv.Itoa(int(id)) + separator + reqId
}

func NewAPITask(fetchType APIType, id int64, reqId string) (*asynq.Task, string, error) {
	taskId := GetNewAPITaskKey(fetchType, id, reqId)
	payload, err := json.Marshal(APIPayload{
		Type:   fetchType,
		Id:     id,
		TaskId: taskId,
	})
	if err != nil {
		return nil, "", err
	}
	return asynq.NewTask(APITaskName, payload), taskId, nil
}
