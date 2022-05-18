package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//用户模型
type User struct {
	gorm.Model           //内含用户唯一id(主键),用户账号建立时间,更新时间
	Name          string `gorm:"unique"` //用户名
	Password      string //密码
	FollowCount   int64  //关注总数
	FollowerCount int64  //粉丝总数
}

//设置密码
func (user *User) SetPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //非明文存储密码
	user.Password = string(bytes)
}

//校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
