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

type Stage3 struct {
	StopCh chan bool
}

func main() {
	stage1 := &Stage1{StopCh: make(chan bool)}
	go stage1.Run()
	<-stage1.StopCh
}

func (self *Stage1) Run() {
	stage2 := &Stage2{StopCh: make(chan bool)}
	fmt.Println("========== Stage1-1 ==========")
	go stage2.Run()
	<-stage2.StopCh
	fmt.Println("========== Stage1-2 ==========")
	go stage2.Run()
	<-stage2.StopCh
	fmt.Println("========== Stage1-3 ==========")
	go stage2.Run()
	self.StopCh <- <-stage2.StopCh
}

func (self *Stage2) Run() {
	stage3 := &Stage3{StopCh: make(chan bool)}

	fmt.Println("===== Stage2-1 =====")
	go stage3.Run()
	<-stage3.StopCh

	fmt.Println("===== Stage2-2 =====")
	go stage3.Run()
	self.StopCh <- <-stage3.StopCh
}

func (self *Stage3) Run() {
	for i := 0; i < 10; i++ {
		fmt.Println("Stage3", i)
		time.Sleep(1 * time.Second)

		if i == 6 {
			self.StopCh <- true
			return
		}
	}
}
