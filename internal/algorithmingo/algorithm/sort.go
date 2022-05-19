package algorithm

import (
	"math/rand"
)

func QuickSort(nums []int, start, end int) {
	if end > start {
		pivot := Partition(nums, start, end)
		QuickSort(nums, start, pivot-1)
		QuickSort(nums, pivot+1, end)
	}
}

func Partition(nums []int, start, end int) int {
	randPos := rand.Intn(end-start+1) + start
	nums[randPos], nums[end] = nums[end], nums[randPos]
	slow := start - 1
	for fast := start; fast < end; fast++ {
		if nums[fast] < nums[end] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
	slow++
	nums[slow], nums[end] = nums[end], nums[slow]
	return slow
}

func MergeSort(nums []int) {
	length := len(nums)
	src := nums
	dst := make([]int, len(nums))
	for seg := 1; seg < length; seg += seg {
		for start := 0; start < length; start += seg * 2 {
			mid := min(start+seg, length)
			end := min(start+seg*2, length)
			i, j, k := start, mid, start
			for i < mid && j < end {
				if src[i] < src[j] {
					dst[k], i, k = src[i], i+1, k+1
				} else {
					dst[k], j, k = src[j], j+1, k+1
				}
			}
			for i < mid {
				dst[k], i, k = src[i], i+1, k+1
			}
			for j < end {
				dst[k], j, k = src[j], j+1, k+1
			}
		}
		copy(src, dst)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
