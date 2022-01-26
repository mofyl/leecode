package v1

import (
	"fmt"
	"strings"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	//num := 3
	//num := 4
	//num := 58
	num := 1994

	res := intToRoman(num)

	fmt.Println(res)
}

func intToRoman(num int) string {

	romanInt := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanStr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""

	for i := 0; i < len(romanInt); i++ {

		if num == 0 {
			break
		}
		if romanInt[i] <= num {
			// 若 romanInt[i] 比num 大 那么count就为0
			count := num / romanInt[i]
			num = num - romanInt[i]*count

			for count > 0 {
				res = strings.Join([]string{res, romanStr[i]}, "")
				count--
			}
		}
	}

	return res

}
