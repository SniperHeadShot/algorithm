package main

import (
	"container/list"
)

func getMinimumDifference(root *TreeNode) int {
	// 中序遍历获取结果
	stack := list.New()
	node := root
	valSlice := make([]int, 0, 0)

	for stack.Len() > 0 || node != nil {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			element := stack.Back()
			if v, ok := element.Value.(*TreeNode); ok {
				valSlice = append(valSlice, v.Val)
				node = v.Right
			}
			stack.Remove(element)
		}
	}

	result := int(^uint(0) >> 1)
	// 遍历获取差值
	for i := 1; i < len(valSlice); i++ {
		tmp := valSlice[i] - valSlice[i-1]
		if tmp < result {
			result = tmp
		}
	}
	return result
}
