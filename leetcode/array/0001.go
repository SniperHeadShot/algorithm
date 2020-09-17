package main

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	numsLen := len(nums)
	// 使用 map 一次遍历即可
	numsMap := make(map[int]int)

	for i := 0; i < numsLen; i++ {
		tempIndex, existFlag := numsMap[target-nums[i]]
		if existFlag {
			return []int{tempIndex, i}
		} else {
			numsMap[nums[i]] = i
		}
	}
	return nil
}
