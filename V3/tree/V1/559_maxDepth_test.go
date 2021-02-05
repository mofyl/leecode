package V1

import (
	"fmt"
	"testing"
)

func TestMaxDepth_559(t *testing.T) {
	l1 := &NNode{Val: 1}
	l2 := &NNode{Val: 2}
	l3 := &NNode{Val: 3}
	l4 := &NNode{Val: 4}
	l5 := &NNode{Val: 5}
	l6 := &NNode{Val: 6}
	l7 := &NNode{Val: 7}
	l8 := &NNode{Val: 8}
	l9 := &NNode{Val: 9}
	l10 := &NNode{Val: 10}
	l11 := &NNode{Val: 11}
	l12 := &NNode{Val: 12}
	l13 := &NNode{Val: 13}
	l14 := &NNode{Val: 14}

	l1.Children = append(l1.Children, l2, l3, l4, l5)
	l3.Children = append(l3.Children, l6, l7)
	l4.Children = append(l4.Children, l8)
	l5.Children = append(l5.Children, l9, l10)

	l7.Children = append(l7.Children, l11)
	l11.Children = append(l11.Children, l14)

	l8.Children = append(l8.Children, l12)

	l9.Children = append(l9.Children, l13)

	deep := maxDepth_559(l1)

	fmt.Println(deep)
}

func maxDepth_559(root *NNode) int {

	if root == nil {
		return 0
	}
	deep := 0
	for i := 0; i < len(root.Children); i++ {
		curDeep := maxDepth_559(root.Children[i])
		if deep < curDeep {
			deep = curDeep
		}
	}

	return deep + 1
}
