package logger

import (
	"GoLearning/server/demo14/array"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

// log level
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

type Logger struct {
	logger *log.Logger
	level  int
	// formatter []string
}

func NewLogger(outputs ...string) *Logger {
	l := &Logger{}
	fmt.Println("#outputs: ", len(outputs))

	if len(outputs) == 0 {
		l.logger = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		writers := make([]io.Writer, len(outputs))
		sa := array.StringArray{Elements: outputs}
		if sa.Contains("default") || sa.Contains("Default") {
			d := sa.Remove("default")

			if d == nil {
				d = sa.Remove("Default")
			}

			if d != nil {
				writers = append(writers, os.Stdout)
			}
		}
		for _, path := range sa.Elements {
			f, _ := os.Create(path)
			writers = append(writers, f)
		}

		l.logger = log.New(io.MultiWriter(writers...), "", log.LstdFlags)
	}
	l.level = DebugLevel
	return l
}

func (l *Logger) SetLogLevel(level int) {
	l.level = level
}

func (l *Logger) Debug(message string) {
	if l.level > DebugLevel {
		return
	}

	pc, file, line, ok := runtime.Caller(1)

	if ok {
		// fileName := path.Base(file)
		funcName := runtime.FuncForPC(pc).Name()
		l.logger.Printf("[%s] %s | %s (%d)\n", funcName, message, file, line)
	} else {
		l.logger.Println(message)
	}
}

func (l *Logger) Info(message string) {
	if l.level > InfoLevel {
		return
	}

	pc, file, line, ok := runtime.Caller(1)

	if ok {
		// fileName := path.Base(file)
		funcName := runtime.FuncForPC(pc).Name()
		l.logger.Printf("[%s] %s | %s (%d)\n", funcName, message, file, line)
	} else {
		l.logger.Println(message)
	}
}

func (l *Logger) Warning(message string) {
	if l.level > WarningLevel {
		return
	}

	pc, file, line, ok := runtime.Caller(1)

	if ok {
		// fileName := path.Base(file)
		funcName := runtime.FuncForPC(pc).Name()
		l.logger.Printf("[%s] %s | %s (%d)\n", funcName, message, file, line)
	} else {
		l.logger.Println(message)
	}
}

func (l *Logger) Error(message string) {
	pc, file, line, ok := runtime.Caller(1)

	if ok {
		// fileName := path.Base(file)
		funcName := runtime.FuncForPC(pc).Name()
		l.logger.Printf("[%s] %s | %s (%d)\n", funcName, message, file, line)
	} else {
		l.logger.Println(message)
	}
}
