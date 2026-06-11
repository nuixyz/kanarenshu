package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)

type Logger struct {
	file *os.File
	std  *log.Logger
}

// package-level logger instance
var global *Logger

func LogFilePath() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("Could not determine home directory: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}
	dir := filepath.Join(dataHome, "kanarenshu")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("Could not create log directory %s: %w", dir, err)
	}
	return filepath.Join(dir, "debug.log"), nil
}

func Init() (func(), error) {
	path, err := LogFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Could not open log file %s: %w", path, err)
	}

	std := log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)
	global = &Logger{file: f, std: std}

	global.Info("Logger initialized, log file: %s", path)

	cleanup := func() {
		global.Info("Application shutting down")
		_ = f.Close()
	}
	return cleanup, nil
}

func InitDiscard() {
	global = &Logger{
		file: nil,
		std:  log.New(io.Discard, "", 0),
	}
}

func Info(format string, args ...any) {
	if global == nil {
		return
	}
	global.Info(format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.std.Output(2, "[INFO] "+fmt.Sprintf(format, args...))
}

func Error(format string, args ...any) {
	if global == nil {
		return
	}
	global.Error(format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.std.Output(2, "[ERROR] "+fmt.Sprintf(format, args...))
}

func Debug(format string, args ...any) {
	if global == nil {
		return
	}
	global.Debug(format, args...)
}

func (l *Logger) Debug(format string, args ...any) {
	l.std.Output(2, "[DEBUG] "+fmt.Sprintf(format, args...))
}

func RecoverAndLog(onCrash func(reason string)) {
	if r := recover(); r != nil {
		stack := debug.Stack()
		Error("PANIC: %v\n%s", r, stack)

		if global != nil && global.file != nil {
			_ = global.file.Sync()
		}

		if onCrash != nil {
			path, _ := LogFilePath()
			onCrash(fmt.Sprintf("kanarenshu crashed: %v\nSee %s for details.", r, path))
			return
		}

		panic(r)
	}
}
