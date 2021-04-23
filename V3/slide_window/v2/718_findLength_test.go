package v2

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestFindLength(t *testing.T) {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}

	//A := []int{0, 0, 0, 0, 1}
	//B := []int{1, 0, 0, 0, 0}

	res := findLength(A, B)

	fmt.Println(res)
}

func findLength(A []int, B []int) int {

	// 若有一个等于0 那么说明没有 重复子数组
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	// 这里首先将两个数组进行对齐，然后 保持A不动，移动B 进行比较
	n, m := len(A), len(B)
	res := math.MinInt64
	for i := 0; i < n; i++ {

		minLen := tools.Min(m, n-i)
		curMax := findLengthHelper(A, B, i, 0, minLen)

		res = tools.Max(curMax, res)

	}

	for i := 0; i < m; i++ {
		minLen := tools.Min(n, m-i)
		curMax := findLengthHelper(B, A, i, 0, minLen)
		res = tools.Max(curMax, res)
	}

	if res == math.MinInt64 {
		return 0
	}

	return res

}

// 返回 相同 子序列的最大长度
func findLengthHelper(A []int, B []int, aStart, bStart int, moveLen int) int {

	k := 0
	res := math.MinInt64
	for i := 0; i < moveLen; i++ {

		if A[aStart+i] == B[bStart+i] {
			k++
			res = tools.Max(res, k)
		} else {
			k = 0
		}

	}
	return res
}
