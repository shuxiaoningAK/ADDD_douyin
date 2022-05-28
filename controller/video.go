package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
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

	if res, err := service.Feed(latestTime); err != nil {
		c.JSON(http.StatusOK, serializer.ConvertErr(err))
	} else {
		c.JSON(http.StatusOK, serializer.FeedResponse{
			Response:  serializer.Success,
			VideoList: serializer.PackVideos(res),
			NextTime:  time.Now().Unix(), // fixme 本次最早时间作为下次请求时间？
		})
	}
}
