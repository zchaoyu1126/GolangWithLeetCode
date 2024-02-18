package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

type MultiListNode struct {
	Val   int
	Prev  *MultiListNode
	Next  *MultiListNode
	Child *MultiListNode
}

type LinkedList struct {
	Size      int
	DummyHead *ListNode
}

func NewLinkedList(arr ...int) LinkedList {
	if len(arr) == 0 {
		return newLinkedListDefalut()
	}
	return newLinkedListWithArr(arr)
}

func newLinkedListDefalut() LinkedList {
	return LinkedList{0, &ListNode{-1, nil}}
}

func newLinkedListWithArr(arr []int) LinkedList {
	dummyHead := &ListNode{}
	cur := dummyHead
	for i := 0; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return LinkedList{len(arr), dummyHead}
}

func (l *LinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val, Next: l.DummyHead.Next}
	l.DummyHead.Next = node
	l.Size++
}

func (l *LinkedList) AddAtTail(val int) {
	cur := l.DummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	l.Size++
}

func (l *LinkedList) AddAtIndex(idx, val int) bool {
	cur := l.DummyHead
	if idx == 0 {
		l.AddAtHead(val)
	} else if idx == l.Size {
		l.AddAtTail(val)
	} else if idx > l.Size || idx < 0 {
		return false
	} else {
		for cur.Next != nil && idx > 0 {
			cur = cur.Next
			idx--
		}
		node := &ListNode{val, cur.Next}
		cur.Next = node
		l.Size++
	}
	return true
}

func (l *LinkedList) Get(idx int) int {
	if idx < 0 || idx >= l.Size {
		return -1
	}
	cur := l.DummyHead
	for cur.Next != nil && idx >= 0 {
		cur = cur.Next
		idx--
	}
	return cur.Val
}

func (l *LinkedList) Len() int {
	cur := l.DummyHead
	res := 0
	for cur.Next != nil {
		cur = cur.Next
		res++
	}
	return res
}

func (l *LinkedList) DeleteAtIndex(idx int) bool {
	if idx < 0 || idx >= l.Size {
		return false
	}
	cur := l.DummyHead
	for cur.Next != nil && idx > 0 {
		cur = cur.Next
		idx--
	}
	cur.Next = cur.Next.Next
	l.Size--
	return true
}

func (l *LinkedList) DeleteVal(val int) {
	cur := l.DummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}

type DLinkedList struct {
	Size      int
	DummyHead *DListNode
	DummyTail *DListNode
}

type DListNode struct {
	Val  int
	Prev *DListNode
	Next *DListNode
}

func NewDLinkedList() DLinkedList {
	node1, node2 := &DListNode{Val: -1}, &DListNode{Val: -1}
	node1.Next = node2
	node2.Prev = node1
	return DLinkedList{0, node1, node2}
}

func (dl *DLinkedList) Get(index int) int {
	if index < 0 || index >= dl.Size {
		return -1
	}
	if index <= dl.Size/2 {
		cur := dl.DummyHead
		for cur.Next != nil && index >= 0 {
			cur = cur.Next
			index--
		}
		return cur.Val
	} else {
		cur := dl.DummyTail
		cnt := dl.Size - index - 1
		for cur.Prev != nil && cnt >= 0 {
			cur = cur.Prev
			cnt--
		}
		return cur.Val
	}
}

func (dl *DLinkedList) AddAtHead(val int) {
	node := &DListNode{Val: val}
	node.Prev = dl.DummyHead
	node.Next = dl.DummyHead.Next
	dl.DummyHead.Next.Prev = node
	dl.DummyHead.Next = node
	dl.Size++
}

func (dl *DLinkedList) AddAtTail(val int) {
	node := &DListNode{Val: val}
	node.Next = dl.DummyTail
	node.Prev = dl.DummyTail.Prev
	dl.DummyTail.Prev.Next = node
	dl.DummyTail.Prev = node
	dl.Size++
}

func (dl *DLinkedList) AddAtIndex(index int, val int) bool {
	if index == 0 {
		dl.AddAtHead(val)
	} else if index == dl.Size {
		dl.AddAtTail(val)
	} else if index > dl.Size || index < 0 {
		return false
	} else {
		if index <= dl.Size/2 {
			cur := dl.DummyHead
			for cur.Next != nil && index > 0 {
				cur = cur.Next
				index--
			}
			node := &DListNode{Val: val}
			node.Prev = cur
			node.Next = cur.Next
			cur.Next.Prev = node
			cur.Next = node
		} else {
			cur := dl.DummyTail
			// 这里的cnt值要特别注意，和别的都不一样
			// 这是因为从后往前插入的时候，是插入在下标cnt的后面，所以要多一次移动
			cnt := dl.Size - index
			for cur.Prev != nil && cnt > 0 {
				cur = cur.Prev
				cnt--
			}
			node := &DListNode{Val: val}
			node.Prev = cur.Prev
			node.Next = cur
			cur.Prev.Next = node
			cur.Prev = node
		}
		dl.Size++
	}
	return true
}

func (dl *DLinkedList) DeleteAtIndex(index int) bool {
	if index < 0 || index >= dl.Size {
		return false
	}
	if index <= dl.Size/2 {
		cur := dl.DummyHead
		for cur.Next != nil && index > 0 {
			cur = cur.Next
			index--
		}
		cur.Next = cur.Next.Next
		cur.Next.Prev = cur
	} else {
		cur := dl.DummyTail
		cnt := dl.Size - index - 1
		for cur.Prev != nil && cnt > 0 {
			cur = cur.Prev
			cnt--
		}
		cur.Prev = cur.Prev.Prev
		cur.Prev.Next = cur
	}
	dl.Size--
	return true
}
