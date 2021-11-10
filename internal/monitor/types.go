package monitor

type IMonitor interface {
	Start() error

	Stop() error
}
