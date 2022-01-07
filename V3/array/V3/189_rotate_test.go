/*
 * @Author: lirui
 * @Date: 2022-01-01 12:05:46
 * @LastEditTime: 2022-01-03 14:00:14
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\189_rotate_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestRotate_189(t *testing.T) {

	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3

	// nums := []int{-1, -100, 3, 99}
	// k := 2

	nums := []int{-1}
	k := 2

	rotate_189(nums, k)

	fmt.Println(nums)

}

func rotate_189(nums []int, k int) {

	numsLen := len(nums)
	n := numsLen

	// 如果k 是 len(nums) 的整数倍，那么相当于就没有旋转
	if k%numsLen == 0 {
		return
	}

	pos := 0
	i := 0
	nextNum := nums[0]
	// 这里选择倒着来进行循环。 没有等于， > 0 才需要进行交换，
	// 正着来也是可以的 不过无法判断 起始的 nums 中 是否有元素
	// for ; n < numsLen ; n++
	for ; n > 0; n-- {

		next := (pos + k) % numsLen

		tmp := nums[next]

		nums[next] = nextNum

		nextNum = tmp
		pos = next
		// 这里如果 len(n) 是 k 的 整数倍，那么循环一圈之后一定会回来，就会出现无限循环
		// 这里直接让 pos 指向下一个，但是 nextNum 也要进行调整
		if pos == i {
			i++
			pos = i
			nextNum = nums[pos]
		}

	}

}
