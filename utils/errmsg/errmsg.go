package errmsg

const (
	ERROR_USERNAME_USED = 1001
)

var codeMsg = map[int]string{
	ERROR_USERNAME_USED: "用户已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
