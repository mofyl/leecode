package V2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTree2str(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4

	fmt.Println(tree2str(l1))
}

func tree2str(t *TreeNode) string {

	b := strings.Builder{}

	tree2strLoop(t, &b)

	return b.String()
}

func tree2strLoop(root *TreeNode, b *strings.Builder) {

	if root == nil {
		return
	}

	b.WriteString(strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		return
	}

	b.WriteString("(")
	tree2strLoop(root.Left, b)
	b.WriteString(")")

	if root.Right != nil {
		b.WriteString("(")
		tree2strLoop(root.Right, b)
		b.WriteString(")")
	}

}
