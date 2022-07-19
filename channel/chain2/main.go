package main

import "fmt"

const (
	Stop     byte = 0
	Request  byte = 1
	Response byte = 2
)

type Top struct {
	stopChan chan byte
}

type Mid struct {
	stopChan chan byte
}

type Btm struct {
	stopChan chan byte
}

func main() {
	top := &Top{stopChan: make(chan byte)}
	go top.Run()
	<-top.stopChan
}

func (self *Top) Run() {
	go func() {
		fmt.Println("===== Top 1 =====")
		mid := &Mid{stopChan: make(chan byte)}
		mid.Run(self.stopChan)
	}()
	<-self.stopChan

	// go func() {
	// 	fmt.Println("===== Top 2 =====")
	// 	mid := &Mid{stopChan: make(chan byte)}
	// 	mid.Run(self.stopChan)
	// }()
	// <-self.stopChan

	// go func() {
	// 	fmt.Println("===== Top 3 =====")
	// 	mid := &Mid{stopChan: make(chan byte)}
	// 	mid.Run(self.stopChan)
	// }()
	// <-self.stopChan
}

func (self *Mid) Run(parentCh chan<- byte) {
	// go func() {
	// 	fmt.Println("===== Mid 1 =====")
	// 	btm := &Btm{stopChan: make(chan byte)}
	// 	btm.Run(self.stopChan)
	// }()
	// <-self.stopChan

	go func() {
		fmt.Println("===== Mid 2 =====")
		btm := &Btm{stopChan: make(chan byte)}
		btm.Run(self.stopChan)
	}()
	parentCh <- <-self.stopChan
}

func (self *Btm) Run(parentCh chan<- byte) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Btm ", i)
		}
	}()

	// b := <-self.stopChan
	parentCh <- <-self.stopChan
}
