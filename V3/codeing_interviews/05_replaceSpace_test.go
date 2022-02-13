package codeing_interviews

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplaceSpace(t *testing.T) {

	s := "We are happy."

	res := replaceSpace(s)

	fmt.Println(res)
}

func replaceSpace(s string) string {

	sL := len(s)
	res := strings.Builder{}
	for i := 0; i < sL; i++ {

		if s[i] == ' ' {
			res.WriteString("%20")
			continue
		}

		res.WriteByte(s[i])

	}

	if res.Len() == 0 {
		return s
	}
	return res.String()
}
