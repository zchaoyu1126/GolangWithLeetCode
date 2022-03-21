package leetcode

// leetcode1920
func BuildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

// leetcode1929
func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i], ans[i+len(nums)] = nums[i], nums[i]
	}
	return ans
}
