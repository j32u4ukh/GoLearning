package main

import (
	lproto "GoLearning/proto"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("starting tcp client")
	timestamp := time.Now().Unix()
	log.Printf("timestamp: %d\n", timestamp)
	// conn, err := net.Dial("tcp", ":8080")

	// 昕晨 192.168.100.140
	conn, err := net.Dial("tcp", "192.168.100.142:3306")
	checkError(err)

	defer conn.Close()

	data := make([]byte, 4096)
	length, err := conn.Read(data)
	log.Printf("length: %d\n", length)
	log.Printf("data: %+v\n", data[:length])
	checkError(err)

	messagePb := lproto.Message{}
	err = proto.Unmarshal(data[:length], &messagePb)
	checkError(err)

	log.Printf("received message: %s, timestamp: %v", messagePb.Text, messagePb.Timestamp)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
