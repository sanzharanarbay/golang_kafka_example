package services

import (
	"encoding/json"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/models"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaService struct {
	kafkaConn *kafka.Conn
}

func NewKafkaService(kafkaConn *kafka.Conn) *KafkaService {
	return &KafkaService{
		kafkaConn: kafkaConn,
	}
}

type KafkaServiceInterface interface {
	PushMessage(student *models.Student) (bool, error)
}

func (k *KafkaService) PushMessage(student *models.Student) (bool, error) {
	err := k.kafkaConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		panic(err)
	}
	body, _ := json.Marshal(student)
	_, err = k.kafkaConn.WriteMessages(
		kafka.Message{Value: body},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return false, err
	}

	//if err := k.kafkaConn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//	return false, err
	//}

	return true, nil
}
