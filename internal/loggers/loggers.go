package loggers

import (
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/sirupsen/logrus"
)

const (
	APP = "app"
)

var w *loggerWrapper

type loggerWrapper struct {
	loggers map[string]*logrus.Entry
}

func Initialize(config *repo.Config) {
	m := make(map[string]*logrus.Entry)

	for mod, level := range config.Log.Module {
		m[mod] = log.NewWithModule(mod)
		m[mod].Logger.SetLevel(log.ParseLevel(level))
	}

	w = &loggerWrapper{loggers: m}
}

func Logger(name string) logrus.FieldLogger {
	logger, ok := w.loggers[name]
	if !ok {
		logger = log.NewWithModule(name)
		logger.Logger.SetLevel(logrus.InfoLevel)
		w.loggers[name] = logger
	}

	return logger
}
