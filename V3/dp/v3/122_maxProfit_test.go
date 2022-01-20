package v3

import (
	"fmt"
	"testing"
)

func TestMaxProfit_122(t *testing.T) {
	// nums := []int{7, 1, 5, 3, 6, 4}
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{7, 6, 4, 3, 1}
	res := maxProfit_122(nums)

	fmt.Println(res)
}

// 这里求 多次交易的 最大利润，其实就是在求 prices 峰值的和
// 只要求出 prices 的 峰值和 就是最终的结果
// 但是直接求峰值不太好求，转换一下   eg： []int{1, 2, 3, 4, 5}
// 5是一个峰值， 5-1 的值 其实等于  （2-1)+(3-2)+(4-3)+(5-4) 的值
// 如此一来 每次当 price[i] > price[i-1]的 时候 我们就知道 峰值会出现(但是当前不一定是峰值)，
// 我们就去用  price[i] - price[i-1]
// 最终的和 就是结果
func maxProfit_122(prices []int) int {

	res := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}

	}

	return res
}
