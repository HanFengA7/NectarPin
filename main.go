package main

import (
	"NectarPin/constant"
	"NectarPin/internal/Init"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	Init.Conf()
	constant.Log = Init.Logger()
	constant.DB = Init.Gorm()

	for {
		constant.Log.Info("数据库连接成功")
		logrus.Infoln("数据库连接成功")
		time.Sleep(20 * time.Second)
	}

}
