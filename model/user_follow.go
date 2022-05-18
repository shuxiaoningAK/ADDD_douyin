package model

import (
	"gorm.io/gorm"
	"time"
)

type UserFollow struct {
	User     User `gorm:"foreignKey:Uid;"`
	Follower User `gorm:"foreignKey:Fid"`
	Uid      uint `gorm:"primaryKey"`
	Fid      uint `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
