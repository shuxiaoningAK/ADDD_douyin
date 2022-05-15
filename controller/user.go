package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
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

func UserInfo(c *gin.Context) {
	var userInfoService service.UserInfoService
	tokenString := c.GetHeader("Authorization")[len("bearer "):] //TODO 本行代码针对客户端可能需要做出改变
	if tokenString == "" {
		c.JSON(http.StatusOK, serializer.UserInfoResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "token为空",
			},
		})
		c.Abort()
		return
	}
	token, claims, err := util.ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusOK, serializer.UserInfoResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "Token失效",
			},
		})
		c.Abort()
		return
	}
	if err := c.ShouldBind(&userInfoService); err == nil {
		res := userInfoService.UserInfo(claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.UserInfoResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "绑定方法出错",
			},
		})
	}
}
