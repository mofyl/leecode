package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestSplitListToParts(t *testing.T) {

	head := &tools.ListNode{}
	tools.AddNode(head, 1)
	tools.AddNode(head, 2)
	tools.AddNode(head, 3)
	tools.AddNode(head, 4)
	tools.AddNode(head, 5)
	tools.AddNode(head, 6)
	tools.AddNode(head, 7)
	tools.AddNode(head, 8)
	tools.AddNode(head, 9)
	tools.AddNode(head, 10)

	k := 3

	res := splitListToParts(head.Next, k)
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		tools.PrintNode(res[i])
		fmt.Println("===============")
	}

}

func splitListToParts(root *tools.ListNode, k int) []*tools.ListNode {

	if k == 0 {
		return nil
	}

	length := 0

	p := root

	for p != nil {
		length++
		p = p.Next
	}

	// 计算出平均的个数
	avg := length / k
	// 若有余数，那么每个位置分担1个 直到mod 为0
	mod := length % k

	res := make([]*tools.ListNode, k)
	cur := root
	for i := 0; i < k; i++ {
		count := avg
		if mod > 0 {
			count++
		}
		p = cur
		for j := 0; p != nil && j < count-1; j++ {
			p = p.Next
		}

		if p == nil {
			break
		}

		next := p.Next
		p.Next = nil
		res[i] = cur
		cur = next

		if mod > 0 {
			mod--
		}
	}

	return res
}
