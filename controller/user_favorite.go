package controller

import (
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 点赞
func FavoriteAction(ctx *gin.Context) {
	var action service.UserFavoriteAction
	if err := ctx.ShouldBindQuery(&action); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	token, claim, err := util.ParseToken(action.Token)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusOK, serializer.InvalidToken)
		ctx.Abort()
		return
	}

	if action.UserId == 0 {
		action.UserId = claim.Id
	}

	if err = action.Action(); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.Success)
}

// 喜欢列表
func FavoriteList(ctx *gin.Context) {
	var videos []*model.Video
	u64, err := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	userId := uint(u64)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}
	if videos, err = service.FavoriteList(userId); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	ctx.JSON(http.StatusOK, serializer.FeedResponse{
		Response: serializer.Success,
		// videos已经是like，不进行检查
		VideoList: serializer.PackVideos(videos, userId, false, true),
	})

}
