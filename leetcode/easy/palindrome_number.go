package leetcode_easy

import "strconv"

func IsPalindrome(x int) bool {
	switch {
	case x < 0:
		return false
	case x == 0:
		return true
	default:
		s := strconv.Itoa(x)
		length := len(s)

		for i := 0; i < length; i++ {
			if s[i] != s[length-i-1] {
				return false
			}
		}

		return true
	}
}
