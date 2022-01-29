package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxArea(t *testing.T) {

	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	//height := []int{1, 2, 1}

	height := []int{1, 8, 100, 2, 100, 4, 8, 3, 7}

	res := maxArea(height)

	fmt.Println(res)
}

func maxArea(height []int) int {

	l, r := 0, len(height)-1
	area := 0
	for l < r {

		h := tools.Min(height[l], height[r])
		w := r - l

		area = tools.Max(area, h*w)

		if height[l] > height[r] {
			r--
		} else if height[l] <= height[r] {
			l++
		}

	}

	return area

}
