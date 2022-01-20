package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_309(t *testing.T) {

	nums := []int{1, 2, 3, 0, 2}

	res := maxProfit_309(nums)

	fmt.Println(res)
}

func maxProfit_309(prices []int) int {
	// return maxProfit_309Helper(prices, 0, 2)

	return maxProfit_309HelperDP(prices)

}

// status == 1 表示 持有， status == 2 表示不持有
func maxProfit_309Helper(prices []int, cur int, status int) int {

	if cur >= len(prices) {
		return 0
	}

	// 无论 status 是什么，我们都可以什么都不做
	empty := maxProfit_309Helper(prices, cur+1, status)

	if status == 1 {
		// 表示当前已经持有了 股票，此时我们可以 卖出,这里是 cur+2是因为卖出后 要有一天的冷冻期
		sell := maxProfit_309Helper(prices, cur+2, 2) + prices[cur]
		return tools.Max(empty, sell)
	} else if status == 2 {
		// 表示之前 没有持有 现在就可以买入了
		buy := maxProfit_309Helper(prices, cur+1, 1) - prices[cur]
		return tools.Max(empty, buy)
	}

	return 0
}

// status == 1 表示 持有， status == 2 表示不持有
func maxProfit_309HelperDP(prices []int) int {

	dp := make([][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {
		// 0 , 1 , 2
		dp[i] = make([]int, 3)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		// 持有
		if i+2 >= len(dp) {
			// 若 i+2 越界 那么 我们应该去比较 持有 和 prices[i]
			dp[i][1] = tools.Max(dp[i+1][1], prices[i])
		} else {
			dp[i][1] = tools.Max(dp[i+2][2]+prices[i], dp[i+1][1])
		}

		// 不持有
		dp[i][2] = tools.Max(dp[i+1][1]-prices[i], dp[i+1][2])
	}

	return dp[0][2]

}
