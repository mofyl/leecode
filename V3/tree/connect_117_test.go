package tree

import "testing"

func TestConnect_117(t *testing.T) {
	root := &Node{Val: 1}
	l2 := &Node{Val: 2}
	l3 := &Node{Val: 3}
	l4 := &Node{Val: 4}
	l5 := &Node{Val: 5}
	l7 := &Node{Val: 7}

	root.Left = l2
	root.Right = l3

	l2.Left = l4
	l2.Right = l5

	l3.Right = l7

	connect_117(root)
}

func handler_117(cur, last, nextStart *Node) {
	// last.Next 若 上一个不为 nil，那么则 链接上一个和 当前节点
	if last != nil {
		last.Next = cur
	}
	// nextStart 为空 则表示 当前节点是第一次过来，那么就应该去记录
	if nextStart == nil {
		nextStart = cur
	}
	// 移动last
	last = cur
}

/*
	这里是从 root节点开始遍历，每次遍历都是去为 下一层的next赋值
	然后 使用 last 和 nextStart 来做层序遍历
*/
func connect_117(root *Node) *Node {

	if root == nil {
		return root
	}

	start := root

	var last, nextStart *Node

	do := func(cur *Node) {
		// last.Next 若 上一个不为 nil，那么则 链接上一个和 当前节点
		if last != nil {
			last.Next = cur
		}
		// nextStart 为空 则表示 当前节点是第一次过来，那么就应该去记录
		if nextStart == nil {
			nextStart = cur
		}
		// 移动last
		last = cur
	}
	for start != nil {
		// last 表示 上一个已经遍历过的 node。nextStart 表示 下一层第一个节点
		last = nil
		nextStart = nil
		for p := start; p != nil; p = p.Next {
			if p.Left != nil {
				do(p.Left)
			}
			if p.Right != nil {
				do(p.Right)
			}
		}
		start = nextStart
	}

	return root
}
