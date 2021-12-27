/*
 * @Author: lirui38
 * @Date: 2021-12-25 16:26:24
 * @LastEditTime: 2021-12-25 18:06:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\117_connect_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

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
	l7 := &Node{Val: 7}

	l2.Left = l4
	l2.Right = l5

	l3.Right = l7

	l1.Left = l2
	l1.Right = l3

	connect(l1)

	fmt.Print("asd")
}

/*
	每次处理一层
	这里引入了每层的 头结点 stepHead ， 方便处理
	q 开始指向 stepHead ， 每处理一个节点，q 就向右 滑动一个

*/
func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	head := root

	stepHead := &Node{}

	for head != nil {

		q := stepHead

		for p := head; p != nil; p = p.Next {

			if p.Left != nil {
				q.Next = p.Left
				q = p.Left
			}

			if p.Right != nil {
				q.Next = p.Right
				q = p.Right
			}
		}

		head = stepHead.Next
		stepHead.Next = nil

	}

	return root

}
