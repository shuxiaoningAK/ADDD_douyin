package serializer

import "ADDD_DOUYIN/model"

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
	User
}

func PackUser(u *model.User, isFollowed bool) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:            u.ID,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      isFollowed,
	}
}
