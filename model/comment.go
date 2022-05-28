package model

import "gorm.io/gorm"

//评论模型
type Comment struct {
	gorm.Model
	User    User
	UserId  uint
	VideoId uint
	Content string `gorm:"type:longtext"`
}
