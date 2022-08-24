package prob

func dailyTemperatures(temperatures []int) []int {
	return []int{}
}

func onlyStack(arr []int) {
	stk := make([]int, 0)
	for _, v := range arr {
		//满足条件将栈顶元素出栈，将当前元素入栈。
		stk = append(stk, v)
	}

}
