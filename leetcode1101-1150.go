package leetcode

// leetcode1109
func CorpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		first, last, seats := booking[0], booking[1], booking[2]
		diff[first-1] += seats
		diff[last] -= seats
	}
	sum := make([]int, n+2)
	for i := 1; i <= len(diff); i++ {
		sum[i] = sum[i-1] + diff[i-1]
	}
	return sum[1 : len(sum)-1]
}
