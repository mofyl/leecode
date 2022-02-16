package codeing_interviews

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var table = map[string]string{

	"0":  "a",
	"1":  "b",
	"2":  "c",
	"3":  "d",
	"4":  "e",
	"5":  "f",
	"6":  "g",
	"7":  "h",
	"8":  "i",
	"9":  "j",
	"10": "k",
	"11": "l",
	"12": "m",
	"13": "n",
	"14": "o",
	"15": "p",
	"16": "q",
	"17": "r",
	"18": "s",
	"19": "t",
	"20": "u",
	"21": "v",
	"22": "w",
	"23": "x",
	"24": "y",
	"25": "z",
}

func TestTranslateNum(t *testing.T) {

	num := 12258

	//num := 12

	res := translateNum(num)

	fmt.Println(res)
}

func translateNum(num int) int {

	s := strconv.FormatInt(int64(num), 10)

	res := make([]string, 0)
	translateNumHelper(s, 0, "", &res)

	return len(res)
}

func translateNumHelper(s string, idx int, cur string, res *[]string) {

	if idx >= len(s) {
		tmp := strings.Builder{}
		tmp.WriteString(cur)
		*res = append(*res, tmp.String())
		return
	}

	// 只选择自己
	v := table[string(s[idx])]
	translateNumHelper(s, idx+1, cur+v, res)

	// 尝试和后面一个字符结合
	if idx < len(s)-1 {
		// 包含 idx+1 这个字符， 所以是 idx+2
		v, ok := table[s[idx:idx+2]]
		if ok {
			translateNumHelper(s, idx+2, cur+v, res)
		}
	}

}
