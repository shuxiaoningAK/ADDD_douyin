package model

import (
	"ADDD_douyin/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	Id            int64  `gorm:"primaryKey;type int;not null" json:"user_id,omitempty"`
	Username      string `gorm:"type: varchar(100); not null" json:"username,omitempty"`
	Password      string `gorm:"type: varchar(100);not null" json:"password"`
	FollowCount   int64  `gorm:"type: int; not null" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"type: int; not null" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"type: tinyint(1); not null" json:"is_follow,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// 查询用户是否存在
func CheckUser(token string) (state int) {
	var users User
	db.Where("token = ?", token).First(&users)
	if users.Name == "" {
		return errmsg.SUCCESS
	} else {
		return errmsg.ERROR_USERNAME_USED
	}
}

// 密码加密

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
