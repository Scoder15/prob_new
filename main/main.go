package main

import (
	"fmt"

	"example.com/m/prob"
)

func main() {
	// s1 := "100"
	// s2 := "101"
	// fmt.Println(prob.AddBinary(s1, s2))
	//fmt.Println(prob.CountBits(5))
	//fmt.Println(prob.SingleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	//fmt.Println(prob.SingleNumber([]int{-2, -2, -4, -2}))
	fmt.Println(prob.MaxProduct([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}))
	// m := make(map[rune]bool, 0)
	// m['a'] = true
	// m['b'] = true
	// m['c'] = true
	// m['w'] = true
	// fmt.Println(prob.IsHaveSameStr(m, "baz"))
}
