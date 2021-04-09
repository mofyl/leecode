package V2

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	//nums := []int{7, 1, 5, 3, 6, 4}
	nums := []int{7, 6, 4, 3, 1}

	res := maxProfit(nums)
	fmt.Println(res)
}

// 该算法的思路就是 找到 最小的买入价格，然后 用最小的买入价格 来算
func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	minBuyIn := prices[0]
	finalPro := 0

	for i := 1; i < len(prices); i++ {

		// 这里表示 以当前 prices[i]为 卖出价格，以minV为买入价格 获取的利润
		profit := prices[i] - minBuyIn
		// 若 当前的买入价格 比 当前 prices[i] 低 表示 以prices[i] 为买入价格
		// 可能获取的利润更高
		if minBuyIn > prices[i] {
			minBuyIn = prices[i]
		}
		// 若之前的利润 比 当前利润低 就更换
		if finalPro < profit {
			finalPro = profit
		}

	}

	return finalPro

}
