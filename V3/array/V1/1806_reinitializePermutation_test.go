package V1

import (
	"fmt"
	"testing"
)

func TestReinitializePermutation(t *testing.T) {

	n := 6

	res := reinitializePermutation(n)

	fmt.Println(res)
}

// n 是一个偶数
func reinitializePermutation(n int) int {

	if n <= 2 {
		return 1
	}

	// 这里表示经过1次转换后的数组
	// 该数组有以下特点： 0和n-1 下标的元素 位置不变
	// 转换后数组组成  前半部分都是偶数，剩下的都是奇数
	// 那么只要 1这个数字 回到了 下标为1 的位置，那么就表示 倒转公式完成
	nextArr := make([]int, 0, n)

	// 对 nextArr 数组进行整合
	for i := 0; i < n; i += 2 {
		nextArr = append(nextArr, i)
	}

	for i := 1; i < n; i += 2 {
		nextArr = append(nextArr, i)
	}

	// 由于第一个元素0 是不会动的，所以不用考虑
	count := 1
	for i := 1; ; {

		i = nextArr[i]

		if i == 1 {
			return count
		}
		count++
		if count >= len(nextArr) {
			return 0
		}
	}

}
