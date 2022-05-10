package prob

func TwoSum1(numbers []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		m[numbers[i]] = i
	}
	//fmt.Println(m)
	for i := 0; i < len(numbers); i++ {

		if _, ok := m[target-numbers[i]]; ok {
			if i == m[target-numbers[i]] {
				continue

			}
			return []int{i, m[target-numbers[i]]}
		}
	}
	return []int{-1, -1}
}

func TwoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i, j}
		}

	}
}
