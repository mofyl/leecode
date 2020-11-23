package V3

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	area := trap(height)
	fmt.Println(area)
}

func trap(height []int) int {
	// 这里采用面积来算
	// 利用重复面积来消除 除了雨水外 多余的部分

	sL := 0
	maxL := 0
	sR := 0
	maxR := 0

	n := len(height)
	for i := 0; i < n; i++ {
		if height[i] > maxL {
			maxL = height[i]
		}
		if height[n-1-i] > maxR {
			maxR = height[n-1-i]
		}
		// 若是上面没有大于的 那么就会一直保持 最大的那个数一直加
		sL += maxL
		sR += maxR
	}
	// 上面会有重复计算的面积
	// S积水面积 = S1 + S2 - 矩形面积 - 柱子面积
	return sL + sR - maxL*n - sum(height)
}

func sum(s []int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			res += s[i]
		}
	}

	return res
}
