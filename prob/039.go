package prob

func LargestRectangleArea(heights []int) int {
	index := 0
	maxArea := 0
	for _, v := range heights {
		area := v
		left := index - 1
		for left > -1 {
			if v <= heights[left] {
				area += heights[left]
				left--
			}
		}
		right := index + 1
		for right < len(heights) {
			if v <= heights[right] {
				area += heights[right]
				right++
			}
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
