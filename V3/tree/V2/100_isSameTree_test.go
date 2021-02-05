package V2

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}

	l1.Left = l2

	l11 := &TreeNode{Val: 1}
	l22 := &TreeNode{Val: 2}

	l11.Right = l22

	fmt.Println(isSameTree(l1, l11))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

	return false

}
