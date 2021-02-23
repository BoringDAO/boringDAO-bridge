package log

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func newRotateHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	baseLogName := path.Join(logPath, logFileName)

	writer, err := rotatelogs.New(
		baseLogName+"%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(baseLogName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %s", err)
		return nil
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
}
