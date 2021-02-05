package V2

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}

	l1.Left = l3
	l1.Right = l2

	l3.Left = l5

	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	r4 := &TreeNode{Val: 4}
	r7 := &TreeNode{Val: 7}

	r2.Left = r1
	r2.Right = r3

	r1.Right = r4
	r3.Right = r7

	res := mergeTrees(l1, r2)
	s := inorderTraversal(res)
	fmt.Println(s)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	t1.Left = mergeTrees(t1.Left, t2.Left)

	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}
