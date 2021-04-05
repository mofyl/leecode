package V2

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {

	s := "06"

	fmt.Println(numDecodings(s))

}

func numDecodings(s string) int {

	if s == "" {
		return 0
	}

	table := map[string]struct{}{

		"1":  struct{}{},
		"2":  struct{}{},
		"3":  struct{}{},
		"4":  struct{}{},
		"5":  struct{}{},
		"6":  struct{}{},
		"7":  struct{}{},
		"8":  struct{}{},
		"9":  struct{}{},
		"10": struct{}{},
		"11": struct{}{},
		"12": struct{}{},
		"13": struct{}{},
		"14": struct{}{},
		"15": struct{}{},
		"16": struct{}{},
		"17": struct{}{},
		"18": struct{}{},
		"19": struct{}{},
		"20": struct{}{},
		"21": struct{}{},
		"22": struct{}{},
		"23": struct{}{},
		"24": struct{}{},
		"25": struct{}{},
		"26": struct{}{},
	}

	count := 0
	for i := 0; i < len(s); i++ {

		left, right := i, i
		curCount := 0
		for right < len(s) {
			tmp := s[left:right]
			if left == right {
				tmp = string(s[left])
			}

			if tmp == "0" {
				// 那么这里就要和前面的数字组合才行
				if left > 0 && (s[left-1] == '1' || s[left-1] == '2') {
					curCount++
				}
			} else {
				_, ok := table[tmp]
				if ok {
					curCount++
				}
			}

			right++
		}

		if curCount != 0 {
			count++
		}

	}

	return count
}
