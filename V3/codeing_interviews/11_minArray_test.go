package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestMinArray(t *testing.T) {

	//nums := []int{3, 4, 5, 1, 2}

	//nums := []int{2, 2, 2, 0, 1}

	nums := []int{3, 1, 3}

	res := minArray_v2(nums)

	fmt.Println(res)
}

func minArray(numbers []int) int {

	// 这里直接找到峰值，峰值的下一个就是最小的元素

	max := 0

	for i := 1; i < len(numbers); i++ {

		if numbers[i] >= numbers[max] {
			max = i
		} else {
			break
		}

	}

	// 若等于最后一个元素 则表示 当前数据是升序的，第一个元素是最小的
	if max == len(numbers)-1 {
		return numbers[0]
	}

	return numbers[max+1]

}

func minArray_v2(numbers []int) int {

	l := 0
	r := len(numbers) - 1

	for l < r {

		mid := l + (r-l)>>1

		if numbers[mid] > numbers[r] {
			// 这里说明 mid 左边是 旋转后的数字， 右边是 未旋转的部分， 最小的元素在右边
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			// 说明 [l , mid] 有序  ，最小的在左边, 这里不能有 -1
			r = mid
		} else {
			// 说明 numbers[mid] ==  numbers[r]
			r--
		}

	}

	return numbers[l]

}
