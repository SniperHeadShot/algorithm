package main

import (
	"container/list"
	"fmt"
)

func main() {
	list1 := list.New()

	var str string
	str = fmt.Sprintf("%p", &list1)

	fmt.Println(str)
}
