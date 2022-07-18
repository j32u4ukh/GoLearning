package main

import (
	"fmt"
	"reflect"

	lproto "GoLearning/proto"

	// "github.com/golang/protobuf/proto"
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

	// filedIds := descriptor.Fields()

	// for i := 0; i < filedIds.Len(); i++ {
	// 	filedId := filedIds.Get(i)
	// 	fmt.Printf("Has %s? %v\n", filedId.Name(), ref.Has(filedId))
	// }

	// teacher := &lproto.Teacher{}
	// ref = teacher.ProtoReflect()
	// descriptor = ref.Descriptor()
	// fmt.Printf("Name: %s\n", descriptor.Name())

}

// func newMessage(name string) (proto.Message, error) {
// 	reflectType, _ := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(name))
// 	if reflectType == nil {
// 		return nil, fmt.Errorf("protolog: no Message registered for name: %s", name)
// 	}

// 	return reflect.New(reflectType.Elem()).Interface().(proto.Message), nil
// }

func makeInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
}
