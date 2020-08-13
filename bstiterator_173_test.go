package main

import (
	"fmt"
	"testing"
)

// 这里从1 开始
type heap struct {
	ele    []int
	length int // 保存当前元素的位置
	cap    int
}

func newHeap() *heap {
	return &heap{
		ele:    make([]int, 10+1),
		length: 0,
		cap:    10,
	}
}

func (p *heap) IsEmpty() bool {
	return p.length == 0
}

func (p *heap) AddEle(e int) {

	if p.length >= p.cap {
		tmp := make([]int, p.cap*2+1)
		for i := 1; i <= p.length; i++ {
			tmp[i] = p.ele[i]
		}
		p.ele = tmp
		p.cap = p.cap * 2
	}

	p.length++
	pos := p.length
	for ; pos > 0 && p.ele[pos/2] > e; pos = pos / 2 {
		p.ele[pos] = p.ele[pos/2]
	}
	p.ele[pos] = e

}

func (p *heap) Pop() int {

	if p.length < 1 {
		return -1
	}

	lastEle := p.ele[p.length]
	e := p.ele[1]
	parent := 1
	child := 0
	for ; parent*2 < p.length; parent = child {
		child = parent * 2

		if p.ele[child] > p.ele[child+1] {
			child++
		}

		if p.ele[child] < lastEle {
			p.ele[parent] = p.ele[child]
		} else {
			break
		}

	}
	p.ele[parent] = lastEle
	p.length--
	return e

}

func TestHeap(t *testing.T) {

	h := newHeap()

	h.AddEle(7)
	h.AddEle(3)
	h.AddEle(15)
	h.AddEle(9)
	h.AddEle(20)
	k := 0
	for k = h.Pop(); !h.IsEmpty(); {
		fmt.Println(k)
		k = h.Pop()
	}

	fmt.Println(k)

}

type BSTIterator struct {
	// h *heap
	ele   []int
	index int
}

func Constructor_V2(root *TreeNode) BSTIterator {
	b := BSTIterator{
		// h: newHeap(),
		ele: make([]int, 0),
	}
	b.build(root)
	return b
}

func (this *BSTIterator) build(root *TreeNode) {
	if root == nil {
		return
	}

	this.build(root.Left)
	this.ele = append(this.ele, root.Val)
	this.build(root.Right)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {

	e := this.ele[this.index]
	this.index++

	return e
	// return this.h.Pop()
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	// return !this.h.IsEmpty()
	return this.index < len(this.ele)
}
