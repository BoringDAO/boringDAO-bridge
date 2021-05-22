package loggers

import (
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/sirupsen/logrus"
)

const (
	APP  = "app"
	BSC  = "bsc"
	OKEX = "okex"
	ETH  = "eth"
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
	m[OKEX] = log.NewWithModule(OKEX)
	m[OKEX].Logger.SetLevel(log.ParseLevel(config.Log.Module.OKEX))
	m[ETH] = log.NewWithModule(ETH)
	m[ETH].Logger.SetLevel(log.ParseLevel(config.Log.Module.ETH))

	w = &loggerWrapper{loggers: m}
}

func Logger(name string) logrus.FieldLogger {
	return w.loggers[name]
}
