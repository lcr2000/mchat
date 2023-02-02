package client

import (
	"log"
	"net"
)

func Dial(username string) {
	rawConn, err := net.Dial("tcp", ":8090")
	if err != nil {
		log.Fatalf("Dial failed, err=%v", err)
	}
	conn = NewClientConn(username, rawConn)
	go conn.read()
}
