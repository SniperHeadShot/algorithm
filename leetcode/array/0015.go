package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sum := threeSum(nums)
	fmt.Println(len(sum))
}

// TODO 未完成
func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0, len(nums))

	mp := make(map[string]int8)

	duplicate := func(intSlice []int) bool {
		sort.Ints(intSlice)

		var key string
		for _, v := range intSlice {
			key += fmt.Sprintf("%d", v)
		}

		if _, ok := mp[key]; ok {
			return false
		} else {
			mp[key] = 1
			return true
		}
	}

	// 暴力遍历
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					intSlice := []int{nums[i], nums[j], nums[k]}
					if duplicate(intSlice) {
						result = append(result, intSlice)
					}
				}
			}
		}
	}
	return result
}
