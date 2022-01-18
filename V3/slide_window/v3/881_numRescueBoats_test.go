package v3

import (
	"fmt"
	"sort"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	people := []int{1, 2}
	limit := 3

	//people := []int{3, 2, 2, 1}
	//limit := 3

	//people := []int{3, 5, 3, 4}
	//limit := 5

	res := numRescueBoats(people, limit)

	fmt.Println(res)
}

func numRescueBoats(people []int, limit int) int {

	sort.Ints(people)

	l := 0
	r := len(people) - 1
	res := 0

	for l < r {

		weight := people[l] + people[r]
		if weight <= limit {
			res++
			l++
			r--
		} else if weight > limit {

			if people[l] >= limit {
				res += r - l + 1
				return res
			}

			res++
			r--
		}

	}

	if l == r && people[l] <= limit {
		res++
	}

	return res
}

func numRescueBoats_V2(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	out := 0
	for l <= r {

		weight := people[l] + people[r]

		if weight <= limit {
			l++
		}
		r--
		out++
	}
	return out
}
