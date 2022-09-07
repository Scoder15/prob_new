package prob

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stk []int
	for i, t := range temperatures {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < t {
			j := stk[len(stk)-1]
			res[j] = i - j
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return res
}

func onlyStack(arr []int) {
	stk := make([]int, 0)
	for i, v := range arr {
		//满足条件将栈顶元素出栈，将当前元素入栈。
		stk = append(stk, v)
		//pop
		stk = stk[:len(stk)-1]
		stk = append(stk, i)
	}

}
