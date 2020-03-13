package wechat

import (
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
)

type user int

var User user

func (user) login(c *gin.Context) {
	form := &model.User{}
	if err := c.Bind(form); err != nil {
		c.String(400, "参数错误")
		c.Abort()
	}
	if err := service.User.Login(form); err != nil {

		c.String(500, "内部服务器错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) Register(r *gin.RouterGroup) {
	r.POST("/v1/user/login", User.login)
}
