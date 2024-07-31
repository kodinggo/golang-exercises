package main

import (
	"encoding/json"
	"fmt"
	"log"
	"queue-messaging/task"
	"sync"

	"github.com/hibiken/asynq"
)

func main() {
	// redis opt
	redisOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	// Init client
	client := asynq.NewClient(redisOpt)

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go sendEmail(&wg, client, i)
	}

	wg.Wait()

	log.Println("success publish an event")
}

func sendEmail(wg *sync.WaitGroup, client *asynq.Client, i int) {
	defer wg.Done()

	// prepare payloao
	payload, _ := json.Marshal(task.SendEmail{
		From:    "foo@gmail.com",
		To:      "bar@gmail.com",
		Subject: fmt.Sprintf("Testing %d", i),
		Content: "Testing",
	})

	// init task
	task := asynq.NewTask(task.TaskSendEmail, payload)

	// enqueue (publish)
	client.Enqueue(task, asynq.Queue("critical"), asynq.MaxRetry(3))
}

func uploadFile(client *asynq.Client) {
	// prepare payloao
	payload, _ := json.Marshal(task.UplaodFile{
		Origin:      "/path/to/origin/1.jpg",
		Destination: "/path/to/destination/1.jpg",
	})

	// init task
	task := asynq.NewTask(task.TaskUploadFile, payload)

	// enqueue (publish)
	client.Enqueue(task)
}
