package codeing_interviews

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestPermutation(t *testing.T) {

	s := "aab"
	res := permutation(s)

	fmt.Println(res)

}

func permutation(s string) []string {

	res := make([]string, 0)
	used := make([]bool, len(s))

	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	permutationHelper(string(b), 0, "", &res, used)

	return res
}

func permutationHelper(s string, idx int, cur string, res *[]string, used []bool) {

	if idx >= len(s) {
		sb := strings.Builder{}
		sb.WriteString(cur)

		*res = append(*res, sb.String())

		return
	}

	for i := 0; i < len(s); i++ {

		if used[i] {
			continue
		}

		if i > 0 && s[i] == s[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		permutationHelper(s, idx+1, cur+string(s[i]), res, used)
		used[i] = false
	}

}
