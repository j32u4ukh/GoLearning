package ans

// golang實現帶有心跳檢測的tcp長連接
// server
import (
	lproto "GoLearning/proto"
	"fmt"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

// message struct:
// c#d

var (
	Close       byte = 0
	registerReq byte = 1 // 1 --- c register cid
	registerRes byte = 2 // 2 --- s response

	heartBeatReq byte = 3 // 3 --- s send heartbeat req
	heartBeatRes byte = 4 // 4 --- c send heartbeat res

	Req byte = 5 // 5 --- c/s send data
	Res byte = 6 // 6 --- c/s send ack

	protobufReq byte = 7 // 7 --- c/s send protobuf data
	protobufRes byte = 8 // 8 --- c/s send ack
)

// connection struct
type Conn struct {
	// read
	Rch chan []byte

	// write
	Wch chan []byte

	// shutdown
	Dch chan bool

	// uid
	u string
}

func NewConn(uid string) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), u: uid}
}

type Server struct {
	ConnMap  map[string]*Conn
	listener *net.TCPListener
}

func (s *Server) Init(ip string, port int) error {
	var err error
	s.ConnMap = make(map[string]*Conn)
	s.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})
	return err
}

// send message to client every 15 seconds
func (s *Server) pushGRT() {
	data := []byte{protobufReq, 1}
	var content []byte
	var err error
	var messageProto lproto.Message

	for {
		time.Sleep(7 * time.Second)
		messageProto = lproto.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
		content, err = proto.Marshal(&messageProto)

		if err != nil {
			continue
		}

		for k, v := range s.ConnMap {
			fmt.Println("push msg to user:" + k)
			v.Wch <- append(data, content...)
		}
	}
}

func (s *Server) Run() {
	// send message to client every 15 seconds
	go s.pushGRT()

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
	defer conn.Write([]byte{registerRes, '#', 'e', 'r'})

	data := make([]byte, 128)
	var uid string
	var C *Conn

	for {
		conn.Read(data)
		fmt.Println("客戶端發來數據:", string(data))

		// register of client
		if data[0] == registerReq {
			// register
			conn.Write([]byte{registerRes, '#', 'o', 'k'})
			uid = string(data[2:])
			fmt.Println("uid:", uid)
			C = NewConn(uid)
			s.ConnMap[uid] = C
			break
		} else {
			conn.Write([]byte{registerRes, '#', 'e', 'r'})
		}
	}

	//	WHandler
	go s.wHandler(conn, C)

	//	RHandler
	go s.rHandler(conn, C)

	//	Worker
	go s.work(C)

	// Wait for shutdown command
	if <-C.Dch {
		fmt.Println("close handler goroutine")
	}
}

// 正常寫數據
// 定時檢測 conn die => goroutine die
func (s *Server) wHandler(conn net.Conn, C *Conn) {
	// 讀取業務Work 寫入Wch的數據
	ticker := time.NewTicker(20 * time.Second)

	for {
		select {
		case d := <-C.Wch:
			conn.Write(d)
		case <-ticker.C:
			if _, ok := s.ConnMap[C.u]; !ok {
				fmt.Println("conn die, close WHandler")
				return
			}
		}
	}
}

// 讀客戶端數據 + 心跳檢測
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
			// 可能是來自客戶端的消息確認
			//           	     數據消息
			fmt.Println("data:", string(data))

			if data[0] == Res {
				fmt.Println("recv client data ack")
			} else if data[0] == Req {
				fmt.Println("recv client data")
				fmt.Println(data)
				conn.Write([]byte{Res, '#'})
				// C.Rch <- data
			}

			continue
		}

		conn.Write([]byte{heartBeatReq, '#'})
		fmt.Println("send ht packet")

		// Update the time that will kill the connection
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		if _, herr := conn.Read(data); herr == nil {
			// fmt.Println(string(data))
			fmt.Println("resv ht packet ack")
		} else {
			delete(s.ConnMap, C.u)
			fmt.Println("delete user!\nherr: ", herr)
			return
		}
	}
}

func (s *Server) work(C *Conn) {
	time.Sleep(5 * time.Second)
	C.Wch <- []byte{Req, '#', 'h', 'e', 'l', 'l', 'o'}

	time.Sleep(5 * time.Second)
	data := []byte{protobufReq, 2}
	teacher := &lproto.Teacher{Name: "Bob", Age: 32}
	msg, _ := proto.Marshal(teacher)
	data = append(data, msg...)

	C.Wch <- data
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

func main() {
	server := &Server{}
	err := server.Init("127.0.0.1", 8080)

	if err != nil {
		fmt.Println("監聽端口失敗:", err.Error())
		return
	}

	fmt.Println("已初始化連接，等待客戶端連接...")

	// listen to client
	server.Run()
}
