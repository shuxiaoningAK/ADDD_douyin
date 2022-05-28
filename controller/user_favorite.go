package controller

import (
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(ctx *gin.Context) {
	var action service.UserFavoriteAction
	if err := ctx.ShouldBindQuery(&action); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	if err := action.Action(); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.Success)
}

func FavoriteList(ctx *gin.Context) {
	var videos []*model.Video
	var err error
	if videos, err = service.FavoriteList(ctx.Query("user_id")); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.FeedResponse{
		Response:  serializer.Success,
		VideoList: serializer.PackVideos(videos),
	})

}
