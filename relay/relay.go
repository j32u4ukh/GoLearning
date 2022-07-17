package relay

import "google.golang.org/protobuf/reflect/protoreflect"

// 路由器(TODO: 單例)
var router *IRouter

type IRouter interface {
	Run()

	// 傳送 Protobuf 訊息
	SendMessage(protoreflect.ProtoMessage)

	// // 供外部物件註冊，當收到註冊的 ProtoMessage 時通知該物件
	// // 或許可以提供一個介面給所有收發訊息的物件來鑲嵌，只有鑲嵌了該物件才能註冊
	// RegisterCallback(protoreflect.ProtoMessage)
}

type Router struct {
	// connection map | key: uid; value: connection
	ConnMap map[string]*Conn
}

// Request Protocol
// [type 0]: 0 heartbeat; 1 tcp; 2 http
// [code 1]: request type

// Response Protocol
// [type 0]: 0 heartbeat; 1 tcp; 2 http
// [code 1]: request type

// connection struct
type Conn struct {
	// read
	Rch chan []byte

	// write
	Wch chan []byte

	// exit
	Ech chan bool

	// Uid
	Uid string
}

func init() {

}

func NewConn(uid string) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), Uid: uid}
}
