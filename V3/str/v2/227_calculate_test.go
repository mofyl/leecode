package v2

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {

	//s := "3+2*2"

	s := " 3+5 / 2 "

	res := calculate(s)

	fmt.Println(res)
}

func calculate(s string) int {

	s = strings.ReplaceAll(s, " ", "")

	num := 0
	var lastOp byte = '+'
	que := make([]int, 0)
	for i := 0; i < len(s); i++ {

		if isNumber(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		if !isNumber(s[i]) || i == len(s)-1 {
			// 不是 数字，或者 到了最后一个字符的时候
			switch lastOp {
			case '+':
				que = append(que, num)
			case '-':
				que = append(que, -num)
			case '*':
				num1 := que[len(que)-1]
				que[len(que)-1] = num1 * num
			case '/':
				num1 := que[len(que)-1]
				que[len(que)-1] = num1 / num
			}
			lastOp = s[i]
			num = 0
		}

	}

	res := 0
	for i := 0; i < len(que); i++ {
		res += que[i]
	}

	return res
}

func isNumber(b byte) bool {

	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
