package str

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {

	//s := "25525511135"
	//s := "0000"
	//s := "1111"
	//s := "010010"
	s := "101023"
	res := restoreIpAddresses(s)

	fmt.Println(res)
}

func restoreIpAddresses(s string) []string {

	if len(s) < 4 {
		return nil
	}
	res := make([]string, 0)
	ipAddrHelper(s, 0, 0, []int{}, &res)

	return res
}

func ipAddrHelper(s string, start int, level int, cur []int, res *[]string) {

	if level == 4 {
		if len(cur) > 0 {
			str := strings.Builder{}
			for i := 0; i < len(cur); i++ {
				str.WriteString(strconv.Itoa(cur[i]))
				if i < len(cur)-1 {
					str.WriteString(".")
				}
			}

			if str.Len()-3 == len(s) {
				*res = append(*res, str.String())
			}

		}
		return
	}

	if start >= len(s) {
		return
	}
	// 这里表示当前start 若为0 的话那么 就只能是单独的 不能和别人组合
	if s[start] == '0' {
		// 这里没有回溯是 因为 cur 是传递的副本 后面直接return了，不会影响到上层
		cur = append(cur, 0)
		ipAddrHelper(s, start+1, level+1, cur, res)
		return
	}

	num := 0

	for i := start; i < len(s); i++ {

		num = num*10 + int(s[i]-'0')

		if num < 256 {
			cur = append(cur, num)
			ipAddrHelper(s, i+1, level+1, cur, res)
			cur = cur[:len(cur)-1]
		} else {
			break
		}

	}

}
