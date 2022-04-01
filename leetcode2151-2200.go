package leetcode

import (
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
	"sort"
	"strconv"
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

func isPalindrome(head *ListNode) bool {
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
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			newNum1 := append(tmp[:i], tmp[i+1:]...)
			// 把j删除
			copy(tmp, nums)
			newNum2 := append(tmp[:j], tmp[j+1:]...)
			return isParlindrome(newNum1) || isParlindrome(newNum2)
		}
	}
	return true
}

// 需要记录用户参加过哪些活动，参加了多少次
// 需要计算，参加哪个活动，优惠力度最大
type DiscountSystem struct {
	Acts    map[int]DiscoutActivity
	Records map[int]ConsumerRecords
}

type DiscoutActivity struct {
	priceLimit  int
	discount    int
	number      int
	userLimit   int
	hasFinished bool
}

type ConsumerRecords struct {
	cnt map[int]int
}

func NewDiscountSystem() DiscountSystem {
	return DiscountSystem{}
}

func (ds *DiscountSystem) AddActivity(actId int, priceLimit int, discount int, number int, userLimit int) {
	ds.Acts[actId] = DiscoutActivity{priceLimit, discount, number, userLimit, false}
}

func (ds *DiscountSystem) RemoveActivity(actId int) {
	act := ds.Acts[actId]
	act.hasFinished = true
}

func (ds *DiscountSystem) Consume(userId int, cost int) int {
	record := ds.Records[userId]
	curId, curCost := -1, -1
	for id, act := range ds.Acts {
		if !act.hasFinished && cost > act.priceLimit && act.userLimit > 0 && record.cnt[id] < act.userLimit {
			// 活动未结束 满足消费下限 活动人数未满 消费次数未满
			if act.discount > curCost {
				curId = id
			} else if act.discount == curCost {
				if curId > id {
					curId = id
				}
			}
		}
	}
	act := ds.Acts[curId]
	act.userLimit--
	record.cnt[curId]++
	return curCost
}

// 输入：
// ["DiscountSystem","addActivity","addActivity","consume","removeActivity","consume","consume","consume","consume"]
// [[],[1,10,6,3,2],[2,15,8,8,2],[101,13],[2],[101,17],[101,11],[102,16],[102,21]]
// 输出：
// [null,null,null,7,null,17,11,10,21]
// 预期：
// [null,null,null,7,null,11,11,10,21]

// 输入：
// ["DiscountSystem","addActivity","consume","consume","consume","addActivity","addActivity","consume","addActivity","consume"]
// [[],[3,55,20,4,4],[6,98],[2,37],[6,55],[8,45,12,5,4],[1,55,16,4,4],[6,100],[0,40,13,7,2],[10,43]]
// 输出：
// [null,null,78,37,55,null,null,80,null,30]
// 预期：
// [null,null,78,37,35,null,null,80,null,30]

func MaxInvestment(product []int, limit int) int {
	sort.Ints(product)
	for i, j := 0, len(product)-1; i < j; i, j = i+1, j-1 {
		product[i], product[j] = product[j], product[i]
	}
	curN, sum := 1, 0
	var m int = 1e9 + 7
	for i := 0; i < len(product)-1; i++ {
		cnt := product[i] - product[i+1]
		if cnt*curN <= limit {
			sum = (sum%m + (product[i]+product[i+1])*cnt/2*curN%m) % m
			limit -= cnt
			curN++
		} else {
			//如果cnt*curN > limit，那么说明到头了要
			// limit/curN 均摊到每个curN
			tmp := int(limit / curN)
			sum = (sum + (product[i]+product[i]-tmp+1)*tmp/2*curN) % m
			remain := limit - tmp
			sum = (sum + (product[i]-tmp)*remain) % m
		}
	}
	return sum
}

// [2,1,5,8,7]
// 10
// 输出：
// 55
// 预期：
// 57
// 标准输出：
// 10
// 8
// 9
// 32
// 5
// 47 1
// 55
// 0

func FindDifference(nums1 []int, nums2 []int) [][]int {
	mp1 := make(map[int]bool)
	mp2 := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		mp1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		mp2[nums2[i]] = true
	}
	ans1 := []int{}
	ans2 := []int{}
	for i := 0; i < len(nums1); i++ {
		if _, has := mp2[nums1[i]]; !has {
			ans1 = append(ans1, nums1[i])
			mp2[nums1[i]] = true
		}
	}
	for i := 0; i < len(nums2); i++ {
		if _, has := mp1[nums2[i]]; !has {
			ans2 = append(ans2, nums2[i])
			mp1[nums2[i]] = true
		}
	}
	return [][]int{ans1, ans2}
}

func MinDeletion(nums []int) int {
	// 4
	// 0 1 2 3
	// nums[i] != nums[i+1]
	return 2
}

func KthPalindrome(queries []int, intLength int) []int64 {
	var ans []int64 = []int64{}
	var build func(n, length int) int64

	powerForTen := func(cnt int) int {
		res := 1
		for i := 1; i <= cnt; i++ {
			res *= 10
		}
		return res
	}

	build = func(n, length int) int64 {
		cnt := int((length+1)/2) - 1
		if n > 9*powerForTen(cnt) {
			return -1
		}

		base := powerForTen(cnt) + n - 1
		copyBase := base
		firstFlag := true
		for copyBase != 0 {
			remain := copyBase % 10
			copyBase /= 10
			if firstFlag && length%2 != 0 {
				firstFlag = false
				continue
			} else {
				base = base*10 + remain
			}
		}
		return int64(base)
	}

	for i := 0; i < len(queries); i++ {
		ans = append(ans, build(queries[i], intLength))
	}
	return ans
}
