package main

import (
	"algorithm/datastructure/tree"
	"fmt"
)

func main() {
	// 二叉搜索树相关测试
	binarySearchTreeTest()
}

func binarySearchTreeTest() {
	var mdArr = []int{15, 6, 23, 4, 7, 19, 71}

	binarySearchTreeContainer := new(tree.BinarySearchTreeContainer)
	for _, val := range mdArr {
		binarySearchTreeContainer.Add(val)
	}

	fmt.Println(fmt.Sprintf("21 是否存在: %t", binarySearchTreeContainer.Exist(21)))
	fmt.Println(fmt.Sprintf("22 是否存在: %t", binarySearchTreeContainer.Exist(22)))

	fmt.Println(fmt.Sprintf("树的最大深度: %d", binarySearchTreeContainer.MaxDepth()))

	// 遍历
	binarySearchTreeContainer.Traverse()
}
