package leetcode

import (
	"fmt"
	"programs/kit/utils"
	"strconv"
	"strings"
)

// leetcode1154
func DayOfYear(date string) int {
	data := strings.Split(date, "-")
	year, _ := strconv.Atoi(data[0])
	month, _ := strconv.Atoi(data[1])
	day, _ := strconv.Atoi(data[2])
	flag, res := false, 0
	num := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%100 != 0 && year%4 == 0 || year%400 == 0 {
		flag = true
	}
	for i := 0; i < month-1; i++ {
		res += num[i]
	}
	res += day
	if flag && month > 2 {
		res++
	}
	return res
}

// leetcode1155
func NumRollsToTarget(n int, k int, target int) int {
	var h int = 1e9 + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for p := 1; p <= k && p <= j; p++ {
				dp[i][j] = (dp[i][j]%h + dp[i-1][j-p]%h) % h
			}
		}
	}
	return dp[n][target]
}

func NumRollsToTarget2(n int, k int, target int) int {
	// 价值在外面，物品在里面，这个是求排列数
	var h int = 1e9 + 7
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for p := utils.MinNum(k, j); p >= 0; p-- {
				dp[j] = (dp[j]%h + dp[j-p]%h) % h
			}
		}
	}
	return dp[target]
}

// leetcode1173
func Tribonacci(n int) int {
	preOne, preTwo, preThree := 0, 1, 1
	res := 0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	for n-3 >= 0 {
		res = preOne + preTwo + preThree
		preOne = preTwo
		preTwo = preThree
		preThree = res
	}
	return res
}

// leetcode1185
func DayOfTheWeek(day int, month int, year int) string {
	y := year - 1971
	totalDay := 0
	days := map[int]int{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 30}
	for i := 1; i < month; i++ {
		totalDay += days[i]
		if i == 2 && year%400 == 0 || (year%100 != 0 && year%4 == 0) {
			totalDay++
		}
	}
	totalDay += day
	if y == 0 {
		totalDay += 0
	} else if y == 1 {
		totalDay += 365
	} else {
		totalDay += (365 + 366)
		totalDay += ((y-2)/4)*(365*4+1) + (y-2)%4*365
	}
	fmt.Println(totalDay)
	res := totalDay % 7
	mp := map[int]string{0: "Thursday", 1: "Friday", 2: "Saturday", 3: "Sunday", 4: "Monday", 5: "Tuesday", 6: "Wednesday"}
	return mp[res]
}
