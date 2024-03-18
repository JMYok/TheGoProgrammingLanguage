package main

import "fmt"

type MyLinkedList struct {
	size int
	head *LinkedNode
	tail *LinkedNode
}

type LinkedNode struct {
	val  int
	prev *LinkedNode
	next *LinkedNode
}

func Constructor() MyLinkedList {
	l := MyLinkedList{
		size: 0,
		head: initNode(-2),
		tail: initNode(-2),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.size-1 || index < 0 {
		return -1
	}
	pa := this.head.next
	cnt := 0
	for cnt < index {
		pa = pa.next
		cnt++
	}
	return pa.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		return
	}
	node := initNode(val)
	var lastNode *LinkedNode
	if index == this.size {
		lastNode = this.tail.prev

		lastNode.next = node
		node.next = this.tail
		node.prev = lastNode
		this.tail.prev = node
	} else {
		idx := 0
		idxNode := this.head.next
		for idx < index {
			idxNode = idxNode.next
			idx++
		}
		lastNode = idxNode.prev
		lastNode.next = node
		node.next = idxNode
		node.prev = lastNode
		idxNode.prev = node
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size-1 || index < 0 {
		return
	}
	pa := this.head.next
	cnt := 0
	for cnt < index {
		pa = pa.next
		cnt++
	}
	pa.prev.next = pa.next
	pa.next.prev = pa.prev
	this.size--
}

func initNode(val int) *LinkedNode {
	return &LinkedNode{val: val, next: nil, prev: nil}
}

func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(7)
	myLinkedList.AddAtHead(2)
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtIndex(3, 0)
	myLinkedList.DeleteAtIndex(2)
	myLinkedList.AddAtHead(6)
	myLinkedList.AddAtTail(4)
	fmt.Println(myLinkedList.Get(4))
}
