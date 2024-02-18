package leetcode

import (
	"sort"
	"strconv"
)

// leetcode1402
func MaxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	sum := make([]int, len(satisfaction)+1)
	for i := 0; i < len(satisfaction); i++ {
		sum[i+1] = satisfaction[i] + sum[i-1]
	}
	res := 0
	for i := 0; i < len(satisfaction); i++ {
		if satisfaction[i] < sum[len(satisfaction)]-sum[i+1] {
			res += sum[len(satisfaction)] - sum[i]
		}
	}
	return res
}

// leetcode1418
func DisplayTable(orders [][]string) [][]string {
	tableIndex := make(map[string]int)
	tableOrder := []map[string]int{}
	foodmp := map[string]bool{}
	foodlist := []string{}
	tablemp := map[string]bool{}
	tablelist := []int{}
	cnt := 0
	for _, order := range orders {
		food := order[2]
		table := order[1]
		if _, ok := foodmp[food]; !ok {
			foodmp[food] = true
			foodlist = append(foodlist, food)
		}
		if _, ok := tablemp[table]; !ok {
			tablemp[table] = true
			tableNum, _ := strconv.Atoi(table)
			tablelist = append(tablelist, tableNum)
		}
		if _, ok := tableIndex[table]; !ok {
			tableIndex[table] = cnt
			tableOrder = append(tableOrder, map[string]int{})
			cnt++
		}
		index := tableIndex[table]

		if value, ok := tableOrder[index][food]; !ok {
			tableOrder[index][food] = 1
		} else {
			tableOrder[index][food] = value + 1
		}
	}
	sort.Strings(foodlist)
	sort.Ints(tablelist)
	res := make([][]string, len(tablelist)+1)
	for i := 0; i < len(tablelist)+1; i++ {
		res[i] = []string{}
	}
	res[0] = append(res[0], "Table")
	for i := 0; i < len(foodlist); i++ {
		res[0] = append(res[0], foodlist[i])
	}
	for i := 0; i < len(tablelist); i++ {
		tableString := strconv.Itoa(tablelist[i])
		res[i+1] = append(res[i+1], tableString)
		index := tableIndex[tableString]
		for j := 0; j < len(foodlist); j++ {
			res[i+1] = append(res[i+1], strconv.Itoa(tableOrder[index][foodlist[j]]))
		}
	}
	//fmt.Println(tableOrder)
	return res
}

// leetcode1446
func MaxPower(s string) int {
	length := len(s)
	cnt := 1
	max := 0
	for i := 1; i < length; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 1
		}
	}
	if cnt > max {
		max = cnt
	}
	//fmt.Println(max)
	return max
}

// leetcode1450 差分数组
func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	diff := make([]int, 1005)
	sum := make([]int, 1005)
	for i := 0; i < n; i++ {
		start, end := startTime[i], endTime[i]
		diff[start]++
		diff[end+1]--
	}
	for i := 1; i <= 1004; i++ {
		sum[i] = sum[i-1] + diff[i]
	}
	return sum[queryTime]
}
