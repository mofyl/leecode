package V1

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrefix(t *testing.T) {
	str := []string{"flower", "flow", "flight"}
	// str := []string{"cacb", "ac", "bac", "a"}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {

	lenStr := len(strs)
	if lenStr == 0 {
		return ""
	}

	if lenStr == 1 {
		return strs[0]
	}

	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) < len(strs[j]) {
			return true
		}
		return false
	})

	res := ""
	len1 := len(strs[0])
	len2 := len(strs[1])
	minNum := getMin(len1, len2)

	for i := 0; i < minNum; i++ {
		if strs[0][i] == strs[1][i] {
			res += string(strs[0][i])
		} else {
			break
		}
	}

	if res == "" {
		return res
	}

	for i := 2; i < lenStr; i++ {
		if res == "" {
			return res
		}
		for j := 0; j < len(res); j++ {
			// 这里出现了 不相同的情况 就该把 res从不相同的地方截断
			if res[j] != strs[i][j] {
				res = res[:j]
				break
			}
		}

	}

	return res

}

func getMin(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}
