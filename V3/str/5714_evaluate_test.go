package str

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	//
	//s := "(name)is(age)yearsold"
	//knowledge := [][]string{
	//	{"name", "bob"},
	//	{"age", "two"},
	//}

	//s := "hi(name)"
	//knowledge := [][]string{
	//	{"a", "b"},
	//}

	//s := "(a)(a)(a)aaa"
	//knowledge := [][]string{
	//	{"a", "yes"},
	//}

	s := "c"
	knowledge := [][]string{}

	res := evaluate(s, knowledge)

	fmt.Println(res)
}

func evaluate(s string, knowledge [][]string) string {

	if len(s) < 1 {
		return s
	}

	knowledgeMap := make(map[string]string, len(knowledge))

	for i := 0; i < len(knowledge); i++ {
		knowledgeMap[knowledge[i][0]] = knowledge[i][1]
	}

	l, r := 0, 0
	start := 0
	b := []byte(s)
	has := false
	res := make([]byte, 0, len(b))
	for i := 0; i < len(b); i++ {

		if b[i] == '(' {
			has = true
			l = i
		} else if b[i] == ')' {
			r = i
			content := string(b[l+1 : r])
			v, ok := knowledgeMap[content]

			res = append(res, b[start:l]...)
			if !ok {
				v = "?"
			}
			res = append(res, []byte(v)...)
			start = r + 1
			fmt.Println(start)
		}

	}

	if !has {
		return s
	}

	// 表示最后一个括号后还有内容 则应该将最后一部分的内容补上来
	if r != len(b)-1 {
		res = append(res, b[r+1:]...)
	}

	return string(res)
}
