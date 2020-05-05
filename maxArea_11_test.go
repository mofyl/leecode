package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}

func maxArea(height []int) int {
	n := len(height)

	if n < 2 {
		return height[0]
	}

	left := 0
	right := n - 1
	res := -1
	num1 := 0
	mul := 0
	for left < right {

		dis := right - left
		num1 = height[left]
		if height[left] > height[right] {
			num1 = height[right]
			// 这里是让 较低的那边去移动: 既然要面积最大，那么就应该是最长的宽和最长的高
			right--
		} else {
			left++
		}
		mul = dis * num1

		if res < mul {
			res = mul
		}
	}

	return res

}
