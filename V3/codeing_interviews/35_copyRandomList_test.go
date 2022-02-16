package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {

	l1 := &Node{Val: 1}
	l2 := &Node{Val: 2}

	l1.Next = l2
	l2.Random = l2

	resH := copyRandomListV2(l1)

	fmt.Println(resH)
}

func copyRandomList(head *Node) *Node {

	resH := &Node{}
	res := resH

	p := head
	nodeToRandom := make(map[*Node]*Node)
	oldNodeToNewNode := make(map[*Node]*Node)

	for p != nil {
		tmp := &Node{Val: p.Val}

		res.Next = tmp

		nodeToRandom[p] = p.Random
		oldNodeToNewNode[p] = tmp

		res = res.Next
		p = p.Next
	}

	for k, v := range nodeToRandom {

		if v == nil {
			continue
		}

		newNode := oldNodeToNewNode[k]
		randomNode := oldNodeToNewNode[v]

		newNode.Random = randomNode

	}

	return resH.Next

}

func copyRandomListV2(head *Node) *Node {

	p := head

	oldNodeToNewNode := make(map[*Node]*Node)

	for p != nil {
		oldNodeToNewNode[p] = &Node{Val: p.Val}
		p = p.Next
	}
	p = head
	for p != nil {

		oldNodeToNewNode[p].Next = oldNodeToNewNode[p.Next]
		oldNodeToNewNode[p].Random = oldNodeToNewNode[p.Random]
		p = p.Next
	}

	//res := oldNodeToNewNode[head]
	return oldNodeToNewNode[head]
}
