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
	filedIds := descriptor.Fields()

	for i := 0; i < filedIds.Len(); i++ {
		filedId := filedIds.Get(i)
		// fmt.Printf("%s %s %v\n", filedId.Name(), filedId.FullName(), filedId.Kind())
		fmt.Printf("Has %s? %v\n", filedId.Name(), ref.Has(filedId))
	}
}
