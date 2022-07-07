package leetcode_easy

func RemoveDuplicates(nums []int) int {
	var point, number int
	length := len(nums)

	if length > 0 {
		point = nums[0]
		number = 1
	} else {
		return 0
	}

	for i := 1; i < length; i++ {
		if nums[i] != point {
			point = nums[i]
			nums[number] = point
			number++
		}
	}

	return number
}
