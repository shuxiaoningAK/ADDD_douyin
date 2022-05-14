package service

import (
	"douyin/conf"
	"douyin/model"
	"douyin/serializer"
	"douyin/util"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//UserRegisterService 用户注册服务
func (service *UserService) Register(username, password string) *serializer.UserRegisterResponse {
	var user model.User
	count := 0
	conf.DB.Model(&model.User{}).Where("name=?", username).First(&user).Count(&count)
	if count == 1 { //查找到了用户，表示该用户名已被使用
		return &serializer.UserRegisterResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "用户名已被使用",
			},
		}
	}
	user.Name = username
	//设置密码
	if err := user.SetPassword(password); err != nil {
		fmt.Println(err)
		return &serializer.UserRegisterResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "加密时出现问题",
			},
		}
	}
	//创建用户
	if err := conf.DB.Create(&user).Error; err != nil {
		fmt.Println(err)
		return &serializer.UserRegisterResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "创建用户时数据库出现问题",
			},
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		fmt.Println(err)
		return &serializer.UserRegisterResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "生成token出错",
			},
		}
	}
	return &serializer.UserRegisterResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserId: int64(user.ID),
		Token:  token,
	}
}

//Login 用户登陆服务
func (service *UserService) Login() serializer.UserLoginResponse {
	var user model.User
	if err := conf.DB.Where("name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
			return serializer.UserLoginResponse{
				Response: serializer.Response{StatusCode: 1,
					StatusMsg: "用户不存在",
				},
			}
		}
		fmt.Println(err)
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "数据库出错",
			},
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "密码不匹配",
			},
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		fmt.Println(err)
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1,
				StatusMsg: "生成token出错",
			},
		}
	}
	return serializer.UserLoginResponse{ //登录成功
		Response: serializer.Response{StatusCode: 0,
			StatusMsg: "登录成功",
		},
		UserId: int64(user.ID),
		Token:  token,
	}
}
