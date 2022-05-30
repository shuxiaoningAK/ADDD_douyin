package service

import (
	"ADDD_DOUYIN/conf"
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/util"
)

type UserFavoriteAction struct {
	UserId     uint   `form:"user_id"`
	Token      string `form:"token"`
	VideoId    uint   `form:"video_id"`
	ActionType int    `form:"action_type"`
}

func (f *UserFavoriteAction) Action() error {
	switch f.ActionType {
	case 1:
		return f.create()
	case 2:
		return f.delete()
	}
	return nil
}

// 客户端有个bug，未登陆状态下点击喜欢并登陆一个账号（已经喜欢过该视频），会导致显示的喜欢数比实际多1
func (f *UserFavoriteAction) create() error {
	//uf := model.UserFavorite{
	//	UserId:     f.UserId,
	//	FavoriteId: f.VideoId,
	//}
	//
	//err := conf.DB.First(&uf).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return conf.DB.Create(&uf).Error
	//} else {
	//	return nil
	//}
	return util.Redis.Like(f.UserId, f.VideoId)
}

func (f *UserFavoriteAction) delete() error {
	//return conf.DB.Where("user_id = ? and favorite_id = ?", f.UserId, f.VideoId).Delete(&model.UserFavorite{}).Error
	return util.Redis.Unlike(f.UserId, f.VideoId)
}

func FavoriteList(userId uint) ([]*model.Video, error) {
	//fs := make([]*model.UserFavorite, 0)
	//err := conf.DB.Where("user_id = ?", userId).Preload("Favorite").Find(&fs).Error
	//videos := make([]*model.Video, len(fs))
	//for i, v := range fs {
	//	videos[i] = &v.Favorite
	//}
	//return videos, err
	videoIdList, err := util.Redis.RangeLike(userId, 0, -1)
	if err != nil {
		return nil, err
	}
	videos := make([]*model.Video, len(videoIdList))
	if len(videoIdList) == 0 {
		return videos, err
	}
	if err = conf.DB.Model(&model.Video{}).Find(&videos, videoIdList).Error; err != nil {
		return nil, err
	}
	return videos, err
}
