package model

import "gorm.io/gorm"

//评论模型
type Comment struct {
	gorm.Model
	User    User
	UserId  uint
	Uid     uint   `gorm:"not null"`
	Content string `gorm:"type:longtext"`
}
