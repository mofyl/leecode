package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	// nums := []int{7, 1, 5, 3, 6, 4}

	nums := []int{7, 6, 4, 3, 1}
	res := maxProfit(nums)

	fmt.Println(res)
}

func maxProfit(prices []int) int {

	res := 0
	minPrices := 100000

	for i := 0; i < len(prices); i++ {

		minPrices = tools.Min(prices[i], minPrices)

		res = tools.Max(res, prices[i]-minPrices)

	}

	return res
}
