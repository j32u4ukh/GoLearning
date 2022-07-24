package logger

import (
	"GoLearning/array"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
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

	l.Logout(fmt.Sprintf("[Debug] %s", message))
}

func (l *Logger) Info(message string) {
	if l.level > InfoLevel {
		return
	}

	l.Logout(fmt.Sprintf("[Info] %s", message))
}

func (l *Logger) Warning(message string) {
	if l.level > WarningLevel {
		return
	}

	l.Logout(fmt.Sprintf("[Warning] %s", message))
}

func (l *Logger) Error(message string) {
	l.Logout(fmt.Sprintf("[Error] %s", message))
}

func (l *Logger) Logout(message string) {
	pc, file, line, ok := runtime.Caller(2)

	if ok {
		// fileName := path.Base(file)
		funcName := runtime.FuncForPC(pc).Name()
		names := strings.Split(funcName, ".")

		if len(names) == 2 {
			l.logger.Printf("[%s] %s | %s\n%s (%d)\n", names[0], names[1], message, file, line)
		} else {
			l.logger.Printf("[%s] %s | %s\n%s (%d)\n", names[1], names[2], message, file, line)
		}
	} else {
		l.logger.Println(message)
	}
}
