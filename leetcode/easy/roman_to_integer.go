package leetcode_easy

import "strings"

func RomanToInt(s string) int {
	roman_len := 7
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	romans := []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}

	length := len(s)

	switch {
	case length == 0:
		return 0
	case length == 1:
		return m[s[0]]
	default:
		value := 0
		t := 0

		for i := 0; i < roman_len; i++ {
			c := romans[i]
			t = strings.Index(s, string(c))

			if t != -1 {
				value = m[c]
				break
			}
		}

		left := s[0:t]
		right := s[t+1 : length]
		return value - RomanToInt(left) + RomanToInt(right)
	}
}
