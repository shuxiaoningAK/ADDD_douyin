package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentList(ctx *gin.Context) {
	//if token, _, err := util.ParseToken(ctx.Query("token")); err != nil || !token.Valid {
	//	ctx.JSON(http.StatusOK, serializer.InvalidToken)
	//	ctx.Abort()
	//	return
	//}

	if res, err := service.CommentList(ctx.Query("video_id")); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.CommentListResponse{
			Response:    serializer.Success,
			CommentList: serializer.PackComments(res),
		})
	}
}

func CommentAction(ctx *gin.Context) {
	var c service.CommentAction
	if err := ctx.ShouldBind(&c); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
		return
	}

	//if token, _, err := util.ParseToken(c.Token); err != nil || !token.Valid {
	//	ctx.JSON(http.StatusOK, serializer.InvalidToken)
	//	ctx.Abort()
	//	return
	//}

	if res, err := c.Action(); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.CommentResponse{
			Response: serializer.Success,
			Comment:  *serializer.PackComment(res),
		})
	}

}
