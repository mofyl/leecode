package V1

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	// n1 := &TreeNode{Val: 1}
	// n2 := &TreeNode{Val: 2}
	// n3 := &TreeNode{Val: 2}
	// n4 := &TreeNode{Val: 3}
	// n5 := &TreeNode{Val: 3}
	// // n6 := &TreeNode{Val: 4}
	// // n7 := &TreeNode{Val: 4}

	// n1.Left = n2
	// n1.Right = n3

	// n2.Left = nil
	// n2.Right = n4

	// n3.Left = n5
	// n3.Right = nil

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 0}

	n1.Left = n2

	fmt.Println(isSymmetric(n1))
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetricStep(root, root)
}

func isSymmetricStep(p, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSymmetricStep(p.Left, q.Right) && isSymmetricStep(p.Right, q.Left)
}

func TestTemp(t *testing.T) {

	s := make([]int, 0, 2)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(cap(s))

	s = s[:0]
	s = append(s, 3)
	fmt.Println(cap(s))
	fmt.Println(s)
}
