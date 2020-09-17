package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 4, 1, 7, 8, 3}
	// 递归方式
	fmt.Print(rec_opt(&arr, 6))

	// 遍历方式
	fmt.Print(findMax(&arr))
}

func findMax(arr *[]int) int {
	// 状态转移方程式 f(n) = max(arr[n]+f(n-2), f(n-1))
	arrLen := len(*arr)
	dp := make([]int, arrLen)

	dp[0] = (*arr)[0]
	dp[1] = getMaxInt((*arr)[0], (*arr)[1])

	for i := 2; i < arrLen; i++ {
		// 选当前
		a := (*arr)[i] + dp[i-2]
		// 不选当前
		b := dp[i-1]
		dp[i] = getMaxInt(a, b)
	}
	return dp[arrLen-1]
}

func rec_opt(arr *[]int, index int) int {
	if index == 0 {
		return (*arr)[0]
	}
	if index == 1 {
		return getMaxInt((*arr)[0], (*arr)[1])
	}
	// 选当前
	a := (*arr)[index] + rec_opt(arr, index-2)
	// 不选当前
	b := rec_opt(arr, index-1)
	return getMaxInt(a, b)
}

func getMaxInt(i1 int, i2 int) int {
	if i1 >= i2 {
		return i1
	} else {
		return i2
	}
}
