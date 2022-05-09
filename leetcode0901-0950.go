package leetcode

import "programs/kit/common"

// leetcode904
func TotalFruit1(fruits []int) int {
	mp, res := make(map[int]int), 0
	for l, r := 0, 0; r < len(fruits); r++ {
		if _, ok := mp[fruits[r]]; !ok {
			// fruits[r] 不存在有两种情况
			// 初始化窗口的时候
			if len(mp) < 2 {
				mp[fruits[r]] = r
				if res < r-l+1 {
					res = r - l + 1
				}
			} else {
				// 结算的时候
				if r-l > res {
					res = r - l
				}
				keys, vals := []int{}, []int{}
				for key, val := range mp {
					keys = append(keys, key)
					vals = append(vals, val)
				}
				if vals[0] < vals[1] {
					delete(mp, keys[0])
					l = vals[0] + 1
				} else {
					delete(mp, keys[1])
					l = vals[1] + 1
				}
				mp[fruits[r]] = r
			}
		} else {
			// 更新最后出现的下标位置
			mp[fruits[r]] = r
			if r-l+1 > res {
				res = r - l + 1
			}
		}
	}
	return res
}

func TotalFruit2(fruits []int) int {
	mp, res := make(map[int]int), 0
	for l, r := 0, 0; r < len(fruits); r++ {
		mp[fruits[r]]++
		for len(mp) > 2 {
			mp[fruits[l]]--
			if mp[fruits[l]] == 0 {
				delete(mp, fruits[l])
			}
			l++
		}

		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}

// leetcode908
func SmallestRangeI(nums []int, k int) int {
	var minV int = 1e5
	var maxV int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
		if nums[i] < minV {
			minV = nums[i]
		}
	}
	return common.LargerNumber(0, maxV-minV-2*k)
}

// leetcode933
type RecentCounter struct {
	// 我基本的思路是类似滑动窗口，每次Ping的时候，对当前的窗口进行移动
	// 移动到max(0,t-3000)
	// 这么写会有内存性能的问题吗？
	// 预先分配好3000的容量
	cnt   int
	left  int
	times []int
}

func NewRecentCounter() RecentCounter {
	return RecentCounter{0, 0, make([]int, 0, 3000)}
}

func (r *RecentCounter) Ping(t int) int {
	r.cnt++
	r.times = append(r.times, t)
	v := common.LargerNumber(0, t-3000)
	for i := 0; i < len(r.times); i++ {
		if r.times[i] < v {
			r.cnt--
		} else {
			r.left = i
			break
		}
	}
	r.times = r.times[r.left:]
	return r.cnt
}

// leetcode942
func DiStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	l, r := 0, n

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			// res[i] < res[i+1]
			res[i] = l
			l++
		} else if s[i] == 'D' {
			// res[i] > res[i+1]
			res[i] = r
			r--
		}
	}
	res[n] = l
	return res
}
