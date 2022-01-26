package v1

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	//haystack := "hello"
	//needle := "ll"
	//haystack := "aaaaa"
	//needle := "bba"
	//
	//haystack := "aaa"
	//needle := "aaaa"

	haystack := "mississippi"
	needle := "issip"

	res := strStr(haystack, needle)

	fmt.Println(res)
}

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	if haystack == "" {
		return -1
	}

	if len(needle) > len(haystack) {
		return -1
	}
	// 暴力破解
	//for i := 0; i < len(haystack); {
	//	j := 0
	//	start := i
	//	if haystack[i] == needle[j] {
	//		for start < len(haystack) && j < len(needle) && haystack[start] == needle[j] {
	//			start++
	//			j++
	//			if j == len(needle) {
	//				return i
	//			}
	//		}
	//
	//	}
	//	i++
	//}
	//return -1

	// 这里表示 i只需要移动 目标串和模式串 的差值就好了
	// 多余的比较没有用
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 这里直接比较
		if haystack[i:i+len(needle)] == needle {
			return i
		}

	}
	return -1
}
