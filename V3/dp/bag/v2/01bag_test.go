package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func Test01Bag(t *testing.T) {

	N := 3
	V := 4
	v := []int{4, 2, 3}
	w := []int{4, 2, 3}

	res := Bag01(N, V, v, w)

	fmt.Println(res)
}

/*
	N 表示有 N 件物品， V表示 背包的容量，
	v 保存了 每件物品的价值
	w 保存了 每件物品的重量

	每件物品 只有一件， 一件物品不可重复选择

*/
func Bag01(N, V int, v, w []int) int {
	// return Bag01Helper(N, V, v, w, 0)

	// return Bag01DP(N, V, v, w)

	return Bag01DPV2(N, V, v, w)
}

func Bag01Helper(N, V int, v, w []int, idx int) int {

	if idx >= len(v) {
		return 0
	}

	if V <= 0 {
		return 0
	}

	// 当前物品可以选 也可以不选

	noOp := Bag01Helper(N, V, v, w, idx+1)

	if V >= w[idx] {
		choice := Bag01Helper(N, V-w[idx], v, w, idx+1) + v[idx]

		return tools.Max(noOp, choice)
	}

	return 0

}

func Bag01DP(N, V int, v, w []int) int {

	dp := make([][]int, N+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, V+1)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		// 这里表示 当前剩余的 重量
		for j := V; j >= 0; j-- {

			if j >= w[i] {
				dp[i][j] = tools.Max(dp[i+1][j], dp[i+1][j-w[i]]+v[i])
			}

		}

	}

	return dp[0][V]

}

func Bag01DPV2(N, V int, v, w []int) int {

	dp := make([]int, V+1)

	for i := len(v) - 1; i >= 0; i-- {

		// 这里表示 当前剩余的 重量
		for j := V; j >= w[i]; j-- {

			dp[j] = tools.Max(dp[j], dp[j-w[i]]+v[i])

		}

	}

	return dp[V]

}
