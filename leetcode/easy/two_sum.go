package leetcode_easy

func TwoSum(nums []int, target int) []int {
	length := len(nums)
	var i, j int

	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
