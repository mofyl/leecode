package v1

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	//s1 := "ab"
	//s2 := "eidbaooo"

	s1 := "hello"
	s2 := "ooolleoooleh"

	res := checkInclusion(s1, s2)

	fmt.Println(res)
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) < 1 {
		return true
	}

	table := [26]int{}
	window := [26]int{}

	for i := 0; i < len(s1); i++ {
		table[s1[i]-'a']++
	}

	l, r := 0, 0

	for r < len(s2) {
		rC := s2[r] - 'a'

		window[rC]++

		for window[rC] > table[rC] {
			window[s2[l]-'a']--
			l++
		}

		if r-l+1 == len(s1) {
			return true
		}
		r++
	}

	return false

}
