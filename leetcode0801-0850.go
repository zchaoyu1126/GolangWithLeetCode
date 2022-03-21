package leetcode

import "sort"

// leetcode807
func MaxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// row[i]
			if grid[i][j] > rowMax[i] {
				rowMax[i] = grid[i][j]
			}
			// col[j]
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += minNumber(rowMax[i], colMax[j]) - grid[i][j]
		}
	}
	return res
}
func minNumber(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode844
func BackspaceCompare(s string, t string) bool {
	bytes1, bytes2 := []byte(s), []byte(t)
	slow1, fast1 := 0, 0
	for ; fast1 < len(bytes1); fast1++ {
		if bytes1[fast1] == '#' {
			if slow1 >= 1 {
				slow1--
			}
		} else {
			bytes1[slow1] = bytes1[fast1]
			slow1++
		}
	}

	slow2, fast2 := 0, 0
	for ; fast2 < len(bytes2); fast2++ {
		if bytes2[fast2] == '#' {
			if slow2 >= 1 {
				slow2--
			}
		} else {
			bytes2[slow2] = bytes2[fast2]
			slow2++
		}
	}
	return string(bytes1[:slow1]) == string(bytes2[:slow2])
}

// leetcode846
func IsNStraightHand(hand []int, groupSize int) bool {
	num := len(hand)
	if num%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	mp := make(map[int]int)
	arr := []int{}
	for i := 0; i < num; i++ {
		mp[hand[i]]++
	}
	for i := 0; i < num; i++ {
		if len(arr) == 0 {
			arr = append(arr, hand[i])
		} else if hand[i] == arr[len(arr)-1] {
			continue
		} else {
			arr = append(arr, hand[i])
		}
	}
	for i := 0; i < num/groupSize; i++ {
		if len(arr) < groupSize {
			return false
		}
		prev := arr[0] - 1
		for j := 0; j < groupSize; j++ {
			if arr[j] != prev+1 && mp[arr[j]] == 0 {
				return false
			}
			mp[arr[j]]--
			prev = arr[j]
		}
		for len(arr) > 0 && mp[arr[0]] == 0 {
			arr = arr[1:]
		}
	}
	return true
}
