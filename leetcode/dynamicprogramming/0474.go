package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "10001"
	i1, i2 := getZeroAndOneCount(&str)
	fmt.Printf("%d  %d", i1, i2)
}

// TODO unfinished
func findMaxForm(strs []string, m int, n int) int {
	strslen := len(strs)
	zeroArr := make([]int, strslen)
	oneArr := make([]int, strslen)

	for index, val := range strs {
		zeroCount, oneCount := getZeroAndOneCount(&val)
		zeroArr[index] = zeroCount
		oneArr[index] = oneCount
	}

	// 状态转移方程  f(n) = max(f(n), f(n-1))
	dp := make([]int, strslen)

	if m-zeroArr[0] > 0 && n-oneArr[0] > 0 {
		dp[0] = 1
	}
	return 0
}

func getZeroAndOneCount(str *string) (int, int) {
	if str == nil || *str == "" {
		return 0, 0
	}
	var zeroCount, oneCount int
	splitArr := strings.Split(*str, "")
	for _, val := range splitArr {
		if val == "0" {
			zeroCount++
		} else if val == "1" {
			oneCount++
		}
	}
	return zeroCount, oneCount
}
