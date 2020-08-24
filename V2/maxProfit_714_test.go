package V2

import (
	"fmt"
	"testing"
)

func TestMaxProfit_714(t *testing.T) {
	s := []int{1, 3, 2, 8, 4, 9}

	fmt.Println(maxProfit(s, 2))
}

func maxProfit(prices []int, fee int) int {

	if prices == nil {
		return 0
	}

	sLen := len(prices)

	if sLen <= 1 {
		return 0
	}

	//  0 位置表示前一天没有持股， 1表示前一天持股
	/*

		今天的盈利状态取决于昨天的状态
		我们将昨天持股标记为1 昨天未持股标记为 0
		状态一共有三种：持有 买入 和 卖出。
		那么 状态转移方程为：
			今天持股的收益:= max(昨天持股，昨天未持股但是今天买进了)
			今天未持股的收益 := max(昨天未持股，昨天持股但是今天卖出了)

			在买入和卖出的时候才会有收益
			买入的时候 我们的本金是减少了的 ，所以要减去
			卖出的时候 我们的本金会增加，所以是增加
			这里还有手续费用，在买入和卖出 随便那个地方减去就好了
			dp[k][1] := max (dp[k-1][1] , dp[k-1][0]-prices[k])
			dp[k][0] := max (dp[k-1][0] , dp[k-1][1]+prices[k])

	*/
	dp := []int{0, 0}               // 这里保存前一天的持股结果
	dp[1] = dp[0] - prices[0] - fee // 这里其实是 dp[0]-prices[0] -fee  由于dp[0]是0 所以省略
	for i := 1; i < sLen; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], dp[0]-prices[i]-fee) // 这里是持股 可以看到持股需要从本金中减去今天的股价
	}
	// 不持股是收益最大的，因为持股要花钱
	return dp[0]

}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
