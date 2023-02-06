package tcp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lcr2000/mchat/internal/model"
)

func SendAll(b []byte) {
	for _, conn := range connMgr.GetAll() {
		if _, err := conn.rawConn.Write(b); err != nil {
			fmt.Println("write to client fail, err=", err)
			continue
		}
	}
}

func SendAllExcept(exceptID string, b []byte) {
	for _, conn := range connMgr.GetAll() {
		if exceptID == conn.uid {
			continue // No need to forward to sender.
		}
		if _, err := conn.rawConn.Write(b); err != nil {
			fmt.Println("write to client fail, err=", err)
			continue
		}
	}
}

func buildSysChatMsgBytes(data string) []byte {
	bytes, _ := json.Marshal(&model.ChatMsg{
		FromID:   model.SYStemName,
		FromName: model.SYStemName,
		MsgID:    uuid.New().String(),
		Data:     data,
		ServerTs: time.Now().Unix(),
	})
	marshal, _ := json.Marshal(model.BuildServerPacket(model.CmdChat, model.ErrCodeSuccess, string(bytes)))
	return marshal
}
