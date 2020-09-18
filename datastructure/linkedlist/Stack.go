package linkedlist

// 基于链表的栈实现
type Stack struct {
	head *linkedListNode

	tail *linkedListNode
}

func (self *Stack) Empty() bool {
	return self.head == nil
}

// 获取栈顶元素
func (self *Stack) Pop() interface{} {
	var result *linkedListNode
	if self.head == self.tail {
		result = self.head
		self.head = nil
	} else {
		result = self.head
		// 头节点可能为空
		if self.head != nil {
			self.head = result.next
		}
	}
	if result == nil {
		return nil
	}
	return result.Val
}

// 向栈中添加元素
func (self *Stack) Push(val interface{}) {
	node := new(linkedListNode)
	node.Val = val

	if self.head == nil {
		self.head = node
		self.tail = node
	} else {
		self.tail.next = node
		self.tail = node
	}
}
