package tree

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	//
	//l3 := &TreeNode{Val: 3}
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l4 := &TreeNode{Val: 4}
	//
	//l3.Left = l1
	//l3.Right = l4
	//
	//l1.Right = l2
	//
	//k := 1
	//res := kthLargest(l3, k)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l5.Left = l3
	l5.Right = l6

	l3.Left = l2
	l3.Right = l4

	l2.Left = l1

	k := 3
	res := kthLargest(l5, k)

	fmt.Println(res)
}

var (
	kRes int
)

func kthLargest(root *TreeNode, k int) int {

	if root == nil || k < 1 {
		return 0
	}
	res := 0
	//kRes = k
	//kthLargest_loop(root, &res)
	curRes := 0
	kthLargest_loop_v2(root, &res, &curRes, k)
	return res
}

func kthLargest_loop(root *TreeNode, res *int) {

	if root == nil {
		return
	}

	kthLargest_loop(root.Right, res)

	if kRes == 0 {
		return
	}
	kRes--
	if kRes == 0 {
		*res = root.Val
	}

	kthLargest_loop(root.Left, res)
}

func kthLargest_loop_v2(root *TreeNode, res *int, cur *int, k int) {

	if root == nil {
		return
	}

	if *res != 0 {
		return
	}

	kthLargest_loop_v2(root.Right, res, cur, k)
	// 这里一定要先对  *cur++
	*cur++
	if *cur == k {
		*res = root.Val
		return
	}
	kthLargest_loop_v2(root.Left, res, cur, k)
}
