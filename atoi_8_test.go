package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("-91283472332"))
}

func myAtoi_v1(str string) int {

	reg, err := regexp.Compile("^[+-]?\\d+")

	if err != nil {
		fmt.Println("err is ", err)
		return 0
	}
	temp := reg.FindString(strings.TrimLeft(str, " "))

	num, _ := strconv.ParseInt(temp, 10, 64)
	var res int64

	if num > math.MaxInt32 {
		res = math.MaxInt32
	} else {
		res = num
	}

	if res < math.MinInt32 {
		res = math.MinInt32
	}
	return int(res)

}

func myAtoi(str string) int {

	str = strings.TrimLeft(str, " ")
	if len(str) < 1 {
		return 0
	}
	var sum int
	switch {
	case isNum(str[0]):
		for i := 0; i < len(str) && isNum(str[i]); i++ {
			sum = sum*10 + int(str[i]-'0')
			if sum > math.MaxInt32 {
				return math.MaxInt32
			}
		}
	case str[0] == '+':
		for i := 1; i < len(str) && isNum(str[i]); i++ {
			sum = sum*10 + int(str[i]-'0')
			if sum > math.MaxInt32 {
				return math.MaxInt32
			}
		}

	case str[0] == '-':
		for i := 1; i < len(str) && isNum(str[i]); i++ {
			sum = sum*10 - int(str[i]-'0')
			if sum < math.MinInt32 {
				return math.MinInt32
			}
		}
	default:
		return 0
	}

	return sum
}

func isNum(i byte) bool {
	return '0' <= i && i <= '9'
}
