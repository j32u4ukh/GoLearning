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

type Client12 struct {
	Dch  chan bool
	Rch  chan []byte
	Wch  chan []byte
	conn *net.TCPConn
}

func (c *Client12) Init(ip string, port int) {
	c.Dch = make(chan bool)
	c.Rch = make(chan []byte)
	c.Wch = make(chan []byte)

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	var err error
	c.conn, err = net.DialTCP("tcp", nil, addr)

	if err != nil {
		fmt.Println("連接服務端失敗:", err.Error())
		return
	}

	fmt.Println("已連接服務器")
}

func (c *Client12) Run() {
	defer c.conn.Close()
	go c.handler(c.conn)

	if <-c.Dch {
		fmt.Println("關閉連接")
	}
}

func (c *Client12) handler(conn *net.TCPConn) {
	data := make([]byte, 128)

	// 直到register ok
	for {
		conn.Write([]byte{registerReq, '#', '2'})
		conn.Read(data)
		//		fmt.Println(string(data))
		if data[0] == registerRes {
			break
		}
	}

	go c.hHandler(conn)
	go c.wHandler(conn)
	go c.work()
}

func (c *Client12) hHandler(conn *net.TCPConn) {
	var err error

	for {
		// 心跳包,回覆ack
		data := make([]byte, 2)
		length, _ := conn.Read(data)

		if length == 0 {
			c.Dch <- true
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
			fmt.Printf("length: %d\n", length)
			c.Rch <- data[2:]
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

func (c *Client12) wHandler(conn net.Conn) {

	for {
		if msg := <-c.Wch; msg != nil {
			fmt.Printf("send code %v data: %v", msg[0], string(msg[1:]))
			conn.Write(msg)
		}
	}

}

func (c *Client12) work() {
	for {
		if msg := <-c.Rch; msg != nil {
			fmt.Println("work recv " + string(msg))
			c.Wch <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func connect8080() {
	client := &Client12{}
	client.Init("127.0.0.1", 8080)
	client.Run()
}

func main() {
	connect8080()
}
