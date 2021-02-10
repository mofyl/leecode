package V2

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	l3.Left = l1
	l3.Right = l4

	l1.Right = l2

	res := kthSmallest(l3, 4)

	fmt.Println(res)
}

func kthSmallest(root *TreeNode, k int) int {

	if root == nil {
		return 0
	}
	prev := 0

	kthSmallestHelper(root, &k, &prev)

	return prev
}

func kthSmallestHelper(root *TreeNode, k *int, prev *int) {

	if root == nil {
		return
	}

	if *k == 0 {
		return
	}

	kthSmallestHelper(root.Left, k, prev)
	if *k == 0 {
		return
	}
	*prev = root.Val
	*k--

	kthSmallestHelper(root.Right, k, prev)
	if *k == 0 {
		return
	}
}
