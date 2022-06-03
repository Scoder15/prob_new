package prob

import "math"

func MinSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	ret := math.MaxInt32
	for right = 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= target {
			if ret > right-left+1 {
				ret = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	return ret
}
