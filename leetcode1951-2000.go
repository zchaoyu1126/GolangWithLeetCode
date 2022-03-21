package leetcode

// leetcode1995
func CountQuadruplets(nums []int) int {
	cnt := len(nums)
	res := 0
	ans := []int{}
	for i := 0; i < cnt; i++ {
		ans = append(ans, nums[i])
		for j := i + 1; j < cnt; j++ {
			ans = append(ans, nums[j])
			for k := j + 1; k < cnt; k++ {
				ans = append(ans, nums[k])
				sum := ans[0] + ans[1] + ans[2]
				for p := k + 1; p < cnt; p++ {
					if sum == nums[p] {
						res++
					}
				}
				ans = ans[:2]
			}
			ans = ans[:1]
		}
		ans = []int{}
	}
	return res
}

func CountQuadruplets2(nums []int) int {
	mp := make(map[int]int)
	cnt := len(nums)
	res := 0
	for c := cnt - 2; c >= 2; c-- {
		mp[nums[c+1]]++
		for a := 0; a < c; a++ {
			for b := a + 1; b < c; b++ {
				sum := nums[a] + nums[b] + nums[c]
				res += mp[sum]
			}
		}
	}
	return res
}

func CountQuadruplets3(nums []int) int {
	mp := make(map[int]int)
	cnt, res := len(nums), 0
	for b := cnt - 3; b >= 1; b-- {
		for d := b + 2; d < cnt; d++ {
			mp[nums[d]-nums[b+1]]++
		}
		for a := 0; a < b; a++ {
			res += mp[nums[b]+nums[a]]
		}
	}
	return res
}
