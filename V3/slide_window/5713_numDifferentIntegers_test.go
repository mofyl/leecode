package slide_window

import (
	"fmt"
	"testing"
)

func TestNumDifferentIntegers(t *testing.T) {

	word := "a1b01c001"
	res := numDifferentIntegers(word)

	fmt.Println(res)
}

func numDifferentIntegers(word string) int {

	if len(word) < 1 {
		return len(word)
	}

	num := -1
	res := make(map[int]struct{}, 10)
	for i := 0; i < len(word); {
		num = -1
		if isDigit(word[i]) {
			num = 0
			for i < len(word) && isDigit(word[i]) {
				num = num*10 + int(word[i]-'0')
				i++
			}
			if num != -1 {
				res[num] = struct{}{}
			}
		} else {
			i++
		}

	}

	return len(res)
}

func isDigit(c uint8) bool {

	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
