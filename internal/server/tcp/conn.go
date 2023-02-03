package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/lcr2000/mchat/internal/model"
)

type Conn struct {
	rawConn      net.Conn      // 连接对象
	cid          string        // 链接ID，程序内部自增生成的，
	ip           string        // IP
	areaCode     string        // IP 所在国家/地区码
	country      string        // IP 所在国家/地区
	province     string        // IP 所在省份
	city         string        // IP 所在城市
	closeChan    chan struct{} // close chanel
	uid          string        // 用户ID
	status       uint          // 连接状态 0-初始连接; 1-正常登录; 2-被踢出
	lastActiveTs int64         // 最后活跃时间
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		rawConn:      conn,
		cid:          uuid.New().String(),
		ip:           conn.RemoteAddr().String(),
		areaCode:     "",
		country:      "",
		province:     "",
		city:         "",
		closeChan:    nil,
		uid:          "",
		lastActiveTs: time.Now().Unix(),
	}
}

func (c *Conn) process() {
	defer func() {
		defer c.rawConn.Close()
		defer connMgr.Remove(c.uid)
	}()

	for {
		var buf [128]byte
		n, err := c.rawConn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect fail, err= %v\n", err)
			break
		}
		// fmt.Printf("receive from client, data: %v\n", string(buf[:n]))
		err = c.distribute(buf[:n])
		if err != nil {
			fmt.Printf("processHandle fail, err= %v\n", err)
			if err = c.serverError(); err != nil {
				break
			}
		}
	}
}

func (c *Conn) distribute(b []byte) error {
	var p *model.ClientPacket
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Printf("distribute Unmarshal fail, err= %v\n", err)
		return nil
	}

	switch p.Cmd {
	case model.CmdReady:
		packet := model.BuildServerPacket(p.Cmd, model.ErrCodeSuccess, "Ready.")
		marshal, _ := json.Marshal(packet)
		if _, err := c.rawConn.Write(marshal); err != nil {
			fmt.Printf("write to client fail, err= %v\n", err)
			return err
		}
	case model.CmdChatEnter:
		username := p.Data.(string)
		c.uid = username
		connMgr.Add(username, c)
		packet := model.BuildServerPacket(p.Cmd, model.ErrCodeSuccess, "Success.")
		marshal, _ := json.Marshal(packet)
		if _, err := c.rawConn.Write(marshal); err != nil {
			fmt.Printf("write to client fail, err= %v\n", err)
			return err
		}
		// Send welcome message.
		chatMsg := &model.ChatMsg{
			FromID:   model.SYStemName,
			FromName: model.SYStemName,
			MsgID:    uuid.New().String(),
			Data:     fmt.Sprintf("Welcome %s.", username),
			ServerTs: time.Now().Unix(),
		}
		bytes, _ := json.Marshal(chatMsg)
		packet = model.BuildServerPacket(model.CmdChat, model.ErrCodeSuccess, string(bytes))
		marshal, _ = json.Marshal(packet)
		for _, conn := range connMgr.GetAll() {
			if _, err := conn.rawConn.Write(marshal); err != nil {
				fmt.Printf("write to client fail, err= %v\n", err)
				continue
			}
		}
	case model.CmdChat:
		c.lastActiveTs = time.Now().Unix()
		chatMsg := &model.ChatMsg{
			FromID:   c.uid,
			FromName: c.uid,
			MsgID:    uuid.New().String(),
			Data:     p.Data,
			ServerTs: time.Now().Unix(),
		}
		bytes, _ := json.Marshal(chatMsg)
		packet := model.BuildServerPacket(p.Cmd, model.ErrCodeSuccess, string(bytes))
		marshal, _ := json.Marshal(packet)
		for _, conn := range connMgr.GetAll() {
			if c.uid == conn.uid {
				continue // No need to forward to sender.
			}
			if _, err := conn.rawConn.Write(marshal); err != nil {
				fmt.Printf("write to client fail, err= %v\n", err)
				continue
			}
		}
	}

	return nil
}

func (c *Conn) serverError() error {
	p := &model.ServerPacket{
		ErrCode: model.ErrCodeServerError,
		Data:    nil,
		Ts:      0,
	}
	marshal, _ := json.Marshal(p)
	if _, err := c.rawConn.Write(marshal); err != nil {
		fmt.Printf("process write to client fail, err= %v\n", err)
		return err
	}
	return nil
}

func (c *Conn) GetCity() string {
	return c.city
}

func (c *Conn) GetUID() string {
	return c.uid
}

func (c *Conn) GetLastActiveTs() int64 {
	return c.lastActiveTs
}
