package prob

import (
	"math"
)

func Divide(a, b int) int {
	sign := 1
	if a*b < 0 {
		sign = -1

	}
	a = abs(a)
	b = abs(b)
	tot := 0
	for a > b {
		cnt := 0
		for a >= b<<(cnt+1) {
			cnt += 1
		}
		tot += 1 << cnt
		a -= b << cnt
	}

	ans := sign * tot
	if ans >= math.MinInt32 && ans <= math.MaxInt32 {
		return ans

	}
	return math.MaxInt32
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// func qpow(a, n int) int {
// 	ans := 1
// 	for a > 0 {
// 		if n&1 > 0 {
// 			ans *= a
// 		}
// 		a *= a
// 		n >>= 1
// 	}
// 	return ans
// }
