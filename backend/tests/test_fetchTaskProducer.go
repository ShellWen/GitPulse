package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func main() {
	var err error

	pusher := kq.NewPusher([]string{"localhost:9092"}, "fetcher-task", kq.WithAllowAutoTopicCreation())

	task := tasks.FetcherTask{
		Type: tasks.FetchDeveloper,
		Id:   105362324,
	}

	var marshalString string

	if marshalString, err = jsonx.MarshalToString(task); err != nil {
		panic(err)
	}

	fmt.Println(marshalString)

	if err = pusher.Push(context.Background(), marshalString); err != nil {
		panic(err)
	}

	println("pushed")
}
