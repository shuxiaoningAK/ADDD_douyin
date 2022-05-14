package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/service"
	"net/http"
)

func UserInfo(ctx *gin.Context) {
	token := ctx.Query("token")
	ctx.JSON(http.StatusOK, service.UserInfo(token))

}

func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	ctx.JSON(http.StatusOK, service.Register(username, password))
}

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	ctx.JSON(http.StatusOK, service.Login(username, password))
}
