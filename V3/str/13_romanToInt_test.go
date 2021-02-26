package str

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	//s := "III"
	//s := "IV"
	//s := "IX"
	//s := "LVIII"
	s := "MCMXCIV"
	res := romanToInt(s)

	fmt.Println(res)
}

func romanToInt(s string) int {

	roman := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	ans := 0
	flag := false
	for i := 0; i < len(s); {
		flag = false

		if i+1 < len(s) {
			v, ok := roman[s[i:i+2]]
			if ok {
				ans += v
				i += 2
				flag = true
			}
		}

		if !flag {
			v, ok := roman[string(s[i])]
			if ok {
				ans += v
				i++
			}
		}

	}

	return ans
}
