package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_714(t *testing.T) {

	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	// prices := []int{1, 3, 7, 5, 10, 3}
	// fee := 3

	res := maxProfit_714(prices, fee)

	fmt.Println(res)
}

func maxProfit_714(prices []int, fee int) int {

	// return maxProfit_714Helper(prices, fee, 1, 0)

	return maxProfit_714DP(prices, fee)
}

// 1 表示 不持有， 2表示持有
func maxProfit_714Helper(prices []int, fee int, status int, cur int) int {

	if cur >= len(prices) {
		return 0
	}

	none := maxProfit_714Helper(prices, fee, status, cur+1)

	if status == 1 {
		// 当前不持有，那么就可以买入
		// 买入需要 减去 当前的价格 和 手续费
		buy := maxProfit_714Helper(prices, fee, 2, cur+1) - prices[cur] - fee
		return tools.Max(none, buy)
	} else if status == 2 {
		// 表示持有，那么就可以卖出
		sell := maxProfit_714Helper(prices, fee, 1, cur+1) + prices[cur]
		return tools.Max(sell, none)
	}

	return 0
}

func maxProfit_714DP(prices []int, fee int) int {

	dp := make([][]int, len(prices)+1)

	// 1 表示 不持有， 2表示持有
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		dp[i][1] = tools.Max(dp[i+1][2]-prices[i]-fee, dp[i+1][1])

		dp[i][2] = tools.Max(dp[i+1][1]+prices[i], dp[i+1][2])

	}

	return dp[0][1]

}
