package controller

import (
	"ADDD_DOUYIN/serializer"
	"ADDD_DOUYIN/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Feeds(c *gin.Context) {
	var feedService service.FeedService
	if err := c.ShouldBind(&feedService); err != nil {
		c.JSON(http.StatusOK, serializer.FeedResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "binding方法出错",
			},
		})
	} else {
		res := feedService.FeedService()
		c.JSON(http.StatusOK, res)
	}
}
