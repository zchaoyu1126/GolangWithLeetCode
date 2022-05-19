package luogu

import (
	"fmt"
	"math"
)

func P5735() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scanln(&x3, &y3)
	res := 0.0
	distance := func(x1, y1, x2, y2 float64) float64 {
		return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	}
	res += distance(x1, y1, x2, y2)
	res += distance(x1, y1, x3, y3)
	res += distance(x2, y2, x3, y3)
	fmt.Printf("%.2f", res)
}

func P5736() {
	var n, tmp int
	isPrime := func(x int) bool {
		if x <= 1 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if isPrime(tmp) {
			fmt.Printf("%d ", tmp)
		}
	}
}

func P5737() {
	var a, b int
	fmt.Scan(&a, &b)
	cnt := 0
	res := []int{}
	for i := a; i <= b; i++ {
		if i%400 == 0 {
			cnt++
			res = append(res, i)
		} else if i%4 == 0 && i%100 != 0 {
			cnt++
			res = append(res, i)
		}
	}
	fmt.Println(cnt)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
}
