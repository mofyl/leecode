package str

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	digits := "23"

	res := letterCombinations(digits)

	fmt.Println(res)
}

func letterCombinations(digits string) []string {

	if len(digits) < 1 {
		return nil
	}

	letter := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	strs := make([]string, 0, len(digits))

	for i := 0; i < len(digits); i++ {

		v, ok := letter[digits[i]]
		if ok {
			strs = append(strs, v)
		}
	}

	res := make([]string, 0)

	letterCombinationsHelper(strs, 0, "", &res)

	return res
}

func letterCombinationsHelper(strs []string, start int, curRes string, res *[]string) {

	if start == len(strs) {
		*res = append(*res, curRes)
		return
	}

	for i := 0; i < len(strs[start]); i++ {
		curRes += string(strs[start][i])
		letterCombinationsHelper(strs, start+1, curRes, res)
		curRes = curRes[:len(curRes)-1]
	}

}
