package log

import "time"

type config struct {
	reportCaller bool
	persist      bool
	filePath     string
	fileName     string
	maxSize      int64
	maxAge       time.Duration
	rotationTime time.Duration
}

type Option func(*config)

func WithReportCaller(reportCaller bool) Option {
	return func(c *config) {
		c.reportCaller = reportCaller
	}
}

func WithPersist(persist bool) Option {
	return func(c *config) {
		c.persist = persist
	}
}

func WithFilePath(filePath string) Option {
	return func(c *config) {
		c.filePath = filePath
	}
}

func WithFileName(fileName string) Option {
	return func(c *config) {
		c.fileName = fileName
	}
}

func WithMaxSize(maxSize int64) Option {
	return func(c *config) {
		c.maxSize = maxSize
	}
}

func WithMaxAge(maxAge time.Duration) Option {
	return func(c *config) {
		c.maxAge = maxAge
	}
}

func WithRotationTime(rotationTime time.Duration) Option {
	return func(c *config) {
		c.rotationTime = rotationTime
	}
}

func defaultConfig() *config {
	return &config{
		reportCaller: false,
		persist:      false,
		filePath:     "./",
		fileName:     "log",
		maxSize:      2 * 1024 * 1024,
		maxAge:       2 * time.Hour,
		rotationTime: 24 * time.Hour,
	}
}

func generateConfig(opts ...Option) *config {
	config := defaultConfig()

	for _, opt := range opts {
		opt(config)
	}

	return config
}
