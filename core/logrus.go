package core

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"service/global"
)

func Logrus() *logrus.Logger {
	return &logrus.Logger{
		Out:          setLogrusOut(global.GVA_SETTING.Logrus.OutputType),
		Formatter:    setLogrusFormatter(global.GVA_SETTING.Logrus.Formatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        setLogrusLevel(global.GVA_SETTING.Logrus.Level),
		ReportCaller: global.GVA_SETTING.Logrus.ReportCaller,
	}
}

// 默认为InfoLevel
func setLogrusLevel(level string) logrus.Level {
	if level == "debug" {
		return logrus.DebugLevel
	} else if level == "warn" {
		return logrus.WarnLevel
	} else if level == "error" {
		return logrus.ErrorLevel
	} else if level == "fatal" {
		return logrus.FatalLevel
	} else {
		return logrus.InfoLevel
	}
}

func setLogrusFormatter(formatter string) logrus.Formatter {
	if formatter == "json" {
		return &logrus.JSONFormatter{}
	} else {
		return &logrus.TextFormatter{}
	}
}

func setLogrusOut(outputType int) io.Writer {
	if outputType == 0 {
		return os.Stderr
	} else if outputType == 1 {
		//TODO: 日志切片存放
		file, err := os.OpenFile(global.GVA_SETTING.System.LogPath+"logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return os.Stderr
		}
		// BUG: 关闭文件的时机？
		return file
	} else {
		file, err := os.OpenFile(global.GVA_SETTING.System.LogPath+"logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return os.Stderr
		}
		multiWriter := io.MultiWriter(os.Stderr, file)
		return multiWriter
	}
}
