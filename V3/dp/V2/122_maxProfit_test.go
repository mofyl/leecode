package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit_122(t *testing.T) {

	num := []int{7, 1, 5, 3, 6, 4}
	//num := []int{1, 2, 3, 4, 5}
	//num := []int{7, 6, 4, 3, 1}

	res := maxProfit_122(num)
	fmt.Println(res)
}

func maxProfit_122(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	//dp := make([][]int, len(prices))
	//
	//// 初始化 dp
	//for i := 0; i < len(prices); i++ {
	//	// 0 表示 今天没有买入，1表示今天买入了
	//	dp[i] = make([]int, 2)
	//}

	maxPro := maxProfit122Helper(prices, 0, 0, false)

	return maxPro

}

//
//func dpWay(prices []int) int {
//
//	if len(prices) < 2 {
//		return 0
//	}
//
//	dp := make([][]int, len(prices))
//
//	// 初始化 dp
//	for i := 0; i < len(prices); i++ {
//		// 0 表示 今天没有买入，1表示今天买入了
//		dp[i] = make([]int, 2)
//	}
//
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//
//	for i := len(prices); i > 0; i-- {
//
//		// 今天没有买入的 最大利润
//		dp[i][0] = tools.Max(dp[i-1][0])
//
//	}
//
//}

// 暴力递归会超时
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
