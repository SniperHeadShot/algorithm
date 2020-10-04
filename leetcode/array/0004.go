package main

func main() {

}

// TODO 未完成
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var calcMid = func(nums []int) float64 {
		len := len(nums)
		if len == 1 {
			return float64(nums[0])
		}

		if len%2 == 0 {
			mid := len / 2
			return (float64(nums[mid]) + float64(nums[mid-1])) / 2
		} else {
			return float64(nums[len/2])
		}
	}

	if len(nums1) == 0 {
		if len(nums2) == 0 {
			return 0
		} else {
			return calcMid(nums2)
		}
	} else {
		if len(nums2) == 0 {
			return calcMid(nums1)
		} else {
			return (calcMid(nums1) + calcMid(nums2)) / 2
		}
	}
}
