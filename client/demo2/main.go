package main

import (
	lproto "GoLearning/proto"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	log.Println("starting tcp client")

	conn, err := net.Dial("tcp", ":8080")
	checkError(err)

	defer conn.Close()

	messageProto := lproto.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
	data, err := proto.Marshal(&messageProto)
	checkError(err)

	length, err := conn.Write(data)
	checkError(err)

	log.Printf("Hello world sent, length %d bytes", length)

	keep := true
	timeAfterTrigger := time.After(time.Second * 1)

	for keep {
		select {
		case curTime := <-timeAfterTrigger:
			go func() {
				messageProto = lproto.Message{Text: "Hello World 2", Timestamp: time.Now().Unix()}
				data, err = proto.Marshal(&messageProto)
				checkError(err)
				length, err := conn.Write(data)
				checkError(err)

				log.Printf("Hello world sent, length %d bytes", length)

				// print current time
				log.Println(curTime.Format("2006-01-02 15:04:05"))
				keep = false
			}()
		default:
			// fmt.Println("Default")
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
