package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	s := "ABAB"
	k := 2
	res := characterReplacement(s, k)
	fmt.Println(res)
}

func characterReplacement(s string, k int) int {

	if len(s) <= 1 {
		return len(s)
	}

	window := [26]int{}

	l, r := 0, 0
	maxCount := 0
	for r = 0; r < len(s); r++ {

		c := s[r] - 'A'

		window[c]++

		// 寻找当前窗口中 出现次数最多的字符
		maxCount = tools.Max(maxCount, window[c])

		// 这里 表示窗口内部可替换的字符大于k个了
		if r-l+1 > maxCount+k {
			window[s[l]-'A']--
			l++
		}

	}

	return r - l

}
