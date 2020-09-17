package main

func main() {

}

func massage(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[numsLen-1]
	}
	if numsLen == 2 {
		return getMaxInt(nums[numsLen-1], nums[numsLen-2])
	}

	dp := make([]int, numsLen)
	dp[0] = nums[0]
	dp[1] = getMaxInt(nums[0], nums[1])

	for i := 2; i < numsLen; i++ {
		a := dp[i-2] + nums[i]
		b := dp[i-1]
		dp[i] = getMaxInt(a, b)
	}

	return dp[numsLen-1]
}

func getMaxInt(i1 int, i2 int) int {
	if i1 >= i2 {
		return i1
	} else {
		return i2
	}
}
