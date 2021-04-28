package binary_search

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {

	x := 2.0
	n := 10

	res := myPow(x, n)

	fmt.Println(res)
}

func myPow(x float64, n int) float64 {
	// 快速幂

	if n >= 0 {
		return myPowHelperV2(x, n)
	}

	return 1.0 / myPowHelperV2(x, -n)

}

func myPowHelper(x float64, n int) float64 {

	if n == 0 {
		return 1.0
	}

	y := myPowHelper(x, n/2)

	// 次数是2的整数倍 则 返回 y*y
	if n%2 == 0 {
		return y * y
	} else {
		// 基数次则 还需要乘以 x
		// 应为这里 %2  余数只能是1, 所以补一个 *x 就好了
		return y * y * x
	}

}

// 非递归版快速幂

/*
	也就是 x   x^2   x^2*x^2 = x^4	 x^4 * x^4  = x^8 等等

	将 tmp 不断的进行 相乘

	若 n 当前位 为1 则 将tmp 计入res 中

*/
func myPowHelperV2(x float64, n int) float64 {

	var res float64 = 1

	tmp := x

	for n > 0 {
		// 表示当前最低位为1 ，因为若当前最低位
		if n&1 == 1 {
			res *= tmp
		}
		tmp *= tmp
		n = n >> 1
	}
	return res
}
