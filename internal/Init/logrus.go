package Init

import (
	"NectarPin/constant"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

var color int

type FileLogFormatter struct {
	ForMatTime string
}

type MyFormatter struct {
	ForMatTime string
}

func (f1 FileLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//设置Buffer缓冲区
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	//设置时间格式
	formatTime := entry.Time.Format(f1.ForMatTime)

	//设置行号
	filVal := fmt.Sprintf("path=\"%s:%d\"", path.Base(entry.Caller.File), entry.Caller.Line)

	//设置格式
	if constant.Config.Logger.ShowLine == true {
		_, _ = fmt.Fprintf(b,
			"time=\"%s\" %s level=\"%s\" msg=\"%s\" \n",
			formatTime,
			filVal,
			strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
	} else {
		_, _ = fmt.Fprintf(b,
			" [%s] level=\"%s\" msg=\"%s\" \n",
			formatTime,
			strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
	}

	return b.Bytes(), nil
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//设置Level颜色
	switch entry.Level {
	case logrus.FatalLevel:
		color = constant.ColorRedA
	case logrus.ErrorLevel:
		color = constant.ColorRedA
	case logrus.WarnLevel:
		color = constant.ColorYoungA
	case logrus.InfoLevel:
		color = constant.ColorGreenA
	case logrus.DebugLevel:
		color = constant.ColorYoungA
	default:
		color = constant.ColorGreenA
	}

	//设置Buffer缓冲区
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	//设置时间格式
	formatTime := entry.Time.Format(f.ForMatTime)

	//设置行号
	filVal := fmt.Sprintf("\u001B[46m%s\u001B[0m\u001B[47m:%d\u001B[0m", path.Base(entry.Caller.File), entry.Caller.Line)

	//设置格式
	if constant.Config.Logger.ShowLine == true {
		_, _ = fmt.Fprintf(b,
			" [%s] [%s] \u001B[44m %s \u001B[0m \u001B[4%dm\u001B[37m %s \u001B[0m\u001B[0m %s \n",
			formatTime,
			filVal,
			constant.Config.Logger.Prefix,
			color,
			strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
	} else {
		_, _ = fmt.Fprintf(b,
			" [%s] \u001B[44m %s \u001B[0m \u001B[4%dm\u001B[37m %s \u001B[0m\u001B[0m %s \n",
			formatTime,
			constant.Config.Logger.Prefix,
			color, strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
	}

	return b.Bytes(), nil
}

func FileLogger() {
	//日志写成log文件
	logDir := constant.Config.Logger.Director
	logFile, _ := os.OpenFile(logDir+"/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetReportCaller(constant.Config.Logger.ShowLine)
	logrus.SetFormatter(&FileLogFormatter{ForMatTime: "2006-01-02 15:04:06"})
	logrus.SetOutput(logFile)
}

func Logger() *logrus.Logger {
	FileLogger()
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetLevel(logrus.DebugLevel)
	mLog.SetReportCaller(constant.Config.Logger.ShowLine)
	mLog.SetFormatter(&MyFormatter{ForMatTime: "2006-01-02 15:04:06"})
	return mLog
}
