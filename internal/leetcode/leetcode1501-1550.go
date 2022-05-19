package leetcode

func NumWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		changeNum := numBottles / numExchange
		res += changeNum
		numBottles = changeNum + numBottles%numExchange
	}
	return res
}
