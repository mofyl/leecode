package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	思路：
		先去找起点，也就是相同的节点
			找到后就从该节点开始递归剩下的所有节点，
			这里若是重复节点很多 那么效率就很差

	isSubPath 就是在找相同的节点
	dfs就是在从这个节点开始遍历 剩下的树
*/

func isSubPath(head *ListNode, root *TreeNode) bool {
	// 该函数负责找起点
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)

}

func dfs(listNode *ListNode, treeNode *TreeNode) bool {

	if listNode == nil {
		return true
	}

	if treeNode == nil {
		return false
	}

	if listNode.Val != treeNode.Val {
		return false
	}

	return dfs(listNode.Next, treeNode.Left) || dfs(listNode.Next, treeNode.Right)
}
