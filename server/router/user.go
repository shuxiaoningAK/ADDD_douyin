package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/controller"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userRouter.GET("/user/", controller.UserInfo)
	userRouter.POST("/user/register/", controller.Register)
	userRouter.POST("/user/login/", controller.Login)
}
