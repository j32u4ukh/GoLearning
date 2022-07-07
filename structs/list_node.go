package structs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (self *ListNode) Add(val int) {
	curr := self
	temp := curr.Next

	for temp != nil {
		curr = temp
		temp = temp.Next
	}

	curr.Next = &ListNode{
		Val: val,
	}
}

func (self *ListNode) GetDepth() int {
	depth := 0

	if self != nil {
		depth++
		next := self.Next

		for next != nil {
			depth++
			next = next.Next
		}
	}

	return depth
}

func (self ListNode) String() string {
	s := fmt.Sprintf("ListNode(%d", self.Val)
	next := self.Next

	for next != nil {
		s += fmt.Sprintf(" -> %d", next.Val)
		next = next.Next
	}

	s += ")"

	return s
}

func (self *ListNode) IsSame(other *ListNode) bool {
	s := self
	o := other

	for s != nil && o != nil {
		if s.Val != o.Val {
			return false
		}

		s = s.Next
		o = o.Next
	}

	return s == nil && o == nil
}
