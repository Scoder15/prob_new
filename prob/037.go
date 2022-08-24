package prob

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, v := range asteroids {
		for len(stack) > 0 && stack[len(stack)-1] > 0 && v < 0 { //碰撞条件
			a := stack[len(stack)-1]
			b := abs(v)
			stack = stack[:len(stack)-1]
			if a > b {
				v = a
			} else if a < b {
				continue
			} else {
				v = 0
			}
		}
		if v != 0 {
			stack = append(stack, v)
		}
	}
	return stack
}
