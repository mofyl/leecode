package v1

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {

	//s := "()"
	//s := "()[]{}"
	//s := "(}"
	//s := "([)]"
	s := "{[]}"

	res := isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	que := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {

		if len(que) > 0 && isMatch(que[len(que)-1], string(s[i])) {
			que = que[:len(que)-1]
			continue
		}

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			que = append(que, string(s[i]))
			continue
		}

		return false
	}

	if len(que) > 0 {
		return false
	}
	return true
}

func isMatch(a, b string) bool {

	if a == "{" && b == "}" {
		return true
	} else if a == "(" && b == ")" {
		return true
	} else if a == "[" && b == "]" {
		return true
	}

	return false
}
