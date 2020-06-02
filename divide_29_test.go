package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	// fmt.Println(2 ^ 4)
	fmt.Println(divide(-2147483648, 1))
}

func divide(dividend int, divisor int) int {

	isNegative := (dividend ^ divisor) < 0

	dividendA := abs(dividend)
	divisorA := abs(divisor)

	if dividend == 0 {
		return 0
	}

	// 做特殊处理
	if dividendA == 2147483648 && divisorA == 1 {
		if dividend < 0 && divisor < 0 {
			return int(dividendA) - 1
		} else {
			return dividend
		}
	}

	/*
		n << i 就相当于 n * 2^i
		n >> i 就相当于 n / 2^i

		本题原理：
			被除数= 商 * 除数 + 余数  => 除数 = 被除数 / 商 + 余数
		我们这里就是去求 被除数 / 商 >= 除数 , 说明当前的商小了，需要增加,直到
			被除数 / 商 < 除数 此时的商就对了
		通过不断的去逼近除数 来检验 商
			只不过我们这里的商是  2^i 次方

		这里直接从31 开始是为了 减少循环次数
	*/

	bit := 0

	for i := 0; divisorA<<i < dividendA; i++ {
		bit++
	}

	var res int
	for i := bit; i >= 0; i-- {
		if (dividendA >> i) >= divisorA {
			pow := math.Pow(2, float64(i))
			res += int(pow)
			dividendA -= int(pow) * divisorA // 这里要用 2 ^ i * divisorA 的意思是 被除数 - 本次的被除数   本次的被除数 = 本次的商 * 除数
		}
	}

	if isNegative {
		return -res
	} else {
		return res
	}

}

func abs(num int) int {
	// num >> 31 ^ num 表示连同 符号位一起取反，
	// 正数右移在前面补0，负数右移补1
	/*

		对于正数： num >> 31 就变成了0， 0 对num 取异或 num不变，num-0，则还是num
		对于负数： num >> 31 每一位上都是1 16进制为0xfffff , 对num取异或，就相当于取反，这里符号位是会参与运算的。异或就是相同为0，不同为1，
		连同符号位取反后 - num>> 31，实际上就是在 +1，num>>31就变成了-1
		减去-1就变成了+1

	*/
	return (num >> 31) ^ num - num>>31

}

func divide_v1(dividend int, divisor int) int {

	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend < math.MaxInt32 {
			return -dividend
		}
		return -math.MaxInt32
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
