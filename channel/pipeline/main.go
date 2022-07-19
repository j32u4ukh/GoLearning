package main

import (
	"fmt"
	"time"
)

type Data int
type Square int
type Result int

func main() {
	chanSiae := 3
	rawChan := make(chan int, chanSiae)
	dataChan := make(chan Data, chanSiae)
	squareChan := make(chan Square, chanSiae)

	go func() {
		for i := 0; i < 10; i++ {
			rawChan <- i
		}
	}()

	go StageData(rawChan, dataChan)
	go StageSquare(dataChan, squareChan)
	go StageResult(squareChan)

	time.Sleep(3 * time.Second)
}

func StageData(raw <-chan int, data chan<- Data) {
	for {
		select {
		case rawData := <-raw:
			value := Data(rawData)
			data <- value
			fmt.Printf("[StageData] rawData: %d, data: %d\n", rawData, value)
		}
	}
}

func StageSquare(data <-chan Data, square chan<- Square) {
	for {
		select {
		case d := <-data:
			sq := d * d
			square <- Square(sq)
			fmt.Printf("[StageSquare] data: %d, square: %d\n", d, sq)
		}
	}
}

func StageResult(square <-chan Square) {
	for {
		select {
		case squared := <-square:
			fmt.Printf("Square: %d, Result: %d\n", squared, 3*squared)
		}
	}
}
