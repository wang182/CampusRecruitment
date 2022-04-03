package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	logrus.FieldLogger
}

type MyLogger struct {
	*logrus.Logger
}

func (l *MyLogger) Write(p []byte) (n int, err error) {
	return l.Out.Write(p)
}

var (
	defaultLogger Logger
)

func newLogger() *MyLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		PadLevelText:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000",
	})
	return &MyLogger{Logger: logger}
}

func init() {
	defaultLogger = newLogger()
}

func Get() Logger {
	return defaultLogger
}

func Writer() io.Writer {
	return defaultLogger.(io.Writer)
}
