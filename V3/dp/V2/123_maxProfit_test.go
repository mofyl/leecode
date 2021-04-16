package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit123(t *testing.T) {

	//prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	prices := []int{1, 2, 3, 4, 5}

	res := maxProfit123(prices)
	fmt.Println(res)
}

func maxProfit123(prices []int) int {

	if len(prices) < 1 {
		return 0
	}

	//return maxProfit123Helper(prices, 0, 0, 2)
	return maxProfit123DpWay(prices)

}

// 返回 从 prices[idx...] 获取的最大利润
// dealNum 表示剩余的交易次数， 卖出后算一次交易
// status 0表示没有持有，1表示 持有
func maxProfit123Helper(prices []int, idx int, status int, dealNum int) int {

	if idx >= len(prices) || dealNum == 0 {
		return 0
	}

	var (
		noOp  int
		buyIn int
		sell  int
	)

	noOp = maxProfit123Helper(prices, idx+1, status, dealNum)

	if status == 0 {
		// 这里表示可以买入 有买入的机会
		buyIn = maxProfit123Helper(prices, idx+1, 1, dealNum) - prices[idx]
	} else if status == 1 {
		// 表示已经卖出了
		sell = maxProfit123Helper(prices, idx+1, 0, dealNum-1) + prices[idx]
	}

	return tools.Max(noOp, tools.Max(buyIn, sell))
}

func maxProfit123DpWay(prices []int) int {

	// dp 第一维度表示 每天的利润，
	// 第二维度表示 当天的状态 0 表示不持有，1表示持有
	// 第三维度表示 当天是否有交易资格，共有 0，1，2 从2 -> 0 变化， 0的时候就不能买了

	dp := make([][][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 2)

		for j := 0; j < len(dp[i]); j++ {
			// 这里用到了 2 这个值 所以要给3
			dp[i][j] = make([]int, 3)
		}

	}

	for i := len(dp) - 2; i >= 0; i-- {

		/*
			有交易次数
				买入
				卖出
			没有交易次数
				持有
		*/
		// 有交易次数
		// 当前未持有 -> 已持有
		dp[i][0][1] = tools.Max(dp[i+1][1][1]-prices[i], dp[i+1][0][1])
		dp[i][0][2] = tools.Max(dp[i+1][1][2]-prices[i], dp[i+1][0][2])

		// 有交易次数
		// 保持当前状态 -> 未持有    有两次交易机会
		dp[i][1][2] = tools.Max(dp[i+1][1][2], dp[i+1][0][1]+prices[i])
		// 保持当前状态 -> 未持有    有1次交易机会
		dp[i][1][1] = tools.Max(dp[i+1][1][1], dp[i+1][0][0]+prices[i])

	}

	return dp[0][0][2]

}
