package ans

import (
	"GoLearning/server/demo14/config"
	"fmt"
	"net"
	"time"
)

type Anser struct {
	ConnMap map[string]*Conn
	// 監聽連線物件
	listener *net.TCPListener
	// 連線位置
	laddr *net.TCPAddr
	// 斷線 channel
	stopCh chan bool
}

// 連線到 Anser 的使用者
type Conn struct {
	// user id
	uid string
	// 連線物件
	conn *net.TCPConn
	// 讀、寫、斷線 等功能使用 chan，應該是要利用其阻塞的特性
	Dch chan bool
	Rch chan []byte
	Wch chan []byte
}

func NewAnser(ip net.IP, port int) *Anser {
	a := &Anser{ConnMap: make(map[string]*Conn), laddr: &net.TCPAddr{IP: ip, Port: port, Zone: ""}}
	return a
}

func NewConn(uid string) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), uid: uid}
}

func (a *Anser) ListenTCP() error {
	var err error
	a.listener, err = net.ListenTCP("tcp", a.laddr)
	return err
}

// 等待使用者連入
func (a *Anser) Run() {
	for {
		conn, err := a.listener.AcceptTCP()

		if err != nil {
			fmt.Println("接受客戶端連接異常:", err.Error())
			continue
		}

		fmt.Println("客戶端連接來自:", conn.RemoteAddr().String())

		// handler goroutine
		go a.handler(conn)
	}
}

func (a *Anser) handler(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 128)
	var uid string
	var C *Conn

	// bind
	for {
		conn.Read(data)
		fmt.Println("客戶端首次連接數據:", string(data))

		// register of client
		if data[0] == config.GetSendCode().RegisterReq {
			// register
			conn.Write([]byte{config.GetSendCode().RegisterRes, '#', 'o', 'k'})
			uid = string(data[2:])
			fmt.Println("uid:", uid)
			C = NewConn(uid)
			a.ConnMap[uid] = C
			break
		} else {
			conn.Write([]byte{config.GetSendCode().RegisterRes, '#', 'e', 'r'})
		}
	}

	//	WHandler
	go a.wHandler(C)

	//	RHandler
	go a.rHandler(C)

	// Wait for shutdown command
	if <-C.Dch {
		fmt.Println("close handler goroutine")
	}
}

// 正常寫數據
// 定時檢測 conn die => goroutine die
func (a *Anser) wHandler(C *Conn) {
	// 讀取業務Work 寫入Wch的數據
	ticker := time.NewTicker(20 * time.Second)

	for {
		select {
		case d := <-C.Wch:
			C.conn.Write(d)
		case <-ticker.C:
			if _, ok := a.ConnMap[C.uid]; !ok {
				fmt.Println("conn die, close WHandler")
				return
			}
		}
	}
}

// 讀客戶端數據 + 心跳檢測
func (s *Anser) rHandler(C *Conn) {
	// 心跳ack
	// 業務數據 寫入Wch

	for {
		data := make([]byte, 128)

		// setReadTimeout
		// Update the time that will kill the connection
		err := C.conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		if err != nil {
			fmt.Println(err)
		}

		// Data from client, maybe request or response
		if _, derr := C.conn.Read(data); derr == nil {
			if data[0] == config.GetSendCode().Res {
				fmt.Println("recv client data ack")
			} else if data[0] == config.GetSendCode().Req {
				fmt.Println("recv client data")
				fmt.Println(data)
				C.conn.Write([]byte{config.GetSendCode().Res, '#'})
				// C.Rch <- data
			}

			continue
		}

		C.conn.Write([]byte{config.GetSendCode().HeartBeatReq, '#'})
		fmt.Println("send ht packet")

		// Update the time that will kill the connection
		C.conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		if _, herr := C.conn.Read(data); herr == nil {
			// fmt.Println(string(data))
			fmt.Println("resv ht packet ack")
		} else {
			delete(s.ConnMap, C.uid)
			fmt.Println("delete user!\nherr: ", herr)
			return
		}
	}
}
