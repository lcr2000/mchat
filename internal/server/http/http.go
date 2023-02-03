package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lcr2000/mchat/internal/config"
)

func InitHTTPServer() {
	router := gin.Default()
	router.POST("/login", loginHandle)
	router.GET("/get_online_users", getOnlineUsersHandle)

	err := router.Run(fmt.Sprintf(":%s", config.Cfg.HTTPPort))
	if err != nil {
		log.Fatalf("InitHTTPServer fail, err=%v", err)
	}
}
