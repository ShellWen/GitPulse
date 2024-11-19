package main

import (
	"fmt"
	"github.com/hibiken/asynq"
)

func main() {
	AsynqInspector := asynq.NewInspector(&asynq.RedisClientOpt{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	queues, err := AsynqInspector.Queues()
	if err != nil {
		return
	}

	fmt.Println(queues)
}
