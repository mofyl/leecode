package V2

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {

	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(h))
}

func maxArea(height []int) int {

	if len(height) < 2 {
		return height[0]
	}

	l := 0
	r := len(height) - 1

	res := 0
	dis := 0
	num1 := 0
	mul := 0
	// 由于n最小等于 所以这里没有等于
	for l < r {
		dis = r - l
		num1 = height[l]
		if height[l] < height[r] {
			l++
		} else {
			num1 = height[r]
			r--
		}

		mul = dis * num1
		if mul > res {
			res = mul
		}
	}

	return res

}
