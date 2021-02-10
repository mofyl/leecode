package V2

import "testing"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func TestConnect(t *testing.T) {
	l1 := &Node{Val: 1}
	l2 := &Node{Val: 2}
	l3 := &Node{Val: 3}
	l4 := &Node{Val: 4}
	l5 := &Node{Val: 5}
	l6 := &Node{Val: 6}
	l7 := &Node{Val: 7}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4
	l2.Right = l5

	l3.Left = l6
	l3.Right = l7

	connect_v2(l1)

}

func connect_v2(root *Node) *Node {

	if root == nil {
		return root
	}

	node := root

	for node.Left != nil {
		nextStart := node.Left

		for node != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}

		node = nextStart
	}
	return root
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	connectHelper(root)
	return root
}

func connectHelper(root *Node) {

	if root == nil {
		return
	}

	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}

	connectHelper(root.Left)
	connectHelper(root.Right)

}
