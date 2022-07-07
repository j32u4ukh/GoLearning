package leetcode_easy

func IsValid(s string) bool {
	length := len(s)

	if length == 0 {
		return true
	}

	if length&0x1 == 1 {
		// fmt.Printf("總個數為奇數個, length: %d\n", length)
		return false
	}

	var right int

	switch s[0] {
	case '(':
		right = findRightParentheses('(', s, ')')
	case '[':
		right = findRightParentheses('[', s, ']')
	case '{':
		right = findRightParentheses('{', s, '}')
	default:
		// fmt.Println("No left parentheses")
		return false
	}

	// fmt.Printf("left: %v, right index: %d\n", string(s[0]), right)

	if right == -1 {
		// fmt.Println("找不到右括弧")
		return false
	}

	// 第一組括弧間沒有其他括弧
	if right == 1 {
		if length == 2 {
			//cout << "true: 第一組括弧間沒有其他括弧，且長度為 2" << endl;
			return true
		}

		// 計算右括弧右方剩餘的括弧個數，若為奇數個，則無法組成有效的括弧
		n_right := length - right - 1

		if n_right&0x1 == 1 {
			// fmt.Printf("右方剩餘的括弧個數為奇數個, length: %d, right: %d, n_right: %d\n", length, right, n_right)
			return false
		}

		return true && IsValid(s[right+1:length])

	} else {
		// 第一組括弧間包含了其他括弧
		// right 為偶數時，表示左右括弧所夾括弧個數為奇數個，無法組成有效的括弧
		if right&0x1 == 0 {
			// fmt.Println("第一組括弧間包含了其他括弧，且右方剩餘的括弧個數為奇數個")
			return false
		}

		// 計算右括弧右方剩餘的括弧個數，若為奇數個，則無法組成有效的括弧
		n_right := length - right - 1

		if n_right&0x1 == 1 {
			// fmt.Println("第一組括弧間包含了其他括弧，且左右括弧所夾括弧個數為奇數個")
			return false
		}

		if n_right == 0 {
			return true && IsValid(s[1:right])
		} else {
			return true && IsValid(s[1:right]) && IsValid(s[right+1:length])
		}
	}
}

func findRightParentheses(left byte, s string, right byte) int {
	var i, length, count int = 0, len(s), 0

	for ; i < length; i++ {
		switch {
		case s[i] == left:
			count--
		case s[i] == right:
			count++
		}

		if count == 0 {
			return i
		}
	}

	return -1
}
