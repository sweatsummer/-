package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"main/service"
	"main/util"
	"net/http"
)

func Recivemessage(c *gin.Context) {
	username := c.PostForm("username")
	if username == " " || username == "" {
		util.RespParamErr(c)
		return
	}
	u, err := service.FindID(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error: %v", err)
		util.ResInternalErr(c)
		return
	}
	if u.OwnerUID == 0 {
		util.NormErr(c, 300, "用户不存在")
		return
	}
	_, err = service.Lookmessage(u.OwnerUID)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	if u.Detail == "" {
		util.NormErr(c, 300, "留言为空")
	}
	c.JSON(http.StatusOK, u.Detail)
	return

}
