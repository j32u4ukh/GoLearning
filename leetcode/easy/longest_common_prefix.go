package leetcode_easy

func LongestCommonPrefix(strs []string) string {
	n_str := len(strs)

	switch n_str {
	case 0:
		return ""
	case 1:
		return strs[0]
	default:
		min_length := len(strs[0])

		for i := 1; i < n_str; i++ {
			if len(strs[i]) < min_length {
				min_length = len(strs[i])
			}
		}

		var b bool
		var c byte
		var i int

		for i = 0; i < min_length; i++ {
			b = true
			c = strs[0][i]

			for j := 1; j < n_str; j++ {
				if strs[j][i] != c {
					b = false
				}
			}

			if !b {
				break
			}
		}

		return strs[0][0:i]
	}

}
