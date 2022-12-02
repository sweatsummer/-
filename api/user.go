package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"main/model"
	"main/service"
	"main/util"
)

func register(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if (UserName == " " || Password == " ") || (UserName == "" || Password == "") {
		util.RespParamErr(c)
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error: %v", err)
		util.ResInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "账户已存在...")
		return
	}
	err = service.CreateUser(model.User{
		UserName: UserName,
		Password: Password,
	})
	if err != nil {
		util.ResInternalErr(c)
		return
	}
	util.RespOk(c)
}

func login(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if (UserName == " " || Password == " ") || (UserName == "" || Password == "") {
		util.RespParamErr(c)
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err %v", err)
			util.ResInternalErr(c)
			return
		}
		return
	}
	if u.Password != Password {
		util.NormErr(c, 200, "密码错误")
		return
	}
	util.OkLogin(c)
	c.SetCookie("name", Password, 3600, "", "/", false, false)
}

func alter(c *gin.Context) {
	UserName := c.PostForm("username")
	NewPassword := c.PostForm("NewPassword")
	if UserName == "" || UserName == " " {
		util.RespParamErr(c)
		return
	}
	_, err := service.UpdataPasseword(NewPassword)
	if err != nil {
		log.Printf("search user err %v", err)
		util.ResInternalErr(c)
		return
	}
	util.RespOk(c)

}
