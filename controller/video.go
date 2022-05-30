package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"ADDD_DOUYIN/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {

	latestTime := c.Query("latest_time")
	if latestTime == "" {
		latestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}

	// 如果登陆了 对于feed流检查like情况 不登陆时默认为false
	userId, check := uint(0), false
	token, claim, err := util.ParseToken(c.Query("token"))
	if err == nil && token.Valid {
		userId = claim.Id
		check = true
	}

	if res, err := service.Feed(latestTime); err != nil {
		c.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		c.JSON(http.StatusOK, serializer.FeedResponse{
			Response:  serializer.Success,
			VideoList: serializer.PackVideos(res, userId, check, false),
			NextTime:  time.Now().Unix(), // fixme 本次最早时间作为下次请求时间？
		})
	}
}
