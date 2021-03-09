package loggers

import (
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/sirupsen/logrus"
)

const (
	APP     = "app"
	BSC     = "bsc"
	OkCHAIN = "okchain"
	ETH     = "eth"
)

var w *loggerWrapper

type loggerWrapper struct {
	loggers map[string]*logrus.Entry
}

func Initialize(config *repo.Config) {
	m := make(map[string]*logrus.Entry)
	m[APP] = log.NewWithModule(APP)
	m[APP].Logger.SetLevel(log.ParseLevel(config.Log.Module.APP))
	m[OkCHAIN] = log.NewWithModule(OkCHAIN)
	m[OkCHAIN].Logger.SetLevel(log.ParseLevel(config.Log.Module.OkChain))
	m[ETH] = log.NewWithModule(ETH)
	m[ETH].Logger.SetLevel(log.ParseLevel(config.Log.Module.ETH))

	w = &loggerWrapper{loggers: m}
}

func Logger(name string) logrus.FieldLogger {
	return w.loggers[name]
}
