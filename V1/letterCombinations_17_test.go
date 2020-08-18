package V1

import (
	"fmt"
	"testing"
)

func TestLeetComb(t *testing.T) {
	// str := "23"
	str := ""

	fmt.Println(letterCombinations(str))

}

func letterCombinations(digits string) []string {
	var numMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res := make([]string, 0)
	// 回溯算法
	backtrace(digits, 0, "", numMap, &res)

	return res
}

func backtrace(digits string, deep int, builder string, numMap map[string]string, res *[]string) {

	if deep == len(digits) {
		if builder != "" {
			*res = append((*res), builder)

		}
		return
	}

	numStr, ok := numMap[string(digits[deep])]

	if ok {

		for i := 0; i < len(numStr); i++ {
			builder += string(numStr[i])
			backtrace(digits, deep+1, builder, numMap, res)
			builder = builder[:len(builder)-1]
		}

	}

}
