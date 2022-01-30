package v2

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {

	x := 2.00000
	n := -2

	res := myPow(x, n)

	fmt.Println(res)
}

// 快速幂结论: 要求n次方，那么只需要在 n 的每个二进制位 为 1 的位置 将 tmp 乘进去就可以了
func myPow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}

	if n > 0 {
		return myPowHelper(x, n)
	} else {
		return 1.0 / myPowHelper(x, -n)
	}

}

func myPowHelper(x float64, n int) float64 {
	var res float64 = 1

	tmp := x

	for n > 0 {

		if n&1 == 1 {
			res *= tmp
		}

		tmp *= tmp
		n = n >> 1

	}

	return res
}
