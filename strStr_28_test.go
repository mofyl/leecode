package main

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {

	next := getNext(needle)

	i, j := 0, 0

	for i < len(haystack)-1 && j < len(needle)-1 {
		if j == -1 || needle[i] == haystack[j] {
			i++
			j++
		} else {
			j = next[j]
		}

	}

	return j
}

func getNext(needle string) []int {

	next := make([]int, len(needle))

	i := 0
	next[0] = -1
	j := -1
	for i < len(needle)-1 {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = next[j]
		} else {
			next[i] = next[j]
		}
	}

	return next

}
