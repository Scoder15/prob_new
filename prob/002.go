package prob

import "strconv"

func AddBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	// aa := []rune(a)
	// bb := []rune(b)
	carry := 0
	ret := ""
	for i >= 0 || j >= 0 {
		digitA := 0
		if i >= 0 {
			digitA = int(a[i] - '0')
		}
		digitB := 0
		if j >= 0 {
			digitB = int(b[j] - '0')
		}
		sum := digitA + digitB + carry

		if sum >= 2 {
			carry = 1
			sum = sum - 2
		} else {
			carry = 0
		}
		ret += strconv.Itoa(sum)
		i--
		j--
	}
	if carry == 1 {
		ret += strconv.Itoa(carry)
	}
	return reverse(ret)
}

func reverse(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1
	for i < j {
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
		i++
		j--
	}
	return string(r)
}
