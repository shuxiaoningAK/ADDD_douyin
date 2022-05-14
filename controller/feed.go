package controller

import (
	"ADDD_douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, model.FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
