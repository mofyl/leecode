package v1

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//numbers := []int{2, 7, 11, 15}
	//target := 9

	//numbers := []int{2, 3, 4}
	//target := 6

	numbers := []int{-1, 0}
	target := -1
	res := twoSum(numbers, target)

	fmt.Println(res)
}

func twoSum(numbers []int, target int) []int {

	if len(numbers) < 1 {
		return nil
	}
	// 基于数据有序，使用双指针
	l, r := 0, len(numbers)-1

	for l < r {

		if numbers[l] > target {
			return nil
		}
		tmp := numbers[l] + numbers[r]
		if tmp > target {
			r--
		} else if tmp < target {
			l++
		} else {
			// 这里+1 是因为 题目要求下标是从1开始计数的
			return []int{l + 1, r + 1}
		}

	}

	return nil

}
