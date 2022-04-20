package prob

import "fmt"

func SingleNumber(nums []int) int {
	ret := int32(0)
	for i := 0; i < 32; i++ {
		cnt := int32(0)
		for _, num := range nums {
			cnt += int32(num) >> i & 1
		}
		cnt %= 3
		ret |= cnt << i
		fmt.Println(ret)
	}
	return int(ret)
}
