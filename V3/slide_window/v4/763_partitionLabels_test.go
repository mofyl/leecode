package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestPartitionLabels(t *testing.T) {

	s := "ababcbacadefegdehijhklij"

	res := partitionLabels(s)

	fmt.Println(res)
}

func partitionLabels(s string) []int {

	data := make([]int, 26)

	// 记录每个字符出现的最远的 idx
	for i := 0; i < len(s); i++ {
		data[s[i]-'a'] = i
	}

	res := make([]int, 0)
	maxIdx := 0
	start := 0
	for i := 0; i < len(s); i++ {

		maxIdx = tools.Max(maxIdx, data[s[i]-'a'])

		if i == maxIdx {
			res = append(res, i-start+1)
			start = i + 1
		}

	}

	return res

}
