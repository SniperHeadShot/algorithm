package linkedlist

// 双向链表节点
type linkedListNode struct {
	Val interface{}

	prev *linkedListNode

	next *linkedListNode
}
