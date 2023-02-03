package client

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/lcr2000/mchat/internal/model"
	"github.com/lcr2000/mchat/internal/utils"
)

var conn *Conn

func GetClientConn() *Conn {
	return conn
}

type Conn struct {
	rawConn      net.Conn // 连接对象
	uid          string   // 用户ID
	lastActiveTs int64    // 最后活跃时间
}

func NewClientConn(uid string, conn net.Conn) *Conn {
	return &Conn{
		rawConn:      conn,
		uid:          uid,
		lastActiveTs: time.Now().Unix(),
	}
}

func (c *Conn) GetConn() net.Conn {
	return c.rawConn
}

func (c *Conn) read() {
	defer c.rawConn.Close()
	for {
		buf := [512]byte{}
		n, err := c.rawConn.Read(buf[:])
		if err != nil {
			fmt.Println("receive data fail, err=", err)
			return
		}
		// fmt.Println(string(buf[:n]))
		c.process(buf[:n])
	}
}

func (c *Conn) process(b []byte) {
	var p *model.ServerPacket
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("Unmarshal fail, err=", err)
		return
	}

	switch p.Cmd {
	case model.CmdReady:
	case model.CmdChat:
		var chatMsg *model.ChatMsg
		err = json.Unmarshal([]byte(p.Data.(string)), &chatMsg)
		if err != nil {
			fmt.Println("Unmarshal fail, err=", err)
			return
		}
		utils.PrintYellow(os.Stdout, fmt.Sprintf("%s %s  ", chatMsg.FromName, utils.TimeFormat(chatMsg.ServerTs)))
		fmt.Println(chatMsg.Data)
	case model.CmdChatEnter: // No need to do anything.
	}
}
