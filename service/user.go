package service

type UserService struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserInfoService struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

var userIdSequence = int64(0)

// 用户注册
