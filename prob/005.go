package prob

func MaxProduct(words []string) int {
	n := len(words)
	ret := 0
	for i := 0; i < n-1; i++ {
		dupM := make(map[rune]bool, 0)
		word := words[i]
		for _, v := range word {
			//fmt.Println("word:", v)
			dupM[(v)] = true
		}
		//fmt.Println("dupM:", dupM)
		for j := i + 1; j < n; j++ {
			if IsHaveSameStr(dupM, words[j]) {
				continue
			}
			l := len(words[i]) * len(words[j])
			if l > ret {

				ret = l
				//fmt.Println("ret:", ret, "wordi:", words[i], words[j])
			}
		}

	}
	return ret

}

func IsHaveSameStr(m map[rune]bool, s string) bool {
	for _, ss := range s {
		//fmt.Println(ss)
		if _, ok := m[(ss)]; ok {
			return true
		}
	}
	return false
}
