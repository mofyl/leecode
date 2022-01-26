package v1

import (
	"fmt"
	"strings"
	"testing"
)

func TestReverseWords(t *testing.T) {

	//s := "  hello world!  "
	s := "a good   example"

	res := reverseWords(s)
	fmt.Println(res, len(res))

}

func reverseWords(s string) string {

	if len(s) == 0 {
		return ""
	}

	word := make([]string, 0, 10)
	str := strings.Builder{}
	str.Reset()

	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			continue
		}

		if str.Len() > 0 {
			str.Reset()
		}

		if isWord(s[i]) {

			for i < len(s) && isWord(s[i]) {
				str.WriteByte(s[i])
				i++
			}

			word = append(word, str.String())

		}

	}

	res := strings.Builder{}

	for i := len(word) - 1; i >= 0; i-- {

		res.WriteString(word[i])

		if i > 0 {
			res.WriteByte(' ')
		}

	}

	return res.String()
}

func isWord(c uint8) bool {
	if c != ' ' {
		return true
	}
	return false
}
