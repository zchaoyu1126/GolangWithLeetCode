package algorithm

// 生成素数
func EratosthenesMakePrimeList(maxNum int) []int {
	isNotPrime := make([]bool, maxNum+5)
	primeList := []int{}
	for i := 2; i <= maxNum; i++ {
		if !isNotPrime[i] {
			primeList = append(primeList, i)
		}
		for j := i + i; j <= maxNum; j += i {
			isNotPrime[j] = true
		}
	}
	return primeList
}

func EulerMakePrimeList(maxNum int) []int {
	isNotPrime := make([]bool, maxNum+5)
	primeList := []int{}
	for i := 2; i <= maxNum; i++ {
		if !isNotPrime[i] {
			primeList = append(primeList, i)
		}
		for j := 0; j < len(primeList); j++ {
			if i*primeList[j] > maxNum {
				break
			}
			isNotPrime[i*primeList[j]] = true
			if i%primeList[j] == 0 {
				break
			}
		}
	}
	return primeList
}
