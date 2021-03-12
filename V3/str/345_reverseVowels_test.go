package str

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {

	//s := "hello"
	s := "leetcode"

	fmt.Println(reverseVowels(s))

}

func reverseVowels(s string) string {

	if len(s) < 2 {
		return s
	}

	volwelsMap := map[uint8]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}

	left := 0
	right := len(s) - 1

	b := []byte(s)

	for left < right {

		_, lOk := volwelsMap[b[left]]
		_, rOk := volwelsMap[b[right]]
		if lOk && rOk {

			if b[left] != b[right] {
				b[left] = b[left] ^ b[right]
				b[right] = b[left] ^ b[right]
				b[left] = b[left] ^ b[right]
			}
			left++
			right--
		} else if lOk {
			right--
		} else {
			left++
		}

	}

	return string(b)

}
