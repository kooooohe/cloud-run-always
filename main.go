package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/start", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	topicID := os.Getenv("TOPIC_ID")

	go func() {
		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Faild to create client %v", err)
		}

		defer client.Close()
		for i := 0; i < 30000; i++ {

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
	fmt.Fprintf(w, "Start!")
}
