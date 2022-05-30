package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentList(ctx *gin.Context) {
	token, claim, err := util.ParseToken(ctx.Query("token"))
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusOK, serializer.InvalidToken)
		ctx.Abort()
		return
	}

	if res, err := service.CommentList(ctx.Query("video_id")); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.CommentListResponse{
			Response:    serializer.Success,
			CommentList: serializer.PackComments(res, claim.Id),
		})
	}
}

func CommentAction(ctx *gin.Context) {
	var action service.CommentAction
	if err := ctx.ShouldBind(&action); err != nil {
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

	if res, err := action.Action(); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.CommentResponse{
			Response: serializer.Success,
			Comment:  *serializer.PackComment(res, claim.Id),
		})
	}

}
