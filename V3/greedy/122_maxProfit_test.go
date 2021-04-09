package V2

import (
	"fmt"
	"testing"
)

func TestMaxProfit_122(t *testing.T) {

	num := []int{7, 4, 1, 2, 3, 1}
	//num := []int{1, 2, 3, 4, 5}
	//num := []int{7, 6, 4, 3, 1}

	res := maxProfit_122(num)
	fmt.Println(res)
}

// 贪心策略：每次使用今天 -  昨天 , 计算得到的利润， 若利润 大于0 则 累加，其他情况不管
func maxProfit_122(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	curPro := 0
	res := 0
	for i := 1; i < len(prices); i++ {

		curPro = prices[i] - prices[i-1]

		if curPro > 0 {
			res += curPro
		}

	}

	return res
}
