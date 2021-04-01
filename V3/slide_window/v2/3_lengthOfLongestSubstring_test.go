package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	//s := "abcabcbb"
	s := "bbbbb"
	num := lengthOfLongestSubstring(s)
	fmt.Println(num)
}

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	table := make(map[uint8]int, len(s))

	l := 0
	maxLen := 0
	for r := 0; r < len(s); r++ {

		v, ok := table[s[r]]
		// 表示已经重复,需要收缩左边
		if ok {
			l = tools.Max(l, v)
		}
		// 这里若放到 上面ok的块中， 若s中一个重复的都没有，就检查不出来了
		maxLen = tools.Max(maxLen, r-l+1)
		table[s[r]] = r + 1
	}

	return maxLen

}
