package test

import (
	leetcode_easy "GoLearning/leetcode/easy"
	"GoLearning/structs"
	"testing"
)

// 測試指令 go test -v .\test\leetcode_easy_test.go
func TestIsValid1(t *testing.T) {
	correct := leetcode_easy.IsValid("()") == true

	if !correct {
		t.Error("fail")
	}
}

func TestIsValid2(t *testing.T) {
	correct := leetcode_easy.IsValid("()[]{}") == true

	if !correct {
		t.Error("fail")
	}
}

func TestIsValid3(t *testing.T) {
	correct := leetcode_easy.IsValid("(]") == false

	if !correct {
		t.Error("fail")
	}
}

func TestIsValid4(t *testing.T) {
	correct := leetcode_easy.IsValid("{[]}") == true

	if !correct {
		t.Error("fail")
	}
}

func TestMergeTwoLists1(t *testing.T) {
	var list1 *structs.ListNode = nil
	var list2 *structs.ListNode = nil
	var answer *structs.ListNode = nil

	list1 = &structs.ListNode{Val: 1}
	list1.Add(2)
	list1.Add(4)

	list2 = &structs.ListNode{Val: 1}
	list2.Add(3)
	list2.Add(4)

	answer = &structs.ListNode{Val: 1}
	answer.Add(1)
	answer.Add(2)
	answer.Add(3)
	answer.Add(4)
	answer.Add(4)

	merged := leetcode_easy.MergeTwoLists(list1, list2)
	correct := merged.IsSame(answer)

	if !correct {
		t.Error("fail")
	}
}

func TestMergeTwoLists2(t *testing.T) {
	var list1 *structs.ListNode = nil
	var list2 *structs.ListNode = nil
	var answer *structs.ListNode = nil

	merged := leetcode_easy.MergeTwoLists(list1, list2)
	correct := merged.IsSame(answer)

	if !correct {
		t.Error("fail")
	}
}

func TestMergeTwoLists3(t *testing.T) {
	var list1 *structs.ListNode = nil
	var list2 *structs.ListNode = nil
	var answer *structs.ListNode = nil

	list2 = &structs.ListNode{Val: 0}
	answer = &structs.ListNode{Val: 0}
	merged := leetcode_easy.MergeTwoLists(list1, list2)
	correct := merged.IsSame(answer)

	if !correct {
		t.Error("fail")
	}
}
