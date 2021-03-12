package str

import (
	"fmt"
	"math"
	"testing"
)

func TestSplitIntoFibonacci(t *testing.T) {

	//res := splitIntoFibonacci("123456579")
	res := splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511")
	//res := splitIntoFibonacci("123")
	fmt.Println(res)
}

func splitIntoFibonacci(S string) []int {

	cur := make([]int, 0)
	if len(S) < 2 {
		return cur
	}

	res := splitIntoFibonacciHelper(S, 0, &cur)

	if res {
		return cur
	}

	return []int{}
}

// TODO 没做完
func splitIntoFibonacciHelper(S string, start int, cur *[]int) bool {

	if start == len(S) && len(*cur) > 2 {
		return true
	}

	if S[start] == '0' {

	}

	num := 0
	for i := start; i < len(S); i++ {

		// 这里表示 01 这种情况  i>start 就是保证是01。
		// 若 单独是 S[start] == '0'，起始为 0 也是可以的
		if S[start] == '0' && i > start {
			break
		}

		num = num*10 + int(S[i]-'0')

		if num > math.MaxInt64 {
			break
		}

		curLen := len(*cur)

		if curLen > 0 && num < (*cur)[curLen-1] {
			continue
		}

		if curLen >= 2 && num > (*cur)[curLen-1]+(*cur)[curLen-2] {
			break
		}

		if curLen < 2 || num == (*cur)[curLen-1]+(*cur)[curLen-2] {
			*cur = append(*cur, num)
			res := splitIntoFibonacciHelper(S, i+1, cur)
			if res {
				return res
			}
			*cur = (*cur)[:len(*cur)-1]
		}

	}

	return false

}
