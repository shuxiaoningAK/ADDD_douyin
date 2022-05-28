package model

import (
	"gorm.io/gorm"
)

type UserFavorite struct {
	User       User
	Favorite   Video
	UserId     uint `gorm:"index:user_favorite"`
	FavoriteId uint `gorm:"index:user_favorite"`

	gorm.Model
}
