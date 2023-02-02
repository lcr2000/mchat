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