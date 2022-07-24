package ask

import (
	lproto "GoLearning/proto"
	"GoLearning/server/demo15/config"
	"GoLearning/server/demo15/utils"
	"fmt"
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
)

type Asker struct {
	// 連線位置
	Addr string
	// 連線物件
	conn *net.TCPConn
	// 是否為連線中
	isConnected bool
	// 是否為斷線
	isShutdown bool
	// 讀、寫、斷線 等功能使用 chan，應該是要利用其阻塞的特性
	Dch chan bool
	Rch chan []byte
	Wch chan []byte
	// key: Response type; value: callback function
	callbackMap map[string]func([]byte)
}

func NewAsker(addr string) *Asker {
	ans := &Asker{Addr: addr}
	ans.isConnected = false
	ans.isShutdown = false
	ans.Dch = make(chan bool)
	ans.Rch = make(chan []byte)
	ans.Wch = make(chan []byte)
	ans.callbackMap = map[string]func([]byte){}
	return ans
}

// func (a *Asker) Init() {
// 	a.isConnected = false
// 	a.isShutdown = false
// 	a.Dch = make(chan bool)
// 	a.Rch = make(chan []byte)
// 	a.Wch = make(chan []byte)
// 	a.callbackMap = map[string]func([]byte){}
// }

func (a *Asker) Connect() error {
	var err error
	addr, _ := net.ResolveTCPAddr("tcp", a.Addr)
	a.conn, err = net.DialTCP("tcp", nil, addr)

	if err != nil {
		a.isConnected = false
		return err
	}

	fmt.Println("Coonected to", a.Addr)
	a.isConnected = true
	return nil
}

func (a *Asker) Run(wg *sync.WaitGroup) {
	defer a.conn.Close()
	go a.handler()

	// 斷線重連時的 wg 將會是 nil
	if wg != nil {
		wg.Done()
	}

	if <-a.Dch {
		fmt.Println("Addr: ", a.Addr)
		fmt.Println("關閉連接")
	}
}

// func (a *Asker) RunServer(ip string, port int, wg *sync.WaitGroup) {
// 	for {
// 		addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
// 		var err error
// 		a.conn, err = net.DialTCP("tcp", nil, addr)

// 		if err != nil {
// 			fmt.Println("連接服務端失敗:", err.Error())
// 			return
// 		}

// 		fmt.Println("已連接服務器")
// 		defer a.conn.Close()
// 		go a.handler(a.conn)
// 		wg.Done()

// 		if <-a.Dch {
// 			fmt.Println("Addr: ", a.Addr)
// 			fmt.Println("關閉連接")
// 		}
// 	}
// }

func (a *Asker) handler() {
	data := make([]byte, 128)

	// bind
	for {
		// 送出註冊請求
		a.conn.Write([]byte{config.GetSendCode().RegisterReq, '#', '2'})

		// 取得註冊成功回應
		a.conn.Read(data)

		if data[0] == config.GetSendCode().RegisterRes {
			fmt.Println("註冊成功")
			break
		}
	}

	go a.rHandler()
	go a.wHandler()
	// go a.work()
}

func (a *Asker) rHandler() {
	var err error

	for {
		// 心跳包,回覆ack
		data := make([]byte, 1)
		length, _ := a.conn.Read(data)
		fmt.Println("data length:", length)

		if data[0] == config.GetSendCode().Close {
			// TODO: 紀錄狀態為結束連線
			a.Dch <- true
			a.isShutdown = true
		}

		if length == 0 {
			// c.Dch <- true
			// TODO: 若非 結束連線，則需再次連線
			fmt.Println("length == 0")
			return
		}

		if data[0] == config.GetSendCode().HeartBeatReq {
			fmt.Println("recv ht pack")
			a.conn.Write([]byte{config.GetSendCode().RegisterRes, 'h', 'i'})
			fmt.Println("send ht pack ack")
		} else if data[0] == config.GetSendCode().Req {
			fmt.Println("recv data pack")
			data = make([]byte, 4096)
			length, _ = a.conn.Read(data)

			fmt.Printf("%v\n", string(data))
			fmt.Printf("length: %d\n", length)
			// a.Rch <- data[2:]
			a.conn.Write([]byte{config.GetSendCode().Res})
		} else if data[0] == config.GetSendCode().ProtobufReq {
			fmt.Println("Recieve protobuf data")
			pbtype := data[1]

			data = make([]byte, 4096)
			length, _ = a.conn.Read(data)
			var pbstring string

			switch pbtype {
			case 1:
				fmt.Println("Pb type: Message")
				pbstring = "Message"
				messagePb := lproto.Message{}
				err = proto.Unmarshal(data[:length], &messagePb)
				utils.CheckError(err)
				fmt.Printf("messagePb.Text: %s, messagePb.Timestamp: %v\n", messagePb.Text, messagePb.Timestamp)
			case 2:
				fmt.Println("Pb type: Teacher")
				pbstring = "Teacher"
				taecher := lproto.Teacher{}
				err = proto.Unmarshal(data[:length], &taecher)
				utils.CheckError(err)
				fmt.Printf("taecher.Name: %s, taecher.Age: %v\n", taecher.Name, taecher.Age)
			default:
				fmt.Println("Pb type: Nothing")
				pbstring = "Nothing"
			}

			if callback, ok := a.callbackMap[pbstring]; ok {
				callback(data[:length])
			}

			// Rch <- data[2:]
			a.conn.Write([]byte{config.GetSendCode().Res})
		}
	}
}

func (a *Asker) wHandler() {

	for {
		if msg := <-a.Wch; msg != nil {
			fmt.Printf("send code %v, data: %v\n", msg[0], string(msg[1:]))
			a.conn.Write(msg)
		}
	}

}

func (a *Asker) send(msg []byte) {
	a.Wch <- msg
}

func (a *Asker) RegisterFunc(msg string, callback func([]byte)) {
	fmt.Printf("Asker RegisterFunc -> %s\n", msg)
	a.callbackMap[msg] = callback
}
