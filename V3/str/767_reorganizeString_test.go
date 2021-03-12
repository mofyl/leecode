package str

import (
	"fmt"
	"testing"
)

func TestReorganizeString(t *testing.T) {

	fmt.Println(reorganizeString("abbabbaaab"))
}

func reorganizeString(S string) string {

	if len(S) < 2 {
		return S
	}

	table := [26]int{}

	for i := 0; i < len(S); i++ {
		table[S[i]-'a']++
	}
	maxCount := 0
	c := 0
	threshold := (len(S) + 1) >> 1
	for i := 0; i < len(table); i++ {
		if maxCount < table[i] {
			maxCount = table[i]
			c = i
		}

		if maxCount > threshold {
			return ""
		}
	}

	b := make([]byte, len(S))
	idx := 0
	// 先将出现次数最多的放到偶数位置上
	for table[c] > 0 {
		b[idx] = uint8(c) + 'a'
		idx += 2
		table[c]--
	}

	// 若偶数位置没有填充满，那么就继续填充
	for i := 0; i < len(table); i++ {

		for table[i] > 0 {
			// 若偶数填充满了，那么就从奇数位置开始放
			if idx >= len(b) {
				idx = 1
			}

			b[idx] = uint8(i) + 'a'
			table[i]--
			idx += 2
		}

	}

	return string(b)

}
