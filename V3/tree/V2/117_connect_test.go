package V2

import (
	"testing"
)

func TestConnect_117(t *testing.T) {

}

func connect_117(root *Node) *Node {

	if root == nil {
		return root
	}

	var lastNode *Node
	var nextStart *Node

	node := root

	f := func(cur *Node) {

		if lastNode != nil {
			lastNode.Next = cur
		}

		if nextStart == nil {
			nextStart = cur
		}

		lastNode = cur
	}

	for node != nil {
		lastNode = nil
		nextStart = nil
		for p := node; p != nil; p = p.Next {

			if p.Left != nil {
				f(p.Left)
			}

			if p.Right != nil {
				f(p.Right)
			}

		}
		node = nextStart
	}
	return root
}
