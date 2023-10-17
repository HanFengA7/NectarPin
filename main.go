package main

import (
	"NectarPin/internal/Init"
	"github.com/sirupsen/logrus"
)

func main() {
	Init.Conf()
	Init.Logger()
	//Init.Gorm()

	logrus.Errorln("23333")
	logrus.Warningln("6666")
	logrus.Infoln("55555")
}
