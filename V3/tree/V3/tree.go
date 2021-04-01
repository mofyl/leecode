package V3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderTraversalLoop(root, &res)
	return res
}

func inorderTraversalLoop(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}

	inorderTraversalLoop(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalLoop(root.Right, res)
}
