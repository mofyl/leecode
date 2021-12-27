package V3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func prevorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	prevOrderTraversalLoop(root, &res)
	return res
}

func prevOrderTraversalLoop(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	prevOrderTraversalLoop(root.Left, res)
	prevOrderTraversalLoop(root.Right, res)
}
