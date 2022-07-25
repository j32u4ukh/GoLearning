package main

import (
	"fmt"
	"net"
	"time"
)

var (
	Req  byte = 1
	Res  byte = 2
	Udid int  = 0
)

// connection struct
type Conn struct {
	// read
	Rch chan []byte

	// write
	Wch chan []byte

	// shutdoen
	Dch chan bool

	// user id
	uid int
}

func NewConn(uid int) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), uid: uid}
}

type Server struct {
	ConnMap  map[int]*Conn
	listener *net.TCPListener
}

func (s *Server) Init(ip string, port int) error {
	var err error
	s.ConnMap = make(map[int]*Conn)
	s.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})
	return err
}

func (s *Server) Run() {
	for {
		conn, err := s.listener.AcceptTCP()

		if err != nil {
			fmt.Println("接受客戶端連接異常:", err.Error())
			continue
		}

		fmt.Println("客戶端連接來自:", conn.RemoteAddr().String())

		// handler goroutine
		go s.handler(conn)
	}
}

func (s *Server) handler(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 128)
	// var C *Conn
	// NOTE: jmeter 中的 TCPClient classname 需設為 org.apache.jmeter.protocol.tcp.sampler.BinaryTCPClientImpl，
	// 表示傳送數據為 Binary
	conn.Read(data)
	fmt.Println("客戶端發來數據:", data, string(data))
	res := []byte{Res, '0'}
	fmt.Println("res:", res)
	conn.Write(res)
	// Udid += 1
	// C = NewConn(Udid)
	// s.ConnMap[Udid] = C

	// for {
	// 	conn.Read(data)
	// 	fmt.Println("客戶端發來數據:", string(data))
	// 	res := []byte{Res, '0'}
	// 	fmt.Println("res:", res)
	// 	conn.Write(res)
	// 	Udid += 1
	// 	C = NewConn(Udid)
	// 	s.ConnMap[Udid] = C
	// 	break

	// 	// // register of client
	// 	// if data[0] == Req {
	// 	// 	val := int8(data[1])
	// 	// 	val += 1
	// 	// 	conn.Write([]byte{Res, byte(val)})
	// 	// 	Udid += 1
	// 	// 	C = NewConn(Udid)
	// 	// 	s.ConnMap[Udid] = C
	// 	// 	break
	// 	// }
	// }

	// //	WHandler
	// go s.wHandler(conn, C)

	// //	RHandler
	// go s.rHandler(conn, C)

	// // Wait for shutdown command
	// if <-C.Dch {
	// 	fmt.Println("close handler goroutine")
	// }
}

// 正常寫數據
// 定時檢測 conn die => goroutine die
func (s *Server) wHandler(conn net.Conn, C *Conn) {
	// 讀取業務 Work 寫入Wch的數據
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case d := <-C.Wch:
			conn.Write(d)
		case <-ticker.C:
			if _, ok := s.ConnMap[C.uid]; !ok {
				fmt.Println("conn die, close WHandler")
				return
			}
		}
	}
}

// 讀客戶端數據
func (s *Server) rHandler(conn net.Conn, C *Conn) {
	// 心跳ack
	// 業務數據 寫入Wch

	for {
		data := make([]byte, 128)

		// setReadTimeout
		// Update the time that will kill the connection
		err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		if err != nil {
			fmt.Println(err)
		}

		// Data from client, maybe request or response
		if _, derr := conn.Read(data); derr == nil {
			res := []byte{Res, '0'}
			fmt.Println("res:", res)
			// data := []byte{Res, '0'}
			C.Wch <- res

			// if data[0] == Req {
			// 	val := int8(data[1])
			// 	val += 1
			// 	data := []byte{Res, byte(val)}
			// 	C.Wch <- data
			// }
		} else {
			delete(s.ConnMap, C.uid)
			// fmt.Println("delete user!")
			return
		}
	}
}

func main() {
	server := &Server{}
	// err := server.Init("3.1.109.13", 6205)
	err := server.Init("https://pergehero.ifunservice.com", 6205)

	if err != nil {
		fmt.Println("監聽端口失敗:", err.Error())
		return
	}

	fmt.Println("已初始化連接，等待客戶端連接...")

	// listen to client
	server.Run()
}
