package routes

import "github.com/gin-gonic/gin"

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

}
