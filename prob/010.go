package prob

//整数数组，有正有负
func SubarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	sum, ans := 0, 0
	//构造前缀和数组
	for _, num := range nums {
		sum += num
		ans += m[sum-k]
		m[sum]++

	}
	return ans
}
