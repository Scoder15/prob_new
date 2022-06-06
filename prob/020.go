package prob

//其实区间[i,j]里面是否是字符串
func CountSubString(s string) int {
	ret := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = false
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] {
				if j-i < 2 || dp[i+1][j-1] {
					dp[i][j] = true
					ret++
				}

			}

		}
	}

	return ret
}
