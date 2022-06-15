package leetcode

import (
	"sort"
)

// leetcode2303
func CalculateTax(brackets [][]int, income int) float64 {
	res := 0.0
	prev := brackets[0][0]
	if income >= prev {
		res += float64(brackets[0][1]) * float64(prev) / 100.0
	} else {
		res += float64(income) * float64(brackets[0][1]) / 100.0
		return res
	}

	for i := 1; i < len(brackets); i++ {
		if brackets[i][0] > income {
			res += float64(income-prev) * float64(brackets[i][1]) / 100.0
			break
		}
		res += float64(brackets[i][0]-prev) * float64(brackets[i][1]) / 100.0
		prev = brackets[i][0]
	}
	return res
}

// leetcode2305
var res2305 int

func DistributeCookies(cookies []int, k int) int {
	res2305 = 0
	sort.Sort(sort.Reverse(sort.IntSlice(cookies)))
	num := make([]int, 8)
	num[0] += cookies[0]
	traceback(cookies, 1, num)
	return res2305
}
func traceback(cookies []int, pid int, num []int) {
	// 下面要开始分配第pid包
	if pid >= len(cookies) {
		maxNum := 0
		for i := 0; i < len(num); i++ {
			if num[i] > maxNum {
				maxNum = num[i]
			}
		}
		if maxNum < res2305 {
			res2305 = maxNum
		}
		return
	}
	for j := 0; j < len(num); j++ {
		// 将pid包分配给了j号小朋友
		num[j] += cookies[pid]
		traceback(cookies, pid+1, num)
		// 取消pid包的分配
		num[j] -= cookies[pid]
	}
}

// leetcode2306
// func DistinctNames(ideas []string) int64 {
// 	mp := make(map[string]int32)
// 	cnt := make(map[string]int)

// 	for _, idea := range ideas {
// 		head := idea[0]
// 		suffix := idea[1:]
// 		mp[suffix] |= 1 << int(head-'a')
// 		cnt[suffix]++
// 	}

// 	var res int64
// 	arr1 := make([]int32, 0, len(mp))
// 	arr2 := make([]int, 0, len(mp))
// 	for k, v := range mp {
// 		arr1 = append(arr1, v)
// 		arr2 = append(arr2, cnt[k])
// 	}

// 	for i := 0; i < len(arr1); i++ {
// 		for j := i + 1; j < len(arr1); j++ {
// 			num := bits.OnesCount(uint(arr1[i] & arr1[j]))
// 			cnt1 := arr2[i] - num
// 			cnt2 := arr2[j] - num
// 			res += int64(cnt1 * cnt2)
// 		}
// 	}
// 	return res * 2
// }
func DistinctNames(ideas []string) int64 {
	group := make(map[string]int)

	for _, idea := range ideas {
		head := idea[0]
		suffix := idea[1:]
		group[suffix] |= 1 << int(head-'a')
	}

	var res int
	var cnt [26][26]int
	for _, mask := range group {
		for i := 0; i < 26; i++ {
			if mask>>i&1 == 0 {
				// 包含 i 不包含j的数目
				// 含i 不含j
				for j := 0; j < 26; j++ {
					if mask>>j&1 > 0 {
						cnt[i][j]++
					}
				}
			} else {
				// 不包含i
				for j := 0; j < 26; j++ {
					// 含j不含i
					if mask>>j&1 == 0 {
						res += cnt[i][j]
					}
				}
			}
		}
	}

	return int64(res * 2)
}
