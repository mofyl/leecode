package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {

	//s := "ABAB"
	//k := 2

	s := "AABABBA"
	k := 1

	res := characterReplacement(s, k)

	fmt.Println(res)
}

func characterReplacement(s string, k int) int {

	data := make([]int, 26)

	l, r := 0, 0
	maxLen := 0
	maxCount := 0
	for ; r < len(s); r++ {

		data[s[r]-'A']++

		maxCount = tools.Max(maxCount, data[s[r]-'A'])

		// 这里表示当前窗口中 可替换的 字符 太多
		if r-l+1 > maxCount+k {
			data[s[l]-'A']--
			l++
		}

		maxLen = tools.Max(maxLen, r-l+1)

	}

	return maxLen

}
