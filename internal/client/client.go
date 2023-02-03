package client

import (
	"fmt"
	"log"
	"net"

	"github.com/lcr2000/mchat/internal/config"
)

func Dial(username string) {
	fmt.Println("Connect to the server.")
	address := fmt.Sprintf("%s:%s", config.Cfg.Address, config.Cfg.TCPPort)
	rawConn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Dial fail, err=%v", err)
	}
	conn = NewClientConn(username, rawConn)
	go conn.read()
}
