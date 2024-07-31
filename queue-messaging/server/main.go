package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"queue-messaging/task"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// redis opt
	redisOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	// init asynq server
	svr := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 1,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			// log.Println("task ID", task.ResultWriter().TaskID())
			log.Println("task payload", string(task.Payload()))
		}),
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(task.TaskSendEmail, func(ctx context.Context, t *asynq.Task) error {
		time.Sleep(1 * time.Second)

		// Mendapatkan paylaod
		var payload task.SendEmail
		_ = json.Unmarshal(t.Payload(), &payload)

		if payload.Subject == "Testing 10" {
			return errors.New("error")
		}

		log.Printf("Sending email, subject %s....", payload.Subject)
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
