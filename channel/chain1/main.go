package main

import (
	"fmt"
	"time"
)

type Stage1 struct {
	StopCh chan bool
}

type Stage2 struct {
	StopCh chan bool
}

func main() {
	stage1 := &Stage1{StopCh: make(chan bool)}
	go stage1.Run()
	<-stage1.StopCh
}

func (self *Stage1) Run() {
	stage2 := &Stage2{StopCh: make(chan bool)}
	go stage2.Run()

	for {
		select {
		case <-stage2.StopCh:
			self.StopCh <- true
		}
	}
}

func (self *Stage2) Run() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)

		if i == 6 {
			self.StopCh <- true
			return
		}
	}
}
