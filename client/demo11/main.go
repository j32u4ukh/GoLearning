package main

import (
	lproto "GoLearning/proto"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	res, err := http.Get("http://192.168.100.142:8080/cat")
	checkError(err)

	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("body: %+v\n", body)
	checkError(err)

	messagePb := lproto.Message{}
	err = proto.Unmarshal(body, &messagePb)

	checkError(err)

	// fmt.Printf("received message: %s, timestamp: %v\n", messagePb.Text, messagePb.Timestamp)

	ref := proto.MessageReflect(&messagePb)
	descriptor := ref.Descriptor()

	filedIds := descriptor.Fields()
	for i := 0; i < filedIds.Len(); i++ {
		filedId := filedIds.Get(i)
		// fmt.Printf("%s %s %v\n", filedId.Name(), filedId.FullName(), filedId.Kind())
		fmt.Printf("Has %s? %v\n", filedId.Name(), ref.Has(filedId))
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
