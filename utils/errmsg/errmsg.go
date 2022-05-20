package errmsg

const (
	SUCCESS             = 0
	ERROR               = 1
	ERROR_USERNAME_USED = 1001
	ERROR_TOKEN_EXIST   = 1002
)

var codeMsg = map[int]string{
	ERROR_USERNAME_USED: "用户已存在",
	SUCCESS:             "OK",
	ERROR:               "错误",
	ERROR_TOKEN_EXIST:   "token错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
