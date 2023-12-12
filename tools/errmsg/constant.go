package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ErrorUsernameUsed      = 1001
	ErrorUserPasswordWrong = 1002
	ErrorUserNotExist      = 1003
	ErrorTokeExist         = 1004
	ErrorTokeRuntime       = 1005
	ErrorTokeWrong         = 1006
	ErrorTokeTypeWrong     = 1007
	ErrorUserNoRight       = 1008

	ErrorCategoryNameUsed = 2001

	ErrorArticleNotExist = 3001
)

var CodeMsg = map[int]string{
	SUCCESS:                "SUCCESS",
	ERROR:                  "FAIL",
	ErrorUsernameUsed:      "用户名已存在",
	ErrorUserPasswordWrong: "用户密码错误",
	ErrorUserNotExist:      "用户不存在",
	ErrorTokeExist:         "TOKE不存在",
	ErrorTokeRuntime:       "TOKE已过期",
	ErrorTokeWrong:         "TOKE不正确",
	ErrorTokeTypeWrong:     "TOKE格式错误",
	ErrorCategoryNameUsed:  "分类名已存在",
	ErrorArticleNotExist:   "文章不存在！",
	ErrorUserNoRight:       "该用户无权限",
}
