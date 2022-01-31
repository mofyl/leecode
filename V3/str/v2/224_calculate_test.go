package v2

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculate_224(t *testing.T) {
	//s := "(6+8)"
	//s := "(1+(4+5+2)-3)+(6+8)"
	s := "- (4 + 5)"
	res := calculate_224(s)

	fmt.Println(res)
}

func calculate_224(s string) int {

	// 去掉空格
	s = strings.ReplaceAll(s, " ", "")

	res := calculate_224Helper(&s)

	return res
}

// 第一个 int 表示 结果，第二个int 表示 下标
func calculate_224Helper(s *string) int {

	var lastOp byte = '+'
	num := 0
	que := make([]int, 0)

	for len(*s) > 0 {

		c := (*s)[0]
		*s = (*s)[1:]

		if isNumber_224(c) {
			num = num*10 + int(c-'0')
		}

		if c == '(' {
			num = calculate_224Helper(s)
		}

		if !isNumber_224(c) || len(*s) == 0 {

			switch lastOp {
			case '+':
				que = append(que, num)
			case '-':
				que = append(que, -num)
			case '*':
				num1 := que[len(que)-1]
				que = append(que, num1*num)
			case '/':
				num1 := que[len(que)-1]
				que = append(que, num1/num)
			}

			num = 0
			lastOp = c

		}

		if c == ')' {
			break
		}

	}

	res := 0
	for i := 0; i < len(que); i++ {

		res += que[i]

	}

	return res

}

func isNumber_224(b byte) bool {

	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
