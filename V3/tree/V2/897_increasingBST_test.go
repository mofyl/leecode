package V2

import "testing"

func TestIncreasingBST(t *testing.T) {

}

var (
	cur *TreeNode
)

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	ans := &TreeNode{}
	cur = ans

	increasingBSTLoop(root)

	return ans.Right
}

func increasingBSTLoop(root *TreeNode) {

	if root == nil {
		return
	}

	increasingBSTLoop(root.Left)

	root.Left = nil
	cur.Right = root
	cur = root

	increasingBSTLoop(root.Right)
}
