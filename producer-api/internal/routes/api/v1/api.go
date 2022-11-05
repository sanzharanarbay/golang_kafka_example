package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/config/kafka"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/controllers"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	kafkaObj := kafka.ConnectKafka()
	apiGroup := router.Group(prefix)
	{
		kafkaGroup := apiGroup.Group("/kafka")
		{
			kafkaService := services.NewKafkaService(kafkaObj)
			kafkaController := controllers.NewKafkaController(kafkaService)

			kafkaGroup.POST("/push", kafkaController.SendMessage)
		}
	}
}
