//go:build linux
// +build linux

package v4

import (
	"fmt"
	"strings"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	strings.Fields()
	//s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	s := "AAAAAAAAAAA"

	res := findRepeatedDnaSequences(s)

	fmt.Println(res)
}

func findRepeatedDnaSequences(s string) []string {

	/*
		以 每个 s[i] 为起点的 10个 s[i : i+10] 字符 为串，存入map中
		若 出现次数 > 1 则加入结果中
	*/
	m := make(map[string]int)
	res := make([]string, 0)
	// 这里要有 等于号
	for i := 0; i <= len(s)-10; i++ {

		tmp := s[i : i+10]
		n, ok := m[tmp]

		if ok {
			m[tmp] = n + 1
		} else {
			m[tmp] = 1
		}

	}

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res

}
