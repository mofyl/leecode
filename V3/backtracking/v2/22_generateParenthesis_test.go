package v2

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	n := 3

	res := generateParenthesis(n)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {

	if n < 1 {
		return nil
	}
	res := make([]string, 0, n)
	generateParenthesisHelper(n, n, "", &res)
	return res
}

func generateParenthesisHelper(l, r int, cur string, res *[]string) {

	if l == 0 && r == 0 {
		*res = append(*res, cur)
		return
	}
	// 通过画 递归树 分析，可得 右括号比左括号多的时候 都不合法
	if l > r {
		return
	}

	if l > 0 {
		generateParenthesisHelper(l-1, r, cur+"(", res)
	}

	if r > 0 {
		generateParenthesisHelper(l, r-1, cur+")", res)
	}

}
