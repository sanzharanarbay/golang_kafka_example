package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/models"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/requests"
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/services"
	"log"
	"net/http"
)

type KafkaController struct {
	kafkaService *services.KafkaService
}

func NewKafkaController(kafkaService *services.KafkaService) *KafkaController {
	return &KafkaController{
		kafkaService: kafkaService,
	}
}

func (k *KafkaController) SendMessage(c *gin.Context) {
	var requestObj models.Student

	if err := c.ShouldBindJSON(&requestObj); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	validErrs := requests.NewStudentRequest(&requestObj).ValidateRequest()
	if len(validErrs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validErrs)
		return
	}

	pushed, err := k.kafkaService.PushMessage(&requestObj)
	if err != nil {
		log.Printf("ERROR - %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if pushed == true {
		log.Println("Message Pushed Successfully")
	}
	c.JSON(http.StatusOK, pushed)
}
