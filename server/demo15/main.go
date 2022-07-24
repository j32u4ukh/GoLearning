package main

import (
	"GoLearning/server/demo15/ans"
	"GoLearning/server/demo15/ask"
	"GoLearning/server/demo15/config"
	"GoLearning/server/demo15/log15"
	"fmt"
	"os"
)

type Service struct {
	// 總管整個服務的關閉流程(可能有不同原因會觸發關閉流程)
	StopCh chan bool
}

func main() {
	service := Service{StopCh: make(chan bool)}
	log15.Logger().Debug(fmt.Sprintf("整合收發服務, Args: %+v", os.Args[1:]))

	if os.Args[1] == "ask" {
		go service.RunAsk()

	} else if os.Args[1] == "ans" {
		go service.RunAns()
	}

	<-service.StopCh
}

func (s Service) RunAsk() {
	ask.Init()

	ask.GetAskerManager().Send(config.GetAddr().GS, []byte{config.GetSendCode().Req, 'F', 'i', 'r', 's', 't'})
	ask.GetAskerManager().Send(config.GetAddr().GS, []byte{config.GetSendCode().Req, 'S', 'e', 'c', 'o', 'n', 'd'})

	end := <-ask.GetAskerManager().Cch
	fmt.Println("End of connection.", end)

	s.StopCh <- true
}

func (s Service) RunAns() {
	ans.Init()
	fmt.Println("After ans.Init()")
	s.StopCh <- true
}
