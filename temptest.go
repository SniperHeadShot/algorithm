package main

import "fmt"

func main() {
	var t1, t2 interface{}
	t1 = test1{
		Id: 1,
	}
	t2 = Test2{
		Id: 2,
	}

	t1, ok1 := t1.(test1)
	fmt.Println("t1.t", t1, "t1.ok", ok1)
	t2, ok2 := t2.(Test2)
	fmt.Println("t2.t", t2, "t2.ok", ok2)
}

type test1 struct {
	Id int
}

type Test2 struct {
	Id int
}
