package controller

import "github.com/shuxiaoningAK/ADDD_DOUYIN/server/service"

type ControllerGroup struct {
	UserController
}

var ControllerGroupApp = new(ControllerGroup)

var (
	userService = service.ServiceGroupApp.UserService
)
