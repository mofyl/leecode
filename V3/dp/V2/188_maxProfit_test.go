package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit188(t *testing.T) {
	//
	prices := []int{2, 4, 1}
	k := 2
	//prices := []int{3, 2, 6, 5, 0, 3}
	//k := 2

	res := maxProfit188(k, prices)

	fmt.Println(res)
}

func maxProfit188(k int, prices []int) int {
	if len(prices) < 1 || k < 1 {
		return 0
	}
	//return maxProfit188Helper(prices, 0, 0, k)
	return maxProfit188DpWay(prices, k)
}

// k 表示交易次数 若 k == 0 则无法进行交易
// status : 0 未持有 1 表示已持有
func maxProfit188Helper(prices []int, idx int, status int, k int) int {

	if idx == len(prices) {
		return 0
	}

	if k == 0 {
		return 0
	}

	var (
		noOp  int
		buyIn int
		sell  int
	)

	// 没有买入没有卖出
	noOp = maxProfit188Helper(prices, idx+1, status, k)

	if status == 0 {
		// 当前未持有 那么就可以买入
		buyIn = maxProfit188Helper(prices, idx+1, 1, k) - prices[idx]
	} else if status == 1 {
		// 表示当前已持有
		sell = maxProfit188Helper(prices, idx+1, 0, k-1) + prices[idx]
	}

	return tools.Max(noOp, tools.Max(buyIn, sell))

}

func maxProfit188DpWay(prices []int, k int) int {

	// dp[n] 表示第n天的价格
	// dp[n][1] 表示第n天价格 对应的状态   0 表示未持有，1表示已持有
	// dp[n][1][i]  表示第n天价格 对应状态 ，然后当前的购买次数
	// 这里可以处理到 prices长度，所以要求  prices+1
	dp := make([][][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {

		dp[i] = make([][]int, 2)

		dp[i][0] = make([]int, k+1)
		dp[i][1] = make([]int, k+1)

	}

	// idx == len(prices) 的时候 为0
	// k == 0 的时候为0

	for i := len(dp) - 2; i >= 0; i-- {

		// k 为0 的时候 为0  所以直接跳过第0列
		for j := k; j >= 1; j-- {

			// 当前为没有持有股票的时候
			dp[i][0][j] = tools.Max(dp[i+1][0][j], dp[i+1][1][j]-prices[i])

			// 当前为 持有股票的时候 选择卖出
			dp[i][1][j] = tools.Max(dp[i+1][1][j], dp[i+1][0][j-1]+prices[i])

		}

	}

	return dp[0][0][k]

}
