package errmsg

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
