package relay

import "google.golang.org/protobuf/reflect/protoreflect"

type Site struct {
}

// 收到註冊的 ProtoMessage 時通知該物件
func (s Site) RegisterCallback(protoreflect.ProtoMessage) {

}
