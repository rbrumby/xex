package xex

type LogLevel uint8

const (
	ERROR LogLevel = iota
	INFO
	DEBUG
)

var LogLevels map[LogLevel]string = map[LogLevel]string{
	ERROR: "ERROR",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

func (ll LogLevel) String() string {
	return LogLevels[ll]
}

type Logger interface {
	Debug(entries ...interface{})
	Debugf(fmtStr string, args ...interface{})
	Info(entries ...interface{})
	Infof(fmtStr string, args ...interface{})
	Error(entries ...interface{})
	Errorf(fmtStr string, args ...interface{})
}

func SetLogger(l Logger) {
	if l == nil {
		logger = &noLogger{}
		return
	}
	logger = l
}

var logger Logger = &noLogger{}

type noLogger struct{}

func (n *noLogger) Debug(entries ...interface{})              {}
func (n *noLogger) Debugf(fmtStr string, args ...interface{}) {}
func (n *noLogger) Info(entries ...interface{})               {}
func (n *noLogger) Infof(fmtStr string, args ...interface{})  {}
func (n *noLogger) Error(entries ...interface{})              {}
func (n *noLogger) Errorf(fmtStr string, args ...interface{}) {}
