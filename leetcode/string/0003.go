package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	strArr := strings.Split(s, "")
	if len(strArr) == 1 {
		return 1
	}

	// 历史最大长度, 当前划窗长度
	var historyMaxLength, currentLength int
	// 上次截断地址
	lastTruncationIndex := -1

	strIndexMap := make(map[string]int)

	for index, val := range strArr {
		strIndex, ok := strIndexMap[val]
		if ok {
			// 重复字符地址是否大于上次截断地址, 大于需要重新计算 currentLength 长度
			if strIndex > lastTruncationIndex {
				if currentLength > historyMaxLength {
					historyMaxLength = currentLength
				}
				currentLength = index - strIndex
				lastTruncationIndex = strIndex
			} else {
				currentLength++
			}
		} else {
			currentLength++
		}
		// 更新字符下标
		strIndexMap[val] = index
	}

	// 没有重复字符串的情况 或 结尾不重复情况
	if currentLength > historyMaxLength {
		historyMaxLength = currentLength
	}

	return historyMaxLength
}
