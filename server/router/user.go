package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/controller"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userController := controller.ControllerGroupApp.UserController
	userRouter.GET("/user/", userController.UserInfo)
	userRouter.POST("/user/register/", userController.Register)
	userRouter.POST("/user/login/", userController.Login)

}
