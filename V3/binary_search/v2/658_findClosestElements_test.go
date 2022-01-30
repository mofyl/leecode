package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFindClosestElements(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5}
	k := 4
	x := -1

	res := findClosestElements(nums, k, x)

	fmt.Println(res)
}

func findClosestElements(arr []int, k int, x int) []int {

	// 这里也是使用二分法
	// 最后二分到 窗口中 只有 k 个元素就好
	// 每次我们去计算 a = nums[r] - x , b =  nums[l]-x 的大小
	// 若 a > b 表示当前的 r 有些大， 还可以再小一点

	l, r := 0, len(arr)-1

	for (r - l + 1) > k {

		a := tools.Abs(int32(x) - int32(arr[l]))
		b := tools.Abs(int32(x) - int32(arr[r]))

		if a > b {
			l++
		} else {
			r--
		}
	}

	return arr[l : l+k]

}
