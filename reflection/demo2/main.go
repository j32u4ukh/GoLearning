package main

import (
	"fmt"
	"reflect"
)

type Member struct {
	id   int
	name string
	age  int
}

func main() {
	member := Member{1, "Adam", 100}

	//將資料本身透過 reflect 轉換為物件結構
	t := reflect.TypeOf(member)  //取得所有元素
	v := reflect.ValueOf(member) //獲得值

	fmt.Println(t) //output main.Member
	fmt.Println(v) //output {1 Adam 100}

	//將資料本身透過 reflect 轉換為物件結構

	for i := 0; i < t.NumField(); i++ {
		file := t.Field(i)
		fmt.Println(file.Name, file.Type)
	}
}
