package model

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}
