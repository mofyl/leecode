package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFullBag(t *testing.T) {

	N := 2
	C := 5
	v := []int{1, 2}
	w := []int{1, 2}

	res := fullBag(N, C, v, w)

	fmt.Println(res)
}

/*
	这里是完全背包问题，
	N 表示 有 N个元素， w表示每个元素的重量
	C 表示总重量， v 表示每个元素的价值

	要求：
		每个元素可以重复使用多次，最终要在 物品重量 不超过 总重量C 的情况下 ，价值总和最大

*/
func fullBag(N, C int, v, w []int) int {
	// return fullBagHelper(N, v, w, C, 0)

	return fullBagDP(N, v, w, C)
}

func fullBagHelper(N int, v, w []int, C int, idx int) int {

	if idx >= len(v) {
		return 0
	}

	var (
		noOp    int
		choice  int
		choice1 int
	)
	// 不选择 当前元素
	noOp = fullBagHelper(N, v, w, C, idx+1)

	// 选择 当前元素
	if C >= v[idx] {
		// 但是 下一个元素 选择 idx+1
		choice = fullBagHelper(N, v, w, C-w[idx], idx+1) + v[idx]
		// 下一个依旧选 当前元素
		choice1 = fullBagHelper(N, v, w, C-w[idx], idx) + v[idx]
	}

	return tools.Max(noOp, tools.Max(choice, choice1))

}

func fullBagDP(N int, v, w []int, C int) int {

	dp := make([][]int, len(v)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, C+1)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		for j := 0; j <= C; j++ {

			// 不选当前物品
			noOp := dp[i][j]
			choice := 0
			// 选择 当前物品，1 件 或多件
			for k := 1; ; k++ {

				if j < k*w[i] {
					break
				}

				choice = tools.Max(choice, dp[i+1][j-k*w[i]]+v[i]*k)

			}

			dp[i][j] = tools.Max(noOp, choice)

		}

	}

	return dp[0][C]

}
