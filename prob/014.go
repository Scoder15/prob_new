package prob

func CheckInclude(s1, s2 string) bool {
	m := make(map[rune]bool, 0)
	for _, v := range s1 {
		m[v] = true
	}
	k := len(m)
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			k--
			if k == 0 {
				return true
			}
		}
	}
	return false
}
