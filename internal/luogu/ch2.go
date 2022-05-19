package luogu

import "fmt"

func P1046() {
	var res, reachHeigth int
	var heights [10]int

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &heights[i])
	}
	fmt.Scanf("\n%d", &reachHeigth)
	for i := 0; i < 10; i++ {
		if heights[i]-30 <= reachHeigth {
			res++
		}
	}
	fmt.Println(res)
}

func P1055() {
	var str string
	fmt.Scan(&str)
	checkSum, cur := 0, 1
	for i := 0; i <= 10; i++ {
		if str[i] != '-' {
			checkSum += int(str[i]-'0') * cur
			cur++
		}
	}
	code := byte('0' + checkSum%11)
	if checkSum%11 == 10 {
		code = 'X'
	}
	if str[len(str)-1] == code {
		fmt.Println("Right")
	} else {
		bytes := []byte(str)
		bytes[len(bytes)-1] = code
		fmt.Println(string(bytes))
	}
}

func P1422() {
	var total int
	fmt.Scan(&total)
	if total >= 401 {
		fmt.Printf("%.1f", 150*0.4463+250*0.4663+float64(total-400)*0.5663)
	} else if total < 400 && total >= 151 {
		fmt.Printf("%.1f", (150*0.4463 + float64(total-150)*0.4663))
	} else {
		fmt.Printf("%.1f", float64(total)*0.4463)
	}
}

func P1424() {
	var x, n int
	fmt.Scanf("%d %d\n", &x, &n)
	sum := (n / 7) * 250 * 5
	remain := n % 7
	for i := 0; i < remain; i++ {
		if (x+i)%7 == 0 || (x+i)%7 == 6 {
			continue
		} else {
			sum += 250
		}
	}
	fmt.Println(sum)
}

func P1805() {
	var school, afterSchool int
	maxV, day := -1, 0
	for i := 0; i < 7; i++ {
		fmt.Scanf("%d %d\n", &school, &afterSchool)
		if school+afterSchool > 8 && school+afterSchool > maxV {
			maxV = school + afterSchool
			day = i + 1
		}
	}
	fmt.Println(day)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func P1888() {
	var a, b, c int
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}
	x := gcd(a, c)
	fmt.Printf("%d/%d", c/x, a/x)
}

func P1909() {
	var n int
	fmt.Scanf("%d\n", &n)
	min := 0xFFFFFFFF
	for i := 0; i < 3; i++ {
		var num, cnt, price int
		fmt.Scanf("%d %d\n", &cnt, &price)
		if n%cnt != 0 {
			num = n/cnt + 1
		} else {
			num = n / cnt
		}
		//fmt.Println(n, num, cnt, price)
		if price*num < min {
			min = num * price
		}
	}
	fmt.Println(min)
}

func P4414() {
	var a, b, c int
	var str string
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	mp := make(map[byte]int)
	mp['A'], mp['B'], mp['C'] = a, b, c
	fmt.Scanf("\n%s", &str)
	fmt.Printf("%d %d %d\n", mp[str[0]], mp[str[1]], mp[str[2]])

}

func P5710() {
	var x int
	fmt.Scan(&x)
	res := 0
	if x%2 == 0 {
		res += 1
	}
	if x <= 12 && x > 4 {
		res += 1
	}
	var a, b, c, d int
	if res == 2 {
		a, b = 1, 1
	} else if res >= 1 {
		b, c = 1, 1
	} else if res == 0 {
		d = 1
	}
	fmt.Println(a, b, c, d)
}

func P5711() {
	var year int
	fmt.Scan(&year)

	if year%100 == 0 && year%400 == 0 {
		fmt.Println(1)
		return
	} else if year%100 != 0 && year%4 == 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(0)
}

func P5712() {
	var cnt int
	fmt.Scan(&cnt)
	if cnt <= 1 {
		fmt.Printf("Today, I ate %d apple.\n", cnt)
	} else {
		fmt.Printf("Today, I ate %d apples.\n", cnt)
	}
}

func P5713() {
	var cnt int
	fmt.Scan(&cnt)
	if cnt*5 < cnt*3+11 {
		fmt.Println("Local")
	} else {
		fmt.Println("Luogu")
	}
}

func P5714() {
	var m, h float64
	fmt.Scanln(&m, &h)
	bmi := m / h / h
	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi >= 18.5 && bmi < 24 {
		fmt.Println("Normal")
	} else {
		fmt.Printf("%.6g\n", bmi)
		fmt.Println("Overweight")
	}
}

func P5715() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Println(a, b, c)
}

func P5716() {
	var year, month int
	fmt.Scan(&year, &month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 2:
		if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	case 4, 6, 9, 11:
		fmt.Println(30)
	}
}

func P5717() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Not triangle")
		return
	}
	if a*a+b*b == c*c {
		fmt.Println("Right triangle")
	} else if a*a+b*b > c*c {
		fmt.Println("Acute triangle")
	} else {
		fmt.Println("Obtuse triangle")
	}
	if b == c || a == b {
		fmt.Println("Isosceles triangle")
	}
	if a == b && a == c {
		fmt.Println("Equilateral triangle")
	}
}
