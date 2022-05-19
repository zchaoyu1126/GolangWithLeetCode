package main

import (
	"fmt"
)

func main() {
	ans := SearchInsert([]int{4, 5, 6}, 7)
	fmt.Println(ans)
}

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
