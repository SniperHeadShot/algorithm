package main

import (
	"algorithm/datastructure/linkedlist"
	"fmt"
)

func main() {
	// 栈测试
	stackTest()
}

func stackTest() {
	var mdArr = []int{9, 12, 5, 8, 3, 22}

	stack := new(linkedlist.Stack)

	for _, val := range mdArr {
		stack.Push(val)
	}

	peek := stack.Peek()
	if v, ok := peek.(int); ok {
		fmt.Println(fmt.Sprintf("Peek获得: %v", v))
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
