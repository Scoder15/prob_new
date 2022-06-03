package prob

func FindMaxLength(nums []int) int {
	sum := 0
	ret := 0
	m := map[int]int{0: -1}
	for i, v := range nums {
		if v == 0 {
			v = -1
		}
		sum += v
		if j, ok := m[sum]; ok {
			ret = max(ret, i-j)
		} else {
			m[sum] = i
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
