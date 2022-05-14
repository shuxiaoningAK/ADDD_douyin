package model

import "github.com/jinzhu/gorm"

//评论模型
type Comment struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:User;AssociationForeignKey:ID"`
	Uid       uint   `gorm:"not null"`
	Content   string `gorm:"type:longtext"`
	DeletedAt int    `gorm:"default:'0'"` //逻辑删除
}
