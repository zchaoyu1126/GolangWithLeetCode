package luogu

import "fmt"

func P1428() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}

func P1427() {
	nums := []int{}
	for {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == 0 {
			break
		}
		nums = append(nums, tmp)
	}
	for j := len(nums) - 1; j >= 0; j-- {
		fmt.Printf("%d ", nums[j])
	}
}

func P5727() {
	var num int
	fmt.Scan(&num)
	res := []int{num}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
			res = append(res, num)

		} else {
			num = num*3 + 1
			res = append(res, num)
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d ", res[i])
	}
}

func P1047() {
	var l, m int
	fmt.Scan(&l, &m)
	for i := 0; i < m; i++ {

	}
}

func P5728() {
	var n int
	fmt.Scan(&n)
	scores := make([][]int, n)
	for i := 0; i < n; i++ {
		scores[i] = []int{}
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		scores[i] = append(scores[i], a, b, c)
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	res := 0
	for i := 0; i < n; i++ {
		totali := scores[i][0] + scores[i][1] + scores[i][2]
		for j := i + 1; j < n; j++ {
			totalj := scores[j][0] + scores[j][1] + scores[j][2]
			flag := abs(scores[i][0]-scores[j][0]) <= 5 && abs(scores[i][1]-scores[j][1]) <= 5 && abs(scores[i][2]-scores[j][2]) <= 5
			if flag && abs(totali-totalj) <= 10 {
				res++
			}
		}
	}
	fmt.Println(res)
}

func P1319() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	cur := true // true for zero
	hasPrint := 0
	for cnt != n*n {
		var tmp, ch int
		fmt.Scan(&tmp)
		cnt += tmp
		if cur {
			ch = 0
		} else {
			ch = 1
		}
		for i := 0; i < tmp; i++ {
			if hasPrint < n {
				fmt.Print(ch)
				hasPrint++
			} else {
				fmt.Println()
				fmt.Print(ch)
				hasPrint = 1
			}
		}
		cur = !cur
	}
}
