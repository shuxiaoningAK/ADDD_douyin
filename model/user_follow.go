package model

import "gorm.io/gorm"

type UserFollow struct {
	User       User
	Follower   User
	UserId     uint `gorm:"index:user_follow"`
	FollowerId uint `gorm:"index:user_follow"`

	gorm.Model
}
