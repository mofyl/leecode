package codeing_interviews

import "testing"

func TestsingleNumber(t *testing.T) {

}

func singleNumber(nums []int) int {

	data := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		data[nums[i]]++
	}

	for k, v := range data {
		if v == 1 {
			return k
		}
	}

}
