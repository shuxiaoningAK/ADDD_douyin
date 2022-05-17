package service

import (
	"ADDD_douyin/model"
	"sync/atomic"
)

type UserService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (service *UserService) RegisterService() model.UserRegisterResponse {
	if len(service.Username) > 32 {
		return model.UserRegisterResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "用户名长度不得大于32个字符",
			},
		}
	}

	if len(service.Password) > 32 {
		return model.UserRegisterResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "密码长度不得大于32个字符",
			},
		}
	}

	token := service.Username + service.Password // 生成用户token

	// 查询数据，发现用户已存在
	if _, exist := model.UsersLoginInfo[token]; exist {
		return model.UserRegisterResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "用户已存在",
			},
		}

	} else { // 成功注册新的用户
		atomic.AddInt64(&model.UserIdSequence, 1) // 生成用户id
		newUser := model.User{
			Id:   model.UserIdSequence,
			Name: service.Username,
		}
		model.UsersLoginInfo[token] = newUser // 创建新用户

		return model.UserRegisterResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "用户注册成功",
			},
			UserId: model.UserIdSequence,
			Token:  token,
		}
	}
}

func (service UserService) LoginService() model.UserLoginResponse {
	token := service.Username + service.Password // 生成用户token

	if user, exist := model.UsersLoginInfo[token]; exist {
		return model.UserLoginResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "登录成功",
			},
			UserId: user.Id,
			Token:  token,
		}
	} else {
		return model.UserLoginResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "用户不存在",
			},
		}
	}
}
