package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe("8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	projectID := "midyear-spot-304113"

	go func() {
		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Faild to create client %v", err)
		}

		defer client.Close()
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Minute * 1)
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
	}()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}
