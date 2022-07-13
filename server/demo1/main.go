package main

import (
	lproto "GoLearning/proto"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("starting tcp server")
	listener, err := net.Listen("tcp", ":3306")
	checkError(err)
	timestamp := time.Now().Unix()
	log.Printf("timestamp: %v", timestamp)

	for {
		conn, err := listener.Accept()

		if err == nil {
			log.Println("client connected")
			handleConn(conn)
		} else {
			log.Printf("%+v", err)
		}
	}
}

func handleConn(conn net.Conn) {
	log.Println("client connected")

	defer conn.Close()

	timestamp := time.Now().Unix()

	log.Println("client connected")
	messageProto := lproto.Message{Text: "Hello World", Timestamp: timestamp}
	data, err := proto.Marshal(&messageProto)
	log.Printf("received message: %s, timestamp: %v", messageProto.Text, messageProto.Timestamp)
	checkError(err)

	length, err := conn.Write(data)
	checkError(err)

	log.Printf("Hello world sent, length %d bytes", length)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
