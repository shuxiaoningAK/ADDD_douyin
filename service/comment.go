package service

import (
	"ADDD_DOUYIN/conf"
	"ADDD_DOUYIN/model"
)

type CommentAction struct {
	UserId      uint   `form:"user_id" json:"user_id"`
	Token       string `form:"token" json:"token"`
	VideoId     uint   `form:"video_id" json:"video_id"`
	ActionType  int    `form:"action_type" json:"action_type"`
	CommentText string `form:"comment_text" json:"comment_text;omitempty"`
	CommentId   uint   `form:"comment_id" json:"comment_id;omitempty"`
}

func (c *CommentAction) Action() (*model.Comment, error) {
	switch c.ActionType {
	case 1:
		return c.create()
	case 2:
		return nil, c.delete()
	}
	return nil, nil
}

func (c *CommentAction) create() (*model.Comment, error) {
	comment := model.Comment{
		UserId:  c.UserId,
		VideoId: c.VideoId,
		Content: c.CommentText,
	}

	if err := conf.DB.Create(&comment).Error; err != nil {
		return nil, err
	}

	err := conf.DB.Preload("User").First(&comment).Error
	return &comment, err
}

func (c *CommentAction) delete() error {
	comment := model.Comment{}
	comment.ID = c.CommentId
	return conf.DB.Delete(&comment).Error
}

func CommentList(videoId string) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	err := conf.DB.Where("video_id = ?", videoId).Preload("User").Find(&comments).Error
	return comments, err
}
