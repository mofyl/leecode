package v2

import (
	"fmt"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	S := "ababcbacadefegdehijhklij"

	res := partitionLabels_v2(S)

	fmt.Println(res)

}

func partitionLabels(S string) []int {

	if len(S) < 1 {
		return nil
	}

	table := [26]int{}
	window := make(map[uint8]int32, 26)
	for i := 0; i < len(S); i++ {
		table[S[i]-'a']++
	}

	res := make([]int, 0)

	l, r := 0, 0

	for r < len(S) {
		c := S[r] - 'a'

		// 若当前字符在后面还有，则表明不能分成片段
		// 直到当前字符 没有了，
		if table[c] > 0 {
			table[c]--
		}

		window[c]++
		if table[c] == 0 {
			// 检查当前窗口内的字符是否都已经为0
			for k, _ := range window {

				if table[k] > 0 {
					goto Con
				}
			}
			// 到这里说明window中的字符在 table中已经没有了
			// 开始分片
			res = append(res, r-l+1)
			l = r + 1
			// 清空当前window
			window = make(map[uint8]int32, 26)
		}
	Con:

		r++
	}
	return res
}

func partitionLabels_v2(S string) []int {

	if len(S) < 1 {
		return nil
	}

	// 这里记录的是 当前字符在 S中出现的最远的位置
	table := make(map[uint8]int, 26)
	for i := 0; i < len(S); i++ {
		table[S[i]-'a'] = i
	}

	res := make([]int, 0)

	l, r := 0, 0
	maxPos := 0
	for r = 0; r < len(S); r++ {
		c := S[r] - 'a'

		pos := table[c]

		if pos > maxPos {
			maxPos = pos
		}

		if r == maxPos {
			// 表示当前窗口内最远的字符 的位置 已经找到了。
			// 那就可以开始记录了
			res = append(res, r-l+1)
			l = r + 1 // 这里r是当前片段的最后一个字符，要跳开
		}

	}
	return res
}
