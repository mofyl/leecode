package v1

import (
	"fmt"
	"sort"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	nums := []int{1, 2}
	limit := 3

	res := numRescueBoats(nums, limit)
	fmt.Println(res)
}

func numRescueBoats(people []int, limit int) int {

	if len(people) < 1 {
		return 0
	}

	// 要求最小的船数， 那么就要先将小的重量 拉走
	// 所以这里需要排序
	// 这里所有的 prople[i] <= limit
	sort.Ints(people)

	r := 0
	count := 0
	curSum := 0
	for r < len(people) {

		curSum += people[r]

		if curSum > limit {
			// 记录次数
			count++
			curSum = people[r]
		}

		r++

	}

	return count
}
