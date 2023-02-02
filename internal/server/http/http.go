package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitHTTPServer() {
	router := gin.Default()
	router.POST("/login", loginHandle)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("InitHTTPServer fail, err=%v", err)
	}
}
