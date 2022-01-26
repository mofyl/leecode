package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestBagWithAmount(t *testing.T) {
	N := 2
	C := 5
	v := []int{1, 2}
	w := []int{1, 2}
	s := []int{2, 1}

	res := bagWithAmount(N, C, v, w, s)

	fmt.Println(res)
}

/*

	表示有N中 物品 和一个 容量为 C 的 背包

	第 i 件物品的 价值为 v[i]
				重量为 w[i]
				数量为 s[i]


*/

func bagWithAmount(N, C int, v, w, s []int) int {

	//return bagWithAmountHelper(N, v, w, s, C, 0)

	//return bagWithAmountDp(N, C, v, w, s)

	return bagWithAmountDpV2(N, C, v, w, s)
}

func bagWithAmountHelper(N int, v, w, s []int, C, idx int) int {

	if C <= 0 || idx >= len(v) {
		return 0
	}

	noOp := bagWithAmountHelper(N, v, w, s, C, idx+1)

	choice := -1
	for k := 1; k <= s[idx] && C >= k*w[idx]; k++ {

		cur := bagWithAmountHelper(N, v, w, s, C-k*w[idx], idx+1) + k*v[idx]

		choice = tools.Max(choice, cur)

	}

	return tools.Max(noOp, choice)

}

func bagWithAmountDp(N, C int, v, w, s []int) int {

	dp := make([][]int, len(v)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, C+1)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		for j := C; j >= w[i]; j-- {

			for k := 1; k <= s[i] && j >= k*w[i]; k++ {

				dp[i][j] = tools.Max(dp[i+1][j], dp[i+1][j-k*w[i]]+k*v[i])

			}

		}

	}

	return dp[0][C]

}

func bagWithAmountDpV2(N, C int, v, w, s []int) int {

	dp := make([]int, C+1)

	for i := len(v) - 1; i >= 0; i-- {

		for j := C; j >= w[i]; j-- {

			for k := 1; k <= s[i] && j >= k*w[i]; k++ {
				dp[j] = tools.Max(dp[j], dp[j-k*w[i]]+k*v[i])
			}

		}

	}

	return dp[C]

}
