package serializer

//序列化的UserLoginResponse
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//序列化的UserRegisterResponse
type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//是的，下列两个结构体一样，但是为了顶层的代码可读性，选择了冗余
