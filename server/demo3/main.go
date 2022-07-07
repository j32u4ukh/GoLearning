package main

import (
	lproto "GoLearning/proto"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("starting tcp server")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
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

	data := make([]byte, 4096)
	length, err := conn.Read(data)
	checkError(err)

	messagePb := lproto.Message{}
	err = proto.Unmarshal(data[:length], &messagePb)
	checkError(err)

	result := fmt.Sprintf("received message: %s, timestamp: %v", messagePb.Text, messagePb.Timestamp)
	conn.Write([]byte(result))
	// response(result)
}

func response(result string) {
	res, err := net.Dial("tcp", "127.0.0.1:8081")
	checkError(err)
	defer res.Close()

	length, err := res.Write([]byte(result))
	checkError(err)
	fmt.Println(length)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
