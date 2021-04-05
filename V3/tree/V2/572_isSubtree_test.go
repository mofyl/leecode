package V2

import (
	"fmt"
	"testing"
)

func TestIsSubTree(t *testing.T) {
	//
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	//l0 := &TreeNode{Val: 0}

	l3.Left = l4
	l3.Right = l5

	l4.Left = l1
	l4.Right = l2
	//
	//l2.Left = l0
	//
	//l41 := &TreeNode{Val: 4}
	//l11 := &TreeNode{Val: 1}
	//l21 := &TreeNode{Val: 2}
	//
	//l41.Left = l11
	//l41.Right = l21

	fmt.Println(isSubtree(l3, l4))

	//l1 := &TreeNode{Val: 1}
	//l11 := &TreeNode{Val: 1}
	//
	//l1.Left = l11
	//
	//fmt.Println(isSubtree(l1, l11))

}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var dfs func(s, t *TreeNode) bool
	dfs = func(main, sub *TreeNode) bool {

		if main == nil {
			if sub == nil {
				return true
			}
			return false
		}

		if sub != nil && main.Val == sub.Val {
			if dfs(main.Left, sub.Left) && dfs(main.Right, sub.Right) {
				return true
			}
		}

		if main.Val == t.Val {
			if dfs(main.Left, t.Left) && dfs(main.Right, t.Right) {
				return true
			}
		}

		return dfs(main.Left, t) || dfs(main.Right, t)

	}

	return dfs(s, t)
}

func isSubtree_V2(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		return isSubTreeHelper(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
	return isSubtree_V2(s.Left, t) || isSubtree_V2(s.Right, t)
}

func isSubTreeHelper(s *TreeNode, t *TreeNode) bool {

	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		return isSubTreeHelper(s.Left, t.Left) && isSubTreeHelper(s.Right, t.Right)
	}
	return false
}
