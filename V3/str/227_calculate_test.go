package str

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	s := " 3+5 / 2 "
	//s := "1-1+1"
	//s := "0-2147483647"
	fmt.Println(calculate_v2(s))
}

// 步骤没问题 但是会超时
func calculate(s string) int {
	if s == "" {
		return 0
	}

	s = strings.ReplaceAll(s, " ", "")

	numQ := make([]int, 0, len(s))
	charQ := make([]uint8, 0, len(s))
	for i := 0; i < len(s); {

		if s[i] >= '0' && s[i] <= '9' {
			num, idx := getNum(i, s)
			numQ = append(numQ, num)
			i = idx
		} else {

			pri := getPriority(s[i])

			if pri == 2 {
				num1 := numQ[len(numQ)-1]
				numQ = numQ[:len(numQ)-1]
				c := s[i]
				i++
				num2, idx := getNum(i, s)
				tmp := calcRes(num1, num2, c)
				numQ = append(numQ, tmp)
				i = idx
			} else if pri == 1 {
				charQ = append(charQ, s[i])
				i++
			}

		}

	}
	// 这里要从左向右算
	for i := 0; i < len(charQ); i++ {

		c := charQ[i]

		num1 := numQ[0]
		numQ = numQ[1:]
		num2 := numQ[0]
		numQ = numQ[1:]

		numQ = append([]int{calcRes(num1, num2, c)}, numQ...)
	}

	return numQ[0]
}

func getPriority(a uint8) int {

	if a == '*' || a == '/' {
		return 2
	} else if a == '+' || a == '-' {
		return 1
	}
	return -1
}

func getNum(start int, s string) (int, int) {
	res := 0
	idx := start
	for ; idx < len(s); idx++ {
		if s[idx] >= '0' && s[idx] <= '9' {
			res = res*10 + int(s[idx]-'0')
		} else {
			break
		}
	}

	return res, idx
}

func calcRes(n1, n2 int, c uint8) int {

	switch c {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	}
	return 0
}

func calculate_v2(s string) int {
	if s == "" {
		return 0
	}

	s = strings.ReplaceAll(s, " ", "")

	numQ := make([]int, 0, len(s))
	// 记录的是上一个数字
	num := 0
	// 记录的是 上一次的 符号
	var lastOpt uint8 = '+'

	for i := 0; i < len(s); i++ {

		if isDigit(s[i]) {
			num = num*10 + int(s[i]-'0')

		}
		// 当遇到符号就表示上一个 子串可以计算了 每次遇到符号就去计算上一个子串
		// 若 i 为 最后一个字符，也要强制算一次
		if !isDigit(s[i]) || i == len(s)-1 {
			switch lastOpt {
			case '+':
				numQ = append(numQ, num)
			case '-':
				numQ = append(numQ, -num)
			case '*':
				num1 := numQ[len(numQ)-1]
				numQ[len(numQ)-1] = num1 * num
			case '/':
				num1 := numQ[len(numQ)-1]
				numQ[len(numQ)-1] = num1 / num
			}

			lastOpt = s[i]
			num = 0

		}

	}

	res := 0
	for i := 0; i < len(numQ); i++ {
		res += numQ[i]
	}
	return res
}

func isDigit(c uint8) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
