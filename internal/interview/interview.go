package interview

import (
	"fmt"
	"programs/kit/utils"
	"strconv"
	"strings"
)

// 面试题17.11
func FindClosest(words []string, word1 string, word2 string) int {
	var res int = 1e6 + 5
	p, q := -1, -1

	for i, word := range words {
		if word == word1 {
			p = i
		} else if word == word2 {
			q = i
		}
		if p != -1 && q != -1 {
			delta, _ := utils.AbsInt(p - q)
			res = utils.MinNum(res, delta)
		}
	}
	return res
}

func DiscountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, "")
	for i, word := range words {
		n := len(word)
		if word[0] == '$' && word[n-1] != '$' {
			price, _ := strconv.ParseFloat(word[1:], 64)
			words[i] = fmt.Sprintf("$%.2g", price*float64(100-discount)/100)
		}
	}
	var stringBuilder strings.Builder
	for i := range words {
		stringBuilder.Write([]byte(words[i]))
		if i != len(words)-1 {
			stringBuilder.Write([]byte{' '})
		}
	}
	return stringBuilder.String()
}

func TotalSteps(nums []int) int {
	// 求一个谷的长度
	// 10 1 2 3 4 5 6 1 2 3
	// 5,3,4,4,7,3,6,11,8,5,11
	// 山峰

	arr := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 && i+1 < len(nums) && nums[i] >= nums[i+1] {
			arr = append(arr, i)
		} else if i == len(nums)-1 && i-1 >= 0 && nums[i] >= nums[i-1] {
			arr = append(arr, i)
		} else if nums[i] >= nums[i+1] && nums[i] >= nums[i-1] {
			arr = append(arr, i)
		}
	}
	res := 0
	n := len(nums)
	arr = append(arr, n-1)
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > res {
			res = arr[i] - arr[i-1]
		}
	}

	fmt.Println(arr)
	return res
}
