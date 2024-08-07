package main

import (
	"encoding/json"
	"log"
	"pubsub-messaging/event"

	"github.com/nats-io/nats.go"
)

func main() {
	// connect to nats server
	nc, _ := nats.Connect("nats://127.0.0.1:4223")

	// create jetstream context from nats connection
	js, _ := nc.JetStream()

	js.AddStream(&nats.StreamConfig{
		Name:     event.ProductsTopic,
		Subjects: []string{event.ProductsEvent},
	})

	publishEvent(js)
}

func publishEvent(js nats.JetStreamContext) {
	byteProduct, _ := json.Marshal(event.Product{ID: 1, Name: "Item 1"})
	js.Publish(event.CreateProductEvent, byteProduct)

	log.Println("successfully send the event")
}
