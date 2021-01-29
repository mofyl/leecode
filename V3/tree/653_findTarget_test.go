package tree

import "testing"

func TestFindTarget(t *testing.T) {

}

func findTarget(root *TreeNode, k int) bool {

	if root == nil {
		return false
	}

	res := inorderTraversal(root)

	l := 0
	r := len(res) - 1

	for l < r {
		v := res[l] + res[r]
		if v > k {
			r--
		} else if v < k {
			l++
		} else {
			return true
		}

	}
	return false
}
