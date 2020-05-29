package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Println(divide(7, -3))
}

func divide(dividend int, divisor int) int {

	if dividend == 0 {
		return 0
	}

	var sum float64
	idx := 0
	divisorAbs := math.Abs(float64(divisor))
	dividendAbs := math.Abs(float64(dividend))
	for sum < dividendAbs {
		sum += divisorAbs

		if sum >= math.MaxInt32 {
			return math.MaxInt32
		}

		idx++
	}

	var res int

	if sum > dividendAbs {
		res = idx - 1
	} else {
		res = idx
	}

	if dividend*divisor < 0 {
		return 0 - res
	}

	return res

}
