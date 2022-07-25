package log16

import (
	"GoLearning/logger"
	"fmt"
	"sync"
)

var lg *logger.Logger
var once sync.Once

func Logger() *logger.Logger {
	if lg == nil {
		once.Do(func() {
			fmt.Println("Creating AnserManager Instance Now")
			lg = logger.NewLogger()
		})
	}

	return lg
}
