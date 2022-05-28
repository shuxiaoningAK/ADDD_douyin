package service

import (
	"ADDD_DOUYIN/conf"
	"ADDD_DOUYIN/model"
	"errors"
	"gorm.io/gorm"
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

func (f *UserFavoriteAction) create() error {
	uf := model.UserFavorite{
		UserId:     f.UserId,
		FavoriteId: f.VideoId,
	}

	err := conf.DB.First(&uf).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return conf.DB.Create(&uf).Error
	} else {
		return nil
	}

}

func (f *UserFavoriteAction) delete() error {
	return conf.DB.Where("user_id = ? and favorite_id = ?", f.UserId, f.VideoId).Delete(&model.UserFavorite{}).Error
}

func FavoriteList(userId string) ([]*model.Video, error) {
	fs := make([]*model.UserFavorite, 0)
	err := conf.DB.Where("user_id = ?", userId).Preload("Favorite").Find(&fs).Error
	videos := make([]*model.Video, len(fs))
	for i, v := range fs {
		videos[i] = &v.Favorite
	}
	return videos, err
}
