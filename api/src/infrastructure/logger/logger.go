package logger

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
		PadLevelText:     true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(&MyOutput{})
}

func Info(v ...interface{}) {
	logrus.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func Debug(v ...interface{}) {
	logrus.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	logrus.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	logrus.Warnf(format, v...)
}

func Error(v ...interface{}) {
	logrus.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logrus.Fatalf(format, v...)
}

func New() *logrus.Logger {
	return logrus.StandardLogger()
}

type MyFormatter struct{}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(entry.Time.Format(time.RFC3339))
	b.WriteString(" ")
	b.WriteString(fmt.Sprintf("[%s]", entry.Level))
	b.WriteString(" ")
	if entry.Data["context"] != nil {
		b.WriteString(fmt.Sprintf("(%v)", entry.Data["context"]))
		b.WriteString(" ")
	}
	if entry.Data["requestId"] != nil {
		b.WriteString(fmt.Sprintf("(%v)", entry.Data["requestId"]))
		b.WriteString(" ")
	}
	b.WriteString(fmt.Sprintf("%v", entry.Message))

	b.WriteByte('\n')
	return b.Bytes(), nil
}

type MyOutput struct{}

func (splitter *MyOutput) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("[debug]")) || bytes.Contains(p, []byte("[info]")) {
		return os.Stdout.Write(p)
	}
	return os.Stderr.Write(p)
}

const loggerKey string = "logger"

func InjectInContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := New().WithFields(logrus.Fields{})
		newCtx := context.WithValue(c.Request.Context(), loggerKey, logger)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}

func FromContext(ctx context.Context) *logrus.Entry {
	if _logger, ok := ctx.Value(loggerKey).(*logrus.Entry); ok {
		return _logger
	}
	return logrus.NewEntry(logrus.New())
}
