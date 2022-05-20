package api

import (
	"ADDD_douyin/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var usersLoginInfo = map[string]model.User{
	//
	//"zhangleidouyin": {
	//	Name:          "zhanglei",
	//	FollowCount:   10,
	//	FollowerCount: 5,
	//	IsFollow:      true,
	//},
}

// 登录用户
func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// 查看用户信息
func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, model.UserResponse{
	//		Response: model.Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, model.UserResponse{
	//		Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

// 注册用户
func Register(c *gin.Context) {
	//
	//username := c.Query("username")
	//password := c.Query("password")

	//token := username + password

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, model.UserLoginResponse{
	//		Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//} else {
	//	atomic.AddInt64(&userIdSequence, 1)
	//	newUser := model.User{
	//		Id:   userIdSequence,
	//		Name: username,
	//	}
	//	usersLoginInfo[token] = newUser
	//	c.JSON(http.StatusOK, model.UserLoginResponse{
	//		Response: model.Response{StatusCode: 0},
	//		UserId:   userIdSequence,
	//		Token:    username + password,
	//	})
	//}
}
