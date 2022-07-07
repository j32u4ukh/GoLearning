package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get(":8080/cat")
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
