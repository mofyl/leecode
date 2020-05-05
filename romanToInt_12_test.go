package main

import (
	"fmt"
	"testing"
)

var (
	intMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func TestRomanToint(t *testing.T) {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	// i := len(s) - 1
	// for i >= 0 {
	//   iRoman := intMap[string(s[i])]
	//   if i > 0 {

	//     i1Roman := intMap[string(s[i-1])]
	//     if iRoman > i1Roman {

	//       res += (iRoman - i1Roman)
	//       i -= 2
	//       continue
	//     }
	//   }

	//   res += iRoman
	//   i--

	// }

	var v, last int

	/*
		由于罗马字符代表的数字应该都是从后向前增加的，所以若出现n-1位置的数字 < n位置的数字
		那么这里就出现了特殊规则，我们就应该做减法
		然后每次我们需要记录上一次的数字 用于和下一次进行比较
	*/
	for i := len(s) - 1; i >= 0; i-- {

		v = intMap[string(s[i])]

		if v < last {
			res -= v
		} else {
			res += v
		}

		last = v

	}

	return res

}
