package tcp

import (
	"GoLearning/relay"
	"GoLearning/relay/code"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var instance *TcpRouter
var once sync.Once

type TcpRouter struct {
	relay.IRouter
	router relay.Router
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// public

func GetTcpRouter() *TcpRouter {
	if instance == nil {
		once.Do(func() {
			instance = &TcpRouter{}
		})
	}
	return instance
}

func (p TcpRouter) Run(ip string, port int) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// listen to client
	startServer(listener)
}

// 傳送 Protobuf 訊息
func (p TcpRouter) SendMessage(msg protoreflect.ProtoMessage) {

}

////////////////////////////////////////////////////////////////////////////////////////////////////
// private

func startServer(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		fmt.Println("Addr:", conn.RemoteAddr().String())

		// handler goroutine
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 128)
	var uid string
	var C *relay.Conn

	for {
		conn.Read(data)
		fmt.Println("客戶端發來數據:", string(data))

		// register of client
		if data[0] == code.ReqRegister {
			// register
			conn.Write([]byte{code.ResRegister, '#', 'o', 'k'})
			uid = string(data[2:])
			C = relay.NewConn(uid)
			instance.router.ConnMap[uid] = C
			//			fmt.Println("register client")
			//			fmt.Println(uid)
			break
		} else {
			conn.Write([]byte{code.ResRegister, '#', 'e', 'r'})
		}
	}

	//	WHandler
	go writeHandler(conn, C)

	//	RHandler
	go readHandler(conn, C)

	//	Worker
	go processWork(C)

	// Wait for shutdown command
	if <-C.Ech {
		fmt.Println("close handler goroutine")
	}
}

// 正常寫數據
// 定時檢測 conn die => goroutine die
func writeHandler(conn net.Conn, C *relay.Conn) {
	// 讀取業務Work 寫入Wch的數據
	ticker := time.NewTicker(20 * time.Second)

	for {
		select {
		case d := <-C.Wch:
			conn.Write(d)
		case <-ticker.C:
			if _, ok := instance.router.ConnMap[C.Uid]; !ok {
				fmt.Println("conn die, close WHandler")
				return
			}
		}
	}
}

// 讀客戶端數據 + 心跳檢測
func readHandler(conn net.Conn, C *relay.Conn) {
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
			// 可能是來自客戶端的消息確認
			//           	     數據消息
			fmt.Println(data)

			if data[0] == code.ResHello {
				fmt.Println("recv client data ack")
			} else if data[0] == code.ReqHello {
				fmt.Println("recv client data")
				fmt.Println(data)
				conn.Write([]byte{code.ResHello, '#'})
				// C.Rch <- data
			}

			continue
		}

		conn.Write([]byte{code.ReqHeartBeat, '#'})
		fmt.Println("send ht packet")

		// Update the time that will kill the connection
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		if _, herr := conn.Read(data); herr == nil {
			// fmt.Println(string(data))
			fmt.Println("resv ht packet ack")
		} else {
			delete(instance.router.ConnMap, C.Uid)
			fmt.Println("delete user!")
			return
		}
	}
}

func processWork(C *relay.Conn) {
	time.Sleep(5 * time.Second)
	C.Wch <- []byte{code.ReqHello, '#', 'h', 'e', 'l', 'l', 'o'}

	time.Sleep(15 * time.Second)
	C.Wch <- []byte{code.ReqHello, '#', 'h', 'e', 'l', 'l', 'o'}
	// 從讀ch讀信息
	/*	ticker := time.NewTicker(20 * time.Second)
		for {
			select {
			case d := <-C.Rch:
				C.Wch <- d
			case <-ticker.C:
				if _, ok := CMap[C.u]; !ok {
					return
				}
			}
		}
	*/ // 往寫ch寫信息
}
