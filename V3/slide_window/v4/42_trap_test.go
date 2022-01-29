package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestTrap(t *testing.T) {

	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	height := []int{4, 2, 0, 3, 2, 5}

	res := trap(height)

	fmt.Println(res)
}

func trap(height []int) int {

	lMaxH := 0
	rMaxH := 0
	hLen := len(height)
	sum := 0
	lArea := 0
	rArea := 0

	for i := 0; i < hLen; i++ {

		lMaxH = tools.Max(lMaxH, height[i])
		rMaxH = tools.Max(rMaxH, height[hLen-i-1])

		lArea += lMaxH
		rArea += rMaxH
		sum += height[i]
	}

	return (lArea + rArea) - (hLen * lMaxH) - sum

}
