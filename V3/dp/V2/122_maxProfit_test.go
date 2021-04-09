package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_122(t *testing.T) {

	//num := []int{7, 1, 5, 3, 6, 4}
	//num := []int{1, 2, 3, 4, 5}
	num := []int{7, 6, 4, 3, 1}

	res := maxProfit_122(num)
	fmt.Println(res)
}

func maxProfit_122(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	//maxPro := maxProfit122Helper(prices, 0, 0, false)
	//maxPro := maxProfit122HelperV2(prices, 0, 0)
	maxPro := maxProfit122DpWay(prices)

	return maxPro

}

// 递归函数意义：从 第 idx天 到 len(prices)天 在 status 状态下 获取的最大利润
// status 0 表示不持有， 1表示持有
func maxProfit122HelperV2(prices []int, idx int, status int) int {

	// 当 到了 len(prices) 天后  在往后 idx就越界了，就无法获取利润了
	// 所以返回的利润为0
	if idx == len(prices) {
		return 0
	}

	// 每天可以选择 持有，可以选择 买入 卖出

	// 每次都可以选择不变，
	noOpProfit := maxProfit122HelperV2(prices, idx+1, status)
	var (
		sellProfit int
		buyProfit  int
	)
	if status == 1 {
		// 表示已经持有了股票
		// 那么这里就可以卖出了 得到了卖出的利润
		sellProfit = maxProfit122HelperV2(prices, idx+1, 0) + prices[idx]
	} else {
		// 这里表示已经没有持有股票了，那么这里就可以买入
		buyProfit = maxProfit122HelperV2(prices, idx+1, 1) - prices[idx]
	}

	return tools.Max(noOpProfit, tools.Max(sellProfit, buyProfit))
}

func maxProfit122DpWay(prices []int) int {

	dp := make([][]int, len(prices)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	// 当 idx == len(price) 的时候 无论什么状态都是0
	// 由于 0 是默认值 就不需要给了
	// 0 位置表示不持有 股票
	// 1 位置 表示持有股票
	for i := len(dp) - 2; i >= 0; i-- {
		// 当前不持有股票的利润是
		// i+1天 保持不持有股票的利润， 和 i+1天持有获取的利润- price[i]
		// i+1天持有了股票 说明今天买入了，就要 减去 price[i]  买入股票是有价格的，这个价格不是手续费
		dp[i][0] = tools.Max(dp[i+1][0], dp[i+1][1]-prices[i])
		// 当前持有股票的利润是 : i+1天 卖出股票获取的利润+ price[i]  和 i+1天继续持有获取的利润的最大值
		// i+1天没有持有股票 说明 i天已经将股票卖出了，所以就要  +price[i]
		dp[i][1] = tools.Max(dp[i+1][1], dp[i+1][0]+prices[i])

	}

	return dp[0][0]

}

// 这里就是一个具有后效性的列子，idx天 的状态 和 buyIn有依赖关系，这里是不行的
// 应该是 给定一个 idx 和 status 就可以确定今天的利润
// status 表示昨天的状态 true 表示昨天买入， false 表示昨天卖出
// 该函数的意思是 在 idx...len-1 范围内 选择 买入或者卖出  得到的最大利润
func maxProfit122Helper(prices []int, idx int, buyIn int, status bool) int {

	if idx >= len(prices) {
		return 0
	}

	noOp := maxProfit122Helper(prices, idx+1, buyIn, status)
	buyOp := 0
	sellOp := 0
	// 表示没有买入
	if !status {
		// 那么当前就可以买入也可以不买入
		// 当前不买入
		// 当前买入
		buyOp = maxProfit122Helper(prices, idx+1, prices[idx], true)
	} else {
		// 表示已经买入了

		// 那么当前可以卖出 也可以不卖出, 但是卖出是有条件的
		// 当前的价格比买入的低 才能卖出

		// 选择卖出
		if prices[idx] > buyIn {
			curPro := prices[idx] - buyIn
			nextSell := maxProfit122Helper(prices, idx+1, 0, false)
			sellOp = curPro + nextSell
		}
	}

	return tools.Max(tools.Max(noOp, buyOp), sellOp)

}
