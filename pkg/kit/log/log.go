package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type loggerContext struct {
	loggers map[string]logrus.FieldLogger
	config  *config
	hooks   []logrus.Hook
}

var loggerCtx = defaultLoggerContext()

func defaultLoggerContext() *loggerContext {
	return &loggerContext{
		loggers: make(map[string]logrus.FieldLogger),
		config:  defaultConfig(),
		hooks:   make([]logrus.Hook, 0),
	}
}

func New() *logrus.Logger {
	logger := logrus.New()

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05.000",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := filepath.Split(f.File)
			return "", fmt.Sprintf("%12s:%-4d", filename, f.Line)
		},
	}

	logger.SetFormatter(formatter)
	logger.SetReportCaller(loggerCtx.config.reportCaller)
	logger.SetOutput(os.Stdout)

	for _, hook := range loggerCtx.hooks {
		logger.AddHook(hook)
	}

	return logger
}

func NewWithModule(name string) *logrus.Entry {
	logger := New()

	l := logger.WithField("module", name)

	loggerCtx.loggers[name] = l

	return l
}

func ParseLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	return lvl
}

// Initialize initializes a logger instance with given
// level, filepath, filename, maxSize, maxAge and rotationTime.
func Initialize(opts ...Option) error {
	config := generateConfig(opts...)

	loggerCtx.config = config

	if err := os.MkdirAll(config.filePath, os.ModePerm); err != nil {
		return fmt.Errorf("create file path: %w", err)
	}

	if config.persist {
		rotation := newRotateHook(config.filePath, config.fileName, config.maxAge, config.rotationTime)

		loggerCtx.hooks = append(loggerCtx.hooks, rotation)
	}

	return nil
}
