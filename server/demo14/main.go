package main

import (
	"GoLearning/server/demo14/utils"
	"fmt"
	"os"
	"time"
)

type Service struct {
	// 總管整個服務的關閉流程(可能有不同原因會觸發關閉流程)
	StopCh chan bool
}

func chanelTest(c chan string) {
	c <- "A"
	time.Sleep(2 * time.Second)
	c <- "B"
}

func (s Service) channelListener(c chan string) {
	for {
		if d := <-c; d != "" {
			fmt.Println(d)

			if d == "C" {
				s.StopCh <- true
				return
			}
		}
	}
}

func main() {
	fmt.Println("整合收發服務")
	service := Service{StopCh: make(chan bool)}

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	config := utils.GetConfig()
	fmt.Println("SendCode.Req:", config.SendCode.Req)
	fmt.Println("Addr.GS:", config.Addr.GS)
	fmt.Println("Conn.Alternal:", config.Conn.Alternal)

	c := make(chan string, 3)
	go chanelTest(c)
	go func() {
		c <- "A"
		time.Sleep(2 * time.Second)
		c <- "B"
		time.Sleep(2 * time.Second)
		c <- "C"
	}()
	go service.channelListener(c)

	<-service.StopCh
}
