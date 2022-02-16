package codeing_interviews

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetLeastNumbers(t *testing.T) {

	nums := []int{3, 2, 1}
	k := 2

	//nums := []int{0, 1, 2, 1}
	//k := 1

	res := getLeastNumbers(nums, k)

	fmt.Println(res)
}

func getLeastNumbers(arr []int, k int) []int {

	if k == 0 || len(arr) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	l, r := 0, len(arr)-1
	target := k - 1

	for l < r {

		midIdx := getLeastNumbersHelper(arr, l, r)

		if midIdx == target {
			// 由于是 左开右闭 所以要 l+1
			return arr[:midIdx+1]
		}

		if midIdx > target {
			r = midIdx - 1
		} else {
			l = midIdx + 1
		}

	}
	// 由于是 左开右闭 所以要 l+1
	return arr[:l+1]

}

func getLeastNumbersHelper(nums []int, l, r int) int {

	midIdx := l + rand.Intn(r-l+1)

	nums[midIdx], nums[l] = nums[l], nums[midIdx]

	midN := nums[l]

	for l < r {

		for l < r && nums[r] > midN {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}

		for l < r && nums[l] < midN {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			r--
		}
	}

	nums[l] = midN

	return l

}
