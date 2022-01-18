/*
 * @Author: lirui
 * @Date: 2022-01-07 14:54:29
 * @LastEditTime: 2022-01-07 15:29:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\22_generateParenthesis_test.go
 */
package v3

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

	res := make([]string, 0)

	generateParenthesisHelper(n, n, "", &res)

	return res
}

func generateParenthesisHelper(left int, right int, cur string, res *[]string) {

	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		cur += "("
		// generateParenthesisHelper(left-1, right, cur+"(", res)
		generateParenthesisHelper(left-1, right, cur, res)
		cur = cur[0 : len(cur)-1]
	}

	if right > 0 {
		cur += ")"
		// generateParenthesisHelper(left, right-1, cur+")", res)
		generateParenthesisHelper(left, right-1, cur, res)
		cur = cur[0 : len(cur)-1]
	}

}
