package leetcode

// 第一题：模拟

// 第二题：前缀树
// 给定一组字符串，判断是否存在一个字符串是其他字符串的前缀。
type MyTrieNode struct {
	son    map[byte]*MyTrieNode
	count  int
	prefix int
}

type MyTrie struct {
	root *MyTrieNode
}

func NewRuneTrie() *MyTrie {
	return &MyTrie{root: &MyTrieNode{son: make(map[byte]*MyTrieNode), count: 0, prefix: 0}}
}

func (t *MyTrie) Insert(str string) {
	if t.root == nil || str == "" {
		return
	}
	bytes := []byte(str)
	root := t.root
	for i := 0; i < len(bytes); i++ {
		if _, has := root.son[bytes[i]]; !has {
			// doesn't exist
			root.son[bytes[i]] = &MyTrieNode{}
		}
		root.son[bytes[i]].prefix++
		root = root.son[bytes[i]]
	}
	root.count++
}

func (t *MyTrie) Find(str string) (bool, bool) {
	if t.root == nil || str == "" {
		return false, false
	}
	root := t.root
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if _, has := root.son[bytes[i]]; !has {
			return false, false
		}
		root = root.son[bytes[i]]
	}
	if root.count != 0 {
		return true, true
	}
	return true, false
}

func T2(strs []string) bool {
	runeTrie := NewRuneTrie()
	// 将所有的字符串插入字典树中
	for _, str := range strs {
		runeTrie.Insert(str)
	}
	// 在字典树中进行查找，而且必须isEnd必须得是false状态
	for _, str := range strs {
		exist, isEnd := runeTrie.Find(str)
		if exist && !isEnd {
			return true
		}
	}
	return false
}

// 第三题：背包
// 给定n个任务需要花费的时间和产生的价值，最多只能花费m个时间，且只能完成两个任务，若不足2个任务，则返回0.
// 2 < n <= 1e6, 0 <= m <= 1e6
// 输入：
// 4 6
// 1 8
// 2 1
// 4 3
// 6 4
// 输出
// 11

// 第四题：
// 给出一个数组，最多删除一个连续子数组，求剩下数组的严格递增连续子数组的最大长度。
// n < 1e6.
// 暴力做的n^2，python只过了25%，吐了，题目说的40% n < 1e3。这题怎么做啊，有大佬说下嘛？
// 输入
// 9
// 5 3 4 9 2 8 6 7 1
// 输出
// 4
