package str

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	//s := []byte("hello")
	s := []byte("Hannah")
	reverseString(s)

	fmt.Println(string(s))
}

func reverseString(s []byte) {

	if len(s) <= 1 {
		return
	}

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			s[left] = s[left] ^ s[right]
			s[right] = s[left] ^ s[right]
			s[left] = s[left] ^ s[right]
		}

		left++
		right--
	}

}
