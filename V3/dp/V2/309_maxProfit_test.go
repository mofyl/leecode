package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_309(t *testing.T) {

	prices := []int{1, 2, 3, 0, 2}

	res := maxProfit_309(prices)
	fmt.Println(res)
}

func maxProfit_309(prices []int) int {

	if len(prices) < 1 {
		return 0
	}

	//return maxProfit309Helper(prices, 0, 0)
	return maxProfit309DpWay(prices)

}

// 递归的含义为: 找到 prices[idx ...] 从idx 后 获取的最大利润
// status ： 0 为 不持有，1为持有， 2为 冷冻期
func maxProfit309Helper(prices []int, idx int, status int) int {

	if idx >= len(prices) {
		return 0
	}

	// 这里是无操作的情况,无论是持有,不持有,冷冻期 都可以无操作

	var (
		noOp  int
		buyIn int
		sell  int
	)

	noOp = maxProfit309Helper(prices, idx+1, status)

	if status == 0 {
		// 当前不持有 ,那么今天就可以买入
		buyIn = maxProfit309Helper(prices, idx+1, 1) - prices[idx]
	} else if status == 1 {
		// 当前已经持有了
		// 当前可以选择卖出 ,但是卖出了, 后面一天就无法买入了，所以实际上是 idx+2 之后的利润
		sell = maxProfit309Helper(prices, idx+2, 0) + prices[idx]
	}

	return tools.Max(noOp, tools.Max(buyIn, sell))

}

func maxProfit309DpWay(prices []int) int {

	dp := make([][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}

	for i := len(dp) - 2; i >= 0; i-- {

		dp[i][0] = tools.Max(dp[i+1][0], dp[i+1][1]-prices[i])

		if i+2 >= len(dp) {
			// 这里表示 卖出后 天数不够 i+2。这里就应该去取 i+1保持持有 和 当前利润 的max
			dp[i][1] = tools.Max(dp[i+1][1], prices[i])
		} else {
			dp[i][1] = tools.Max(dp[i+1][1], dp[i+2][0]+prices[i])
		}
	}

	return dp[0][0]
}
