package model

import "time"

type ClientPacket struct {
	Cmd  Cmd         `json:"cmd"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

type Ready struct{}

type ServerPacket struct {
	Cmd     Cmd         `json:"cmd"`
	ErrCode ErrCode     `json:"err_code"`
	Data    interface{} `json:"data"`
	Ts      int64       `json:"ts"`
}

func BuildClientPacket(cmd Cmd, data interface{}) *ClientPacket {
	return &ClientPacket{
		Cmd:  cmd,
		Data: data,
		Ts:   time.Now().Unix(),
	}
}

func BuildServerPacket(cmd Cmd, errCode ErrCode, data interface{}) *ServerPacket {
	return &ServerPacket{
		Cmd:     cmd,
		ErrCode: errCode,
		Data:    data,
		Ts:      time.Now().Unix(),
	}
}

type ChatMsg struct {
	FromID   string      `json:"from_id"`   // 发送人 ID
	FromName string      `json:"from_name"` // 发送方的昵称
	MsgID    string      `json:"msg_id"`    // 消息唯一id
	Type     string      `json:"type"`      // 类型
	Data     interface{} `json:"data"`      // Data
	ServerTs int64       `json:"server_ts"` // 时间戳
}
