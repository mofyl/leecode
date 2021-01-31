package dp

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {

}

func maxProfit(prices []int) int {

	if len(prices) < 1 {
		return 0
	}

	minV := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {

		v := prices[i] - minV

		if prices[i] < minV {
			minV = prices[i]
		}

		if v > res {
			res = v
		}

	}
	return res
}
