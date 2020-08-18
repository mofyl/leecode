package V1

import "container/list"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//var ele = make([]int, 0)

func preorderTraversal(root *TreeNode) []int {

	ele := make([]int, 0)

	preorderTree(root, &ele)

	return ele
}

func preorderTree(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}

	(*res) = append((*res), root.Val)

	preorderTree(root.Left, res)
	preorderTree(root.Right, res)

}

func inorderTraversal(root *TreeNode) []int {
	ele := make([]int, 0)

	inorderTree(root, &ele)

	return ele
}

func inorderTree(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderTree(root.Left, res)
	(*res) = append((*res), root.Val)
	inorderTree(root.Right, res)
}

func postorderTraversal(root *TreeNode) []int {
	ele := make([]int, 0)

	postorderTree(root, &ele)

	return ele
}

func postorderTree(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postorderTree(root.Left, res)
	postorderTree(root.Right, res)
	(*res) = append((*res), root.Val)
}

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

	l := list.New()

	if root != nil {
		l.PushBack(root)
	}

	for l.Len() != 0 {
		curLen := l.Len()
		level := make([]int, 0, curLen)
		for i := 0; i < curLen; i++ {
			ele := l.Front()
			l.Remove(ele)
			v := ele.Value.(*TreeNode)
			level = append(level, v.Val)

			if v.Left != nil {
				l.PushBack(v.Left)
			}
			if v.Right != nil {
				l.PushBack(v.Right)
			}
		}
		res = append(res, level)
	}

	return res
}

//func level(root *TreeNode) [][]int {
//
//}
//
//func levelOrderRecursion(root *TreeNode) []int {
//
//	if root == nil {
//		return nil
//	}
//
//}
