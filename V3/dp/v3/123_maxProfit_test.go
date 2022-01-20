package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_123(t *testing.T) {

	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}

	// prices := []int{1, 2, 3, 4, 5}

	res := maxProfit_123(prices)

	fmt.Println(res)
}

func maxProfit_123(prices []int) int {

	// 该题目用dp way 计算的时候 要注意：只能交易两次 这个条件 必须放在 卖出的时候，也就是 在卖出的时候去减交易次数
	// 递归的时候没有这个限制
	// return maxProfit_123Helper(prices, 2, 0, 1)

	return maxProfit_123DP(prices)
}

// 1 表示未持有， 2 表示 已经持有
func maxProfit_123Helper(prices []int, count int, cur int, status int) int {

	if cur >= len(prices) || count == 0 {
		return 0
	}

	none := maxProfit_123Helper(prices, count, cur+1, status)

	if status == 1 && count > 0 {
		// 表示未持有
		// buy := maxProfit_123Helper(prices, count-1, cur+1, 2) - prices[cur] // 这样也可以
		buy := maxProfit_123Helper(prices, count, cur+1, 2) - prices[cur]

		return tools.Max(buy, none)
	} else if status == 2 {
		// 表示已经持有
		sell := maxProfit_123Helper(prices, count-1, cur+1, 1) + prices[cur]
		return tools.Max(sell, none)
	}
	// 这里就是 count 已经不够了， 那么就返回0 就可以
	return 0
}

// 1 表示未持有， 2 表示 已经持有
func maxProfit_123DP(prices []int) int {

	// 第一纬度表示 共有 prices+1 个元素
	// 第二维度表示 共有 2 中状态 ，1 表示 未持有， 2表示已经持有
	// 第三维度表示  对于 未持有的状态 可以选择  买入 或者 不买入
	dp := make([][][]int, len(prices)+1)
	count := 3 // 这里可以取到2 所以给了3

	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 3)

		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, count)
		}
	}

	for i := len(dp) - 2; i >= 0; i-- {

		// 未持有的时候 是随时可以买卖的
		dp[i][1][2] = tools.Max(dp[i+1][2][2]-prices[i], dp[i+1][1][2])
		// 未持有 有 1 次的时候 是可以直接买的
		dp[i][1][1] = tools.Max(dp[i+1][2][1]-prices[i], dp[i+1][1][1])

		// 已持有 -> 未持有
		// 卖出的时候 就要减交易次数
		dp[i][2][1] = tools.Max(dp[i+1][1][0]+prices[i], dp[i+1][2][1])
		dp[i][2][2] = tools.Max(dp[i+1][1][1]+prices[i], dp[i+1][2][2])
	}

	return dp[0][1][2]
}
