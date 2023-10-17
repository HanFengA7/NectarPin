package main

import (
	"NectarPin/constant"
	"NectarPin/internal/Init"
	"github.com/sirupsen/logrus"
)

func main() {
	Init.Conf()
	constant.Log = Init.Logger()
	//Init.Gorm()
	constant.Log.Errorln("23333")
	logrus.Errorln("23333")
}
