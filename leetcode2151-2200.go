package leetcode

import (
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
	"sort"
	"strconv"
	"strings"
)

// leetcode2164
func SortEvenOdd(nums []int) []int {
	even, odd, res := []int{}, []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	for i, j := 0, len(even)-1; i < j; i, j = i+1, j-1 {
		even[i], even[j] = even[j], even[i]
	}
	for i, j, k := 0, 0, 0; i < len(nums); i++ {
		if i%2 == 0 {
			res = append(res, even[j])
			j++
		} else {
			res = append(res, odd[k])
			k++
		}
	}
	return res
}

// leetcode2165
func SmallestNumber(num int64) int64 {
	arr := []int{}
	origin := num
	cntZero := 0
	if origin < 0 {
		num = -num
	}
	for num != 0 {
		if num%10 == 0 {
			cntZero++
		}
		arr = append(arr, int(num%10))
		num /= 10
	}
	sort.Ints(arr)
	res := 0
	if origin < 0 {
		// big--small
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		for i := 0; i < len(arr); i++ {
			res = res*10 + arr[i]
		}
		res *= -1
	} else {
		if cntZero < len(arr) {
			res = arr[cntZero]
		}
		for i := 0; i < cntZero; i++ {
			res = res * 10
		}
		for i := cntZero + 1; i < len(arr); i++ {
			res = res*10 + arr[i]
		}
	}
	return int64(res)

}

// leetcode2166
// 用下面这种方式失败了，原因在于int64比数据范围小得多
// 要表示的数，其位数为10的5次方个！
// type Bitset struct {
// 	size int
// 	value int
// 	all int
// }

// func NewBitset(size int) Bitset {
// 	all := 1
// 	for i := 0; i < size; i++ {
// 		all *= 2
// 	}
// 	all -= 1
// 	return Bitset{size: size, value: 0, all:all}
// }

// func (bs *Bitset) Fix(idx int)  {
// 	mask := bs.all
// 	flag := 1
// 	for i := 0; i < bs.size-idx-1; i++ {
// 		flag *= 2
// 	}
// 	mask = mask &^ flag
// 	bs.value |= mask
// }

// func (bs *Bitset) Unfix(idx int)  {
// 	mask := bs.all
// 	// 把idx变成0
// 	flag := 1
// 	for i := 0; i < bs.size-idx-1; i++ {
// 		flag *= 2
// 	}
// 	mask -= flag
// 	bs.value &= mask
// }

// func (bs *Bitset) Flip()  {
// 	bs.value = ^bs.value
// }

// func (bs *Bitset) All() bool {
// 	return bs.value == bs.all
// }

// func (bs *Bitset) One() bool {
// 	return bs.value == 0
// }

// func (bs *Bitset) Count() int {
// 	tmp := bs.value
// 	cnt := 0
// 	for tmp != 0 {
// 		if tmp & 1 != 0 {
// 			cnt++
// 		}
// 		tmp >>= 1
// 	}
// 	return cnt
// }

// func (bs *Bitset) ToString() string {
//     tmp := bs.value
//     str := make([]byte, bs.size)
//     for i := 0; i < len(str); i++ {
// 		str[i] = '0'
// 	}
// 	k := bs.size-1
// 	for tmp != 0 {
// 		if tmp & 1 == 1 {
// 			str[k] = '1'
// 		}
// 		k--
// 		tmp >>= 1
// 	}
// 	return string(str)
// }

type Bitset struct {
	arr      []int
	cnt      int
	reversed int
}

func NewBitset(size int) Bitset {
	return Bitset{arr: make([]int, size), cnt: 0}
}

func (bs *Bitset) Fix(idx int) {
	// 不等于1的变成1
	if bs.arr[idx]^bs.reversed != 1 {
		bs.arr[idx] ^= 1
		bs.cnt++
	}
}

func (bs *Bitset) Unfix(idx int) {
	if bs.arr[idx]^bs.reversed != 0 {
		bs.arr[idx] ^= 1
		bs.cnt--
	}
}

func (bs *Bitset) Flip() {
	bs.reversed ^= 1
	bs.cnt = len(bs.arr) - bs.cnt
}

func (bs *Bitset) All() bool {
	return bs.cnt == len(bs.arr)
}

func (bs *Bitset) One() bool {
	return bs.cnt >= 1
}

func (bs *Bitset) Count() int {
	return bs.cnt
}

func (bs *Bitset) ToString() string {
	bytes := []byte{}
	for i := 0; i < len(bs.arr); i++ {
		bytes = append(bytes, byte(bs.arr[i]^bs.reversed)+'0')
	}
	return string(bytes)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
// leetcode2167
func MinimumTime(s string) int {
	minCostLeft := make([]int, len(s)+5)
	minCostRight := make([]int, len(s)+5)
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '0' {
			minCostLeft[i] = minCostLeft[i-1]
		} else {
			minCostLeft[i] = common.SmallerNumber(minCostLeft[i-1]+2, i)
		}
		if s[len(s)-i] == '0' {
			minCostRight[len(s)-i+1] = minCostRight[len(s)-i+2]
		} else {
			minCostRight[len(s)-i+1] = common.SmallerNumber(minCostRight[len(s)-i+2]+2, i)
		}
	}
	res := len(s)
	for i := 0; i < len(s); i++ {
		res = common.SmallerNumber(res, minCostLeft[i]+minCostRight[i+1])
	}
	return res
}

// leetcode2172
func MaximumANDSum(nums []int, numSlots int) int {
	slots := []int{}
	for i := 0; i < numSlots; i++ {
		slots = append(slots, i+1)
		slots = append(slots, i+1)
	}
	scope := make([][]int, numSlots*2)
	for i := range scope {
		scope[i] = make([]int, numSlots*2)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < numSlots*2; j++ {
			scope[i][j] = nums[i] & slots[j]
		}
	}

	km := algorithm.NewKmAlgo(numSlots*2, true, scope)
	return km.KuhnMunkras()
}

// leetcode2185
func PrefixCount(words []string, pref string) int {
	cnt := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			cnt++
		}
	}
	return cnt
}

// leetcode2194
func CellsInRange(s string) []string {
	left, right := s[0], s[3]
	low, high := int(s[1]-'0'), int(s[4]-'0')
	res := []string{}
	for left <= right {
		i := low
		for i <= high {
			res = append(res, string(left)+strconv.Itoa(i))
			i++
		}
		left++
	}
	return res
}

// leetcode2195
func MinimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	arr := []int{0}
	arr = append(arr, nums...)
	var sum int64 = 0
	for i := 1; i <= len(nums); i++ {
		first := arr[i-1] + 1
		last := arr[i] - 1
		if first > last {
			continue
		}
		cnt := last - first + 1
		if k >= cnt {
			sum += int64((last + first) * cnt / 2)
			k -= cnt
		} else {
			sum += int64((first + first + k - 1) * k / 2)
			k -= k
			break
		}
	}

	if k != 0 {
		start := arr[len(nums)] + 1
		sum += int64((start + start + k - 1) * k / 2)
	}
	return sum
}

// leetcode2196
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CreateBinaryTree(descriptions [][]int) *algorithm.TreeNode {
	mp := make(map[int]*algorithm.TreeNode)
	visit := make(map[int]int)
	for _, desc := range descriptions {
		parent, child, isLeft := desc[0], desc[1], desc[2]
		if _, ok := mp[parent]; !ok {
			mp[parent] = &algorithm.TreeNode{Val: parent}
			visit[parent] = 0
		}
		if _, ok := mp[child]; !ok {
			mp[child] = &algorithm.TreeNode{Val: child}
			visit[child] = 0
		}
		parentNode := mp[parent]
		childNode := mp[child]
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
		visit[child]++

	}
	for key, val := range visit {
		//fmt.Println(key, val)
		if val == 0 {
			return mp[key]
		}
	}
	return nil
}

// leetcode2197
func ReplaceNonCoprimes(nums []int) []int {
	res := []int{nums[0]}
	for cur := 1; cur < len(nums); cur++ {
		replace := nums[cur]
		for len(res) > 0 && Gcd(res[len(res)-1], replace) > 1 {
			top := res[len(res)-1]
			res = res[:len(res)-1]
			replace = top * replace / Gcd(top, replace)
		}
		res = append(res, replace)
	}
	return res
}

func Gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	temp := a % b
	for temp != 0 {
		a = b
		b = temp
		temp = a % b
	}
	return b
}

func IsPalindrome2(head *algorithm.ListNode) bool {
	nums := []int{}
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	isParlindrome := func(arr []int) bool {
		for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
			if arr[i] != arr[j] {
				return false
			}
		}
		return true
	}
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			// 把i删除
			newNum1 := append(nums[:i], nums[i+1:]...)
			// 把j删除
			newNum2 := append(nums[:j], nums[j+1:]...)
			return isParlindrome(newNum1) || isParlindrome(newNum2)
		}
	}
	return true
}
