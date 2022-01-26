package v1

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 1

	res := generateParenthesis(n)

	fmt.Println(res)
}

func generateParenthesis(n int) []string {

	if n < 1 {
		return nil
	}

	res := make([]string, 0)

	generateParenthesisHelper("", n, n, &res)

	return res

}

func generateParenthesisHelper(cur string, left, right int, res *[]string) {

	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	if left > right {
		return
	}

	if left > 0 {
		generateParenthesisHelper(cur+"(", left-1, right, res)
	}

	if right > 0 {
		generateParenthesisHelper(cur+")", left, right-1, res)
	}

}
