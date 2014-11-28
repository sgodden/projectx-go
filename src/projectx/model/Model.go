package model

type logger interface {
	Log(msg string)
}

var (
	thelogger logger
)

func SetLogger(logger logger) {
	thelogger = logger;
}
