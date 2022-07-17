package main

import (
	"fmt"

	lproto "GoLearning/proto"

	// "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	msg := &lproto.Message{}
	msg.Text = *proto.String("Hello Reflection")

	ref := msg.ProtoReflect()
	descriptor := ref.Descriptor()
	fmt.Printf("Name: %s\n", descriptor.Name())

	// filedIds := descriptor.Fields()

	// for i := 0; i < filedIds.Len(); i++ {
	// 	filedId := filedIds.Get(i)
	// 	fmt.Printf("Has %s? %v\n", filedId.Name(), ref.Has(filedId))
	// }

	teacher := &lproto.Teacher{}
	ref = teacher.ProtoReflect()
	descriptor = ref.Descriptor()
	fmt.Printf("Name: %s\n", descriptor.Name())

	str := "dog dog dog"
	fmt.Printf("str: %v\n", str)
	bs := []byte(str)
	fmt.Printf("bs: %v\n", bs)
	bstr := string(bs)
	fmt.Printf("bstr: %v\n", bstr)
}
