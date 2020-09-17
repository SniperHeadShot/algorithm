package main

import "math"

// 动态规划：状态转移方程 f(n) = f(n-1) > 0 ? arr[n] + f(n-1) : arr[n] 这里为什么要和0比，我有很多问号
func maxSubArray(nums []int) int {
	numsLen := len(nums)

	if numsLen == 0 {
		return 0
	}

	dp := make([]int, numsLen)

	dp[0] = nums[0]
	for i := 1; i < numsLen; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}

	return findMax(&dp)
}

func findMax(nums *[]int) int {
	max := math.MinInt32
	for _, val := range *nums {
		if val > max {
			max = val
		}
	}
	return max
}
