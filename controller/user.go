package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use to store user info, and key is username+possword for demo
// user data will be cleared every time the server starts
// test data: username = zhanglei, possword  = douyin

var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, user)
	}
}
