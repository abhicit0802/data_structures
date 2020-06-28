package linkedlist

import "fmt"

type singlyNode struct {
	data int
	next *singlyNode
}

type singlyLinkedList struct {
	head *singlyNode
}

func NewSinglyList() *singlyLinkedList {
	return new(singlyLinkedList)
}

func (l *singlyLinkedList) Add(data int) {
	if l.head == nil {
		l.head = &singlyNode{data, nil}
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &singlyNode{data, nil}
}

func (l *singlyLinkedList) AddAll(vals []int) {
	p := l.head
	for {
		if p == nil || p.next == nil {
			break
		}
		p = p.next
	}
	for _, val := range vals {
		if p == nil {
			p = &singlyNode{val, nil}
			continue
		}
		p.next = &singlyNode{val, nil}
		p = p.next
	}
}

func (l *singlyLinkedList) Print() {
	p := l.head
	for {
		if p == nil {
			return
		}
		fmt.Print(p.data, " ")
		p = p.next
	}
}

func (l *singlyLinkedList) IsPresent(data int) bool {
	p := l.head
	for {
		if p == nil {
			return false
		}
		if p.data == data {
			return true
		}
		p = p.next
	}

}

func (l *singlyLinkedList) PrintReverse() {
	reverse(l.head)
}

func reverse(p *singlyNode) {
	if p == nil {
		return
	}
	reverse(p.next)
	fmt.Print(p.data, " ")
}
