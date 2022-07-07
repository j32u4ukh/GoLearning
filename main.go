package main

import (
	leetcode_easy "GoLearning/leetcode/easy"
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := leetcode_easy.RemoveElement(nums, 2)

	for i, v := range nums[:n] {
		fmt.Println(i, v)
	}
}
