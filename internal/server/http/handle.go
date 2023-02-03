package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lcr2000/mchat/internal/model"
	"github.com/lcr2000/mchat/internal/server/tcp"
)

func loginHandle(c *gin.Context) {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		model.ErrorResponse(c, model.ErrCodeServerError, "params bind fail")
		return
	}
	if req.Username == "" {
		model.ErrorResponse(c, model.ErrCodeBadParams, "username is required")
		return
	}
	_, isExist := tcp.GetConnMgr().Get(req.Username)
	if isExist {
		model.ErrorResponse(c, model.ErrCodeAccountExist, "account exist")
		return
	}
	model.SuccessResponse(c, "success")
}

func getOnlineUsersHandle(c *gin.Context) {
	conns := tcp.GetConnMgr().GetAll()
	users := make([]*model.User, 0, len(conns))
	for _, conn := range conns {
		users = append(users, &model.User{
			City:         conn.GetCity(),
			UID:          conn.GetUID(),
			LastActiveTs: conn.GetLastActiveTs(),
		})
	}
	model.SuccessResponse(c, users)
}
