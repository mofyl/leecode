/*
 * @Author: lirui
 * @Date: 2022-01-03 16:26:07
 * @LastEditTime: 2022-01-04 13:47:56
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\list.go
 */

package codeing_interviews

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNode(l *ListNode) {
	if l == nil {
		fmt.Println("nil")
	}
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
