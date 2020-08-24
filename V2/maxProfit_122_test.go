package V2

import (
	"fmt"
	"testing"
)

func TestMaxProfit2(t *testing.T) {
	// s := []int{7, 1, 5, 3, 6, 4}
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(maxProfit2(s))
}

func maxProfit2(prices []int) int {

	if prices == nil {
		return 0
	}
	sLen := len(prices)
	if sLen < 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		tmp := prices[i+1] - prices[i]
		if tmp > 0 {
			res += tmp
		}

	}

	return res

}
