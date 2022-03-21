package leetcode

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
