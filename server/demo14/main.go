package main

import (
	"GoLearning/server/demo14/ans"
	"GoLearning/server/demo14/array"
	"GoLearning/server/demo14/ask"
	"GoLearning/server/demo14/config"
	"GoLearning/server/demo14/logger"

	"github.com/pkg/errors"

	// "errors"
	"fmt"
	"os"
)

var lg *logger.Logger

type Service struct {
	// 總管整個服務的關閉流程(可能有不同原因會觸發關閉流程)
	StopCh chan bool
}

type Array[T string | int] struct {
	Elements []T
}

func (a *Array[string]) Len() int { return len(a.Elements) }

func main() {
	// fmt.Println("整合收發服務")
	service := Service{StopCh: make(chan bool)}
	// // fmt.Println(os.Args[1:])
	// lg = logger.NewLogger()
	// lg.Debug("This is debug message.")
	// service.test()

	// fmt.Println("打印运行中的函数名")
	// test1()
	// test2()

	if len(os.Args) == 1 {
		fmt.Println("未輸入服務名稱，結束服務")
		return
	}

	if os.Args[1] == "ask" {
		go service.RunAsk()
	} else if os.Args[1] == "ans" {
		go service.RunAns()
	} else if os.Args[1] == "log" {
		go service.RunLog()
	} else if os.Args[1] == "err" {
		go service.RunErr()
	} else {
		fmt.Println("服務名稱錯誤，結束服務")
		service.StopCh <- true
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
	server := &ans.Server{}
	err := server.Init("127.0.0.1", 8080)

	if err != nil {
		fmt.Println("監聽端口失敗:", err.Error())
		return
	}

	fmt.Println("已初始化連接，等待客戶端連接...")

	// listen to client
	server.Run()
}

func (s Service) RunLog() {
	arr := Array[string]{Elements: []string{"test", "default"}}
	fmt.Println(arr.Len())
	outputs := array.StringArray{Elements: []string{"test", "default"}}
	b := outputs.Contains("default")
	fmt.Println(b)
	s.StopCh <- true
}

func doMath() (int, error) {
	d, err := divide(1, 0)

	return d, errors.Wrap(err, "Do math function")
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cann't divided by 0.")
	}

	return a / b, nil
}

func (s Service) RunErr() {
	val, err := doMath()

	if err != nil {
		fmt.Printf("val: %d, err: %+v", val, err)
	}

	s.StopCh <- true
}
