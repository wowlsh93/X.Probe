/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package logging

import (
	"bitbucket.org/Monitoring/gaemi/core/collector/config"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var logger = logrus.New()

func GetLogger() *logrus.Logger {
	return logger
}

func loatefilehook(configpath string, level logrus.Level) logrus.Hook {

	Formatter := new(logrus.JSONFormatter)
	Formatter.TimestampFormat = "Jan _2 15:04:05.0000"

	rotateFileHook := NewRotateFileHook(RotateFileConfig{
		Filename:   configpath + "gaemi.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     7,
		Level:      level,
		Formatter:  logrus.Formatter(Formatter),
	})

	return rotateFileHook
}

func getLoggingLevel(conf *config.Configuration) logrus.Level {
	switch level := conf.LOG.Loglevel; level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
func InitLog(conf *config.Configuration) {

	logger.AddHook(filename.NewHook())
	logger.SetFormatter(&logrus.TextFormatter{})

	var logLevel = getLoggingLevel(conf)
	logger.SetLevel(logLevel)

	if conf.LOG.Logoutput == "both" {
		rotateFileHook := loatefilehook(conf.LOG.Logpath, logLevel)
		logger.AddHook(rotateFileHook)
		logger.SetOutput(os.Stdout)

	} else if conf.LOG.Logoutput == "file" {
		rotateFileHook := loatefilehook(conf.LOG.Logpath, logLevel)
		logger.AddHook(rotateFileHook)
		logger.SetOutput(ioutil.Discard)

	} else {
		logger.SetOutput(os.Stdout)
	}
}
