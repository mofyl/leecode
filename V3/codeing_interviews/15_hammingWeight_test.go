package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {

	num := uint32(4294967293)

	res := hammingWeight(num)

	fmt.Println(res)
}

func hammingWeight(num uint32) int {

	res := 0
	for num != 0 {

		if num&1 == 1 {
			res++
		}

		num = num >> 1
	}

	return res

}
