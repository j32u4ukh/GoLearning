package ans

import "net"

// 連線到 Anser 的使用者
type Conn struct {
	// user id
	uid string
	// 連線物件
	conn net.Conn
	// 讀、寫、斷線 等功能使用 chan，應該是要利用其阻塞的特性
	Dch chan bool
	Rch chan []byte
	Wch chan []byte
}

func NewConn(uid string) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), uid: uid}
}
