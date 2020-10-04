package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "str1"
	str2 := "str11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111adaedaedaedae"

	sli := make([]int, 0, 10)
	mp1 := make(map[string]int)
	mp2 := make(map[string]test1)

	mp3 := map[string]string{}

	test1 := test1{
		Id1: 1,
		Id2: 2,
		Id3: 3,
		Id4: 4,
	}

	fmt.Println(unsafe.Sizeof(str1))
	fmt.Println(unsafe.Sizeof(str2))
	fmt.Println(unsafe.Sizeof(sli))
	fmt.Println("====================")
	fmt.Println(unsafe.Sizeof(mp1))
	fmt.Println(unsafe.Sizeof(mp2))
	fmt.Println(unsafe.Sizeof(mp3))
	fmt.Println(unsafe.Sizeof(&test1))
	fmt.Println(unsafe.Sizeof(test1))
	fmt.Println(unsafe.Sizeof(int8(8)))
	fmt.Printf("%d", unsafe.Sizeof(&test1))
}

type test1 struct {
	Id1 int8
	Id2 int16
	Id3 int32
	Id4 int8
	Id5 float64
}

type Test2 struct {
	Id int
}
