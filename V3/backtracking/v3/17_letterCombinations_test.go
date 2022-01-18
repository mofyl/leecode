/*
 * @Author: lirui
 * @Date: 2022-01-08 16:06:46
 * @LastEditTime: 2022-01-08 16:23:43
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\17_letterCombinations_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	digits := ""

	res := letterCombinations(digits)

	fmt.Println(res)
}

func letterCombinations(digits string) []string {

	if digits == "" {
		return []string{}
	}

	digitsMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	in := make([]string, 0, len(digits))

	for i := 0; i < len(digits); i++ {

		v, ok := digitsMap[digits[i]]
		if !ok {
			return nil
		}

		in = append(in, v)
	}

	res := make([]string, 0)

	letterCombinationsHelper(in, 0, "", &res)

	return res
}

func letterCombinationsHelper(in []string, idx int, cur string, res *[]string) {

	if idx == len(in) {
		*res = append(*res, cur)
		return
	}

	for i := 0; i < len(in[idx]); i++ {
		letterCombinationsHelper(in, idx+1, cur+string(in[idx][i]), res)
	}

}
