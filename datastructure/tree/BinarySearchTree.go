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

// 三种遍历方式
// 深度优先DFS
func (self *BinarySearchTreeContainer) TraverseBFS() {
	if self.root == nil {
		return
	}

	var resultArr []int
	stack := new(linkedlist.Stack)
	stack.Push(self.root)

	for !stack.Empty() {
		pop := stack.Pop()
		switch v := pop.(type) {
		case *treeNode:
			var popElement *treeNode
			popElement = v

			resultArr = append(resultArr, popElement.Val)

			if popElement.Right != nil {
				stack.Push(popElement.Right)
			}

			if popElement.Left != nil {
				stack.Push(popElement.Left)
			}
		}
	}

	fmt.Println(fmt.Sprintf("深度遍历DFS: %v", resultArr))
}
