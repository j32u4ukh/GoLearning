package main

// golang實現帶有心跳檢測的tcp長連接
// server

import (
	"fmt"
	"net"

	lproto "GoLearning/proto"

	"google.golang.org/protobuf/proto"
)

var (
	registerReq byte = 1 // 1 --- c register cid
	registerRes byte = 2 // 2 --- s response

	heartBeatReq byte = 3 // 3 --- s send heartbeat req
	heartBeatRes byte = 4 // 4 --- c send heartbeat res

	Req byte = 5 // 5 --- c/s send data
	Res byte = 6 // 6 --- c/s send ack

	protobufReq byte = 7 // 7 --- c/s send protobuf data
	protobufRes byte = 8 // 8 --- c/s send ack
)

var Dch chan bool
var Rch chan []byte
var Wch chan []byte

func main() {
	Dch = make(chan bool)
	Rch = make(chan []byte)
	Wch = make(chan []byte)
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, addr)
	//	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("連接服務端失敗:", err.Error())
		return
	}
	fmt.Println("已連接服務器")
	defer conn.Close()
	go Handler(conn)

	if <-Dch {
		fmt.Println("關閉連接")
	}
	// select {
	// case <-Dch:
	// 	fmt.Println("關閉連接")
	// }
}

func Handler(conn *net.TCPConn) {
	// 直到register ok
	data := make([]byte, 128)
	for {
		conn.Write([]byte{registerReq, '#', '2'})
		conn.Read(data)
		//		fmt.Println(string(data))
		if data[0] == registerRes {
			break
		}
	}
	//	fmt.Println("i'm register")
	go RHandler(conn)
	go WHandler(conn)
	go Work()
}

func RHandler(conn *net.TCPConn) {
	var err error

	for {
		// 心跳包,回覆ack
		data := make([]byte, 2)
		length, _ := conn.Read(data)

		if length == 0 {
			Dch <- true
			return
		}

		if data[0] == heartBeatReq {
			fmt.Println("recv ht pack")
			conn.Write([]byte{registerRes, '#', 'h'})
			fmt.Println("send ht pack ack")
		} else if data[0] == Req {
			fmt.Println("recv data pack")
			data = make([]byte, 4096)
			length, _ = conn.Read(data)

			fmt.Printf("%v\n", string(data))
			Rch <- data[2:]
			conn.Write([]byte{Res, '#'})
		} else if data[0] == protobufReq {
			fmt.Println("Recieve protobuf data")
			data = make([]byte, 4096)
			length, _ = conn.Read(data)

			messagePb := lproto.Message{}
			err = proto.Unmarshal(data[:length], &messagePb)
			checkError(err)

			fmt.Printf("received message: %s, timestamp: %v\n", messagePb.Text, messagePb.Timestamp)
			// Rch <- data[2:]
			conn.Write([]byte{Res, '#'})
		}
	}
}

func WHandler(conn net.Conn) {

	for {
		if msg := <-Wch; msg != nil {
			fmt.Println((msg[0]))
			fmt.Println("send data after: " + string(msg[1:]))
			conn.Write(msg)
		}
		// select {
		// case msg := <-Wch:
		// 	fmt.Println((msg[0]))
		// 	fmt.Println("send data after: " + string(msg[1:]))
		// 	conn.Write(msg)

		// }
	}

}

func Work() {
	for {
		if msg := <-Rch; msg != nil {
			fmt.Println("work recv " + string(msg))
			Wch <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		}

		// select {
		// case msg := <-Rch:
		// 	fmt.Println("work recv " + string(msg))
		// 	Wch <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		// }
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
