package list

import (
	"fmt"
	"testing"
)

func TestSplitListToParts(t *testing.T) {

	head := &ListNode{}
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 4)
	AddNode(head, 5)
	AddNode(head, 6)
	AddNode(head, 7)
	AddNode(head, 8)
	AddNode(head, 9)
	AddNode(head, 10)

	k := 3

	res := splitListToParts(head.Next, k)
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		PrintNode(res[i])
		fmt.Println("===============")
	}

}

func splitListToParts(root *ListNode, k int) []*ListNode {

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

	res := make([]*ListNode, k)
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
