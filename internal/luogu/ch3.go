package luogu

import (
	"fmt"
	"math"
	"math/big"
	"programs/internal/algorithmingo/algorithm"
)

func P51718() {
	var n int
	fmt.Scanf("%d\n", &n)
	var tmp, res int
	res = 1000 + 5
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		if tmp < res {
			res = tmp
		}
	}
	fmt.Println(res)
}

func P5719() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	sum1, sum2, cnt1, cnt2 := 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			sum1 += i
			cnt1++
		} else {
			sum2 += i
			cnt2++
		}
	}
	fmt.Printf("%.1f %.1f", float64(sum1)/float64(cnt1), float64(sum2)/float64(cnt2))
}

func P5720() {
	var length int = 0
	var res int = 0
	fmt.Scanf("%d\n", &length)
	for length != 0 {
		res++
		length >>= 1
	}
	fmt.Println(res)
}

func P5721() {
	var n int = 0
	fmt.Scanf("%d", &n)
	cur := 1
	for i := n; i >= 1; i-- {
		// n rows
		for j := i; j >= 1; j-- {
			// n cols
			fmt.Printf("%02d", cur)
			cur++
		}
		fmt.Println()
	}
}

func P1009() {
	var n int
	fmt.Scanf("%d", &n)
	res := big.NewInt(0)
	for i := 1; i <= n; i++ {
		res = res.Add(res, factorial(i))
	}
	fmt.Printf("%d\n", res)
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// // 这份代码对x为0时，根本无法处理
// func P1980() {
// 	mp := map[int]int{10: 1, 100: 20, 1000: 300, 10000: 4000, 100000: 50000, 1000000: 600000, 10000000: 7000000}
// 	var n, x int
// 	fmt.Scanf("%d %d\n", &n, &x)
// 	placeCnt, tmp := 0, n
// 	for tmp > 0 {
// 		tmp /= 10
// 		placeCnt++
// 	}

// 	cnt := 0
// 	division := powInt(10, placeCnt-1)
// 	for n > 0 {
// 		if x != 0 {
// 			if n/division > x {
// 				cnt += mp[division]*(n/division) + division
// 				fmt.Println(cnt, 1)
// 			} else if n/division == x {
// 				cnt += mp[division]*(n/division) + n%division + 1
// 				fmt.Println(cnt, 2)
// 			} else {
// 				cnt += mp[division] * (n / division)
// 				fmt.Println(cnt, 3)
// 			}
// 		}
// 		n %= division
// 		division /= 10

// 	}
// 	fmt.Println(cnt)
// }

// func powInt(x, num int) int {
// 	res := 1
// 	for i := 0; i < num; i++ {
// 		res *= x
// 	}
// 	return res
// }

// // better one
func P1980() {
	var n, x int
	fmt.Scanf("%d %d\n", &n, &x)
	m, ans := 1, 0
	for m <= n {
		a, b, c := n/(m*10), n/m%10, n%m
		if x != 0 {
			if b > x {
				ans += (a + 1) * m
			} else if b == x {
				ans += a*m + c + 1
			} else {
				ans += a * m
			}
		} else {
			if b != 0 {
				ans += a * m
			} else {
				ans += (a-1)*m + c + 1
			}
		}
		m *= 10
	}
	fmt.Println(ans)
}

func P1035() {
	var k int
	fmt.Scanf("%d\n", &k)
	ans := 0.0
	cur := 1.0
	for ans <= float64(k) {
		ans += 1.0 / cur
		cur += 1.0
	}
	fmt.Println(int(cur - 1.0))
}

func P2669() {
	var k int
	fmt.Scanf("%d\n", &k)
	l, r := 1, int(math.Ceil(math.Sqrt(float64(k)*2)))

	for l < r {
		m := (l + r + 1) / 2
		if m*(m+1)/2 > k {
			r = m - 1
		} else {
			l = m
		}
	}
	res := 0
	for i := 1; i <= l; i++ {
		res += i * i
	}
	res += (k - (l*l+l)/2) * (l + 1)
	fmt.Println(res)
}

func P5722() {
	var n int
	fmt.Scanf("%d\n", &n)
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(res)
}

func P5723() {
	primes := algorithm.EulerMakePrimeList(100000)
	var l int
	fmt.Scanf("%d\n", &l)
	if l <= 1 {
		fmt.Println(0)
		return
	}
	cnt, res := 0, 0
	for res+primes[cnt] <= l {
		res += primes[cnt]
		fmt.Println(primes[cnt])
		cnt++
	}
	fmt.Println(cnt)
}

func generatePalindromes(cnt int) []int {
	res := []int{}
	if cnt == 1 {
		res = append(res, 5, 7)
	}
	if cnt == 2 {
		for i := 1; i <= 9; i += 2 {
			res = append(res, i*10+i)
		}
	}
	if cnt == 3 {
		for d1 := 1; d1 <= 9; d1++ {
			for d2 := 0; d2 <= 9; d2++ {
				res = append(res, d1*100+d2*10+d1)
			}
		}
	}
	if cnt == 4 {
		for d1 := 1; d1 <= 9; d1 += 2 {
			for d2 := 0; d2 <= 9; d2++ {
				res = append(res, d1*1000+d2*100+d2*10+d1)
			}
		}
	}
	if cnt == 5 {
		for d1 := 1; d1 <= 9; d1 += 2 {
			for d2 := 0; d2 <= 9; d2++ {
				for d3 := 0; d3 <= 9; d3++ {
					res = append(res, d1*10000+d2*1000+d3*100+d2*10+d1)
				}
			}
		}
	}
	if cnt == 6 {
		for d1 := 1; d1 <= 9; d1 += 2 {
			for d2 := 0; d2 <= 9; d2++ {
				for d3 := 0; d3 <= 9; d3++ {
					res = append(res, d1*100000+d2*10000+d3*1000+d3*100+d2*10+d1)
				}
			}
		}
	}
	if cnt == 7 {
		for d1 := 1; d1 <= 9; d1 += 2 {
			for d2 := 0; d2 <= 9; d2++ {
				for d3 := 0; d3 <= 9; d3++ {
					for d4 := 0; d4 <= 9; d4++ {
						res = append(res, d1*1000000+d2*100000+d3*10000+d4*1000+d3*100+d2*10+d1)
					}
				}
			}
		}
	}
	if cnt == 8 {
		for d1 := 1; d1 <= 9; d1 += 2 {
			for d2 := 0; d2 <= 9; d2++ {
				for d3 := 0; d3 <= 9; d3++ {
					for d4 := 0; d4 <= 9; d4++ {
						res = append(res, d1*10000000+d2*1000000+d3*100000+d4*10000+d4*1000+d3*100+d2*10+d1)
					}
				}
			}
		}
	}
	return res
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func P1217() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	intPlaceCnt := func(x int) int {
		cnt := 0
		for x > 0 {
			cnt++
			x >>= 1
		}
		return cnt
	}
	cnt := intPlaceCnt(b)
	for i := 1; i <= cnt; i++ {
		nums := generatePalindromes(i)
		for j := 0; j < len(nums); j++ {
			if a <= nums[j] && nums[j] <= b && isPrime(nums[j]) {
				fmt.Println(nums[j])
			}
		}
	}
}

func P1423() {
	var x float64
	fmt.Scanf("%f", &x)
	sum, cur, cnt := 2.0, 2.0, 1
	for sum < x {
		cur *= 0.98
		sum += cur
		cnt++
	}
	fmt.Println(cnt)
}

func P1307() {
	var x int
	fmt.Scanf("%d", &x)
	if x == 0 {
		fmt.Println(0)
		return
	}
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	flag := false
	if x < 0 {
		flag = true
	}
	x = abs(x)
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	if flag {
		fmt.Println(-res)
	} else {
		fmt.Println(res)
	}
}

func P1720() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 2 {
		fmt.Println(1.00)
	}
	prev1, prev2 := 1.00, 1.00
	res := 0.0
	for i := 3; i <= n; i++ {
		res = prev1 + prev2
		prev1 = prev2
		prev2 = res
	}
	fmt.Printf("%.2f", res)
}

func P5724() {
	var cnt, tmp int
	max, min := 0, 1000
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&tmp)
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max - min)
}

func P1420() {
	var n, prev, cur, cnt, max int
	fmt.Scanln(&n)
	fmt.Scan(&prev)
	cnt, max = 1, 1
	for i := 1; i < n; i++ {
		fmt.Scan(&cur)
		if cur == prev+1 {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 1
		}
		prev = cur
	}
	fmt.Println(max)
}

func P1075() {
	var x int
	fmt.Scan(&x)
	for i := 2; i <= x; i++ {
		if x%i == 0 && isPrime(x/i) {
			fmt.Println(x / i)
			return
		}
	}
}

func P5725() {
	var x int
	fmt.Scan(&x)
	cur := 1
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%02d", cur)
			cur++
		}
		fmt.Println()
	}

	fmt.Println()
	cur = 1
	length := x * 2
	for i := 0; i < x; i++ {
		for j := 0; j < length-(i+1)*2; j++ {
			fmt.Print(" ")
		}
		for j := x - i - 1; j < x; j++ {
			fmt.Printf("%02d", cur)
			cur++
		}
		fmt.Println()
	}
}

func P5726() {
	var num, tmp int
	var max, min, average int
	max, min = 0, 10
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fmt.Scan(&tmp)
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
		average += tmp
	}
	average -= max
	average -= min
	fmt.Printf("%.2f", float64(average)/float64(num-2))
}

func P4956() {
	var n int
	fmt.Scanf("%d", &n)
	n /= 52
	n /= 7
	x := n - 3
	for x > 100 {
		x -= 3
	}
	k := (n - x) / 3
	fmt.Println(x)
	fmt.Println(k)
}

func P1089() {
	var budget [12]int
	for i := 0; i < 12; i++ {
		fmt.Scanln(&budget[i])
	}
	remain, restore := 0, 0
	for i := 0; i < 12; i++ {
		if remain+300 < budget[i] {
			fmt.Printf("-%d", i+1)
			return
		} else {
			remain = remain + 300 - budget[i]
			if remain >= 100 {
				restore += remain / 100 * 100
				remain -= remain / 100 * 100
			}
		}
	}
	fmt.Println(float64(restore)*1.2 + float64(remain))
}
