package v3

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {

	s := "ADOBECODEBANC"
	target := "ABC"

	res := minWindow(s, target)

	fmt.Println(res)
}

func minWindow(s string, t string) string {

	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	data := make([]int, 256)

	for i := 0; i < len(t); i++ {
		data[t[i]]++
	}

	l := 0
	r := 0
	cnt := len(t)
	minLen := len(s) + 1
	start := 0
	for ; r < len(s); r++ {

		if data[s[r]] > 0 {
			cnt--
		}
		data[s[r]]--

		if cnt == 0 {

			// 小于 0 的 都是 非目标字符
			// 这里 没有等于 因为 到了这里 窗口中一定要有值
			for l < r && data[s[l]] < 0 {
				// 还原次数
				data[s[l]]++
				l++
			}
			// 到了这里 还差 data[l] 这个元素没有还原
			if minLen > r-l+1 {
				start = l
				minLen = r - l + 1
			}

			// 这里是因为 我们既然要还原 。在第一次循环那里，我们对 data[l]做了-1.虽然 data[l] >= 0 但是也是 -1之后的结果。
			// 既然是 >= 0 那么上面的循环就无法进行还原。所以这里就进行补偿
			// 此时 l 指向位置的元素 一定是 t 中字符。那么既然 其他t中字符已经都在 r-l+1范围内了。那么我们下面只需要在去找 l位置的元素就好了
			data[s[l]]++
			// 为了寻找 l 后面l位置的元素。所以将当前位置跳过。
			l++
			cnt++

		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
