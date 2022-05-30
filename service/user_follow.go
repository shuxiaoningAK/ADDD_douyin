package service

import (
	"ADDD_DOUYIN/conf"
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/util"
)

type RelationAction struct {
	UserId     uint   `form:"user_id" json:"user_id"`
	Token      string `form:"token" json:"token"`
	ToUserId   uint   `form:"to_user_id" json:"to_user_id"`
	ActionType int    `form:"action_type" json:"action_type"`
}

func (r *RelationAction) Action() error {
	switch r.ActionType {
	case 1:
		return r.follow()
	case 2:
		return r.unfollow()
	}
	return nil
}

func (r *RelationAction) follow() error {
	return util.Redis.Follow(r.UserId, r.ToUserId)
}

func (r *RelationAction) unfollow() error {
	return util.Redis.UnFollow(r.UserId, r.ToUserId)
}

func Follower(userId uint) ([]*model.User, error) {
	ids, err := util.Redis.RangeFollower(userId, 0, -1)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(ids))
	if len(ids) == 0 {
		return users, nil
	}
	err = conf.DB.Find(&users, ids).Error
	return users, err
}

func Followee(userId uint) ([]*model.User, error) {
	ids, err := util.Redis.RangeFollowee(userId, 0, -1)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(ids))
	if len(ids) == 0 {
		return users, nil
	}
	err = conf.DB.Find(&users, ids).Error
	return users, err
}
