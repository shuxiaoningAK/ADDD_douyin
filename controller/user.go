package controller

import (
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
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
	username := c.Query("username")
	password := c.Query("password")
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(username, password)
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
	tokenString := c.Query("token")
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
		var user serializer.User
		user.Id = res.Id
		user.Name = res.Name
		user.FollowCount = res.FollowCount
		user.FollowerCount = res.FollowerCount
		user.IsFollow = res.IsFollow
		c.JSON(http.StatusOK, gin.H{
			"status_code": res.StatusCode,
			"status_msg":  res.StatusMsg,
			"user":        user,
		})
	} else {
		c.JSON(http.StatusOK, serializer.UserInfoResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "绑定方法出错",
			},
		})
	}
}

func Publish(ctx *gin.Context) {
	video := &model.Video{}
	token, claims, err := util.ParseToken(ctx.PostForm("token"))
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusOK, serializer.InvalidToken)
		ctx.Abort()
		return
	}

	video.Title = ctx.PostForm("title")
	video.AuthorId = claims.Id
	video.CoverUrl = "https://iph.href.lu/400x400?text=%E6%97%A0&fg=666666&bg=cccccc" // fixme

	var data *multipart.FileHeader
	if data, err = ctx.FormFile("data"); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	if err = util.UploadVideo(util.NextUuid(), data, video); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	if err = service.Publish(video); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.Success)

}

func PublishList(ctx *gin.Context) {
	token, claim, err := util.ParseToken(ctx.Query("token"))
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusOK, serializer.InvalidToken)
		ctx.Abort()
		return
	}

	res, err := service.PublishList(claim.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.FeedResponse{
		Response:  serializer.Success,
		VideoList: serializer.PackVideos(res, claim.Id, true, false),
	})

}
