package tools

import (
	"NectarPin/constant"
	"github.com/sirupsen/logrus"
)

func Log(level any, msg any) {
	switch level {
	case "panic":
		constant.Log.Panicln(msg)
		logrus.Panicln(msg)
	case "fatal":
		constant.Log.Fatalln(msg)
		logrus.Fatalln(msg)
	case "error":
		constant.Log.Errorln(msg)
		logrus.Errorln(msg)
	case "warn", "warning":
		constant.Log.Warningln(msg)
		logrus.Warningln(msg)
	case "info":
		constant.Log.Infoln(msg)
		logrus.Infoln(msg)
	case "debug":
		constant.Log.Debugln(msg)
		logrus.Debugln(msg)
	case "trace":
		constant.Log.Traceln(msg)
		logrus.Traceln(msg)
	default:
		constant.Log.Infoln(msg)
		logrus.Infoln(msg)
	}
}
