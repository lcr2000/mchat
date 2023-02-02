package tcp

import (
	"fmt"
	"log"
	"net"
)

func InitTCPServer() {
	defer connMgr.Close()
	// 使用 net.Listen 监听连接的地址与端口
	var err error
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("listen fail, err= %v\n", err)
	}

	for {
		var rawConn net.Conn
		rawConn, err = listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		conn := NewConn(rawConn)
		go conn.process()
	}
}
