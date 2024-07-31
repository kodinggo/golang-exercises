package main

import (
	"context"
	"log"
	"queue-messaging/task"

	"github.com/hibiken/asynq"
)

func main() {
	// redis opt
	redisOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	// init asynq server
	svr := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(task.TaskSendEmail, func(ctx context.Context, t *asynq.Task) error {
		log.Println("Sending email....")
		return nil
	})

	mux.HandleFunc(task.TaskUploadFile, func(ctx context.Context, t *asynq.Task) error {
		log.Println("Uploading file....")
		return nil
	})

	if err := svr.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
