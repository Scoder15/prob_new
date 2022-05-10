package prob

import "sort"

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				ret = append(ret, []int{i, left, right})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}

		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}

	}
	return ret
}
