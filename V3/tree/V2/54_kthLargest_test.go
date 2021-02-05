package V2

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	l3.Left = l1
	l3.Right = l4

	l1.Right = l2

	res := kthLargest(l3, 1)

	fmt.Println(res)
}

func kthLargest(root *TreeNode, k int) int {

	res := 0
	tmpK := k
	kthLargestLoop(root, &tmpK, &res)
	return res
}

func kthLargestLoop(root *TreeNode, k *int, res *int) {

	if root == nil {
		return
	}

	if *k == 0 {
		return
	}

	kthLargestLoop(root.Right, k, res)

	*k--
	if *k == 0 {
		*res = root.Val
		return
	}

	kthLargestLoop(root.Left, k, res)

}
