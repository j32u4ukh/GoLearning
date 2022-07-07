package leetcode_easy

import "GoLearning/structs"

func MergeTwoLists(list1 *structs.ListNode, list2 *structs.ListNode) *structs.ListNode {
	var merge_list, node, node1, node2 *structs.ListNode

	if list1 == nil || list2 == nil {
		if list1 == nil && list2 == nil {
			return nil
		} else if list2 == nil {
			return list1
		} else {
			return list2
		}
	}

	if list1.Val <= list2.Val {
		node = &structs.ListNode{Val: list1.Val}
		node1 = list1.Next
		node2 = list2
	} else {
		node = &structs.ListNode{Val: list2.Val}
		node1 = list1
		node2 = list2.Next
	}

	merge_list = node

	for (node1 != nil) || (node2 != nil) {
		if (node1 != nil) && (node2 != nil) {
			// node1, node2 都不為空
			if node1.Val <= node2.Val {
				node.Next = &structs.ListNode{Val: node1.Val}
				node1 = node1.Next
			} else {
				node.Next = &structs.ListNode{Val: node2.Val}
				node2 = node2.Next
			}
		} else {
			// node1, node2 其中一個為空

			if node2 == nil {
				// node1 不為空
				node.Next = &structs.ListNode{Val: node1.Val}
				node1 = node1.Next
			} else {
				// node2 不為空
				node.Next = &structs.ListNode{Val: node2.Val}
				node2 = node2.Next
			}
		}

		node = node.Next
	}

	return merge_list
}
