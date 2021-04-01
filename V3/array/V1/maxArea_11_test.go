package V1

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	num := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(num))
}

func maxArea(height []int) int {

	start, end, res := 0, len(height)-1, 0
	for start < end {
		w := end - start
		h := 0
		if height[start] < height[end] {
			h = height[start]
			start++
		} else {
			h = height[end]
			end--
		}

		area := w * h
		if area > res {
			res = area
		}

	}
	return res
}
