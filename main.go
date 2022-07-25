package main

import (
	"fmt"
	"reflect"

	lproto "GoLearning/proto"

	// "github.com/golang/protobuf/proto"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"google.golang.org/protobuf/proto"
)

type MyStruct struct {
	A int
	B int
}

var typeRegistry = make(map[string]reflect.Type)

func main() {
	msg := &lproto.Message{}
	msg.Text = *proto.String("Hello Reflection")

	ref := msg.ProtoReflect()
	descriptor := ref.Descriptor()
	fmt.Printf("Name: %s\n", descriptor.Name())

	data, _ := proto.Marshal(msg)
	fmt.Printf("data: %+v\n", data)

	messagePb := &lproto.Message{}
	_ = proto.Unmarshal(data, messagePb)
	fmt.Printf("messagePb: %+v, FullName: %+v\n", messagePb, descriptor.FullName())

	b := []byte{1, 2}
	fmt.Println("b:", b)
	Keyboard()
}

func Keyboard() {
	robotgo.EventHook(hook.KeyDown, []string{}, func(e hook.Event) {
		fmt.Println(e.Keycode, e.Keychar)
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
