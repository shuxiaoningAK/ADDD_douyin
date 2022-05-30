package serializer

import (
	"ADDD_DOUYIN/model"
	"ADDD_DOUYIN/util"
)

//序列化的UserLoginResponse
type UserLoginResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//序列化的UserRegisterResponse
type UserRegisterResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//是的，上面两个结构体一样，但是为了顶层的代码可读性，选择了进行冗余

//序列化用户信息
type UserInfoResponse struct {
	Response
	*User
}

type UserList struct {
	Response
	UserList []*User `json:"user_list"`
}

func PackUser(u *model.User, userId uint, check, defaultTo bool) *User {
	if u == nil {
		return nil
	}

	followerCount, followeeCount, isFollow := int64(0), int64(0), defaultTo
	followerCount, _ = util.Redis.CountFollower(u.ID)
	followeeCount, _ = util.Redis.CountFollowee(u.ID)
	if check {
		isFollow, _ = util.Redis.IsFollow(userId, u.ID)
	}

	return &User{
		Id:            u.ID,
		Name:          u.Name,
		FollowCount:   followeeCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}
}

func PackUsers(vs []*model.User, userId uint, check, defaultTo bool) []*User {
	users := make([]*User, 0)
	for _, v := range vs {
		if v2 := PackUser(v, userId, check, defaultTo); v2 != nil {
			users = append(users, v2)
		}
	}
	return users
}
