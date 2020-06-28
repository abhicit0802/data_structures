package linkedlist

import "fmt"

type doublyNode struct {
	data int
	next *doublyNode
	prev *doublyNode
}

type doublyList struct {
	head *doublyNode
	tail *doublyNode
}

func NewDoublyList() *doublyList {
	return &doublyList{}
}

func (dl *doublyList) Add(val int) {
	if dl.head == nil {
		dl.head = &doublyNode{
			data: val,
		}
		dl.tail = dl.head
		return
	}
	p := dl.head
	for {
		if p.next == nil {
			p.next = &doublyNode{
				data: val,
				prev: p,
			}
			dl.tail = p.next
			return
		}
		p = p.next
	}
}

func (dl *doublyList) AddAll(vals []int) {
	p := dl.tail
	for _, val := range vals {
		if p == nil {
			p = &doublyNode{
				data: val,
			}
			dl.tail = p
			dl.head = p
			continue
		}
		p.next = &doublyNode{
			data: val,
			prev: p,
		}
		p = p.next
	}
	dl.tail = p
}

func (dl *doublyList) Print() {
	p := dl.head
	for {
		if p == nil {
			return
		}
		fmt.Print(p.data, " ")
		p = p.next
	}
}

func (dl *doublyList) IsPresent(val int) bool {
	p := dl.head
	for {
		if p == nil || p.data != val {
			return false
		} else if p.data == val {
			return true
		}
		p = p.next
	}
}

func (dl *doublyList) PrintReverse() {
	p := dl.tail
	for {
		if p == nil {
			return
		}
		fmt.Print(p.data, " ")
		p = p.prev
	}
}
