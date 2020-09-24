package tree

import (
	"algorithm/datastructure/linkedlist"
	"fmt"
)

// 二叉搜索树 - 树节点
type treeNode struct {
	Left *treeNode

	Right *treeNode

	Val int
}

// 二叉搜索树 - 集合
type BinarySearchTreeContainer struct {
	root *treeNode
}

// 增方法
func (self *BinarySearchTreeContainer) Add(val int) bool {
	if self.root == nil {
		root := new(treeNode)
		root.Val = val
		self.root = root
	}
	return recursionAdd(self.root, val)
}

func recursionAdd(node *treeNode, val int) bool {
	// 当前树节点大于要插入的值
	if node.Val > val {
		if node.Left == nil {
			left := new(treeNode)
			left.Val = val
			node.Left = left
		} else {
			return recursionAdd(node.Left, val)
		}
	}
	// 当前树节点小于要插入的值
	if node.Val < val {
		if node.Right == nil {
			right := new(treeNode)
			right.Val = val
			node.Right = right
		} else {
			return recursionAdd(node.Right, val)
		}
	}
	// 当前树节点等于要插入的值
	return false
}

// 查是否存在
func (self *BinarySearchTreeContainer) Exist(val int) bool {
	return recursionFind(self.root, val)
}

func recursionFind(node *treeNode, val int) bool {
	// 当前树节点大于要插入的值
	if node == nil {
		return false
	}

	// 当前树节点大于要插入的值
	if node.Val > val {
		return recursionFind(node.Left, val)
	}
	// 当前树节点小于要插入的值
	if node.Val < val {
		return recursionFind(node.Right, val)
	}
	// 当前树节点等于要插入的值
	return true
}

// 获取深度
func (self *BinarySearchTreeContainer) MaxDepth() int {
	if self.root == nil {
		return 0
	}
	return recursionCalculateDepth(self.root)
}

func recursionCalculateDepth(node *treeNode) int {
	if node == nil {
		return 0
	}

	return maxInt(recursionCalculateDepth(node.Left), recursionCalculateDepth(node.Right)) + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 遍历
func (self *BinarySearchTreeContainer) Traverse() {
	fmt.Print("深度优先遍历 - 前序遍历 - 递归 :")
	self.dfsPreOrderRecursionTraversal(self.root)
	fmt.Println()

	fmt.Print("深度优先遍历 - 前序遍历 - 非递归 :")
	self.dfsPreOrderNotRecursionTraversal()
	fmt.Println()

	fmt.Print("深度优先遍历 - 中序遍历 - 递归 :")
	self.dfsMidOrderRecursionTraversal(self.root)
	fmt.Println()

	fmt.Print("深度优先遍历 - 中序遍历 - 非递归 :")
	self.dfsMidOrderNotRecursionTraversal()
	fmt.Println()

	fmt.Print("深度优先遍历 - 后序遍历 - 递归 :")
	self.dfsPostOrderRecursionTraversal(self.root)
	fmt.Println()

	fmt.Print("深度优先遍历 - 后序遍历 - 非递归 :")
	self.dfsPostOrderNotRecursionTraversal_1()
	fmt.Println()

	fmt.Print("深度优先遍历 - 后序遍历 - 非递归 - 双栈法 :")
	self.dfsPostOrderNotRecursionTraversal_2()
	fmt.Println()

	fmt.Print("广度优先遍历 - 递归 :")
	nodes := []*treeNode{}
	nodes = append(nodes, self.root)
	self.bfsRecursionTraversal(&nodes)
	fmt.Println()
}

// 深度优先遍历 - 前序遍历 - 递归
func (self *BinarySearchTreeContainer) dfsPreOrderRecursionTraversal(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Print(" ", node.Val)
	self.dfsPreOrderRecursionTraversal(node.Left)
	self.dfsPreOrderRecursionTraversal(node.Right)
}

// 深度优先遍历 - 中序遍历 - 递归
func (self *BinarySearchTreeContainer) dfsMidOrderRecursionTraversal(node *treeNode) {
	if node == nil {
		return
	}
	self.dfsMidOrderRecursionTraversal(node.Left)
	fmt.Print(" ", node.Val)
	self.dfsMidOrderRecursionTraversal(node.Right)
}

// 深度优先遍历 - 后序遍历 - 递归
func (self *BinarySearchTreeContainer) dfsPostOrderRecursionTraversal(node *treeNode) {
	if node == nil {
		return
	}
	self.dfsPostOrderRecursionTraversal(node.Left)
	self.dfsPostOrderRecursionTraversal(node.Right)
	fmt.Print(" ", node.Val)
}

// 广度优先遍历 - 递归
func (self *BinarySearchTreeContainer) bfsRecursionTraversal(nodes *[]*treeNode) {
	if nodes == nil || len(*nodes) == 0 {
		return
	}

	nextLevelNode := []*treeNode{}
	for _, node := range *nodes {
		fmt.Print(" ", node.Val)
		if node.Left != nil {
			nextLevelNode = append(nextLevelNode, node.Left)
		}
		if node.Right != nil {
			nextLevelNode = append(nextLevelNode, node.Right)
		}
	}

	self.bfsRecursionTraversal(&nextLevelNode)
}

// 深度优先遍历 - 前序遍历 - 非递归
func (self *BinarySearchTreeContainer) dfsPreOrderNotRecursionTraversal() {
	node := self.root
	stack := new(linkedlist.Stack)

	for !stack.Empty() || node != nil {
		if node != nil {
			fmt.Print(" ", node.Val)
			stack.Push(node)
			node = node.Left
		} else {
			pop := stack.Pop()
			switch v := pop.(type) {
			case *treeNode:
				node = v
				node = node.Right
			}
		}
	}
}

// 深度优先遍历 - 中序遍历 - 非递归
func (self *BinarySearchTreeContainer) dfsMidOrderNotRecursionTraversal() {
	node := self.root
	stack := new(linkedlist.Stack)

	for !stack.Empty() || node != nil {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			pop := stack.Pop()
			switch v := pop.(type) {
			case *treeNode:
				node = v
				fmt.Print(" ", node.Val)
				node = node.Right
			}
		}
	}
}

// 深度优先遍历 - 后序遍历 - 非递归 - 传统
func (self *BinarySearchTreeContainer) dfsPostOrderNotRecursionTraversal_1() {
	stack := linkedlist.Stack{}
	mp := make(map[int]int)
	node := self.root

	for !stack.Empty() || node != nil {
		if node != nil {
			stack.Push(node)
			mp[node.Val] = 1
			node = node.Left
		} else {
			peek := stack.Peek()
			if v, ok := peek.(*treeNode); ok {
				if mp[v.Val] == 2 {
					stack.Pop()
					fmt.Printf("%d ", v.Val)
					node = nil
				} else {
					mp[v.Val] = 2
					node = v.Right
				}
			}
		}
	}
}

// 深度优先遍历 - 后序遍历 - 非递归 - 双栈法
func (self *BinarySearchTreeContainer) dfsPostOrderNotRecursionTraversal_2() {
	// 收集节点个给 stack2
	stack1 := linkedlist.Stack{}
	// 真正存储后序遍历节点的栈
	stack2 := linkedlist.Stack{}

	node := self.root
	if node == nil {
		return
	} else {
		stack1.Push(node)
	}

	for !stack1.Empty() {
		pop := stack1.Pop()
		stack2.Push(pop)
		if v, ok := pop.(*treeNode); ok {
			if v.Left != nil {
				stack1.Push(v.Left)
			}
			if v.Right != nil {
				stack1.Push(v.Right)
			}
		}
	}

	for !stack2.Empty() {
		pop := stack2.Pop()
		if v, ok := pop.(*treeNode); ok {
			fmt.Print(" ", v.Val)
		}
	}
}
