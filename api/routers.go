package api

import (
	"github.com/gin-gonic/gin"
	"main/middleware"
)

func Initrouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", register)
		u.POST("/login", login)
		u.PUT("/password", middleware.AuthMiddleWare(), alter)
	}
	m := r.Group("/message")
	{

		m.GET("/message", Recivemessage)
		m.PUT("/message")
		m.DELETE("message")
	}

	r.Run(":8080")
}
