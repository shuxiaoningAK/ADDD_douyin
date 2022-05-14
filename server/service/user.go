package service

import (
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/model"
)

func UserInfo(token string) model.UserInfoResponse {
	//repository.DB.Get(token)
	return model.UserInfoResponse{}
}

func Register(username, password string) model.UserLoginResponse {
	return model.UserLoginResponse{}
}

func Login(username, password string) model.UserLoginResponse {
	return model.UserLoginResponse{}
}
