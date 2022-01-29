package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	//s := "abcabcbb"

	s := "pwwkew"

	res := lengthOfLongestSubstring(s)

	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {

	data := make(map[byte]int)
	res := 0
	l, r := 0, 0

	for ; r < len(s); r++ {

		idx, ok := data[s[r]]

		if ok && idx > l {
			// 说明发生了重复 && 没有回头 则跳过
			l = idx
		}
		data[s[r]] = r + 1
		res = tools.Max(res, r-l+1)
	}

	return res

}
