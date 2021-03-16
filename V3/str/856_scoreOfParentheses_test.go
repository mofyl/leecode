package str

import "testing"

func TestScoreOfParentheses(t *testing.T) {

}

func scoreOfParentheses(S string) int {

}

func scoreOfParenthesesHelper(S string, start int, prev uint8) int {

	if start == len(S) {
		return 0
	}

	if prev == '(' && S[start] == ')' {
		return 1 + scoreOfParenthesesHelper(S, start+1, S[start])
	} else if prev == '(' && S[start] == '(' {
		return 2 * scoreOfParenthesesHelper(S, start+1, S[start])
	} else if prev == ')' && S[start] == '(' {
		return scoreOfParenthesesHelper(S, start+1, S[start])
	}

}
