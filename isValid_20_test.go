package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	// s := "(}"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {

	lenS := len(s)
	if lenS == 0 {
		return true
	}

	if lenS%2 != 0 {
		return false
	}

	buf := ""
	leftBrackets := make(map[string]struct{}, 3)

	leftBrackets["("] = struct{}{}
	leftBrackets["{"] = struct{}{}
	leftBrackets["["] = struct{}{}

	rightBrackets := make(map[string]struct{}, 3)
	rightBrackets[")"] = struct{}{}
	rightBrackets["}"] = struct{}{}
	rightBrackets["]"] = struct{}{}

	for i := 0; i < lenS; i++ {
		curS := string(s[i])
		if _, ok := leftBrackets[curS]; ok {
			buf += curS
		} else if _, ok := rightBrackets[curS]; ok {
			lenBuf := len(buf)
			if lenBuf == 0 {
				return false
			}
			if isMatch(string(buf[lenBuf-1]), curS) {
				buf = buf[:lenBuf-1]
			} else {
				return false
			}
		}
	}

	if len(buf) != 0 {
		return false
	}

	return true

}

func isMatch(s1, s2 string) bool {
	switch s1 {
	case "(":
		return s2 == ")"
	case "{":
		return s2 == "}"
	case "[":
		return s2 == "]"
	}

	return false
}
