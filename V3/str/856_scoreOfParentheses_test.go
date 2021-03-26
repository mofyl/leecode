package str

import (
	"fmt"
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {

	num := scoreOfParentheses("(()(()()))")
	fmt.Println(num)
}

func scoreOfParentheses(S string) int {

	stack := make([]interface{}, 0, len(S))

	for i := 0; i < len(S); i++ {

		if S[i] == '(' {
			stack = append(stack, '(')
		} else {
			// S[i]	== ')'
			sum := 0
			for stack[len(stack)-1] != '(' {
				sum += stack[len(stack)-1].(int)
				stack = stack[:len(stack)-1]
			}

			if sum == 0 {
				// 表示栈顶是 '(', 此时应该:将栈顶元素pop，然后将1 放入栈中
				stack[len(stack)-1] = 1
			} else {
				stack[len(stack)-1] = 2 * sum
			}

		}

	}

	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i].(int)
	}

	return sum
}
