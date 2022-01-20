package v3

import (
	"fmt"
	"testing"
)

func TestChange(t *testing.T) {

	coins := []int{1, 2, 5}
	amount := 5

	res := change(amount, coins)

	fmt.Println(res)
}

func change(amount int, coins []int) int {
	used := make([]bool, len(coins))
	return changeHelper(amount, coins, used)
}

func changeHelper(amount int, coins []int, used []bool) int {

	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}

	nums := 0
	for i := 0; i < len(coins); i++ {

		if !used[i] {
			continue
		}
		used[i] = true
		nums += changeHelper(amount-coins[i], coins, used)
		used[i] = false
	}

	return nums

}
