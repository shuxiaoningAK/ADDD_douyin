package service

import (
	"douyin/conf"
	"douyin/serializer"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type FeedService struct { //TODO 目前用户登录与否都可以返回Feed流，或许以后可以根据用户的特征push一些定制化的Feed流
	LatestTime string `json:"latest_time,omitempty"`
	Token      string `json:"token,omitempty"`
}

//Feed流服务
func (service *FeedService) FeedService() serializer.FeedResponse {
	videoList := make([]serializer.Video, 30) //定义返回Response中的videoList，且预分配内存
	if err := conf.DB.Select("*").Order("created_at DESC").Limit(30).Scan(&videoList).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
			return serializer.FeedResponse{
				Response: serializer.Response{StatusCode: 1,
					StatusMsg: "内容不足30条,请稍后再来",
				},
			}
		}
		//数据库出错
		fmt.Println(err)
		return serializer.FeedResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "数据库出错",
			},
		}
	}
	//正常返回feed流的业务逻辑
	return serializer.FeedResponse{
		Response: serializer.Response{StatusCode: 0,
			StatusMsg: "feed流返回成功",
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}
}
