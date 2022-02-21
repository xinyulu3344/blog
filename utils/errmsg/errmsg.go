package errmsg

const (
    SUCCESS = 200
    ERROR = 500
    ERROR_USERNAME_USED = 1001
    ERROR_PASSWORD_WRONG = 1002
    ERROR_USER_NOT_EXIST = 1003
    ERROR_TOKEN_NOT_EXIST = 1004
    ERROR_TOKEN_RUNTIME = 1005
    ERROR_TOKEN_WRONG = 1006
    ERROR_TOKEN_TYPE_WRONG = 1007
)


var codeMsg = map[int]string{
    SUCCESS:               "OK",
    ERROR:                 "FAIL",
    ERROR_USERNAME_USED:   "用户名已存在!",
    ERROR_PASSWORD_WRONG:  "密码错误",
    ERROR_USER_NOT_EXIST:  "用户不存在",
    ERROR_TOKEN_NOT_EXIST: "TOKEN不存在",
    ERROR_TOKEN_RUNTIME:   "TOKEN已过期",
    ERROR_TOKEN_WRONG:     "TOKEN不正确",
    ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErrMsg(code int) string {
    return codeMsg[code]
}
