package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (c *UserController) UserInfo(ctx *gin.Context) {
	userService.UserInfo()
}

func (c *UserController) Register(ctx *gin.Context) {
	userService.Register()
}

func (c *UserController) Login(ctx *gin.Context) {
	userService.Login()
}
