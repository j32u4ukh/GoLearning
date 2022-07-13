package main

import (
	lproto "GoLearning/proto"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	// messageProto := lproto.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
	// data, err := proto.Marshal(&messageProto)
	// log.Println(data)

	timeAfterTrigger := time.After(time.Second * 2)
	data := []byte{10, 11, 72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 16, 143, 168, 175, 150, 6}
	messagePb := lproto.Message{}
	err := proto.Unmarshal(data, &messagePb)
	log.Printf("received message: %s, timestamp: %v", messagePb.Text, messagePb.Timestamp)
	if err != nil {
		log.Println(err.Error())
	}

	keep := true

	for keep {
		select {
		case curTime := <-timeAfterTrigger:
			// print current time
			fmt.Println(curTime.Format("2006-01-02 15:04:05"))
			keep = false
		default:
			// fmt.Println("Default")
		}
	}
}
