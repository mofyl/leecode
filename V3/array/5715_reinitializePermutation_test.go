package array

import (
	"fmt"
	"testing"
)

func TestReinitializePermutation(t *testing.T) {

	n := 4

	res := reinitializePermutation(n)

	fmt.Println(res)

}

func reinitializePermutation(n int) int {

	if n < 1 {
		return 0
	}

	count := 1

	for i := 1; i < n; i++ {

		tmp := 0
		if i%2 == 0 {
			tmp = i / 2
		} else if i%2 == 1 {
			tmp = n/2 + (i-1)/2
		}

		if i != tmp {
			count++
		}

	}

	return count

}
