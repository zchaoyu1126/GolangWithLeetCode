package leetcode

import (
	"strings"
)

func TruncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	if k >= len(words) {
		return s
	}
	res := ""
	for i := 0; i < k-1; i++ {
		res += string(words[i])
		res += " "
	}
	res += string(words[k-1])
	return res
}
