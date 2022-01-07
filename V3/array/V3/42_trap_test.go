package V3

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {

	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	height := []int{4, 2, 0, 3, 2, 5}

	res := trap(height)

	fmt.Println(res)
}

func trap(height []int) int {
	n := len(height)

	if n < 1 {
		return 0
	}
	// 这里采用 面积的方式 来计算

	// 逐步找到左边最高的和右边最高的， 然后面积就按当前最高的来计算
	maxL := 0
	maxR := 0
	SL := 0
	SR := 0
	sumH := 0
	for i := 0; i < n; i++ {

		if height[i] > maxL {
			maxL = height[i]
		}

		if height[n-1-i] > maxR {
			maxR = height[n-1-i]
		}

		SL += maxL
		SR += maxR
		sumH += height[i]
	}

	// SL + SR = 2S柱子 + 2S积水 + S空白位置

	return SL + SR - maxL*n - sumH

}
