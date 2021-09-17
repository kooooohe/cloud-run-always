package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	projectID := "midyear-spot-304113"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Faild to create client %v", err)
	}

	defer client.Close()

	topicID := "test"

	topic := client.Topic(topicID)
	// topic, err := client.CreateTopic(ctx, topicID)
	// if err != nil {
	// 	log.Fatalf("Faild to create client %v", err)
	// }
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	fmt.Printf("%v\n", topic)
	fmt.Printf("%v\n", res)
	msgID, err := res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgID)
}

