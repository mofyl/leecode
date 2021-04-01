package V2

import (
	"fmt"
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
	maxPro := 0
	maxProfit122Helper(prices, 0, 0, 0, false, &maxPro)

	return maxPro

}

// 暴力递归会超时
// isSell true 表示已经买入了， false表示没有买入
// buyIn 买入的价格
func maxProfit122Helper(prices []int, start int, buyIn int, curProfit int, isBuyIn bool, maxPro *int) {

	if start >= len(prices) {
		if *maxPro < curProfit {
			*maxPro = curProfit
		}
		return
	}

	// 表示没有买入
	if !isBuyIn {
		// 那么当前就可以买入也可以不买入

		// 当前不买入
		maxProfit122Helper(prices, start+1, 0, curProfit, false, maxPro)
		// 当前买入
		maxProfit122Helper(prices, start+1, prices[start], curProfit, true, maxPro)

	} else {
		// 表示已经买入了

		// 那么当前可以卖出 也可以不卖出, 但是卖出是有条件的
		// 当前的价格比买入的低 才能卖出

		// 选择卖出
		if prices[start] > buyIn {
			curPro := prices[start] - buyIn
			maxProfit122Helper(prices, start+1, 0, curProfit+curPro, false, maxPro)
		}

		// 选择 不卖出
		maxProfit122Helper(prices, start+1, buyIn, curProfit, true, maxPro)

	}

}
