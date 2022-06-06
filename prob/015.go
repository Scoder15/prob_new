package prob

func FindAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	n1, n2 := len(p), len(s)
	if n1 > n2 {
		return []int{}
	}
	window := make([]int, 26)
	for i := 0; i < n1; i++ {
		window[p[i]-'a']++
		window[s[i]-'a']--

	}
	if check(window) {
		ret = append(ret, 0)
	}
	for i := n1; i < n2; i++ {
		//入滑动窗口
		window[s[i]-'a']--
		//出滑动窗口
		window[s[i-n1]-'a']++
		if check(window) {
			ret = append(ret, i-n1+1)
		}
	}
	return ret
}

