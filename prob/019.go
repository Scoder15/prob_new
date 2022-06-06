package prob

func vaildPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return isPalin(left, right-1, s) || isPalin(left+1, right, s)
		}

	}
	return true
}

func isPalin(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
