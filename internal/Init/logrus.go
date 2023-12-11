package Init

import (
	"NectarPin/constant"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

var color int

type MyFormatter struct {
	ForMatTime string
}

type FileLogFormatter struct {
	ForMatTime string
}

type LogFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string
	Prefix   string
}

func (f0 MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
	formatTime := entry.Time.Format(f0.ForMatTime)

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

func (f2 *LogFileWriter) Write(data []byte) (n int, err error) {
	if f2 == nil {
		return 0, err
	}
	if f2.file == nil {
		return 0, err
	}
	//判断时间是否要切换
	fileDate := time.Now().Format("2006-01-02")
	if f2.fileDate != fileDate {
		_ = f2.file.Close()
		err := os.MkdirAll(fmt.Sprintf("%s/%s", f2.logPath, fileDate), os.ModePerm)
		if err != nil {
			return 0, err
		}
		fileName := fmt.Sprintf("%s/%s/%s.log", f2.logPath, fileDate, f2.Prefix)
		f2.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			return 0, err
		}
	}
	n, e := f2.file.Write(data)
	return n, e
}

func FileLogger() {
	logDate := time.Now().Format("2006-01-02")
	logDir := constant.Config.Logger.Director

	err := os.MkdirAll(fmt.Sprintf("%s/%s", logDir, logDate), os.ModePerm)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	logFileName := fmt.Sprintf("%s/%s/%s.log", logDir, logDate, constant.Config.Logger.Prefix)

	logFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	logFileWriter := LogFileWriter{
		file:     logFile,
		logPath:  logDir,
		fileDate: logDate,
		Prefix:   constant.Config.Logger.Prefix,
	}
	writer := []io.Writer{&logFileWriter}
	logrus.SetReportCaller(constant.Config.Logger.ShowLine)
	logrus.SetFormatter(&FileLogFormatter{ForMatTime: "2006-01-02 15:04:06"})
	logrus.SetOutput(io.MultiWriter(writer...))
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
