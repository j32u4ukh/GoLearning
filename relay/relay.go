package relay

import "google.golang.org/protobuf/reflect/protoreflect"

// 路由器(TODO: 單例)
var router *Router

// connection map | key: uid; value: connection
var ConnMap map[string]*Conn

type Router interface {
	Run()

	// 傳送 Protobuf 訊息
	SendMessage(protoreflect.ProtoMessage)

	// // 供外部物件註冊，當收到註冊的 ProtoMessage 時通知該物件
	// // 或許可以提供一個介面給所有收發訊息的物件來鑲嵌，只有鑲嵌了該物件才能註冊
	// RegisterCallback(protoreflect.ProtoMessage)
}

// connection struct
type Conn struct {
	// read
	Rch chan []byte

	// write
	Wch chan []byte

	// exit
	Ech chan bool

	// uid
	uid string
}

func init() {

}

func NewConn(uid string) *Conn {
	return &Conn{Rch: make(chan []byte), Wch: make(chan []byte), uid: uid}
}
