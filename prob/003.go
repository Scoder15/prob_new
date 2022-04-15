package prob

func CountBits1(n int) []int {
	ret := make([]int, 0)
	for i := 0; i <= n; i++ {
		ret = append(ret, getBinaryNums(i))
	}
	return ret
}

func CountBits(n int) []int {
	ret := make([]int, n+1)
	//fmt.Println(ret)
	for i := 0; i <= n; i++ {
		ret[i] = ret[i>>1] + (i & 1)
	}
	return ret
}

func getBinaryNums(n int) int {
	sum := 0
	for n > 0 {
		m := n % 2
		n = n >> 1

		if m == 1 {
			sum += 1
		}
	}
	return sum

}
