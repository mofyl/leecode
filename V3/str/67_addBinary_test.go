package str

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	//a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	//b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	a := "1"
	b := "11"
	res := addBinary(a, b)
	fmt.Println(res)
}

func addBinary(a string, b string) string {

	if len(a) < 1 {
		return b
	}

	if len(b) < 1 {
		return a
	}
	sum := 0
	ca := 0
	res := ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
		sum = ca

		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		// 计算 sum 当前进位后的 值为 sum%2
		// 计算 num 在 n 进制下 进位后的值为 num % n 因为若num>n 则需要进位，那么当前位的数字就是  num%n
		// 需要进位的数字 ca = num / n
		res = fmt.Sprintf("%d%s", sum%2, res)
		ca = sum / 2

		i--
		j--
	}

	if ca > 0 {
		res = fmt.Sprintf("%d%s", ca, res)
	}
	return res
}

// 第一个返回的是当前的和  ，第二个返回的是 当前的进位
func addBinaryHelper(a, b, carry byte) (byte, byte) {
	tmp1 := a - '0'
	tmp2 := b - '0'
	tmp3 := carry - '0'
	c := tmp1 + tmp2 + tmp3

	return c % 2, c / 2

}
