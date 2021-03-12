package str

import (
	"fmt"
	"testing"
)

func TestCalculate_224(t *testing.T) {
	s := "(1+(4+5+2)-3)+(6+8)"

	fmt.Println(calculate_224(s))
}

func calculate_224(s string) int {
	if s == "" {
		return 0
	}

	return calcHelper(&s)
}

func calcHelper(s *string) int {

	lastNum := 0
	var lastC uint8 = '+'
	numQ := make([]int, 0)

	for len(*s) > 0 {

		c := (*s)[0]
		*s = (*s)[1:]

		if isDigit(c) {
			lastNum = lastNum*10 + int(c-'0')
		}

		if c == '(' {
			lastNum = calcHelper(s)
		}

		if (!isDigit(c) && c != ' ') || len(*s) == 0 {

			switch lastC {
			case '+':
				numQ = append(numQ, lastNum)
			case '-':
				numQ = append(numQ, -lastNum)
			}
			lastC = c
			lastNum = 0
		}

		if c == ')' {
			break
		}
	}

	res := 0
	for i := 0; i < len(numQ); i++ {
		res += numQ[i]
	}

	return res

}
