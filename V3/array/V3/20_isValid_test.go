/*
 * @Author: lirui
 * @Date: 2022-01-07 15:35:31
 * @LastEditTime: 2022-01-07 16:08:09
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\20_isValid_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {

	// s := "()"

	// s := "()[]{}"

	// s := "{[]}"

	s := "((("

	res := isValid(s)

	fmt.Println(res)
}

func isValid(s string) bool {

	n := len(s)

	// 必须是偶数个
	if n&1 == 1 {
		return false
	}

	backets := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	que := make([]byte, 0, n)

	for i := 0; i < n; i++ {

		if len(que) > 0 && isMatch(que[len(que)-1], s[i], backets) {
			// 如果是匹配的 右边
			que = que[:len(que)-1]
			continue
		}

		_, ok := backets[s[i]]

		if ok {
			// 如果是 左边 就加入 队列
			que = append(que, s[i])
			continue
		}
		// 其他情况
		return false

	}

	// que 中应该都被消耗完，若没有 则 返回false
	return len(que) == 0

}

func isMatch(a, b byte, data map[byte]byte) bool {

	t, ok := data[a]

	if ok && t == b {
		return true
	}

	return false
}
