package V1

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtraceParenthesis(n, n, "", &res)
	return res
}

func backtraceParenthesis(l, r int, builder string, res *[]string) {

	if r < l {
		return
	}

	if r < 0 || l < 0 {
		return
	}

	if l == 0 && r == 0 {
		*res = append(*res, builder)
		return
	}

	builder += "("
	backtraceParenthesis(l-1, r, builder, res)
	builder = builder[:len(builder)-1]

	builder += ")"
	backtraceParenthesis(l, r-1, builder, res)
	builder = builder[:len(builder)-1]
}
