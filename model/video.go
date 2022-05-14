package model

import "github.com/jinzhu/gorm"

//视频模型
type Video struct {
	gorm.Model           //视频唯一id(主键),上传时间,更新时间
	Author        User   `gorm:"ForeignKey:User;AssociationForeignKey:ID"` //视频作者信息
	Uid           uint   `gorm:"not null"`
	PlayUrl       string //视频播放地址
	CoverUrl      string //视频封面地址
	FavoriteCount int64  //`gorm:"index;not null"` //视频的点赞总数
	CommentCount  int64  //视频的评论总数
	IsFavorite    bool   //true已点赞，false未点赞
	DeletedAt     string `gorm:"default:'0'"` //逻辑删除
}
