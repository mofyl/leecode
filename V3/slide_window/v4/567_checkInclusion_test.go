package v4

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {

	s1 := "ab"
	s2 := "eidboaooo"

	res := checkInclusion(s1, s2)
	fmt.Println(res)
}

// s1 是模式串  ， s2 是目标串
func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	data := make([]int, 26)
	window := make([]int, 26)

	for i := 0; i < len(s1); i++ {

		data[s1[i]-'a']++

	}
	l := 0
	for r := 0; r < len(s2); r++ {

		window[s2[r]-'a']++

		// 说明出现了 非 s1 中的字符, 将window 清空
		for window[s2[r]-'a'] > data[s2[r]-'a'] {
			window[s2[l]-'a']--
			l++
		}

		if r-l+1 == len(s1) {
			return true
		}

	}

	return false

}
