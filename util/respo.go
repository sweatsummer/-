package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	status int    `json:"status"`
	Info   string `json:"info"`
}

// 注册//
var Ok = respTemplate{
	status: 200,
	Info:   "registered successfully",
}

var Oklogin = respTemplate{
	status: 200,
	Info:   "login successfully",
}

var paramError = respTemplate{
	status: 300,
	Info:   "param error",
}

var Internalerr = respTemplate{
	status: 500,
	Info:   "internal err",
}

func RespOk(c *gin.Context) {
	c.JSON(http.StatusOK, Ok)
}
func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, paramError)
}

func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Internalerr)
}
func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
func OkLogin(c *gin.Context) {
	c.JSON(http.StatusOK, Oklogin)
}
