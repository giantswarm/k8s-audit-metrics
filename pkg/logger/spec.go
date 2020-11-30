package logger

type Logger interface {
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warning(format string, params ...interface{})
	Error(format string, params ...interface{})
}
