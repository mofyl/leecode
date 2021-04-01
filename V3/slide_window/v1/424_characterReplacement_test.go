package v1

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

	l := 0
	r := 0
	// cnt内保存着 当前窗口中出现的 所有字符的计数
	cnt := [26]int{}
	// maxCount 保存着 s中 出现最多字符的个数
	maxCount := 0

	for r = 0; r < len(s); r++ {

		b := s[r] - 'A'

		cnt[b]++

		maxCount = tools.Max(maxCount, cnt[b])

		if r-l+1 > maxCount+k {
			cnt[s[l]-'A']--
			l++
		}

	}
	// 由于我们是从0开始的，最后出循环以后 r == len(s) ，超过了下标范围，要减一
	// r-1-l+1
	return r - l
}
