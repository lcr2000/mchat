package tcp

import (
	"fmt"
	"log"
	"net"

	"github.com/lcr2000/mchat/internal/config"
)

func InitTCPServer() {
	defer connMgr.Close()
	// 使用 net.Listen 监听连接的地址与端口
	var err error
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Cfg.TCPPort))
	if err != nil {
		log.Fatalf("listen fail, err= %v\n", err)
	}

	for {
		var rawConn net.Conn
		rawConn, err = listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err= %v\n", err)
			continue
		}
		conn := NewConn(rawConn)
		go conn.process()
	}
}
