package algorithm

const INT_MAX = 1061109567
const INT_MIN = -1061109567

// 匈牙利算法， 解决二部图最大权匹配问题 最优匹配
type KmAlgo struct {
	n         int
	weight    [][]int // 权重
	lx        []int   // 项标
	ly        []int
	visx      []bool // 是否访问过，同时也用于标记交错树，标记哪些点属于S和T
	visy      []bool
	match     []int // match[i]记录y[i]与x[match[i]]相对应
	maxWeight bool
}

func NewKmAlgo(n int, maxWeight bool, weight [][]int) KmAlgo {
	ins := KmAlgo{n: n, weight: weight, maxWeight: maxWeight}
	ins.lx = make([]int, ins.n)
	ins.ly = make([]int, ins.n)
	ins.visx = make([]bool, ins.n)
	ins.visy = make([]bool, ins.n)
	ins.match = make([]int, ins.n)

	// 若是求最小匹配，则要将边权取反
	if !maxWeight {
		for i := 0; i < ins.n; i++ {
			for j := 0; j < ins.n; j++ {
				ins.weight[i][j] = -ins.weight[i][j]
			}
		}
	}

	// 初始化顶标
	// Cx[i]设置为max(weight[i][j] | j=0,..,n-1 ), Cy[i]=0;
	// Cy的顶标都是0
	for i := 0; i < ins.n; i++ {
		ins.lx[i] = INT_MIN
		for j := 0; j < ins.n; j++ {
			if ins.lx[i] < ins.weight[i][j] {
				ins.lx[i] = ins.weight[i][j]
			}
		}
	}

	for i := 0; i < ins.n; i++ {
		ins.match[i] = -1
	}

	return ins
}

func (k *KmAlgo) SearchPath(u int) bool {
	// 给x[u]找匹配
	k.visx[u] = true
	for v := 0; v < k.n; v++ {
		if !k.visy[v] && k.lx[u]+k.ly[v] == k.weight[u][v] {
			k.visy[v] = true
			if k.match[v] == -1 || k.SearchPath(k.match[v]) {
				// 若y[v]未被占用， 或者y[v]还能够找到其他可搭配的x点
				// 当发生第二种情况时，由于k.visy[v] = true这句话，已经当前的v排除在外
				// 也就是找k.match[v] 除了v之外有无其他可搭配的点
				k.match[v] = u
				return true
			}
		}
	}
	return false
}

func (k *KmAlgo) KuhnMunkras() int {
	//不断修改顶标，直到找到完备匹配或完美匹配
	for u := 0; u < k.n; u++ {
		//为x里的每个点找匹配
		for {
			for index := range k.visx {
				k.visx[index] = false
			}
			for index := range k.visy {
				k.visy[index] = false
			}

			if k.SearchPath(u) { //x[u]在相等子图找到了匹配,继续为下一个点找匹配
				break
			}

			//若是在相等子图里没有找到匹配，就修改顶标，直到找到匹配为止
			//首先找到修改顶标时的增量inc, min(lx[i] + ly [i] - weight[i][j],inc); 其中lx[i]为搜索过的点，ly[i]是未搜索过的点
			//由于如今是要给u找匹配，因此只须要修改找的过程当中搜索过的点，增长有可能对u有帮助的边
			inc := INT_MAX
			for i := 0; i < k.n; i++ {
				if k.visx[i] {
					for j := 0; j < k.n; j++ {
						if !k.visy[j] && (k.lx[i]+k.ly[j]-k.weight[i][j]) < inc {
							inc = k.lx[i] + k.ly[j] - k.weight[i][j]
						}
					}
				}
			}

			//找不到能够加入的边，返回失败（即找不到完美匹配）
			if inc == INT_MAX {
				return -1
			}

			//找到增量后修改顶标
			// S和T  lx+ly不变
			// S和T' lx+ly会减少,有可能从而符合等式, 从而将边加入到相等子图中
			for i := 0; i < k.n; i++ {
				if k.visx[i] { //若是点x在S集合里
					k.lx[i] -= inc
				}
			}
			for j := 0; j < k.n; j++ {
				if k.visy[j] { //若是点y在T集合里
					k.ly[j] += inc
				}
			}
		}
	}

	sum := 0
	for i := 0; i < k.n; i++ {
		if k.match[i] > -1 {
			sum += k.weight[k.match[i]][i]
		}
	}

	if !k.maxWeight {
		sum = -sum
		return sum
	}
	return sum
}
