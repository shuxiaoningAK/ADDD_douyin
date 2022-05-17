package model

type User struct {
	Id            int64  `gorm:"type:int;not null" json:"user_id,omitempty"`
	Name          string `gorm:"type: varchar(100); not null; primaryKey" json:"username,omitempty"`
	Token         string `gorm:"type: varchar(200);not null" json:"token,omitempty"`
	FollowCount   int64  `gorm:"type: int; not null" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"type: int; not null" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"type: BIT; not null" json:"is_follow,omitempty"`
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
