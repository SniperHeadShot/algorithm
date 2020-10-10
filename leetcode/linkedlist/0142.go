package main

func detectCycle(head *ListNode) *ListNode {
	// 快慢指针法
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
