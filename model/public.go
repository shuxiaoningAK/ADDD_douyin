package model

// 返回状态号 和 状态信息
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Model struct {
}
