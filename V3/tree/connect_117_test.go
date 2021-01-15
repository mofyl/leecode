package tree

import "testing"

func TestConnect_117(t *testing.T) {

}

func connect_1117(root *Node) *Node {

	if root == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}

	if root.Left != nil && root.Right == nil && root.Next != nil {
		p := root
		for p.Next != nil {
			// 右节点为空
			if p.Next.Left != nil {
				root.Left.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				root.Left.Next = p.Next.Right
				break
			}
			p = p.Next
		}

	}

	if root.Right != nil && root.Next != nil {
		p := root
		for p.Next != nil {
			// 右节点为空
			if p.Next.Left != nil {
				root.Left.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				root.Left.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}

	connect_1117(root.Left)
	connect_1117(root.Right)
	return root
}
