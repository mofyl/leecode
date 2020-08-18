package V1

import (
	"fmt"
	"testing"
)

func TestRotateFunc(t *testing.T) {

	a := []int{4, 3, 2, 6}

	fmt.Println(maxRotateFunction(a))
}

func maxRotateFunction(A []int) int {
	// 用归纳法中的 错位相减 归纳出 F(k+1)

	aLen := len(A)

	var sumA, fK, res int64

	for i := 0; i < aLen; i++ {
		sumA += int64(A[i])
		fK += int64(i * A[i])

	}
	res = fK
	for i := aLen - 1; i > 0; i-- {
		fK = sumA + fK - int64(aLen*A[i])

		if res < fK {
			res = fK
		}
	}

	return int(res)

}
