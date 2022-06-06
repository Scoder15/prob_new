package prob

func LengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	n := len(s)
	ans := 0
	left, right := 0, 0
	for right < n {
		ch := s[right]
		right++
		window[ch]++
		for window[ch] > 1 {
			window[s[left]]--
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}
