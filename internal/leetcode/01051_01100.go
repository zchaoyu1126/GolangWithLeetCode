package leetcode

// leetcode1089
func DuplicateZeros(arr []int) {
	var i, j int
	for {
		i++
		j++
		if j+1 > len(arr) {
			break
		} else if arr[i] == 0 {
			j++
		}
	}

	for i >= 0 {
		arr[j] = arr[i]
		if arr[i] == 0 {
			arr[j-1] = 0
			j--
		}
		i--
		j--
	}
}
