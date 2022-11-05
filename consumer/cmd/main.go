package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e.Error())
		log.Fatalf("Error loading .env file")
	}
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    kafkaTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}(reader)

	log.Println("start consuming ... !!")

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}
}
