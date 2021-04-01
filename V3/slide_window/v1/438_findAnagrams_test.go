package v1

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	//s := "cbaebabacd"
	//p := "abc"

	s := "abab"
	p := "ab"

	res := findAnagrams(s, p)
	fmt.Println(res)
}

func findAnagrams(s string, p string) []int {

	if len(s) < 1 || len(p) < 1 {
		return nil
	}

	res := make([]int, 0)

	table := [26]int{}
	window := [26]int{}

	for i := 0; i < len(p); i++ {
		table[p[i]-'a']++
	}

	l, r := 0, 0

	for r < len(s) {
		rC := s[r] - 'a'

		window[rC]++
		// 这里表示当前字符不是table中的字符。因为若是table中的字符 个数一定是 <= table[c]的
		// 出现了无用字符就要收缩窗口左边界.这里其实是一直要将窗口收缩完
		// 最终left 会到达当前right的位置 将当前位置的值减下去
		for window[rC] > table[rC] {
			window[s[l]-'a']--
			l++
		}

		if r-l+1 == len(p) {
			res = append(res, l)
		}
		r++
	}

	return res
}
