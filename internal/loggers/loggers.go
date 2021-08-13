package loggers

import (
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/sirupsen/logrus"
)

const (
	APP   = "app"
	BSC   = "bsc"
	MATIC = "matic"
)

var w *loggerWrapper

type loggerWrapper struct {
	loggers map[string]*logrus.Entry
}

func Initialize(config *repo.Config) {
	m := make(map[string]*logrus.Entry)
	m[APP] = log.NewWithModule(APP)
	m[APP].Logger.SetLevel(log.ParseLevel(config.Log.Module.APP))
	m[BSC] = log.NewWithModule(BSC)
	m[BSC].Logger.SetLevel(log.ParseLevel(config.Log.Module.BSC))
	m[MATIC] = log.NewWithModule(MATIC)
	m[MATIC].Logger.SetLevel(log.ParseLevel(config.Log.Module.MATIC))

	w = &loggerWrapper{loggers: m}
}

func Logger(name string) logrus.FieldLogger {
	return w.loggers[name]
}
