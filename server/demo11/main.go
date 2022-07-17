package main

import (
	lproto "GoLearning/proto"
	"fmt"
	"io"
	"net/http"
	"unsafe"

	"google.golang.org/protobuf/proto"
)

func foo(res http.ResponseWriter, req *http.Request) {
	var data []byte
	req.Body.Read(data)
	io.WriteString(res, "dog dog dog")

	str := "dog dog dog"
	fmt.Printf("str: %v\n", str)
	bs := []byte(str)
	fmt.Printf("bs: %v\n", bs)
	bstr := string(bs)
	fmt.Printf("bstr: %v\n", bstr)
}

func bar(res http.ResponseWriter, req *http.Request) {
	// messageProto := lproto.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
	messageProto := lproto.Message{Flag: false}
	content, _ := proto.Marshal(&messageProto)
	fmt.Printf("content: %+v\n", content)

	// Print variable type and size
	fmt.Printf("content: %T, %d\n", content, unsafe.Sizeof(content))

	res.Write(content)
}

func main() {
	http.HandleFunc("/dog", foo)
	http.HandleFunc("/cat", bar)
	http.ListenAndServe(":8080", nil)
}
