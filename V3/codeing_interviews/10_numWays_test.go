package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestNumWays(t *testing.T) {

	n := 44

	res := numWays(n)

	fmt.Println(res)
}

func numWays(n int) int {

	//return numWaysHelper(n)
	return numWaysDP(n)
}

func numWaysHelper(n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	// 一次跳一层
	one := numWaysHelper(n - 1)
	// 一次跳两层
	two := numWaysHelper(n - 2)

	return one + two

}

func numWaysDP(n int) int {

	nums := make([]int, n+1)

	nums[0] = 1

	for i := 1; i < len(nums); i++ {

		if i-2 >= 0 {
			nums[i] = (nums[i-1] + nums[i-2]) % 1000000007
		} else if i-1 >= 0 {
			nums[i] = nums[i-1] % 1000000007
		}

	}

	return nums[n]
}
