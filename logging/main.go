package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func main() {
	// if os.Getenv("ENV") == "development" {
	log.SetLevel(log.DebugLevel)
	// }

	log.SetFormatter(&log.JSONFormatter{})

	log.Trace("test trace")
	log.Debug("test debug")
	log.Info("test info")
	log.Error("test error")
	// log.Fatal("test fatal")
	// log.Panic("test panic")

	addToCartController(context.Background(), []string{"a", "b"})
}

func addToCartController(ctx context.Context, items []string) {
	logger := log.WithFields(log.Fields{
		"ctx":   ctx,
		"items": items,
	})
	logger.Error("test error")
}
