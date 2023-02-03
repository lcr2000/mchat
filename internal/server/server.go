package server

import (
	"fmt"

	"github.com/lcr2000/mchat/internal/server/http"
	"github.com/lcr2000/mchat/internal/server/tcp"
)

func StartServer() {
	fmt.Println("Start the server.")
	go http.InitHTTPServer()
	go tcp.InitTCPServer()
}
