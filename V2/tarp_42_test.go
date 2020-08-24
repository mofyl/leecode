package V2

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	s := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(s))
}

/*

	这里使用了求面积的方式, 先从左到右求的sLeft ，在从右到左求得sRight
	https://leetcode-cn.com/problems/trapping-rain-water/solution/wei-en-tu-jie-fa-zui-jian-dan-yi-dong-10xing-jie-j/

*/
func trap(height []int) int {

	sLeft := 0
	sRight := 0

	maxLeft := 0
	maxRight := 0
	n := len(height)
	for i := 0; i < n; i++ {

		if height[i] > maxLeft {
			maxLeft = height[i]
		}

		if height[n-i-1] > maxRight {
			maxRight = height[n-i-1]
		}

		sLeft += maxLeft
		sRight += maxRight

	}

	// s积水 = sLeft + sRight - s矩形 - s柱子
	return sLeft + sRight - n*maxLeft - sum(height)

}

func sum(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		res += height[i]
	}
	return res
}
