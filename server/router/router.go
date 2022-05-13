package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	r := gin.Default()
	group := r.Group("")
	RouterGroupApp.InitUserRouter(group)
	return r
}
