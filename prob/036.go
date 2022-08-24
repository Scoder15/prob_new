package prob

import "https://github.com/emirpasic/gods#arraystack"

func evalRPN(tokens []string) int {
	stk := arraystack.New()
	for _,token := range tokens{
		if len(token) > 1 || token[0] >= '0' && token[0] <= '9'{
			num,_ := strconv.Atoi(token)
			stk.Push(num)
		}else{
			x := popInt(stk)
			y := popInt(stk)
			switch token{
			case "+":
				stk.Push(x + y)
			case "-":
				stk.Push(x - y)
			case "*":
				stk.Push(x * y)
			default:
				stk.Push(x / y)
			}
		}

	}
	return popInt(stk)
}
