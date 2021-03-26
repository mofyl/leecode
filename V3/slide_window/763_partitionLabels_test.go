package slide_window

import "testing"

func TestPartitionLabels(t *testing.T) {

}

func partitionLabels(S string) []int {

	if len(S) < 1 {
		return nil
	}

	table := [26]int{}

	for i := 0; i < len(S); i++ {
		table[S[i]-'a']++
	}

	res := make([]int, 0)

	l, r := 0, 0

	for r < len()

}
