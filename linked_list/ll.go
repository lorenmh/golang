// Doubly Linked-List implementation in Go

package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func CreateList() *List {
	return &List{nil, nil}
}

func (l *List) initialInsert(value int) {
	inserted := &Node{value, nil, nil}
	l.head = inserted
	l.tail = inserted
}

func (l *List) Append(value int) {
	if l.head == nil {
		l.initialInsert(value)
		return
	}

	inserted := &Node{value, l.tail, nil}
	l.tail.next = inserted
	l.tail = inserted
}

func (l *List) Prepend(value int) {
	if l.head == nil {
		l.initialInsert(value)
		return
	}

	inserted := &Node{value, nil, l.head}
	l.head.prev = inserted
	l.head = inserted
}

func (l *List) Print() {
	node := l.head

	for node != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}

	fmt.Println("nil")
}

func main() {
	list := CreateList()

	for i := 0; i < 20; i += 1 {
		list.Append(i)
	}

	for i := 0; i < 20; i += 1 {
		list.Prepend(i)
	}

	list.Print()
}
