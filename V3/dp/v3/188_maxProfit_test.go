package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_188(t *testing.T) {

	prices := []int{2, 4, 1}
	k := 2

	// prices := []int{3, 2, 6, 5, 0, 3}
	// k := 2

	res := maxProfit_188(k, prices)

	fmt.Println(res)
}

func maxProfit_188(k int, prices []int) int {
	// return maxProfit_188Helper(k, prices, 0, 1)
	return maxProfit_188Dp(k, prices)
}

// status == 1 未持有， status == 2 已持有
func maxProfit_188Helper(k int, prices []int, cur int, status int) int {

	if cur >= len(prices) || k == 0 {
		return 0
	}

	noOp := maxProfit_188Helper(k, prices, cur+1, status)

	if status == 1 {
		// 未持有，就可以持有
		buy := maxProfit_188Helper(k, prices, cur+1, 2) - prices[cur]

		return tools.Max(noOp, buy)
	} else {
		sell := maxProfit_188Helper(k-1, prices, cur+1, 1) + prices[cur]
		return tools.Max(noOp, sell)
	}

}

func maxProfit_188Dp(k int, prices []int) int {

	dp := make([][][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {

		dp[i] = make([][]int, 3)

		for j := 0; j < len(dp[i]); j++ {

			dp[i][j] = make([]int, k+1)

		}

	}
	// status == 1 未持有， status == 2 已持有
	for i := len(dp) - 2; i >= 0; i-- {

		// 这里进行 k 次
		for j := 1; j <= k; j++ {
			// 未持有 -> 已持有
			dp[i][1][j] = tools.Max(dp[i+1][2][j]-prices[i], dp[i+1][1][j])

			// 已持有 -> 未持有
			dp[i][2][j] = tools.Max(dp[i+1][1][j-1]+prices[i], dp[i+1][2][j])

		}

	}

	return dp[0][1][k]

}
