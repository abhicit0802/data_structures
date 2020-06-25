package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

type LinkedList interface {
	Add(int)
	AddAll([]int)
	Print()
	IsPresent(int) bool
	PrintReverse()
}

func New() *linkedList {
	return new(linkedList)
}

func (l *linkedList) Add(data int) {
	if l.head == nil {
		l.head = &node{data, nil}
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &node{data, nil}
}

func (l *linkedList) AddAll(vals []int) {
	p := l.head
	for {
		if p == nil || p.next == nil {
			break
		}
		p = p.next
	}
	for _, val := range vals {
		if p == nil {
			p = &node{val, nil}
			continue
		}
		p.next = &node{val, nil}
		p = p.next
	}
}

func (l *linkedList) Print() {
	p := l.head
	for {
		if p == nil {
			return
		}
		fmt.Print(p.data, " ")
		p = p.next
	}
}

func (l *linkedList) IsPresent(data int) bool {
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

func (l *linkedList) PrintReverse() {
	reverse(l.head)
}

func reverse(p *node) {
	if p == nil {
		return
	}
	reverse(p.next)
	fmt.Print(p.data, " ")
}
