package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	// Create a new topic with the given name.
	topic, err := client.CreateTopic(ctx, "topicName")
	if err != nil {
		// TODO: Handle error.
	}

	sub, err := client.CreateSubscription(ctx, "subName", pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: 25 * time.Hour,
	})
	if err != nil {
		// TODO: Handle error.
	}

	msg := pubsub.Message{}

	sub.Receive(ctx, f  func(ctx, *msg))
}