package main

import (
	"NectarPin/constant"
	"NectarPin/internal/Init"
	"fmt"
)

func main() {
	Init.Conf()
	constant.DB = Init.Gorm()
	fmt.Println(constant.DB)
}
