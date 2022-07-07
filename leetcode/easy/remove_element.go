package leetcode_easy

func RemoveElement(nums []int, val int) int {
	length := len(nums)
	pre := -1
	index := 0
	curr := nextIndex(&nums, pre, val, length)

	for curr != -1 {
		nums[index] = nums[curr]
		pre = curr
		curr = nextIndex(&nums, pre, val, length)
		index++
	}

	return index
}

func nextIndex(nums *[]int, pre int, target int, length int) int {
	for pre < length {
		pre++

		if pre >= length {
			return -1
		}

		if (*nums)[pre] != target {
			return pre
		}
	}

	return -1
}
