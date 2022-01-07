/*
 * @Author: lirui
 * @Date: 2022-01-04 15:33:55
 * @LastEditTime: 2022-01-04 16:22:03
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\25_reverseKGroup_test.go
 */

package V1

import "testing"

func TestReverseKGroup(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	k := 1

	res := reverseKGroup(l1, k)

	PrintNode(res)

}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k <= 1 {
		return head
	}

	tH := &ListNode{}
	// 记录前一组最后一个节点
	var prev *ListNode
	start := head
	end := head
	for end != nil {

		n := 1
		// 数出 k个节点
		for n < k && end != nil {
			end = end.Next
			n++
		}
		// 若不够 k 个直接返回
		if end == nil {
			return tH.Next
		}
		// 记录 k个节点 后一个节点
		next := end.Next

		// 将 前面的节点 和 start 断开
		if prev != nil {
			prev.Next = nil
		}
		// 将end 和 后面的节点 断开
		end.Next = nil
		// 倒转 [start,end] 之间的节点
		reverseListHelper_143(start)
		// 倒转后 start  就变成了 末尾  end 就变成了开始
		if tH.Next == nil {
			tH.Next = end
		}

		// 将 前面最后一个节点和 倒转后的起始节点接上
		if prev != nil {
			prev.Next = end
		}
		// 将 倒转后的节点 和 后面的节点接上
		start.Next = next
		// 移动 prev 节点
		prev = start
		// 移动 start 和 end 节点
		start = next
		end = start
	}

	return tH.Next
}
