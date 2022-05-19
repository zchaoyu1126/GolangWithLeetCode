package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 新建一个链表
	l := list.New()

	// func (l *List) PushFront(v interface{}) *Element
	// 头插法插入一个节点，并返回新生成的节点
	l.PushFront(20)

	// func (l *List) PushBack(v interface{}) *Element
	// 尾插法插入一个节点，并返回新生成的节点
	l.PushBack(35)

	// func (l *List) Front() *Element
	// func (l *List) Back() *Element
	// 输出链表头节点和尾节点的值
	fmt.Printf("head is %d\n", l.Front().Value)
	fmt.Printf("tail is %d\n", l.Back().Value)

	other := list.New()
	other.PushFront(1)
	other.PushFront(2)
	other.PushFront(3)

	// func (l *List) PushFrontList(other *List)
	// 创建链表other的拷贝, 顺序地以头插法将每个节点插入链表l
	l.PushFrontList(other)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()

	// func (l *List) PushBackList(other *List)
	// 创建链表other的拷贝, 顺序地以尾插法将每个节点插入链表l
	l.PushBackList(other)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()

	// func (l *List) InsertBefore(v interface{}, mark *Element) *Element
	// 将一个值为v的新元素插入到mark前面，并返回生成的新元素。
	// 如果mark不是l的元素，l不会被修改。

	// func (l *List) InsertAfter(v interface{}, mark *Element) *Element
	// 将一个值为v的新元素插入到mark后面，并返回新生成的元素。
	// 如果mark不是l的元素，l不会被修改。
	for e := l.Front(); e != nil; e = e.Next() {
		x := e.Value.(int)
		if x == 35 {
			l.InsertBefore(25, e)
		} else if x == 20 {
			l.InsertAfter(29, e)
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()

	// func (l *List) MoveToFront(e *Element)
	// MoveToFront将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。
	// 注：下面的代码这样是可以的，但是如果将val==1的节点移动到头部，那么将会出现死循环
	// 因为两个1都被移动到了头部，所以一直在执行移动操作，永远没法读到链表的尾部。
	for e := l.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		if val == 35 {
			l.MoveToFront(e)
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()

	// func (l *List) MoveToBack(e *Element)
	// MoveToBack将元素e移动到链表的最后一个位置，如果e不是l的元素，l不会被修改。
	// 同上

	// func (l *List) MoveBefore(e, mark *Element)
	// MoveBefore将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

	// func (l *List) MoveAfter(e, mark *Element)
	// MoveAfter将元素e移动到mark的后面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

	// func (l *List) Remove(e *Element) interface{}
	// Remove删除链表中的元素e，并返回e.Value。
	val := l.Remove(l.Back()).(int)
	fmt.Println(val)
	fmt.Println(l.Len())
	l.Init()
	fmt.Println(l.Len())
}

// package main

// import "fmt"

// var n, m int
// var a []int
// var c []int
// var opt, v1, v2 int

// func lowbit(x int) int {
// 	return x & -x
// }

// func add(x, v int) {
// 	for x <= n {
// 		c[x] += v
// 		x += lowbit(x)
// 	}
// }

// func query(x int) int {
// 	res := 0
// 	for x > 0 {
// 		res += c[x]
// 		x -= lowbit(x)
// 	}
// 	return res
// }
// func main() {
// 	fmt.Scan(&n, &m)
// 	a = make([]int, n+1)
// 	c = make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&a[i])
// 		add(i, a[i])
// 	}
// 	for i := 1; i <= m; i++ {
// 		fmt.Scan(&opt, &v1, &v2)
// 		if opt == 1 {
// 			add(v1, v2)
// 		} else {
// 			fmt.Println(query(v2) - query(v1-1))
// 		}
// 	}
// }

// import "fmt"

// type Edge struct {
// 	next int
// 	to   int
// }

// var head []int
// var edge []Edge
// var visited []bool
// var tot int
// var n, m int

// func main() {
// 	fmt.Scan(&n, &m)
// 	head = make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		head[i] = -1
// 	}
// 	edge = make([]Edge, m)
// 	tot = 0
// 	for i := 0; i < m; i++ {
// 		var from, to int
// 		fmt.Scan(&from, &to)
// 		build(from, to)
// 	}
// 	for i := 1; i <= n; i++ {
// 		visited = make([]bool, n+1)
// 		if i != n {
// 			fmt.Printf("%d ", dfs(i))
// 		} else {
// 			fmt.Printf("%d\n", dfs(i))
// 		}
// 	}
// }

// func dfs(x int) int {
// 	visited[x] = true
// 	maxV := x
// 	for i := head[x]; i != -1; i = edge[i].next {
// 		to := edge[i].to
// 		if visited[to] {
// 			continue
// 		}
// 		maxV = max(maxV, dfs(to))
// 	}
// 	return maxV
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func build(from, to int) {
// 	edge[tot].to = to
// 	edge[tot].next = head[from]
// 	head[from] = tot
// 	tot++
// }

// package main

// import "fmt"

// // 使用链式前向星建树
// type Edge struct {
// 	next int
// 	to   int
// }

// var head []int
// var edge []Edge
// var tot int = 0
// var mp map[int][]int
// var visited []bool
// var minTime int
// var n, m int

// func build(from, to int) {
// 	edge[tot].to = to
// 	edge[tot].next = head[from]
// 	head[from] = tot
// 	tot++
// }

// func LightSticks(height int, width int, indices []int) []int {
// 	n = (height + 1) * (width + 1) //端点数目
// 	m = (height+1)*width + height*(width+1)
// 	minTime = m + 5
// 	visited = make([]bool, n)
// 	tot = 0
// 	head = make([]int, n)
// 	for i := range head {
// 		head[i] = -1
// 	}

// 	edge = make([]Edge, 2*(m-len(indices)))
// 	mp = make(map[int][]int)
// 	for i, k := 0, 0; i < m; i++ {
// 		if k < len(indices) {
// 			if i == indices[k] {
// 				k++
// 				continue
// 			}
// 		}
// 		from, to := calculateNode(height, width, i)
// 		build(from, to)
// 		build(to, from)
// 	}
// 	for i := 0; i < n; i++ {
// 		// 枚举BFS的起点
// 		for k := 0; k < n; k++ {
// 			visited[k] = false
// 		}
// 		queue := []int{i}
// 		visited[i] = true
// 		prevCnt, curCnt, totalCnt := 1, 0, 1
// 		depth := 0
// 		for !(prevCnt == 0 && curCnt == 0) {
// 			for prevCnt != 0 {
// 				prevCnt--
// 				top := queue[0]
// 				queue = queue[1:]
// 				for k := head[top]; k != -1; k = edge[k].next {
// 					child := edge[k].to
// 					if visited[child] {
// 						continue
// 					}
// 					visited[child] = true
// 					queue = append(queue, child)
// 					curCnt++
// 					totalCnt++
// 				}
// 			}
// 			prevCnt = curCnt
// 			curCnt = 0
// 			depth++
// 		}

// 		if totalCnt != m-len(indices) {
// 			continue
// 		}
// 		if depth < minTime {
// 			minTime = depth
// 		}
// 		if _, has := mp[depth]; !has {
// 			mp[depth] = make([]int, 0)
// 		}
// 		mp[depth] = append(mp[depth], i)
// 	}

// 	if minTime == m {
// 		return []int{}
// 	}
// 	fmt.Println(minTime)
// 	return mp[minTime]
// }

// func calculateNode(h, w int, idx int) (int, int) {
// 	row := 0
// 	idx = idx + 1
// 	for idx-2*w-1 > 0 {
// 		idx -= (2*w + 1)
// 		row++
// 	}
// 	if idx <= w {
// 		from, to := row*(w+1)+idx-1, row*(w+1)+idx
// 		return from, to
// 	} else {
// 		idx -= w
// 		from, to := row*(w+1)+idx-1, (row+1)*(w+1)+idx-1
// 		return from, to
// 	}
// }

// func main() {
// 	ans := LightSticks(2, 2, []int{2, 5, 6, 7, 8, 10, 11})
// 	fmt.Println(ans)
// }

// package main

// import "fmt"

// // 使用链式前向星建树
// type Edge struct {
// 	next int
// 	to   int
// }

// var head []int
// var edge []Edge
// var tot int = 0
// var mp map[int][]int
// var visited []bool
// var minTime int
// var n, m int
// var cnt int

// func build(from, to int) {
// 	edge[tot].to = to
// 	edge[tot].next = head[from]
// 	head[from] = tot
// 	tot++
// }

// func LightSticks(height int, width int, indices []int) []int {
// 	n = (height + 1) * (width + 1) //端点数目
// 	m = (height+1)*width + height*(width+1)
// 	minTime = m
// 	tot = 0
// 	head = make([]int, n)
// 	visited = make([]bool, n)
// 	for i := range head {
// 		head[i] = -1
// 	}

// 	edge = make([]Edge, 2*(m-len(indices)))
// 	mp = make(map[int][]int)
// 	for i, k := 0, 0; i < m; i++ {
// 		if k < len(indices) {
// 			if i == indices[k] {
// 				k++
// 				continue
// 			}
// 		}
// 		from, to := calculateNode(height, width, i)
// 		build(from, to)
// 		build(to, from)
// 	}
// 	for i := 0; i < n; i++ {
// 		// 从i出发能遍历到的
// 		cnt = 0
// 		for j := 0; j < n; j++ {
// 			visited[j] = false
// 		}
// 		depth := dfs(i, -1)
// 		if cnt-1 < m-len(indices) {
// 			continue
// 		}
// 		if depth != 0 && depth < minTime {
// 			minTime = depth
// 		}
// 		if _, has := mp[depth]; !has {
// 			mp[depth] = make([]int, 0)
// 		}
// 		mp[depth] = append(mp[depth], i)
// 	}
// 	if minTime == m {
// 		return []int{}
// 	}
// 	return mp[minTime]
// }

// func dfs(root, father int) int {
// 	cnt++
// 	visited[root] = true
// 	var depth int
// 	for i := head[root]; i != -1; i = edge[i].next {
// 		child := edge[i].to
// 		// if child == father {
// 		// 	continue // 防止搜回去
// 		// }
// 		// 这个适合无环图
// 		if visited[child] {
// 			continue
// 		}
// 		depth = max(depth, dfs(child, root)+1)
// 	}
// 	return depth
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func calculateNode(h, w int, idx int) (int, int) {
// 	row := 0
// 	idx = idx + 1
// 	for idx-2*w-1 > 0 {
// 		idx -= (2*w + 1)
// 		row++
// 	}
// 	if idx <= w {
// 		from, to := row*(w+1)+idx-1, row*(w+1)+idx
// 		return from, to
// 	} else {
// 		idx -= w
// 		from, to := row*(w+1)+idx-1, (row+1)*(w+1)+idx-1
// 		return from, to
// 	}
// }

// func main() {
// 	ans := LightSticks(1, 2, []int{3})
// 	fmt.Println(ans)
// }

// var n, m int
// var allocation [][]int
// var need [][]int
// var available []int
// var finished []bool
// var ans []int
// var res [][]int

// func initAndRead() {
// 	fmt.Scan(&n, &m)
// 	allocation = make([][]int, n)
// 	for i := range allocation {
// 		allocation[i] = make([]int, m)
// 	}
// 	need = make([][]int, n)
// 	for i := range need {
// 		need[i] = make([]int, m)
// 	}
// 	available = make([]int, m)
// 	finished = make([]bool, n)
// 	ans = make([]int, 0)
// 	res = make([][]int, 0)

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			fmt.Scan(&allocation[i][j])
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			fmt.Scan(&need[i][j])
// 		}
// 	}
// 	for i := 0; i < m; i++ {
// 		fmt.Scan(&available[i])
// 	}
// }

// func main() {
// 	initAndRead()
// 	backtrace(0)
// 	fmt.Println(res)
// }

// func backtrace(cnt int) {
// 	if cnt == n {
// 		tmp := make([]int, n)
// 		copy(tmp, ans)
// 		res = append(res, tmp)
// 		return
// 	}
// 	for i := 0; i < n; i++ {
// 		if needSatisfy(i) && !finished[i] {
// 			ans = append(ans, i)
// 			allocate(i)
// 			backtrace(cnt + 1)
// 			ans = ans[:len(ans)-1]
// 			cancelAllocate(i)
// 		}
// 	}
// }

// func needSatisfy(x int) bool {
// 	for i := 0; i < m; i++ {
// 		if need[x][i] > available[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func allocate(x int) {
// 	for i := 0; i < m; i++ {
// 		available[i] += allocation[x][i]
// 	}
// 	finished[x] = true
// }

// func cancelAllocate(x int) {
// 	for i := 0; i < m; i++ {
// 		available[i] -= allocation[x][i]
// 	}
// 	finished[x] = false
// }

// type LinkedList struct {
// 	Size      int
// 	DummyHead *ListNode
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {
// 	list := NewLinkedList([]int{2, 2, 1, 2})
// 	list.Print()
// 	list.DeleteAtIndex(2)
// 	list.Print()
// 	list.DeleteValue(2)
// 	list.Print()
// }

// func (l *LinkedList) DetectCycle() *ListNode {
// 	if l.DummyHead.Next == nil {
// 		return nil
// 	}

// 	cur := l.DummyHead
// 	slow, fast := cur.Next, cur.Next
// 	for fast != nil {
// 		slow = slow.Next
// 		if fast.Next == nil {
// 			return nil
// 		}
// 		fast = fast.Next.Next

// 		if slow == fast {
// 			p := cur
// 			for p != slow {
// 				p = p.Next
// 				slow = slow.Next
// 			}
// 			return p
// 		}
// 	}
// 	return nil
// }

// func NewLinkedList(nums []int) LinkedList {
// 	dummyHead := &ListNode{}
// 	cur := dummyHead
// 	for i := 0; i < len(nums); i++ {
// 		cur.Next = &ListNode{Val: nums[i]}
// 		cur = cur.Next
// 	}
// 	return LinkedList{Size: len(nums), DummyHead: dummyHead}
// }

// func (l *LinkedList) InsertAtHead(val int) {
// 	l.DummyHead.Next = &ListNode{Val: val, Next: l.DummyHead.Next}
// 	l.Size++
// }

// func (l *LinkedList) InsertAtTail(val int) {
// 	cur := l.DummyHead
// 	for cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = &ListNode{Val: val, Next: nil}
// 	l.Size++
// }

// func (l *LinkedList) InsertAtIndex(val, idx int) {
// 	// 判断idx是否超出范围
// 	if idx <= 0 {
// 		l.InsertAtHead(val)
// 	} else if idx >= l.Size {
// 		l.InsertAtTail(val)
// 	}
// 	cur := l.DummyHead
// 	// idx是2,l-n0-n1-n2
// 	// idx是2需要移动两次
// 	for idx > 0 {
// 		idx--
// 		cur = cur.Next
// 	}
// 	cur.Next = &ListNode{Val: val, Next: cur.Next}
// 	l.Size++
// }

// func (l *LinkedList) DeleteAtIndex(idx int) {
// 	if idx < 0 || idx >= l.Size {
// 		return
// 	}
// 	cur := l.DummyHead
// 	for idx > 0 {
// 		idx--
// 		cur = cur.Next
// 	}
// 	cur.Next = cur.Next.Next
// }

// func (l *LinkedList) DeleteValue(value int) {
// 	cur := l.DummyHead
// 	for cur.Next != nil {
// 		if cur.Next.Val == value {
// 			cur.Next = cur.Next.Next
// 		} else {
// 			cur = cur.Next
// 		}
// 	}
// }

// func (l LinkedList) Len() int {
// 	return l.Size
// }

// func (l LinkedList) ValueAtIndex(idx int) int {
// 	if idx < 0 || idx >= l.Size {
// 		return -1
// 	}
// 	cur := l.DummyHead
// 	for idx >= 0 {
// 		idx--
// 		cur = cur.Next
// 	}
// 	return cur.Val
// }

// func (l LinkedList) Print() {
// 	cur := l.DummyHead
// 	for cur.Next != nil {
// 		fmt.Printf("%d ", cur.Next.Val)
// 		cur = cur.Next
// 	}
// 	fmt.Println()
// }
