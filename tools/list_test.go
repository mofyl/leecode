package tools

import (
	"leecode/V3/list"
	"testing"
)

func TestListNode_AddNode(t *testing.T) {

	l := list.NewList()
	l.AddNode(12)
	l.PrintList()
}
