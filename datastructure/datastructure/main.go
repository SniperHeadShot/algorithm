package main

import (
	"algorithm/datastructure/linkedlist"
	"algorithm/datastructure/tree"
	"fmt"
)

func main() {
	// 栈测试
	//stackTest()

	// 二叉搜索树相关测试
	binarySearchTreeTest()
}

var mdArr = []int{9, 12, 5, 8, 3, 22}

func stackTest() {
	stack := new(linkedlist.Stack)

	for _, val := range mdArr {
		stack.Push(val)
	}

	for !stack.Empty() {
		pop := stack.Pop()
		switch v := pop.(type) {
		case int:
			var popElement int
			popElement = v
			fmt.Println(fmt.Sprintf("弹栈获得: %v", popElement))
		}
	}
}

func binarySearchTreeTest() {
	binarySearchTreeContainer := new(tree.BinarySearchTreeContainer)
	for _, val := range mdArr {
		binarySearchTreeContainer.Add(val)
	}

	fmt.Println(fmt.Sprintf("21 是否存在: %t", binarySearchTreeContainer.Exist(21)))
	fmt.Println(fmt.Sprintf("22 是否存在: %t", binarySearchTreeContainer.Exist(22)))

	fmt.Println(fmt.Sprintf("树的最大深度: %d", binarySearchTreeContainer.MaxDepth()))

	binarySearchTreeContainer.TraverseBFS()
}
