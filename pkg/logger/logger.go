package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

func NewLogger(outputs []string) Logger {
	return newZap(outputs)
}
