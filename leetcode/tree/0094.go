package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	emptyArr := make([]int, 0, 0)
	collect := valCollect{
		Val: &emptyArr,
	}

	recursionMidTraversal(root, &collect)
	return *collect.Val
}

func recursionMidTraversal(node *TreeNode, collect *valCollect) {
	if node == nil {
		return
	}
	recursionMidTraversal(node.Left, collect)
	collect.put(node.Val)
	recursionMidTraversal(node.Right, collect)
}

type valCollect struct {
	Val *[]int
}

func (self *valCollect) put(val int) {
	newVal := append(*self.Val, val)
	self.Val = &newVal
}
