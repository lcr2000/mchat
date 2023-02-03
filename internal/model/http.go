package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Msg  string      `json:"msg"`
	Code ErrCode     `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func ErrorResponse(ctx *gin.Context, status ErrCode, msg string) {
	respondWithJSON(ctx, HTTPResponse{
		Msg:  msg,
		Code: status,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	respondWithJSON(ctx, HTTPResponse{
		Msg:  "success",
		Code: 0,
		Data: data,
	})
}

func respondWithJSON(ctx *gin.Context, h HTTPResponse) {
	ctx.IndentedJSON(http.StatusOK, h)
}

type LoginRequest struct {
	Username string `json:"username"`
}

type GetOnlineUsersResp struct {
	Users []*User `json:"users"`
}

type User struct {
	IP           string `json:"ip"`             // IP
	AreaCode     string `json:"area_code"`      // IP 所在国家/地区码
	Country      string `json:"country"`        // IP 所在国家/地区
	Province     string `json:"province"`       // IP 所在省份
	City         string `json:"city"`           // IP 所在城市
	UID          string `json:"uid"`            // 用户ID
	LastActiveTs int64  `json:"last_active_ts"` // 最后活跃时间
}
