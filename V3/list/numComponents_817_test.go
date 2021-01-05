package list

import (
	"fmt"
	"testing"
)

func TestNumComponents(t *testing.T) {
	head := &ListNode{}
	AddNode(head, 0)
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 4)

	G := []int{0, 3, 1, 4}

	res := numComponents(head.Next, G)

	fmt.Println(res)
}

func numComponents(head *ListNode, G []int) int {

	if head == nil || len(G) < 1 {
		return 0
	}
	// 使用Hash 表来加速
	gHash := make(map[int]struct{}, len(G))

	for i := 0; i < len(G); i++ {
		gHash[G[i]] = struct{}{}
	}

	p := head
	res := 0
	for p != nil {

		_, ok := gHash[p.Val]
		// 当前Val在G中 但是 下一个节点Val不在G中 那么就认为 这是一个组件
		if ok {
			// 这里若是p是最后一个节点，那么也就认为是一个组件
			if p.Next == nil {
				res++
				break
			}
			if p.Next != nil {
				_, ok := gHash[p.Next.Val]
				if !ok {
					res++
					p = p.Next
				}
			}

		}

		p = p.Next
	}
	return res
}

// 这样更优雅一点 leecode给出来的执行时间表示 这样慢了 2ms 内存多了0.2M
// 内存多了 是因为 这里使用了 bool
func numComponents_v2(head *ListNode, G []int) int {
	if head == nil || len(G) < 1 {
		return 0
	}
	// 使用Hash 表来加速
	gHash := make(map[int]bool, len(G))

	for i := 0; i < len(G); i++ {
		gHash[G[i]] = true
	}

	p := head
	res := 0
	for p != nil {

		// 当前Val在G中 但是 下一个节点Val不在G中 那么就认为 这是一个组件
		if gHash[p.Val] {
			// 这里若是p是最后一个节点，那么也就认为是一个组件
			if p.Next == nil {
				res++
				break
			}
			if p.Next != nil && !gHash[p.Next.Val] {
				res++
				p = p.Next
			}
		}

		p = p.Next
	}
	return res
}
