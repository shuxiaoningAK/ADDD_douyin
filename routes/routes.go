package routes

import (
	"ADDD_DOUYIN/controller"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(sessions.Sessions("douyin-session", cookie.NewStore([]byte("secret-about-douyin"))))
	apiRouter := r.Group("/douyin")
	{

		apiRouter.GET("/ping", ping)
		//   ***************basic apis***************

		//apiRouter.GET("/feed/", controller.Feed)   //TODO不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
		apiRouter.POST("/user/register/", controller.UserRegister) //新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
		apiRouter.POST("/user/login/", controller.UserLogin)       //通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
		apiRouter.GET("/user/", controller.UserInfo)               //TODO获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数

		apiRouter.POST("/publish/action/", controller.Publish)  //TODO登录用户选择视频上传
		apiRouter.GET("/publish/list/", controller.PublishList) //TODO登录用户的视频发布列表，直接列出用户所有投稿过的视频

		//   ***************extra apis - I***************

		//apiRouter.POST("/favorite/action/", controller.FavoriteAction) //TODO登录用户对视频的点赞和取消点赞操作
		//apiRouter.GET("/favorite/list/", controller.FavoriteList)  //TODO登录用户的所有点赞视频
		//apiRouter.POST("/comment/action/", controller.CommentAction) //TODO登录用户对视频进行评论
		//apiRouter.GET("/comment/list/", controller.CommentList)  //TODO查看视频的所有评论，按发布时间倒序

		//    ***************extra apis - II***************

		//apiRouter.POST("/relation/action/", controller.RelationAction) //TODO登录用户对一个用户进行关注
		//apiRouter.GET("/relation/follow/list/", controller.FollowList)  //TODO拉取用户的关注列表
		//apiRouter.GET("/relation/follower/list/", controller.FollowerList) //TODO拉取用户的粉丝列表
	}
	return r
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "PONG")
}
