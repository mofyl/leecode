package V2

import (
	"fmt"
	"math"
	"testing"
)

func TestMaxProfit_121(t *testing.T) {
	// s := []int{7, 1, 5, 3, 6, 4}
	s := []int{2, 4, 1}
	// s := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit1(s))
}
func maxProfit1(prices []int) int {

	// 最小的值
	min := math.MaxInt64
	// 最大的差值
	maxPro := 0

	// 这里每次用当前元素和最小值进行减法，获取最大的差值
	// 然后每次去更新最小的值

	for i := 0; i < len(prices); i++ {

		profit := prices[i] - min

		if profit > maxPro {
			maxPro = profit
		}

		if min > prices[i] {
			min = prices[i]
		}

	}

	return maxPro

}
