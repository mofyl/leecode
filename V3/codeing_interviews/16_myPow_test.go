package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {

	//x := 2.00000
	//n := 10

	//x := 2.10000
	//n := 3

	x := 2.0
	n := -2

	res := myPow(x, n)

	fmt.Println(res)
}

// 快速幂
func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	if n > 0 {
		return myPowHelper(x, n)
	} else {
		return 1.0 / myPowHelper(x, -n)
	}

}

func myPowHelper(x float64, n int) float64 {

	tmp := x
	var res float64 = 1
	for n > 0 {

		if n&1 == 1 {
			res *= tmp
		}

		tmp *= tmp

		n = n >> 1

	}

	return res

}
