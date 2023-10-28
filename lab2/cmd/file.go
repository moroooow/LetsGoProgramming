package main

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func partition(arr []int, low int, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
		}
	}
	tmp := arr[i]
	arr[i] = arr[high]
	arr[high] = tmp
	return arr, i
}

func quickSort(arr []int, low int, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
