package v3

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

	l := 0
	r := 0
	maxCount := 0

	data := make([]int, 26)

	for ; r < len(s); r++ {

		idx := s[r] - 'A'
		data[idx]++

		maxCount = tools.Max(maxCount, data[idx])

		if r-l+1 > maxCount+k {
			data[s[l]-'A']--
			l++
		}
	}
	// 这里 使用 r-l 是因为
	// maxCount+k 是最大的长度范围，若maxCount 出现后 我们一直使用这个区间
	return r - l
}
