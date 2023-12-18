package main

import (
	"NectarPin/constant"
	"NectarPin/internal/Init"
)

func main() {
	Init.Conf()
	constant.Log = Init.Logger()
	constant.DB = Init.Gorm()
	constant.Router = Init.Router()

	//fmt.Println(tools.RandSalt(16))

}
