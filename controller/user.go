package controller

import (
	"douyin/serializer"
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	username := c.Query("username")
	password := c.Query("password")
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(username, password)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "绑定方法出错",
			},
		})
	}
}

func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "绑定方法出错",
			},
		})
	}
}
