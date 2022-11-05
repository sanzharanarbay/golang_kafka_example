package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	v1 "github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/routes/api/v1"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	mode := os.Getenv("GIN_MODE")

	prefix := os.Getenv("ROUTE_PREFIX")
	fmt.Println("Server started at " + port + "...")
	gin.SetMode(mode)
	router := gin.New()
	v1.ApiRoutes(prefix, router)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
