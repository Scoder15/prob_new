package prob

func NumSubarrayProductLessThanK(nums []int, k int) int {
	left, right, sum, n := 0, 0, 1, len(nums)
	ret := 0
	for right < n {
		sum *= nums[right]
		right++
		for left < right && sum >= k {
			sum /= nums[left]
			left++
		}
		ret += right - left
	}

	return ret
}
