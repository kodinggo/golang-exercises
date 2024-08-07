package main

import (
	"encoding/json"
	"log"
	"pubsub-messaging/event"

	"github.com/nats-io/nats.go"
)

func main() {
	quitCh := make(chan bool)

	// connect to nats server
	nc, _ := nats.Connect("nats://127.0.0.1:4223")

	// create jetstream context from nats connection
	js, _ := nc.JetStream()

	js.AddStream(&nats.StreamConfig{
		Name:     event.ProductsTopic,
		Subjects: []string{event.ProductsEvent},
	})

	subscribeEvent(js)

	<-quitCh
}

func subscribeEvent(js nats.JetStreamContext) {
	log.Println("waiting for the incoming events..")
	var paylaod event.Product

	js.QueueSubscribe(event.CreateProductEvent, "test", func(msg *nats.Msg) {
		log.Println(string(msg.Data))

		err := json.Unmarshal(msg.Data, &paylaod)
		if err != nil {
			errHandler(err, dump(paylaod))
		}
	})
}

func errHandler(err error, payload any) {
	// send to slack
}

func dump(payload any) []byte {
	b, _ := json.Marshal(payload)
	return b
}
