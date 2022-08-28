package logutil

import (
	"log"
	"os"
	"sync/atomic"

	"github.com/fatih/color"
)

type Level int32

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var (
	DebugColor = color.New(color.FgMagenta).SprintfFunc()
	InfoColor  = color.New(color.FgGreen).SprintfFunc()
	WarnColor  = color.New(color.FgYellow).SprintfFunc()
	ErrorColor = color.New(color.FgRed).SprintfFunc()
	FatalColor = color.New(color.FgRed, color.Bold).SprintfFunc()
)

type Logger struct {
	logOut *log.Logger
	logErr *log.Logger
	level  Level
}

var DefaultLogger = New(LevelInfo)

func New(level Level) *Logger {
	return &Logger{
		logOut: log.New(os.Stdout, "", log.LstdFlags),
		logErr: log.New(os.Stderr, "", log.LstdFlags),
		level:  level,
	}
}

func (l *Logger) Debugf(format string, a ...any) {
	if l.level <= LevelDebug {
		l.logOut.Println(DebugColor("[DEBUG] "+format, a...))
	}
}

func (l *Logger) Infof(format string, a ...any) {
	if l.level <= LevelInfo {
		l.logOut.Println(InfoColor("[INFO] "+format, a...))
	}
}

func (l *Logger) Warnf(format string, a ...any) {
	if l.level <= LevelWarn {
		l.logErr.Println(WarnColor("[WARN] "+format, a...))
	}
}

func (l *Logger) Errorf(format string, a ...any) {
	if l.level <= LevelError {
		l.logErr.Println(ErrorColor("[ERROR} "+format, a...))
	}
}

func (l *Logger) Fatalf(format string, a ...any) {
	if l.level <= LevelFatal {
		l.logErr.Println(FatalColor("[FATAL] "+format, a...))
	}
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func Debugf(format string, a ...any) {
	DefaultLogger.Debugf(format, a...)
}

func Infof(format string, a ...any) {
	DefaultLogger.Infof(format, a...)
}

func Warnf(format string, a ...any) {
	DefaultLogger.Warnf(format, a...)
}

func Errorf(format string, a ...any) {
	DefaultLogger.Errorf(format, a...)
}

func Fatalf(format string, a ...any) {
	DefaultLogger.Fatalf(format, a...)
}

func SetLevel(level Level) {
	DefaultLogger.SetLevel(level)
}
