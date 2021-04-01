package v2

import (
	"fmt"
	"math"
	"testing"
)

func TestMinWindow(t *testing.T) {

	s := "ADOBECODEBANC"
	tmp := "ABC"

	res := minWindow(s, tmp)
	fmt.Println(res)

}

func minWindow(s string, t string) string {

	if len(s) < 1 || len(t) < 1 {
		return ""
	}
	// 这里给128：由于s和t中 大写小写都有可能出现，所以无法进行映射，只能给一个全量的table

	window := [128]int{}

	for i := 0; i < len(t); i++ {
		window[t[i]]++
	}

	l, r := 0, 0
	minLen := math.MaxInt64
	// 表示当前已经扫描过 t中的字符个数
	cnt := len(t)
	start := 0
	for r < len(s) {
		c := s[r]
		if window[c] > 0 {
			cnt--
		}

		window[c]--
		// 表示此时 t中的字符已经全部出现在了 window中
		if cnt == 0 {
			// 去掉无用的字符,收缩左边界
			for l < r && window[s[l]] < 0 {
				window[s[l]]++
				l++
			}

			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}

			window[s[l]]++
			cnt++
			l++
		}
		r++
	}

	if minLen == math.MaxInt64 {
		return ""
	}

	return s[start : start+minLen]

}
