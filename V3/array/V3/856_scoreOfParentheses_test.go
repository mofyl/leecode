/*
 * @Author: lirui
 * @Date: 2022-01-07 16:48:56
 * @LastEditTime: 2022-01-07 17:37:58
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\856_scoreOfParentheses_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {

	// s := "()"

	s := "(()(()))"

	res := scoreOfParentheses(s)

	fmt.Println(res)

}

func scoreOfParentheses(s string) int {

	stack := make([]interface{}, 0)

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, string(s[i]))
		} else {
			// 表示 s[i] 是 ')'

			//  若 stack[len(stack)-1] 不是 '('
			sum := 0
			for stack[len(stack)-1] != "(" {
				sum += stack[len(stack)-1].(int)
				stack = stack[0 : len(stack)-1]
			}

			// 这里不是 append, 前面操作完成后 stack[len(stack)-1] 位置一定是 '('
			// 直接等于 就把  '(' 覆盖了，也达到了 消除'(' 的作用
			if sum == 0 {
				// 说明 栈顶是  '('
				stack[len(stack)-1] = int(1)
			} else {
				// 栈顶不是 '(' 是数字，那么就是(A) 这种情况
				// 那么 A 翻倍
				stack[len(stack)-1] = int(sum * 2)
			}

		}

	}

	sum := 0

	for i := 0; i < len(stack); i++ {
		sum += stack[i].(int)
	}

	return sum

}
