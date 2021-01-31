package V1

import "testing"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func TestConnect(t *testing.T) {

}

func connect_v2(root *Node) *Node {

	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}

	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connect_v2(root.Left)
	connect_v2(root.Right)

	return root
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}
	connect_loop(root.Left, root.Right)
	return root
}

// 这里是去链接两个 相邻的子节点
func connect_loop(node1, node2 *Node) {

	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	connect_loop(node1.Left, node1.Right)
	connect_loop(node2.Left, node2.Right)
	// 处理 跨父节点的 情况
	connect_loop(node1.Right, node2.Left)
}
