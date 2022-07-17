package main

import (
	lproto "GoLearning/proto"
	"fmt"
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
)

type Client13 struct {
	Pch  *chan string
	Addr string
	Dch  chan bool
	Rch  chan []byte
	Wch  chan []byte
	conn *net.TCPConn
}

// TODO: 連線 和 維持運行 要區分為兩個區塊，方便斷線後重連
// TODO: 先註冊要連線的目標、callback 函式，再一起進行連線。目前利用 sync.WaitGroup 等待所有連線成功
func (c *Client13) Init(ip string, port int) {
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

func (c *Client13) Run(wg *sync.WaitGroup) {
	defer c.conn.Close()
	go c.handler(c.conn)
	wg.Done()

	if <-c.Dch {
		fmt.Println("Addr: ", c.Addr)
		fmt.Println("關閉連接")
	}
}

func (c *Client13) RunServer(ip string, port int, wg *sync.WaitGroup) {
	for {
		addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
		var err error
		c.conn, err = net.DialTCP("tcp", nil, addr)

		if err != nil {
			fmt.Println("連接服務端失敗:", err.Error())
			return
		}

		fmt.Println("已連接服務器")
		defer c.conn.Close()
		go c.handler(c.conn)
		wg.Done()

		if <-c.Dch {
			fmt.Println("Addr: ", c.Addr)
			fmt.Println("關閉連接")
		}
	}
}

func (c *Client13) handler(conn *net.TCPConn) {
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

	go c.rHandler(conn)
	go c.wHandler(conn)
	go c.work()
}

func (c *Client13) rHandler(conn *net.TCPConn) {
	var err error

	for {
		// 心跳包,回覆ack
		data := make([]byte, 2)
		length, _ := conn.Read(data)
		fmt.Println("data length:", length)

		if data[0] == Close {
			// TODO: 紀錄狀態為結束連線
			c.Dch <- true
		}

		if length == 0 {
			// c.Dch <- true
			// TODO: 若非 結束連線，則需再次連線
			fmt.Println("length == 0")
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

func (c *Client13) wHandler(conn net.Conn) {

	for {
		if msg := <-c.Wch; msg != nil {
			fmt.Printf("send code %v data: %v\n", msg[0], string(msg[1:]))
			conn.Write(msg)
		}
	}

}

func (c *Client13) work() {
	for {
		if msg := <-c.Rch; msg != nil {
			fmt.Println("work recv " + string(msg))
			c.Wch <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		}
	}
}

func (c *Client13) send(msg []byte) {
	c.Wch <- msg
}
