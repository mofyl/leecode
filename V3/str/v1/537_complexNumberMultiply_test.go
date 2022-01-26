package v1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	//a := "1+1i"
	//b := "1+1i"

	a := "78+-76i"
	b := "-86+72i"

	res := complexNumberMultiply(a, b)
	fmt.Println(res)
}

func complexNumberMultiply(a string, b string) string {

	reA, imA := getComplexNum(a)
	reB, imB := getComplexNum(b)

	// (reA + imA i ) * (reB + imB i)
	//
	//num1 := reA * reB
	//num2 := reA * imB
	//num3 := imA * reB
	//num4 := -(imA * imB)
	// 这里就是最终的虚部
	imRes := reA*imB + imA*reB
	// 这里就是最终的实部
	reRes := reA*reB + (-(imA * imB))

	str := strings.Join([]string{strconv.Itoa(reRes), "+", strconv.Itoa(imRes), "i"}, "")

	return str
}

// 返回 实部 虚部 i  ; 这里最后一个返回固定是i
func getComplexNum(a string) (int, int) {

	if a == "" {
		return 0, 0
	}

	num := 0
	re := 0
	im := 0
	negIdx := false
	for i := 0; i < len(a); i++ {

		if a[i] == '-' {
			negIdx = true
		}
		if isDigit(a[i]) {
			num = num*10 + int(a[i]-'0')
		} else if a[i] == '+' {
			if negIdx {
				num = -num
			}
			re = num
			num = 0
			negIdx = false
		}

	}

	im = num
	if negIdx {
		im = -im
	}
	return re, im
}
