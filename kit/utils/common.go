package utils

func MaxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func AbsInt(x int) (int, bool) {
	if x < 0 {
		return -x, true
	}
	return x, false
}
