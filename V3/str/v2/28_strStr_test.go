package v2

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {

	//haystack := "hello"
	//needle := "ll"

	//haystack := "aaaaa"
	//needle := "bba"

	haystack := "abc"
	needle := "c"

	res := strStr(haystack, needle)

	fmt.Println(res)
}

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	if len(needle) > len(haystack) {
		return -1
	}

	needleLen := len(needle)

	// 这里一定要有等于
	// haystack="a"  needleLen="a"   若没有等于 就取不到  len(haystack)-needleLen 位置
	for i := 0; i <= len(haystack)-needleLen; i++ {

		if haystack[i:i+needleLen] == needle {
			return i
		}

	}

	return -1

}
