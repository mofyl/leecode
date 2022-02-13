package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {

	n := 7

	res := fib(n)

	fmt.Println(res)
}

func fib(n int) int {

	m1 := 0
	m2 := 1

	if n == 0 {
		return m1
	}

	if n == 1 {
		return m2
	}

	for i := 2; i <= n; i++ {
		tmp := (m1 + m2) % 1000000007
		m1 = m2
		m2 = tmp
	}

	return m2
}
