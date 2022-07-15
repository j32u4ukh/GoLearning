package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now")
				singleInstance = &single{}
			})
		fmt.Println("Single Instance already created-1")
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
