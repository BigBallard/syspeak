package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel uint8

const (
	LogLevelTrace LogLevel = iota
	LogLevelInfo           = iota
	LogLevelDebug          = iota
	LogLevelWarn           = iota
	LogLevelError          = iota
	LogLevelFatal          = iota
)

type Logger interface {
	Trace(message string)
	Tracef(format string, args ...interface{})
	Info(message string)
	Infof(format string, args ...interface{})
	Debug(message string)
	Debugf(format string, args ...interface{})
	Warn(message string)
	Warnf(format string, args ...interface{})
	Error(message string)
	Errorf(format string, args ...interface{})
	Fatal(message string)
	Fatalf(format string, args ...interface{})
}

type syspeakLogger struct {
	name      string
	logLogger *log.Logger
}

func NewSysSpeakLogger(name string) Logger {
	logLogger := log.New(os.Stdout, "", 0)
	return &syspeakLogger{
		name:      name,
		logLogger: logLogger,
	}
}

func (l *syspeakLogger) log(level LogLevel, message string) {
	var levelStr string
	switch level {
	case LogLevelTrace:
		levelStr = "TRACE"
	case LogLevelInfo:
		levelStr = "INFO "
	case LogLevelDebug:
		levelStr = "DEBUG"
	case LogLevelWarn:
		levelStr = "WARN "
	case LogLevelError:
		levelStr = "ERROR"
	default:
		levelStr = "FATAL"
	}
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	l.logLogger.Printf("%s [%s] %s - %s", timeStr, levelStr, l.name, message)
}

func (l *syspeakLogger) Trace(message string) {
	l.log(LogLevelTrace, message)
}

func (l *syspeakLogger) Tracef(format string, args ...interface{}) {
	l.Trace(fmt.Sprintf(format, args...))
}

func (l *syspeakLogger) Info(message string) {
	l.log(LogLevelInfo, message)
}

func (l *syspeakLogger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *syspeakLogger) Debug(message string) {
	l.log(LogLevelDebug, message)
}

func (l *syspeakLogger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *syspeakLogger) Warn(message string) {
	l.log(LogLevelWarn, message)
}

func (l *syspeakLogger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *syspeakLogger) Error(message string) {
	l.log(LogLevelError, message)
}

func (l *syspeakLogger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *syspeakLogger) Fatal(message string) {
	l.log(LogLevelFatal, message)
}

func (l *syspeakLogger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
