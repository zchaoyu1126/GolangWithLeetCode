package luogu

import (
	"fmt"
	"math"
)

func P1000() {
	var str string = `                ********
               ************
               ####....#.
             #..###.....##....
             ###.......######              ###            ###
                ...........               #...#          #...#
               ##*#######                 #.#.#          #.#.#
            ####*******######             #.#.#          #.#.#
           ...#***.****.*###....          #...#          #...#
           ....**********##.....           ###            ###
           ....****    *****....
             ####        ####
           ######        ######
##############################################################
#...#......#.##...#......#.##...#......#.##------------------#
###########################################------------------#
#..#....#....##..#....#....##..#....#....#####################
##########################################    #----------#
#.....#......##.....#......##.....#......#    #----------#
##########################################    #----------#
#.#..#....#..##.#..#....#..##.#..#....#..#    #----------#
##########################################    ############
`
	fmt.Print(str)
}

func P1001() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}

func P1421() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println((a*10 + b) / 19)
}

func P1425() {
	var start1, start2 int
	var end1, end2 int
	fmt.Scan(&start1, &start2, &end1, &end2)
	sum := end1*60 + end2 - (start1*60 + start2)
	fmt.Println(sum/60, sum%60)
}

func P2181() {
	var n uint64
	fmt.Scan(&n)
	if n <= 3 {
		fmt.Println(0)
		return
	}
	res := n * (n - 1) / 2 * (n - 2) / 3 * (n - 3) / 4
	fmt.Printf("%d", res)
}

func P2433() {
	var T int
	fmt.Scan(&T)
	switch T {
	case 1:
		fmt.Println("I love Luogu!")
	case 2:
		fmt.Println("6 4")
	case 3:
		fmt.Println("3")
		fmt.Println("12")
		fmt.Println("2")
	case 4:
		fmt.Printf("%.6g", float64(500)/float64(3))
	case 5:
		fmt.Println("15")
	case 6:
		fmt.Printf("%.6g", math.Sqrt(float64(36+81)))
	case 7:
		fmt.Println("110")
		fmt.Println("90")
		fmt.Println("0")
	case 8:
		fmt.Printf("%.6g\n", 3.141593*10)
		fmt.Printf("%.6g\n", 3.141593*25)
		fmt.Printf("%.6g\n", 3.141593*125*4/3)
	case 9:
		fmt.Println("22")
	case 10:
		fmt.Println(9)
	case 11:
		fmt.Printf("%.6g", float64(100)/float64(3))
	case 12:
		fmt.Println(int('M'-'A') + 1)
		fmt.Printf("%c", 'A'+17)
	case 13:
		sum := 3.141593 * (64 + 1000) * 4 / 3
		fmt.Printf("%d", int(math.Pow(sum, 1.0/3)))
	case 14:
		fmt.Println(50)
	default:
		fmt.Println("error")
	}
}

func P3954() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d", int(a*0.2+b*0.3+c*0.5))
}

func P5703() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a * b)
}

func P5704() {
	var ch byte
	fmt.Scanf("%c\n", &ch)
	fmt.Printf("%c\n", ch-32)
}

func P5705() {
	var num string
	fmt.Scanf("%s", &num)
	bytes := []byte(num)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	fmt.Println(string(bytes))
}

func P5706() {
	var volume float64
	var cnt int
	fmt.Scan(&volume, &cnt)
	fmt.Printf("%.3f\n", volume/float64(cnt))
	fmt.Println(cnt * 2)
}

func P5707() {
	var s, v int
	fmt.Scan(&s, &v)
	costTime := int(math.Ceil(float64(s)/float64(v))) + 10
	h, m := costTime/60, costTime%60
	sh, sm := 0, 0
	if h == 8 && m == 0 {
		fmt.Println("00:00")
		return
	} else if h >= 8 {
		sh = 23 - h + 8
		sm = 60 - m

	} else {
		sh = 7 - h
		sm = 60 - m
	}
	fmt.Printf("%02d:%02d\n", sh, sm)
}

func P5708() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	var p float64 = (a + b + c) / 2
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("%.1g", area)
}

func P5709() {
	var m, t, s int
	fmt.Scanln(&m, &t, &s)
	if t == 0 {
		fmt.Println(0)
		return
	} else {
		canEatNum := int(math.Ceil(float64(s) / float64(t)))
		if canEatNum > m {
			fmt.Println(0)
		} else {
			fmt.Println(m - canEatNum)
		}
	}
}
