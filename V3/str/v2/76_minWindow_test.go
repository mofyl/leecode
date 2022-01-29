package v2

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {

	//s := "ADOBECODEBANC"
	//
	//t1 := "ABC"

	s := "bbaa"

	t1 := "aba"

	res := minWindow(s, t1)

	fmt.Println(res)
}

func minWindow(s string, t string) string {

	if len(t) < 0 || len(s) < len(t) {
		return ""
	}

	if s == t {
		return s
	}

	data := make([]int, 128)

	for i := 0; i < len(t); i++ {
		data[t[i]]++
	}

	l, r := 0, 0
	cnt := len(t)
	minLen := len(s) + 1

	//minLen := 1000000
	start := 0
	for ; r < len(s); r++ {

		if data[s[r]] > 0 {
			cnt--
		}

		data[s[r]]--

		if cnt == 0 {
			// 说明 目标串已经 全部出现在了 模式串中
			// 这里没有等于
			// 还要要求是 非目标串中的字符 。非目标串中的字符 都是小于0 的
			for l < r && data[s[l]] < 0 {
				data[s[l]]++
				l++
			}

			if minLen > r-l+1 {
				minLen = r - l + 1
				start = l
			}

			// 这里是到了 目标串中的字符了，那么下次就去找这个字符就行了。
			// 既然要找这个字符，就需要将 这个字符的值还原。
			// 若 目标串中的一个字符出现了多次，那么在 前面init data 的时候 也同样会 累加多次。这里一样可以将其加到 大于0
			cnt++
			data[s[l]]++

			l++

		}

	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]

}
