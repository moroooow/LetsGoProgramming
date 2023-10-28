package main

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.IntSlice.Sort(nums)
	return nums[k]
}
