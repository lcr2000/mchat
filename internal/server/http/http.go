package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitHTTPServer() {
	router := gin.Default()
	router.POST("/login", loginHandle)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("InitHTTPServer fail, err=%v", err)
	}
}
