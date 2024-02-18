package leetcode

import "sort"

// leetcode1576
func ModifyString(s string) string {
	length := len(s)
	data := []byte(s)
	for i := 0; i < length; i++ {
		if data[i] == '?' {
			data[i] = MinValue(i-1, i+1, data)
		}
	}
	return string(data)
}

func MinValue(i, j int, data []byte) byte {
	// 只有一个?的情况
	if i < 0 && j >= len(data) {
		return 'a'
	}

	if i < 0 { // 第一个
		if data[j] == '?' {
			return 'a'
		} else {
			if data[j]-1 < 'a' {
				return 'z'
			} else {
				return data[j] - 1
			}
		}
	} else if j >= len(data) { // 最后一个
		if data[i]-1 < 'a' {
			return 'z'
		} else {
			return data[i] - 1
		}
	} else {
		// 在中间的，前面肯定不会是?，后面有可能出现?
		if data[j] == '?' {
			if data[i]-1 < 'a' {
				return 'z'
			} else {
				return data[i] - 1
			}
		}
		// 如果都不是，那就取中间的那个
		if data[i] == data[j] || data[i]+1 == data[j] || data[i] == data[j]+1 {
			if (data[i]+data[j])/2+2 > 'z' {
				return 'a'
			} else {
				return (data[i]+data[j])/2 + 2
			}
		} else {
			return (data[i] + data[j]) / 2
		}
	}
}

// leetcode1577
func NumTriplets(nums1 []int, nums2 []int) int {
	res := 0

	mp1 := map[int]int{}
	mp2 := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			mul := nums1[i] * nums1[j]
			mp1[mul]++
		}
	}

	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			mul := nums2[i] * nums2[j]
			mp2[mul]++
		}
	}

	for i := 0; i < len(nums1); i++ {
		p := nums1[i] * nums1[i]
		res += mp2[p]
	}

	for i := 0; i < len(nums2); i++ {
		p := nums2[i] * nums2[i]
		res += mp1[p]
	}

	return res
}

// leetcode1589
func MaxSumRangeQuery(nums []int, requests [][]int) int {
	diff := make([]int, len(nums)+1)
	// 由于题目的特殊性，省去了构建差分数组的步骤
	// diff[0] = nums[0]
	// diff[i] = nums[i] - nums[i-1]

	// 区间查询
	for _, request := range requests {
		l, r := request[0], request[1]
		diff[l] += 1
		diff[r+1] -= 1
	}

	// 求前缀和
	sum := make([]int, len(nums)+2)
	// sum[i]表示的是前面i-1项所有的和
	// i~j项的和表示为， sum[j+1]-sum[i]
	for i := 1; i <= len(diff); i++ {
		sum[i] = sum[i-1] + diff[i]
	}
	// sum[1]代表下标0出现了几次
	// sum[len(nums)] 代表下标len(nums)-1出现了几次
	// 去掉首尾
	sum = sum[1 : len(sum)-1]
	sort.Ints(nums)
	sort.Ints(sum)
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i] * sum[i]
	}
	return ans
}
