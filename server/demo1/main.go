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
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)

	for {
		if conn, err := listener.Accept(); err == nil {
			handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	log.Println("client connected")

	defer conn.Close()

	messageProto := lproto.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
	data, err := proto.Marshal(&messageProto)
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
