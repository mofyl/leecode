package v3

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"

	res := checkInclusion(s1, s2)
	fmt.Println(res)
}

func checkInclusion(s1 string, s2 string) bool {

	table := make([]int, 26)
	window := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		table[s1[i]-'a']++
	}

	l, r := 0, 0

	for ; r < len(s2); r++ {

		idx := s2[r] - 'a'
		window[idx]++
		// 这里用for 循环的原因是 : s1 := ba , s2 := baoba
		// 当r 指向 o 时，前面的字符都要从窗口中去掉
		for window[idx] > table[idx] {
			window[s2[l]-'a']--
			l++
		}

		if r-l+1 == len(s1) {
			return true
		}

	}
	return false
}
