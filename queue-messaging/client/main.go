package main

import (
	"encoding/json"
	"queue-messaging/task"

	"github.com/hibiken/asynq"
)

func main() {
	// redis opt
	redisOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	// Init client
	client := asynq.NewClient(redisOpt)

	// prepare payloao
	payload, _ := json.Marshal(task.SendEmail{
		From:    "foo@gmail.com",
		To:      "bar@gmail.com",
		Subject: "Testing",
		Content: "Testing",
	})

	// init task
	task := asynq.NewTask(task.TaskSendEmail, payload)

	// enqueue (publish)
	client.Enqueue(task)
}
