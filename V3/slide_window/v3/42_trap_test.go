/*
 * @Author: lirui
 * @Date: 2022-01-10 14:20:24
 * @LastEditTime: 2022-01-10 14:34:47
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\42_trap_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {

	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	s := trap(nums)

	fmt.Println(s)
}

func trap(height []int) int {

	if len(height) < 1 {
		return 0
	}
	n := len(height)

	maxL := 0
	maxR := 0
	lS := 0
	rS := 0

	total := 0
	for i := 0; i < n; i++ {

		if height[i] > maxL {
			maxL = height[i]
		}

		if height[n-1-i] > maxR {
			maxR = height[n-1-i]
		}

		lS += maxL
		lS += maxR
		total += height[i]

	}

	S := len(height) * maxL

	return rS + lS - S - total
}
