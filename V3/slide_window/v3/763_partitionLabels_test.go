package v3

import (
	"fmt"
	"testing"
)

func TestPartitionLabels(t *testing.T) {

	S := "ababcbacadefegdehijhklij"

	res := partitionLabels(S)

	fmt.Println(res)
}

func partitionLabels(s string) []int {

	// 先将每个字符出现的最远位置标记出来，
	// 然后找到该最远位置，然后记录当前 窗口的大小
	data := make([]int, 26)

	for i := 0; i < len(s); i++ {
		data[s[i]-'a'] = i
	}

	l := 0
	r := 0
	maxPos := 0
	res := make([]int, 0)
	for ; r < len(s); r++ {

		pos := data[s[r]-'a']

		if pos > maxPos {
			maxPos = pos
		}

		if r == maxPos {
			res = append(res, r-l+1)
			l = r + 1
		}

	}

	return res
}
