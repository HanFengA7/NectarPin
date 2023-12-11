package main

import (
	"NectarPin/constant"
	"NectarPin/internal/init"
)

func main() {
	init.Conf()
	constant.Log = init.Logger()
	constant.DB = init.Gorm()
	constant.Router = init.Router()
}
