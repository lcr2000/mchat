package client

import (
	"log"
	"net"
)

func Dial(address, username string) {
	rawConn, err := net.Dial("tcp", address+":8090")
	if err != nil {
		log.Fatalf("Dial failed, err=%v", err)
	}
	conn = NewClientConn(username, rawConn)
	go conn.read()
}
