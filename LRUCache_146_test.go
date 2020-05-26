package main

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	// h := node{}

	// h.InsertHead(1, 1)
	// n := h.InsertHead(2, 2)
	// h.InsertHead(3, 3)
	// h.Print()
	// fmt.Println("==========")
	// h.removeWithNode(n)
	// h.Print()

	// lru := Constructor(2)

	// // lru.Put(2, 2)
	// // 每次Get 都会影响key的活跃度 拿的越多的key 活跃度就越高，越是在前面
	// fmt.Println(lru.Get(2))
	// lru.Put(2, 6)

	// fmt.Println(lru.Get(1))

	// lru.Put(1, 5)
	// lru.Put(1, 2)

	// fmt.Println(lru.Get(1))
	// fmt.Println(lru.Get(2))

	lru := Constructor(2)

	// lru.Put(2, 2)
	// 每次Get 都会影响key的活跃度 拿的越多的key 活跃度就越高，越是在前面
	lru.Put(2, 1)

	lru.Put(1, 1)

	lru.Put(2, 3)
	lru.Put(4, 1)

	fmt.Println("==========")
	lru.Print()
	fmt.Println("==========")
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))

	// fmt.Println(lru.Get(3))

	// fmt.Println(lru.Get(4))

}

type node struct {
	key  int
	val  int
	pre  *node
	next *node
	last *node
}

func (h *node) Print() {
	l := h
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}

}

func (h *node) removeWithNode(n *node) {

	n.pre.next = n.next

	if n.next != nil {
		n.next.pre = n.pre
	}

}

func (h *node) removeWithVal(key int) {
	l := h
	for l != nil {

		if l.key == key {
			l.pre.next = l.next

			if l.next != nil {
				l.next.pre = l.pre
			}

			return
		}

		l = l.next
	}
}

func (h *node) Update(n *node) {
	// 若当前节点已经是第一个则直接返回
	if h.next.key == n.key {
		return
	}

	// 若当前节点是最后一个 修改最后一个
	if n.key == h.last.key {
		h.last = n.pre
	}

	n.pre.next = n.next
	if n.next != nil {
		n.next.pre = n.pre
	}

	n.next = h.next
	h.next.pre = n
	n.pre = h
	h.next = n

}

func (h *node) InsertHead(key, val int) *node {

	n := &node{
		key: key,
		val: val,
	}

	second := h.next

	if second == nil {
		h.next = n
		n.pre = h
		// 记录最后一个
		h.last = n
		return n
	}

	n.next = second
	n.pre = h
	second.pre = n

	h.next = n

	return n

}

// 双向链表+ hashmap
type LRUCache struct {
	capacity int
	d        *node
	m        map[int]*node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		d:        &node{},
		m:        make(map[int]*node),
	}

	return lru
}

func (this *LRUCache) Print() {
	this.d.Print()

	fmt.Println("=================")

	for k, v := range this.m {
		fmt.Println(k, v)
	}

}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if ok {
		this.d.Update(v)
		return v.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	mLen := len(this.m)

	v, ok := this.m[key]

	if ok {
		v.val = value
		this.d.Update(v)
		return
	}

	if mLen >= this.capacity && this.d.last != nil {

		v, ok := this.m[this.d.last.key]

		if ok {
			this.d.last = v.pre
			this.d.removeWithNode(v)
			delete(this.m, v.key)
		}

	}

	n := this.d.InsertHead(key, value)
	this.m[key] = n
	return

}
