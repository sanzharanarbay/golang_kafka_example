package kafka

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func ConnectKafka() *kafka.Conn {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e.Error())
		log.Fatalf("Error loading .env file")
	}
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, kafkaTopic, partition)
	if err != nil {
		panic(err)
	}

	if err == nil {
		log.Printf("Successfully connected to Kafka !!!")
	}

	return conn
}
