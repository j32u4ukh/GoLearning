package site

import (
	"GoLearning/client/demo13/callback"
	lproto "GoLearning/proto"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type TGS struct {
	Call callback.MessageCallback
	msg  *lproto.Message
}

func (t *TGS) Init() {
	t.Call = callback.MessageCallback{}
	t.Call.SetCallback(t.Callback)
}

func (t *TGS) Callback(data []byte) {
	t.msg = &lproto.Message{}
	_ = proto.Unmarshal(data, t.msg)
	fmt.Println("TGS get data, Msg:", t.msg)
}
