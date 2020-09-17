package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	// 进位
	var carry int

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{carry % 10, nil}
		tail = tail.Next
		carry /= 10
	}

	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head.Next
}
