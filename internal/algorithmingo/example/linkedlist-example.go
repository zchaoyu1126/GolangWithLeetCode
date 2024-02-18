package example

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"strconv"
)

// ********************************************************* \\
// 循环链表——约瑟夫环
func LastRemaining1(n int, m int) int {
	// 如果是1的话，会一直在里面循环
	// 数据量太大的话，会超时
	if m == 1 {
		return n - 1
	}
	head := &algorithm.ListNode{Val: -1, Next: nil}
	cur := head
	for i := 1; i <= n; i++ {
		cur.Next = &algorithm.ListNode{Val: i, Next: nil}
		cur = cur.Next
	}
	cur.Next = head.Next
	cur = head
	for cur.Next != cur {
		cnt := m - 1
		for cnt > 0 {
			cnt--
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
	return cur.Val
}

func LastRemaining2(n int, m int) int {
	// 这样的操作也相当于模拟
	// 当前删除的元素下标是idx，则下一个为(idx+m-1)%n
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	idx := 0
	for n > 1 {
		idx = (idx + m - 1) % n
		arr = append(arr[:idx], arr[idx+1:]...)
		n--
	}
	return arr[0]
}

func LastRemaining3(n int, m int) int {
	// 反推
	// 每一轮都加m并对n取模，得到元素个数为n时的下标值
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}

// ********************************************************* \\
// 一元多项式乘法
// OneLetPoly means one-letter polynomial
type OneLetPoly struct {
	Head *PolyNode
	Size int
}

type PolyInfo struct {
	Coef  int //coefficient
	Power int
}

type PolyNode struct {
	Info PolyInfo
	Next *PolyNode
}

// 输入要求幂是升序的
func NewOneLetPoly(arr [][]int) OneLetPoly {
	if len(arr) == 0 {
		return OneLetPoly{&PolyNode{}, 0}
	}
	head := &PolyNode{}
	cur := head
	for _, val := range arr {
		cur.Next = &PolyNode{PolyInfo{val[0], val[1]}, nil}
		cur = cur.Next
	}

	return OneLetPoly{head, len(arr)}
}

func AddOLP(x, y OneLetPoly) OneLetPoly {
	res := NewOneLetPoly([][]int{})
	cur := res.Head
	cx, cy := x.Head, y.Head
	for cx.Next != nil && cy.Next != nil {
		c1, c2 := cx.Next.Info.Coef, cy.Next.Info.Coef
		p1, p2 := cx.Next.Info.Power, cy.Next.Info.Power
		if p1 == p2 {
			cur.Next = &PolyNode{PolyInfo{c1 + c2, p1}, nil}
			cx = cx.Next
			cy = cy.Next
		} else if p1 < p2 {
			cur.Next = &PolyNode{PolyInfo{c1, p1}, nil}
			cx = cx.Next
		} else {
			cur.Next = &PolyNode{PolyInfo{c2, p2}, nil}
			cy = cy.Next
		}
		cur = cur.Next
	}

	for cx.Next != nil {
		c, p := cx.Next.Info.Coef, cx.Next.Info.Power
		cur.Next = &PolyNode{PolyInfo{c, p}, nil}
		cx = cx.Next
		cur = cur.Next
	}
	for cy.Next != nil {
		c, p := cy.Next.Info.Coef, cy.Next.Info.Power
		cur.Next = &PolyNode{PolyInfo{c, p}, nil}
		cy = cy.Next
		cur = cur.Next
	}
	return res
}

func (p *OneLetPoly) Insert(coef, power int) {
	cur := p.Head
	for cur.Next != nil && cur.Next.Info.Power < power {
		cur = cur.Next
	}
	fmt.Println(coef, power)
	if cur.Next != nil && cur.Next.Info.Power == power {
		cur.Next.Info.Coef += coef
		return
	}
	cur.Next = &PolyNode{PolyInfo{coef, power}, cur.Next}
}

func MulOLP(x, y OneLetPoly) OneLetPoly {
	res := NewOneLetPoly([][]int{})
	cx := x.Head
	for cx.Next != nil {
		cy := y.Head
		c1, p1 := cx.Next.Info.Coef, cx.Next.Info.Power
		for cy.Next != nil {
			c2, p2 := cy.Next.Info.Coef, cy.Next.Info.Power
			res.Insert(c1*c2, p1+p2)
			cy = cy.Next
		}
		cx = cx.Next
	}
	
	return res
}

func (p *OneLetPoly) OutPut() {
	cur := p.Head
	res := ""
	for cur.Next != nil {
		c, p := cur.Next.Info.Coef, cur.Next.Info.Power
		if c == 0 {
			continue
		} else if p == 0 {
			res += strconv.Itoa(c) + "+"
		} else if p == 1 {
			res += strconv.Itoa(c) + "x" + "+"
		} else if p < 0 {
			res += strconv.Itoa(c) + "x^(" + strconv.Itoa(p) + ")+"
		} else {
			res += strconv.Itoa(c) + "x^" + strconv.Itoa(p) + "+"
		}
		cur = cur.Next
	}
	if res[len(res)-1] == '+' {
		res = res[:len(res)-1]
	}
	fmt.Println(res)
}

// package main

// import "programs/algorithmingo"

// func main() {
// 	arr1 := [][]int{{6, -3}, {9, -1}, {2, 4}, {3, 8}, {4, 9}}
// 	arr2 := [][]int{{4, 0}, {3, 2}, {4, 4}, {5, 5}, {9, 8}}
// 	p1 := algorithmingo.NewOneLetPoly(arr1)
// 	p2 := algorithmingo.NewOneLetPoly(arr2)
// 	p1.OutPut()
// 	p2.OutPut()
// 	p3 := algorithmingo.AddOLP(p1, p2)
// 	p4 := algorithmingo.MulOLP(p1, p2)
// 	p3.OutPut()
// 	p4.OutPut()
// }
