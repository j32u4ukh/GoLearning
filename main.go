package main

import (
	leetcode_easy "GoLearning/leetcode/easy"
	"fmt"
	"time"
)

func main() {
	timeAfterTrigger := time.After(time.Second * 2)
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := leetcode_easy.RemoveElement(nums, 2)

	for i, v := range nums[:n] {
		fmt.Println(i, v)
	}

	keep := true

	for keep {
		select {
		case curTime := <-timeAfterTrigger:
			// print current time
			fmt.Println(curTime.Format("2006-01-02 15:04:05"))
			keep = false
		default:
			// fmt.Println("Default")
		}
	}
}
