package controller

import (
	"ADDD_douyin/model"
	"ADDD_douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
var userIdSequence = int64(1)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:   userIdSequence,
			Name: username,
		}
		UsersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}
*/

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := model.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("here")
	userRegister := service.UserService{Username: username, Password: password}
	c.JSON(http.StatusOK, userRegister.RegisterService())
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userLogin := service.UserService{Username: username, Password: password}
	c.JSON(http.StatusOK, userLogin.LoginService())
}
