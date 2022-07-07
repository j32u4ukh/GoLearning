package proto

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	teacher := Teacher{
		Name: "Henry",
		Age:  17,
	}

	// fmt.Println(teacher)
	// m := proto.MessageV1(&teacher)
	data, err := proto.Marshal(proto.MessageV1(&teacher))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	var t Teacher
	err = proto.Unmarshal(data, &t)
	fmt.Println(&t)
}
