package main

// golang實現帶有心跳檢測的tcp長連接
// server

import (
	"fmt"
	"sync"
	"time"
)

var (
	Close       byte = 0
	registerReq byte = 1 // 1 --- c register cid
	registerRes byte = 2 // 2 --- s response

	heartBeatReq byte = 3 // 3 --- s send heartbeat req
	heartBeatRes byte = 4 // 4 --- c send heartbeat res

	Req byte = 5 // 5 --- c/s send data
	Res byte = 6 // 6 --- c/s send ack

	protobufReq byte = 7 // 7 --- c/s send protobuf data
	protobufRes byte = 8 // 8 --- c/s send ack

)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Call to registered server
func call() {
	wg := &sync.WaitGroup{}
	wg.Add(len(GetClient13Manager().Connects))

	for _, c := range GetClient13Manager().Connects {
		go GetClient13Manager().Connect(c, wg)
	}

	wg.Wait()
}

func main() {
	GetClient13Manager().Connects = append(GetClient13Manager().Connects, GS)
	call()
	// ip := "127.0.0.1"
	// port := 8080
	// GetClient13Manager().Connect(GS)
	time.Sleep(3 * time.Second)

	GetClient13Manager().Send(GS, []byte{Req, 'F', 'i', 'r', 's', 't'})
	GetClient13Manager().Send(GS, []byte{Req, 'S', 'e', 'c', 'o', 'n', 'd'})

	end := <-GetClient13Manager().Cch
	fmt.Println("End of connection.", end)
}
