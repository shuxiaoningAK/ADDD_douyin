package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(ctx *gin.Context) {
	var action service.RelationAction
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

	if err := action.Action(); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.Success)
	}

}

func FollowerList(ctx *gin.Context) {
	userId, err := verifyTokenFromQueryAndRetId(ctx)
	if err != nil {
		return
	}
	u64, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	targetId := uint(u64)

	if users, err := service.Follower(targetId); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.UserList{
			Response: serializer.Success,
			UserList: serializer.PackUsers(users, userId, true, false),
		})
	}
}

func FolloweeList(ctx *gin.Context) {
	userId, err := verifyTokenFromQueryAndRetId(ctx)
	if err != nil {
		return
	}
	u64, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
	targetId := uint(u64)

	if users, err := service.Followee(targetId); err != nil {
		ctx.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		ctx.JSON(http.StatusOK, serializer.UserList{
			Response: serializer.Success,
			UserList: serializer.PackUsers(users, userId, true, false),
		})
	}
}

func verifyTokenFromQueryAndRetId(ctx *gin.Context) (uint, error) {
	token, claim, err := util.ParseToken(ctx.Query("token"))
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusOK, serializer.InvalidToken)
		ctx.Abort()
		return 0, fmt.Errorf(serializer.InvalidToken.StatusMsg)
	}
	return claim.Id, nil
}
